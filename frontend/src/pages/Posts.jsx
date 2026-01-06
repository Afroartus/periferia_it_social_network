import {useContext, useEffect, useState} from "react";
import api from "../api";
import {AuthContext} from "../context/AuthContext";
import Profile from "./Profile";



export default function Posts() {
    const  { logout } = useContext(AuthContext);
  //  const [showProfile, setShowProfile] = useState(true);
    const [posts, setPosts] = useState([]);
    const [message, setMessage] = useState("");

    const loadPosts = async () => {
        const response = await api.get("/posts");
        console.log("POSTS RESPONSE:", response.data);
        setPosts(response.data.posts);
    };

    const createPost = async () => {
        if (!message) return;
        await api.post("/posts", { message });
        setMessage("");
        loadPosts();
    };

    const likePost = async (id) => {
        await api.post(`/posts/${id}/like`);
        loadPosts();
    };



    useEffect(() => {
        loadPosts();
    }, []);


    return (
        <div>
            <Profile />

            <button onClick={logout}> logout </button>

            <h2>Posts</h2>

            <input
            placeholder="New post"
            value={message}
            onChange={(e) => setMessage(e.target.value)}
            />
            <button onClick={createPost}>publish</button>

            <hr/>

            {posts.map((p) => (
                <div key={p.id}>
                    <p>{p.message}</p>
                    <button onClick={() => likePost(p.id)}>
                        ❤️ {p.likes}
                    </button>
                </div>
            ))}

        </div>



    );
}
