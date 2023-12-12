package main

import (
	"chat/internal/entity"
	handler "chat/internal/handler/http"
	"chat/internal/pkg/auth"
	"chat/internal/repository"
	repo_sqlite "chat/internal/repository/sqlite"
	"chat/internal/service"
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func checkService(service *service.Service) {
	fmt.Println("####### Checking service")

	fmt.Println("Try to register a user")

	userReg1 := entity.UserRegister{
		UserLogin: entity.UserLogin{
			Email:    "fomin.id@phystech.edu",
			Password: "asdfgfadgs",
		},
		UserName: "foma_ivan",
	}
	service.User.Register(&userReg1, &service.Chat)
	fmt.Println("Register ", userReg1)

	userReg2 := entity.UserRegister{
		UserLogin: entity.UserLogin{
			Email:    "petrov.sa@phytech.edu",
			Password: "hgdfsada",
		},
		UserName: "petrov_stepan",
	}
	service.User.Register(&userReg2, &service.Chat)
	fmt.Println("Register ", userReg2)

	userReg3 := entity.UserRegister{
		UserLogin: entity.UserLogin{
			Email:    "ivanov.da@phytech.edu",
			Password: "kshjkrgfdkj",
		},
		UserName: "ivanov_danila",
	}
	service.User.Register(&userReg3, &service.Chat)
	fmt.Println("Register ", userReg3)

	userReg4 := entity.UserRegister{
		UserLogin: entity.UserLogin{
			Email:    "ollodfs@mail.ru",
			Password: "65jk34563k",
		},
		UserName: "olo_olo",
	}
	service.User.Register(&userReg4, &service.Chat)
	fmt.Println("Register ", userReg4)
	fmt.Println("Ok")
	fmt.Println("########################################")

	fmt.Println("Try to login with CORRECT password")
	userLogin := entity.UserLogin{
		Email:    "ollodfs@mail.ru",
		Password: "65jk34563k",
	}
	_, err := service.User.Login(&userLogin)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Ok")
	}
	fmt.Println()

	fmt.Println("Try to login with INCORRECT password")
	userLogin = entity.UserLogin{
		Email:    "ollodfs@mail.ru",
		Password: "hgfd74hdgf",
	}
	_, err = service.User.Login(&userLogin)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Ok")
	}
	fmt.Println()

	fmt.Println("####### Try to send message in exists inbox")
	var SendMessage *entity.Message
	if SendMessage, err = service.Chat.SendMessageFromTo(1, 2, "Test Message from 1 to 2"); err != nil {
		println("Error")
	} else {
		fmt.Printf("Message: %+v\n", *SendMessage)
		println("Ok")
	}
	fmt.Println()

	fmt.Println("####### Try to send message in not exists inbox")
	if SendMessage, err = service.Chat.SendMessageFromTo(3, 1, "Test Message from 3 to 1"); err != nil {
		fmt.Println("Error")
	} else {
		fmt.Printf("Message: %+v\n", *SendMessage)
		println("Ok")
	}
	fmt.Println()

	fmt.Println("####### GetAllInboxesByUser")
	Inboxes := service.Chat.GetAllInboxesByUserId(1)
	fmt.Println("Inboxes for user1: ", *Inboxes)
	fmt.Println()
	Inboxes = service.Chat.GetAllInboxesByUserId(2)
	fmt.Println("Inboxes for user2: ", *Inboxes)
	fmt.Println()
	Inboxes = service.Chat.GetAllInboxesByUserId(3)
	fmt.Println("Inboxes for user3: ", *Inboxes)
	fmt.Println()
	fmt.Println("####### Ended checking Service")
	fmt.Println()
}

func main() {
	fmt.Println("Start program")

	dbURI, ok := os.LookupEnv("DB_URI")
	if !ok {
		log.Println("cannot get db_uri from ENV")
		dbURI = "test.db"
	}
	db, err := repo_sqlite.NewSQliteDB(dbURI)
	if err != nil {
		log.Panic("Failed to initialize database %s", err.Error())
	} else {
		log.Println("Database is initialize")
	}

	repo := repository.NewRepository(db)
	service_ := service.NewService(repo)

	//signingKey := []byte("TODO-read-from-config")
	signingKey, ok := os.LookupEnv("AUTH_SIGNING_KEY")
	if !ok {
		log.Println("cannot get AUTH_SIGNING_KEY from ENV")
		signingKey = "dfasvhsddddgfds"
	}
	authManager := auth.NewAuthManager([]byte(signingKey))
	checkService(service_)

	h := handler.NewHandler(service_, authManager)

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	srv := &http.Server{
		Addr: ":8000",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      h.NewRouter(), // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
