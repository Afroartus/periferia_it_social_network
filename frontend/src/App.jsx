import { useContext} from "react";
import {AuthContext} from "./context/AuthContext";
import Login from "./pages/Login";
import Posts  from "./pages/Posts";
import './App.css'

function App() {

     const { token } = useContext(AuthContext);
     return token ? <Posts /> : <Login />;

}

export default App
