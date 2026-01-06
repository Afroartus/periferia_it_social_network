import {useContext, useState} from "react";
import {AuthContext} from "../context/AuthContext";
import api from "../api";

export default function Login({goToRegister}){
    const { login } = useContext(AuthContext);
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");


    const submit = async (e) => {
        e.preventDefault();
        const res = await api.post("/login", { username, password});
        login(res.data.token);
    };

    return (
        <form onSubmit={submit}>
            <input
                placeholder="username"
                value={username}
                onChange={e => setUsername(e.target.value)} />
            <input
                type="password"
                placeholder="password"
                value={password}
                onChange={e => setPassword(e.target.value)} />
            <button type="submit">Login</button>
            <button type="button" onClick={goToRegister}>Register</button>
        </form>
    );
}