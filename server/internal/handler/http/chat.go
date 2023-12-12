package handler

import (
	"chat/internal/entity"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) GetInboxesByUserId(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	authClaims, err := h.GetСlaimsFromAuthHeader(r)
	authUserId, _ := strconv.ParseUint((*authClaims)["sub"], 10, 32)

	if err != nil {
		log.Printf("Error when trying to get Inboxes of user, Error: %v\n", err.Error())
		w.WriteHeader(http.StatusTeapot)
		return
	}

	inboxesWithUserName := h.service.Chat.GetAllInboxesWithUserNameByUserId(uint(authUserId))
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(inboxesWithUserName)
	if err != nil {
		log.Printf("Error when encoding inboxes, Error: %v\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *Handler) GetInboxWithUserName(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	vars := mux.Vars(r)
	username := vars["username"]
	inboxId, err := strconv.ParseUint(vars["inbox_id"], 10, 32)
	if err != nil {
		log.Printf("Error when trying to decode inbox_id, Error: %v\n", err.Error())
		w.WriteHeader(http.StatusTeapot)
		return
	}
	tmpInbox, _ := h.service.Chat.GetInboxByInboxId(uint(inboxId))
	result := &entity.InboxWithUserName{
		InboxId:            tmpInbox.InboxId,
		LastMessageContent: tmpInbox.LastMessageContent,
		LastMessageDttm:    tmpInbox.LastMessageDttm,
		LastMessageUser:    tmpInbox.LastMessageUser,
		UserName:           username,
	}
	err = json.NewEncoder(w).Encode(result)
}

// curl -i -X POST -d '{"from": 2, "to": 3, "content": "Test message from 2 to 3 in post req"}' http://localhost:8000/chat/send_message
func (h *Handler) SendMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	var sendMsg entity.SendMessage
	err := json.NewDecoder(r.Body).Decode(&sendMsg)
	if err != nil {
		log.Printf("Error when trying to decode sendMsg, Error: %v\n", err.Error())
		w.WriteHeader(http.StatusTeapot)
		return
	}
	_, err = h.service.Chat.SendMessageFromTo(sendMsg.From, sendMsg.To, sendMsg.Content)
	//h.GetInboxesByUserId(w, r)

	if err != nil {
		log.Printf("Error when trying to use func service.Chat.SendMessageFromTo, Error: %v\n", err.Error())
		w.WriteHeader(http.StatusTeapot)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) GetMessagesByInbox(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	vars := mux.Vars(r)
	inboxId, err := strconv.ParseUint(vars["inbox_id"], 10, 32)
	if err != nil {
		log.Printf("Error when trying to decode inbox_id, Error: %v\n", err.Error())
		w.WriteHeader(http.StatusTeapot)
		return
	}
	messages := h.service.Chat.GetMessagesByInboxId(uint(inboxId))
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(messages)
	if err != nil {
		log.Printf("Error when encoding messages, Error: %v\n", err.Error())
		w.WriteHeader(http.StatusTeapot)
		return
	}
}

func (h *Handler) CreateInbox(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

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

	_, err = h.service.Chat.CreateInboxAndInboxUser(uint(FirstId), uint(SecondId))
	if err != nil {
		log.Printf("Error when trying to CreateInboxAndInboxUser, Error: %v\n", err.Error())
		w.WriteHeader(http.StatusTeapot)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
