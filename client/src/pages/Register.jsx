import React, {useState} from "react";
import { useNavigate, Link} from "react-router-dom";

const Register = () => {
    const navigate = useNavigate()
    const [err, setErr] = useState(false)

    const handleSubmit = (e) => {
        e.preventDefault()
        const userName = e.target[0].value;
        const email = e.target[1].value;
        const password = e.target[2].value;
        const photo = e.target[3].files[0];

        console.log(photo)

        async function SendDataToRegisterNewUser() {
            try {
                const response =
                  await fetch('http:/' + process.env.PUBLIC_URL + ':8000/user/register', {
                      method: 'POST',
                      body: JSON.stringify({
                          "email": email,
                          "password": password,
                          "userName": userName,
                          "userImage": photo
                      })
                  })
                if (response.status === 201) {
                    navigate("/login")
                } else {
                    setErr(true)
                }
                // navigate("/login")
            } catch (err) {
                setErr(true)
                console.log(err)
            }
        }
        SendDataToRegisterNewUser()
    };

    return (
        <div className="formContainer">
            <div className="formWrapper">
                <span className="logo">Chat</span>
                <span className="title">Register</span>
                <form onSubmit={handleSubmit}>
                    <input type="text" placeholder="Your name" />
                    <input type="email" placeholder="email" />
                    <input type="password" placeholder="password" />
                    <input type="file"/>
                    <button>Sign Up</button>
                </form>
                <p>Do you already have an account? <Link to={"/login"}>Login</Link></p>
                {err && <span>Something went wrong</span>}
            </div>
        </div>
    );
}

export default Register