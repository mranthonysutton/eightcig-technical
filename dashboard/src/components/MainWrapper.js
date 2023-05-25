const MainWrapper = ({ children }) => {
    return (

        <main className="min-h-screen w-full bg-gray-100 text-gray-700">
            <header className="flex w-full items-center justify-between border-b-2 border-gray-200 bg-white p-2">
                <div className="flex items-center space-x-2">
                    <button type="button" className="text-3xl"><i className="bx bx-menu"></i></button>
                    <div>Employee Management</div>
                </div>
            </header >

            <div className="flex">
                {children}
            </div>
        </main>
    )
}

export default MainWrapper
