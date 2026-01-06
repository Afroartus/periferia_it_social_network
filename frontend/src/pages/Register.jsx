import {useContext, useState} from "react";
import api from "../api.js";


export default function Register({goToLogin}) {
    const [name, setName] = useState('');
    const [lastName, setLastName] = useState('');
    const [email, setEmail] = useState('');
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [birth_date, setBirth_date] = useState('');


    const submit = async (e) => {
        e.preventDefault();
        const response = await api.post('/user/create', {
            name,
            lastName,
            email,
            username,
            password,
            birth_date,
        });

        goToLogin();
    }

    return (
        <form onSubmit={submit}>
            <input
                placeholder="Name"
                value={name}
                onChange={e => setName(e.target.value)}/>
            <input
                placeholder="Last Name"
                value={lastName}
                onChange={e => setLastName(e.target.value)}/>
            <input
                placeholder="email"
                value={email}
                onChange={e => setEmail(e.target.value)}/>
            <input
                placeholder="username"
                value={username}
                onChange={e => setUsername(e.target.value)}/>
            <input
                placeholder="password"
                value={password}
                onChange={e => setPassword(e.target.value)}/>
            <input
                type="date"
                placeholder="birth_date"
                value={birth_date}
                onChange={e => setBirth_date(e.target.value)}/>
            <button onClick={goToLogin} type="submit">Login</button>
            <button type="submit">Register</button>

        </form>
    )


}