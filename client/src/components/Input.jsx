import React, {useContext, useState} from "react";
import {AuthContext} from "../context/AuthContext";
import {InboxContext} from "../context/InboxContext";

const Input = ({ contactId }) => {
  const { inboxTmp } = useContext(InboxContext)
  const [messageText, setMessageText] = useState('')
  const { currentUser } = useContext(AuthContext)
  const { dispatch } = useContext(InboxContext)
  async function handleSend() {
    let url = 'http:/' + process.env.PUBLIC_URL + ':8000/chat/send_message'
    await fetch(url, {
      method: 'POST',
      body: JSON.stringify({
        "From": +currentUser.UserId,
        "To": contactId,
        "Content": messageText
      })
    })

    url = 'http:/' + process.env.PUBLIC_URL + ':8000/find/inbox/user/'
    const response = await
      fetch(url +
        inboxTmp.inboxCurr?.InboxId + "/" + inboxTmp.inboxCurr?.UserName)
    const UpdateInbox = await response.json()
    console.log("UpdateInbox: ", UpdateInbox)

    const handleSelect = (inb) => {
      dispatch({ type:"CHANGE_INBOX", payload: inb});
    }
    handleSelect(UpdateInbox)
    setMessageText("")
  }
  // console.log("inboxTMP: ", inboxTmp)

  return (
    <div className="input">
      <input
        type="text"
        placeholder="Your message..."
        value={messageText}
        onChange={(e) => {setMessageText(e.target.value)}}
      />
      <div className="send">
        {/*<img src={Attach} alt=""/>*/}
        {/*<input type="file" style={{display: "none"}} id="file"/>*/}
        {/*<label htmlFor="file">*/}
        {/*  <img src={Img} alt=""/>*/}
        {/*</label>*/}
        <button onClick={handleSend}>Send</button>
      </div>
    </div>
  )
}

export default Input