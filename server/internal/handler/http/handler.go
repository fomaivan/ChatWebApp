package handler

import (
	"chat/internal/entity"
	"chat/internal/pkg/auth"
	"chat/internal/service"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Handler struct {
	service *service.Service
	auth    entity.AuthManager
}

func NewHandler(s *service.Service, a *auth.AuthManager) *Handler {
	return &Handler{service: s, auth: a}
}

func (h *Handler) NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", index).Methods(http.MethodGet, http.MethodOptions)

	// return all user
	r.HandleFunc("/all_user", h.GetAllUsers).Methods(http.MethodGet, http.MethodOptions)

	// return user
	r.HandleFunc("/user", h.GetUserById).Methods(http.MethodGet, http.MethodOptions)

	// register user
	r.HandleFunc("/user/register", h.RegisterUser).Methods(http.MethodPost, http.MethodOptions)
	// login user
	r.HandleFunc("/user/login", h.LoginUser).Methods(http.MethodPost, http.MethodOptions)

	// get user by username%  (match by first letters)
	r.HandleFunc("/find/{username}", h.GetUsersByUserName).Methods(http.MethodGet, http.MethodOptions)
	// get user by username (exact match)
	r.HandleFunc("/find/user/{username}", h.GetUserByUserName).Methods(http.MethodGet, http.MethodOptions)
	// get inboxId by 2 users id. If not found -> return 0
	r.HandleFunc("/find/inbox/{first}/{second}", h.GetInboxByTwoUsers).Methods(http.MethodGet, http.MethodOptions)
	// get inbox with username
	r.HandleFunc("/find/inbox/user/{inbox_id}/{username}", h.GetInboxWithUserName).Methods(http.MethodGet, http.MethodOptions)

	// get all inboxes by userId
	r.HandleFunc("/chat/inbox", h.GetInboxesByUserId).Methods(http.MethodGet, http.MethodOptions)
	// create inbox and inboxUser
	r.HandleFunc("/chat/create/inbox/{first}/{second}", h.CreateInbox).Methods(http.MethodGet, http.MethodOptions)
	// send message body={from, to, content}
	r.HandleFunc("/chat/send_message", h.SendMessage).Methods(http.MethodPost, http.MethodOptions)
	// get all messages by inboxId
	r.HandleFunc("/chat/message/{inbox_id}", h.GetMessagesByInbox).Methods(http.MethodGet, http.MethodOptions)

	// create inbox between two users
	// r.HandleFunc("/chat/inbox/create", h.CreateInboxByTwoUsers).Methods("POST")

	r.Use(LoggingMiddleware(r))
	r.Use(CustomCORSMethodMiddleware(r))
	return r
}

func CustomCORSMethodMiddleware(r *mux.Router) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Content-Type", "application/json, application/x-www-form-urlencoded, multipart/form-data, text/plain")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Length, Content-Type, Authorization, authorization, Host, Origin, X-CSRF-Token")
			w.Header().Set("Access-Control-Expose-Headers", "Authorization")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
			next.ServeHTTP(w, req)
		})
	}
}

func LoggingMiddleware(r *mux.Router) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			log.Printf("Origin: %s | Forwarded: %s | Method: %s | RequestURI: %s",
				req.Header.Get("Origin"), req.Header.Get("Forwarded"), req.Method, req.RequestURI)
			next.ServeHTTP(w, req)
		})
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World\n")
}
