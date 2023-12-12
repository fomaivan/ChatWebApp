package handler

import (
	"chat/internal/entity"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) GetUsersByUserName(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	vars := mux.Vars(r)
	name := vars["username"]
	users, err := h.service.User.GetUsersByUserName(name)
	if err != nil {
		log.Printf("Error when GetUsersByUserName, Error: %v", err.Error())
		w.WriteHeader(http.StatusTeapot)
		return
	}
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Printf("Error when encoding users, Error: %v\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *Handler) GetUserByUserName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["username"]
	user, err := h.service.User.GetUserByUserName(name)
	if err != nil {
		log.Printf("Error when GetUsersByUserName, Error: %v", err.Error())
		w.WriteHeader(http.StatusTeapot)
		return
	}
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Printf("Error when encoding users, Error: %v\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *Handler) GetInboxByTwoUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	FirstId, err := strconv.ParseUint(vars["first"], 10, 32)
	if err != nil {
		log.Printf("Error when trying to encode UsersIdx, Error: %v\n", err.Error())
		w.WriteHeader(http.StatusTeapot)
		return
	}

	SecondId, err := strconv.ParseUint(vars["second"], 10, 32)
	if err != nil {
		log.Printf("Error when trying to encode UsersIdx, Error: %v\n", err.Error())
		w.WriteHeader(http.StatusTeapot)
		return
	}

	Users := entity.TwoUserIdx{
		FirstUser:  uint(FirstId),
		SecondUser: uint(SecondId),
	}
	if err != nil {
		log.Printf("Error when trying to decode UsersIdx, Error: %v\n", err.Error())
		w.WriteHeader(http.StatusTeapot)
		return
	}
	InboxId, err := h.service.Chat.GetInboxIdByTwoUsers(&Users)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	Inbox, err := h.service.Chat.GetInboxByInboxId(InboxId)
	fmt.Println(Inbox)
	err = json.NewEncoder(w).Encode(&Inbox)
}

//func (h *Handler) GetContactByUserIdAndInboxId(w http.ResponseWriter, r *http.Request) {
//
//}
