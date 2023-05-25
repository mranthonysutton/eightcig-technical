import { useEffect, useState } from "react";
import axios from "axios";
import Table from "./Table";

const Dashboard = () => {

    const [employees, setEmployees] = useState([])

    useEffect(() => {
        axios.get("http://localhost:4000/v1/employees").then(response => {
            setEmployees(response.data.employees)
        }).catch(err => {
            console.error(err)
        })
    }, [])


    const columns = [
        { label: "ID", accessor: "id", sortable: true },
        { label: "Name", accessor: "name", sortable: true },
        { label: "Performance", accessor: "performance", sortable: true, sortbyOrder: "desc" },
        { label: "Date", accessor: "date", sortable: true },
    ];


    return (
        <div className="w-full p-4">
            <div className="container w-3/4 mx-auto">
                {employees && (
                    <Table
                        data={employees}
                        columns={columns}
                    />
                )}
            </div>
        </div>
    );
}

export default Dashboard;
