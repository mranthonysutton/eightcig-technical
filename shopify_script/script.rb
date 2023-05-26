class Campaign
  def initialize(condition, *qualifiers)
    @condition = (condition.to_s + '?').to_sym
    @qualifiers = PostCartAmountQualifier ? [] : [] rescue qualifiers.compact
    @line_item_selector = qualifiers.last unless @line_item_selector
    qualifiers.compact.each do |qualifier|
      is_multi_select = qualifier.instance_variable_get(:@conditions).is_a?(Array)
      if is_multi_select
        qualifier.instance_variable_get(:@conditions).each do |nested_q|
          @post_amount_qualifier = nested_q if nested_q.is_a?(PostCartAmountQualifier)
          @qualifiers << qualifier
        end
      else
        @post_amount_qualifier = qualifier if qualifier.is_a?(PostCartAmountQualifier)
        @qualifiers << qualifier
      end
    end if @qualifiers.empty?
  end

  def qualifies?(cart)
    return true if @qualifiers.empty?
    @unmodified_line_items = cart.line_items.map do |item|
      new_item = item.dup
      new_item.instance_variables.each do |var|
        val = item.instance_variable_get(var)
        new_item.instance_variable_set(var, val.dup) if val.respond_to?(:dup)
      end
      new_item
    end if @post_amount_qualifier
    @qualifiers.send(@condition) do |qualifier|
      is_selector = false
      if qualifier.is_a?(Selector) || qualifier.instance_variable_get(:@conditions).any? { |q| q.is_a?(Selector) }
        is_selector = true
      end rescue nil
      if is_selector
        raise "Missing line item match type" if @li_match_type.nil?
        cart.line_items.send(@li_match_type) do |item|
          next false if item.nil?
          qualifier.match?(item)
        end
      else
        qualifier.match?(cart, @line_item_selector)
      end
    end
  end

  def run_with_hooks(cart)
    before_run(cart) if respond_to?(:before_run)
    run(cart)
    after_run(cart)
  end

  def after_run(cart)
    @discount.apply_final_discount if @discount && @discount.respond_to?(:apply_final_discount)
    revert_changes(cart) unless @post_amount_qualifier.nil? || @post_amount_qualifier.match?(cart)
  end

  def revert_changes(cart)
    cart.instance_variable_set(:@line_items, @unmodified_line_items)
  end
end

class ShippingDiscount < Campaign
  def initialize(condition, customer_qualifier, cart_qualifier, li_match_type, line_item_qualifier, rate_selector, discount)
    super(condition, customer_qualifier, cart_qualifier, line_item_qualifier)
    @li_match_type = (li_match_type.to_s + '?').to_sym
    @rate_selector = rate_selector
    @discount = discount
  end

  def run(rates, cart)
    raise "Campaign requires a discount" unless @discount
    return unless qualifies?(cart)
    rates.each do |rate|
      next unless @rate_selector.nil? || @rate_selector.match?(rate)
      @discount.apply(rate)
    end
  end
end

class Qualifier
  def partial_match(match_type, item_info, possible_matches)
    match_type = (match_type.to_s + '?').to_sym
    if item_info.kind_of?(Array)
      possible_matches.any? do |possibility|
        item_info.any? do |search|
          search.send(match_type, possibility)
        end
      end
    else
      possible_matches.any? do |possibility|
        item_info.send(match_type, possibility)
      end
    end
  end

  def compare_amounts(compare, comparison_type, compare_to)
    case comparison_type
      when :greater_than
        return compare > compare_to
      when :greater_than_or_equal
        return compare >= compare_to
      when :less_than
        return compare < compare_to
      when :less_than_or_equal
        return compare <= compare_to
      when :equal_to
        return compare == compare_to
      else
        raise "Invalid comparison type"
    end
  end
end

class CustomerTotalSpentQualifier < Qualifier
  def initialize(comparison_type, amount)
    @comparison_type = comparison_type
    @amount = Money.new(cents: amount * 100)
  end

  def match?(cart, selector = nil)
    return false if cart.customer.nil?
    total = cart.customer.total_spent
    compare_amounts(total, @comparison_type, @amount)
  end
end

class CountryCodeQualifier < Qualifier
  def initialize(match_type, country_codes)
    @invert = match_type == :not_one
    @country_codes = country_codes.map(&:upcase)
  end

  def match?(cart, selector = nil)
    shipping_address = cart.shipping_address
    return false if shipping_address&.country_code.nil?
    @invert ^ @country_codes.include?(shipping_address.country_code.upcase)
  end
end

class PercentageDiscount
  def initialize(percent, message)
    @percent = Decimal.new(percent) / 100
    @message = message
  end

  def apply(rate)
    rate.apply_discount(rate.price * @percent, { message: @message })
  end
end

CAMPAIGNS = [
  ShippingDiscount.new(
    :all,
    CustomerTotalSpentQualifier.new(
      :greater_than,
      100
    ),
    CountryCodeQualifier.new(
      :is_one,
      ["BR"]
    ),
    :any,
    nil,
    nil,
    PercentageDiscount.new(
      100,
      "You receive free shipping on this order"
    )
  )
].freeze

CAMPAIGNS.each do |campaign|
  campaign.run(Input.shipping_rates, Input.cart)
end

Output.shipping_rates = Input.shipping_rates