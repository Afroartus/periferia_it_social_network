import {useEffect, useState} from "react";
import api from "../api.js";


export default function Profile(){
    const [user,setUser]=useState(null);


    useEffect(()=>{
        const loadProfile = async () => {
            const resp = await api.get("/me");
            setUser(resp.data);
        };
        loadProfile();
    }, []);

    if (!user) return <div>Loading...</div>;

    return(
        <div>
            <h2>Profile</h2>
            <p><b>Name:</b> {user.name}</p>
            <p><b>Last name:</b> {user.lastname}</p>
            <p><b>Email:</b> {user.email}</p>
            <p><b>Username:</b> {user.username}</p>
        </div>
    )


}