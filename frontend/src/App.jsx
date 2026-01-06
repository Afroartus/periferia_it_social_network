import {useContext, useState} from "react";
import {AuthContext} from "./context/AuthContext";
import Login from "./pages/Login";
import Posts  from "./pages/Posts";
import Register from "./pages/Register";
import './App.css'

function App() {
     const { token } = useContext(AuthContext);
     const [showRegister, setShowRegister] = useState(false);
     if (token) {
          return <Posts />
     }

     return showRegister ? (
         <Register goToLogin={() => setShowRegister(false)} />
     ) : (
         <Login goToRegister={() => setShowRegister(true)} />
     )

}

export default App
