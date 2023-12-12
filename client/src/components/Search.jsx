import React, {useContext, useState} from "react";
import {AuthContext} from "../context/AuthContext";
import {InboxContext} from "../context/InboxContext";

const Search = () => {
    const [userName, setUserName] = useState('')
    const [user, setUser] = useState(null)
    const [err, setErr] = useState(false)
    const {dispatch} = useContext(InboxContext)
    const {currentUser} = useContext(AuthContext)

    const handleSearch = async () => {
        try {
            const url = 'http:/' + process.env.PUBLIC_URL + ':8000/find/'
            const queryUsers = await fetch(url + userName)
            const queryUserData = await queryUsers.json()
            console.log("queryUserData: ", queryUserData)
            queryUserData.forEach((tmpUser) => {
                setUser(tmpUser)
            })
        } catch (err) {
            setErr(true)
        }
    }

    const handleKey = (e) => {
        e.code === "Enter" && handleSearch()
    }

    const handleSelectInbox = (inb) => {
        dispatch({ type: "CHANGE_INBOX", payload: inb });
    }

    const handleSelect = async () => {
        const requestURL = 'http:/' + process.env.PUBLIC_URL + ":8000/find/inbox/" +
          currentUser.UserId + '/' + '' + user.UserId
        try {
            const response = await fetch(requestURL)
            if (response.status == 404) {
                console.log("START: response.status == 404")
                await fetch('http:/' + process.env.PUBLIC_URL + ':8000/chat/create/inbox/' +
                  user.UserId + "/" + "" + currentUser.UserId, {
                    method: 'GET'
                })
                const response = await fetch(requestURL)
                const inbox = await response.json()
                console.log("MY_INBOX:", inbox)
                handleSelectInbox({
                    InboxId: inbox.InboxId,
                    LastMessageContent: inbox.LastMessageContent,
                    LastMessageDttm: inbox.LastMessageDttm,
                    LastMessageUser: inbox.LastMessageUser,
                    UserName: user.UserName
                })

                // await fetch('http://localhost:8000/chat/send_message', {
                //     method: 'POST',
                //     body: JSON.stringify({
                //         "From": currentUser.UserId,
                //         "To": user.UserId,
                //         "Content": "Hello"
                //     })
                // })
                // const response = await fetch(requestURL)
                // const inbox = await response.json()
                // console.log("END IF response.status == 404", inbox)
                // handleSelectInbox({
                //     InboxId: inbox.InboxId,
                //     LastMessageContent: inbox.LastMessageContent,
                //     LastMessageDttm: inbox.LastMessageDttm,
                //     LastMessageUser: inbox.LastMessageUser,
                //     UserName: user.UserName
                // })

            } else {
                const inbox = await response.json()
                handleSelectInbox(inbox)
            }
        } catch (err) {}
        setUser(null)
        setUserName("")
    }

    return (
      <div className="search">
          <div className="searchForm">
              <input
                type="text" placeholder="Find a user"
                onKeyDown={handleKey}
                onChange={e => setUserName(e.target.value)}
                value={userName}
              />
          </div>
          {/*{err && <span>User not found</span>}*/}
          {user && <div className="userChat" onClick={handleSelect}>
              <img
                src="https://images.unsplash.com/photo-1535713875002-d1d0cf377fde?q=80&w=2960&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"
                alt=""/>
              <div className="userChatInfo">
                  <span>{user.UserName}</span>
              </div>
          </div>}
      </div>
    )
}

export default Search