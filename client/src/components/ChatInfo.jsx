import Cam from "../images/cam.png";
import Add from "../images/add.png";
import More from "../images/more.png";
import React from "react";

const ChatInfo = ({contactUsername}) => {
  return (
    <div className="chatInfo">
      <span>{contactUsername}</span>
      <div className="chatIcons">
        <img src={Cam} alt=""/>
        <img src={Add} alt=""/>
        <img src={More} alt=""/>
      </div>
    </div>
  )
}

export default ChatInfo