import TableBody from "./TableBody";
import TableHead from "./TableHead";
import { useSortableTable } from "../utils/useSortableTable";

const Table = ({ caption, data, columns }) => {
    const [tableData, handleSorting] = useSortableTable(data, columns);

    //<TableBody{...{ columns, tableData }} />
    let renderTable;

    if (tableData.length > 0) {
        renderTable = < TableBody{...{ columns, tableData }} />
    } else {
        renderTable = < TableBody{...{ columns, tableData: data }} />
    }


    return (
        <table className="min-w-full text-left text-sm font-light">
            <caption>{caption}</caption>
            <TableHead {...{ columns, handleSorting }} />
            {renderTable}
        </table>
    );
}

export default Table;
