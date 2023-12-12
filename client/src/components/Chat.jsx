import React, {useContext, useEffect, useState} from "react";
import Messages from "./Messages";
import Input from "./Input";
import ChatInfo from "./ChatInfo";
import { InboxContext } from "../context/InboxContext"
import {AuthContext} from "../context/AuthContext";

const Chat = () => {
  const { inboxTmp } = useContext(InboxContext)
  const { currentUser } = useContext(AuthContext)
  const [currentContact, setCurrentContact] = useState({})


  useEffect(() => {
    async function getContact() {


      const response =
        await fetch("http:/" + process.env.PUBLIC_URL + ":8000/find/user/" +
          inboxTmp.inboxCurr?.UserName)
      const contactTmp = await response.json()
      setCurrentContact(contactTmp)
      console.log("contactTmp.UserId = ", contactTmp)
    }
    getContact()
  }, [inboxTmp.inboxCurr?.UserName]);
  /// прееделать contactId
  return (
    <div className="chat">
      <ChatInfo contactUsername={inboxTmp.inboxCurr?.UserName}/>
      <Messages />
      <Input contactId={currentContact.UserId}/>
    </div>
  )
}

export default Chat;