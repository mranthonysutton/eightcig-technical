import { Axios } from "axios";

const useAxios = () => {
    return Axios.create({
        baseURL: "http://localhost:4000/"
    })
}


export default useAxios;
