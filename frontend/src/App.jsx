import { useContext} from "react";
import {AuthContext} from "./context/AuthContext";
import Login from "./pages/Login";
import './App.css'

function App() {

     const { token } = useContext(AuthContext);
     return token ? <div>Autenticado</div> : <Login />;

}

export default App
