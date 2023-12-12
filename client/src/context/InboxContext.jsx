import {createContext, useReducer, useState} from "react"

export const InboxContext = createContext({})
export const InboxContextProvider = ({ children }) => {
  const INITIAL_STATE = {
    inboxCurr: {},
  }
  const InboxReducer = (state, action) => {
    const {type} = action
    switch (type) {
      case "CHANGE_INBOX":
        return {
          inboxCurr: action.payload,
        }
      default:
        return state;
    }
  }
  const [state, dispatch] = useReducer(InboxReducer, INITIAL_STATE);
  return (
    <InboxContext.Provider value={{inboxTmp: state, dispatch}}>
      {children}
    </InboxContext.Provider>
  )
}