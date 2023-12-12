import React, {useState} from "react";

const Message = (props) => {
  const [text, setText] = useState(props.text);
  const [isOwner, setIsOwner] = useState(props.isOwner);

  if (isOwner) {
    return (
      <div className="message owner">
        <div className="messageContent">
          <p>{text}</p>
        </div>

      </div>
    )
  }
  return (
    <div className="message">
      <div className="messageContent">
        <p>{text}</p>
      </div>
    </div>
  )
}

export default Message