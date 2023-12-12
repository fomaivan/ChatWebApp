import React, {useContext, useEffect, useState} from "react";
import { useNavigate } from "react-router-dom";
import {AuthContext} from "../context/AuthContext";

const Navbar = () => {
  const navigate = useNavigate()
  const { currentUser } = useContext(AuthContext)

  function signOut() {
    localStorage.clear()
    navigate('/login')
  }

  return (
    <div className="navbar">
      <span className="logo">Chat</span>
      <div className="user">
        <img
          src="https://images.unsplash.com/photo-1535713875002-d1d0cf377fde?q=80&w=2960&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"
          alt=""/>
        <span> {currentUser.UserName} </span>
        <button onClick={()=>signOut()}>Exit</button>
      </div>
    </div>
  )
}

export default Navbar