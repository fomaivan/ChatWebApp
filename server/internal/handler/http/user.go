package handler

import (
	"chat/internal/entity"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	users, err := h.service.User.GetAll()
	if err != nil {
		log.Printf("Error when getting all users, Error: %v\n", err.Error)
		w.WriteHeader(http.StatusTeapot)
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Printf("Error when encoding users, Error: %v\n", err.Error)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *Handler) GetUserById(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	authClaims, err := h.GetСlaimsFromAuthHeader(r)
	//fmt.Println(*authClaims)
	authUserId, _ := strconv.ParseUint((*authClaims)["sub"], 10, 32)
	//fmt.Println(authUserId)
	if err != nil {
		log.Printf("Error when parsing user_id to uint, Error: %v", err.Error())
		return
	}

	user, err := h.service.User.Get(uint(authUserId))
	if err != nil {
		log.Printf("Error when hetting user_id with id=%v, Error: %v", authUserId, err.Error())
		w.WriteHeader(http.StatusTeapot)
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Printf("Error when encoding users, Error: %v\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *Handler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
	fmt.Println(r.Body)
	userReg := entity.UserRegister{}
	err := json.NewDecoder(r.Body).Decode(&userReg)

	////////////////// file
	//r.ParseMultipartForm(10 << 20)
	//file, handler, err := r.FormFile("myFile")
	//if err != nil {
	//	log.Printf("Error when trying to decode file, Error:%v\n", err.Error())
	//	w.WriteHeader(http.StatusTeapot)
	//	return
	//}
	//
	//defer file.Close()
	//fmt.Printf("UploadingFile: %+v\n", handler.Filename)
	//fmt.Printf("SizeFile: %+v\n", handler.Size)
	//fmt.Printf("HeaderFile: %+v\n", handler.Header)
	//
	//tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	//if err != nil {
	//	log.Printf("Error when trying to uplod file to server, Error:%v\n", err.Error())
	//	w.WriteHeader(http.StatusTeapot)
	//	return
	//}
	//defer tempFile.Close()
	//
	//fileBytes, err := ioutil.ReadAll(file)
	//if err != nil {
	//	log.Printf("Error when trying to ReadALl(file), Error:%v\n", err.Error())
	//	w.WriteHeader(http.StatusTeapot)
	//	return
	//}
	//tempFile.Write(fileBytes)
	////////

	if err != nil {
		log.Printf("Error when trying to decode userReg register a new user, Error: %v\n", err.Error())
		w.WriteHeader(http.StatusTeapot)
		return
	}
	if userReg.UserName == "" || userReg.Password == "" || userReg.Email == "" {
		w.WriteHeader(http.StatusTeapot)
		return
	}
	err = h.service.User.Register(&userReg, &h.service.Chat)
	if err != nil {
		log.Printf("Error when trying to register a new user, Error: %v\n", err.Error())
		w.WriteHeader(http.StatusTeapot)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
	userLogin := entity.UserLogin{}
	err := json.NewDecoder(r.Body).Decode(&userLogin)
	if err != nil {
		log.Printf("Error when trying to decode userLogin register an user, Error: %v\n", err.Error())
		w.WriteHeader(http.StatusTeapot)
		return
	}

	UserId, err := h.service.User.Login(&userLogin)
	if err != nil {
		log.Printf("Error when trying to login an user, Error: %v\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tokenString, _ := h.auth.MakeAuth(UserId)
	w.Header().Add("Authorization", "Bearer "+tokenString)
	User, err := h.service.User.Get(UserId)
	err = json.NewEncoder(w).Encode(User)
	if err != nil {
		log.Printf("Error when trying to Encode an user, Error: %v\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
