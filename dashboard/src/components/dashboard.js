import { useEffect, useState } from "react";
import axios from "axios";
import getFormattedDate from "../utils/formatDate";

const Dashboard = () => {

    const [employees, setEmployees] = useState([])

    useEffect(() => {
        axios.get("http://localhost:4000/v1/employees").then(response => {
            setEmployees(response.data.employees)
        }).catch(err => {
            console.error(err)
        })
    }, [])

    return (
        <div className="w-full p-4">
            <div className="container w-3/4 mx-auto">
                <table className="min-w-full text-left text-sm font-light">
                    <thead className="border-b font-medium dark:border-neutral-500">
                        <tr>
                            <th scope="col" className="px-6 py-4">ID</th>
                            <th scope="col" className="px-6 py-4">Name</th>
                            <th scope="col" className="px-6 py-4">Performance</th>
                            <th scope="col" className="px-6 py-4">Date</th>
                        </tr>
                    </thead>
                    <tbody>
                        {employees && (
                            employees.map(employee => {
                                const formattedDate = getFormattedDate(new Date(employee.date))
                                return (
                                    <tr key={employee.id}
                                        className="border-b transition duration-300 ease-in-out hover:bg-gray-100 dark:border-neutral-500 dark:hover:bg-blue-100">
                                        <td className="whitespace-nowrap px-6 py-4 font-medium">{employee.id}</td>
                                        <td className="whitespace-nowrap px-6 py-4">{employee.name}</td>
                                        <td className="whitespace-nowrap px-6 py-4">{employee.performance}</td>
                                        <td className="whitespace-nowrap px-6 py-4">{formattedDate}</td>
                                    </tr>
                                )
                            }
                            ))}
                    </tbody>
                </table>
            </div>
        </div>
    );
}

export default Dashboard;
