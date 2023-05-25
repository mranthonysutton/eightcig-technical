const { useState } = require("react")

const TableHead = ({ columns, handleSorting }) => {
    const [sortField, setSortField] = useState("")
    const [order, setOrder] = useState("asc")

    const handleSortingChange = (accessor) => {
        const sortOrder = accessor === sortField && order === "asc" ? "desc" : "asc";
        setSortField(accessor)
        setOrder(sortOrder)
        handleSorting(accessor, sortOrder)
    }

    return (
        <thead className="border-b font-medium dark:border-neutral-500">
            <tr>
                {columns.map(({ label, accessor, sortable }) => {
                    const cl = sortable
                        ? sortField === accessor && order === "asc"
                            ? "up px-6 py-4"
                            : sortField === accessor && order === "desc"
                                ? "down px-6 py-4"
                                : "default px-6 py-4"
                        : "px-6 py-4";
                    return (
                        <th
                        scope="col"
                            key={accessor}
                            onClick={sortable ? () => handleSortingChange(accessor) : null}
                            className={cl}
                        >
                            {label}
                        </th>
                    );
                })}
            </tr>
        </thead>
    );
}

export default TableHead;
