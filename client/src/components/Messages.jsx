import React, {useContext, useEffect, useState} from "react";
import Message from "./Message";
import { InboxContext } from "../context/InboxContext";
import { AuthContext } from "../context/AuthContext";

const Messages = () => {
  const [messages, setMessages] = useState([])
  const {inboxTmp} = useContext(InboxContext)
  const {currentUser} = useContext(AuthContext)
  const [isEmptyMessages, setIsEmptyMessages] = useState(true)

  useEffect( () => {
    async function getMessages() {
      if (inboxTmp.inboxCurr?.InboxId === undefined) {
        setMessages([])
      } else {
        const url = 'http:/' + process.env.PUBLIC_URL + ':8000/chat/message/'
        const response = await fetch(url + inboxTmp.inboxCurr?.InboxId)
        const currMessages = await response.json()
        console.log(currMessages)
        setMessages(currMessages)
        if (currMessages.length == 0) {
          setIsEmptyMessages(true)
        } else {
          setIsEmptyMessages(false)
        }
      }
    }
    getMessages()
    return () => {
      setMessages([])
    }
    }, [inboxTmp.inboxCurr?.InboxId, inboxTmp.inboxCurr?.LastMessageDttm]
  )

  console.log("CurrentInbox: ", inboxTmp)
  console.log("CurrentInboxLength", Object.keys(inboxTmp).length)

  return (
    <div className="messages" id="777">
      {isEmptyMessages && Object.keys(inboxTmp.inboxCurr).length > 0 && <span>  Начните диалог с этим пользователем  </span>}
      {isEmptyMessages && Object.keys(inboxTmp.inboxCurr).length === 0 && <span>  Воспользуйтесь поиском и найдите друга  </span>}

      {messages.map((item, index) => (
          <Message
            text={item.Content}
            isOwner={item.UserId == currentUser.UserId}
            key = {index}
          />
        ))
      }
    </div>
  )
}

export default Messages