package api

import (
	"net/http"

	"github.com/YuuHikida/GSC_backend_go/application/service"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func SetRoutes() http.Handler {
	router := mux.NewRouter()

	userService := service.NewUserService()    // UserServiceの初期化
	userHandler := NewUserHandler(userService) // UserHandlerの初期化

	router.HandleFunc("/", userHandler.FindOne).Methods("GET")
	router.HandleFunc("/all", userHandler.AllSelect).Methods("GET")
	router.HandleFunc("/register", userHandler.RegisterUserInfo).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Reactアプリのオリジンを許可
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
	})

	return c.Handler(router)
}
