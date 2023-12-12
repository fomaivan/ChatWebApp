import React, {useEffect, useState, useContext} from "react";
// import UserChat from "./UserChat";
import {AuthContext} from "../context/AuthContext";
import {InboxContext} from "../context/InboxContext";

const Chats = () => {
    const [inboxes, setInboxes] = useState([]);
    const {currentUser} = useContext(AuthContext)
    const {dispatch} = useContext(InboxContext)
    const {inboxTmp} = useContext(InboxContext)

    useEffect(() => {
        (async () => {
            try {
                console.log("process.env.PUBLIC_URL = ", process.env.PUBLIC_URL)
                console.log("process.env.PUBLIC_URL type: ", typeof process.env.PUBLIC_URL)

                const tmpUrl = "http:/" + process.env.PUBLIC_URL + ":8000/chat/inbox"
                const responseInboxes =
                  await fetch(tmpUrl, {
                    method: "GET",
                    headers: {
                        "Access-Control-Request-Method": "POST",
                        "Authorization":
                          "Bearer " + localStorage.getItem('jwtToken').toString(),
                    }
                })
                const data = await responseInboxes.json()
                setInboxes(data)
            } catch (error) {
                console.log("Error response InboxesByUser")
            }
        })()
    }, [currentUser.UserId, inboxTmp.inboxCurr?.LastMessageDttm, inboxes.length])

    console.log("INBOXES: ", inboxes)

    const handleSelect = (inb) => {
        dispatch({type: "CHANGE_INBOX", payload: inb});
    }

    return (
      <div className="chats">
          {inboxes.map((item, index) => (
            <div className="userChat"
                 key={index}
                 onClick={() => handleSelect(item)}>
                <img
                  src="https://images.unsplash.com/photo-1535713875002-d1d0cf377fde?q=80&w=2960&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"
                  alt=""/>
                <div className="userChatInfo">
                    <span>{item.UserName}</span>
                    <p>{item.LastMessageContent}</p>
                </div>
            </div>
          ))
          }
      </div>
    )
}

export default Chats;