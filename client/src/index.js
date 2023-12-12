import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App';
import {AuthContext, AuthContextProvider} from "./context/AuthContext";
import {InboxContextProvider} from "./context/InboxContext";

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <AuthContextProvider>
    <InboxContextProvider>
      <React.StrictMode>
        <App/>
      </React.StrictMode>
    </InboxContextProvider>
  </AuthContextProvider>
);
