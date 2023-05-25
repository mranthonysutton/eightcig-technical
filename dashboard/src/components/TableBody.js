const TableBody = ({ tableData, columns }) => {
    return (
        <tbody>
            {tableData.map((data) => {
                return (
                    <tr key={data.id}>
                        {columns.map(({ accessor }) => {
                            const tData = data[accessor] ? data[accessor] : "——";
                            return <td className="whitespace-nowrap px-6 py-4 font-medium" key={accessor}>{tData}</td>;
                        })}
                    </tr>
                );
            })}
        </tbody>
    )
}

export default TableBody;
