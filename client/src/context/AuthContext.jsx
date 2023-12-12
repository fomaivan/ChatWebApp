import {createContext, useEffect, useState} from "react"

export const AuthContext = createContext()

export const AuthContextProvider = ({ children }) => {
  const [currentUser, setCurrentUser] = useState({})

  useEffect(() => {
    if (!localStorage.getItem("jwtToken")) {
      return
    }
    async function FetchFunc() {
      await fetch("http:/" + process.env.PUBLIC_URL + ":8000/user", {
        method: "GET",
        headers: {
          "Authorization":
            "Bearer " + localStorage.getItem('jwtToken').toString(),
        }
      }).then((response) => response.json())
        .then((userData) => {
          setCurrentUser(userData)
        })
    }
    FetchFunc()
  }, [localStorage.getItem('jwtToken')?.toString() || ""]);

  return (
    <AuthContext.Provider value={{ currentUser }}>
      {children}
    </AuthContext.Provider>
  )
}