import React, {useContext, useState} from "react";
import { useNavigate, Link } from "react-router-dom";
import { Buffer } from "buffer";
import {AuthContext} from "../context/AuthContext";

function jwtToUserId(token) {
    return JSON.parse(Buffer.from(token.split(".")[1], "base64").toString());
}

const Login = () => {
    const navigate = useNavigate()
    const [err, setErr] = useState(false)

    const handleSubmit = (e) => {
        e.preventDefault()
        const email = e.target[0].value;
        const password = e.target[1].value;

        async function TryLogin() {
            try {
                const response = await
                  fetch('http:/' + process.env.PUBLIC_URL + ':8000/user/login', {
                      method: 'POST',
                      body: JSON.stringify({
                          "email": email,
                          "password": password,
                      })
                  })
                const jwtToken = response.headers.get("Authorization").split(' ')[1]
                console.log(jwtToken)
                localStorage.setItem("jwtToken", jwtToken)
                localStorage.setItem("userId", jwtToUserId(jwtToken)['sub'])
                const userData = await response.json()
                localStorage.setItem("userName", userData.UserName)
                navigate("/")
                window.location.reload()
            } catch(err) {
                setErr(true)
                console.log(err)
            }
        }
        TryLogin()
    };

    return (
        <div className="formContainer">
            <div className="formWrapper">
                <span className="logo">Chat</span>
                <span className="title">Login</span>
                <form onSubmit={handleSubmit}>
                    <input type="email" placeholder="email" />
                    <input type="password" placeholder="password" />
                    <button>Login</button>
                </form>
                <p>Don't you have an account yet? <Link to={"/register"}>Register</Link></p>
                {err && <span>Something went wrong</span>}
            </div>
        </div>
    );
}

export default Login