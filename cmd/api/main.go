package main

import (
	"log"
	"net/http"
	"rest-api/internal/database"
	"rest-api/internal/database/repositories"
	"rest-api/internal/handlers"
	"rest-api/internal/middleware"
)

func main() {
	dbUrl := "postgres://postgres:5432@postgres:5432/RestApi?sslmode=disable"

	db, err := database.Connect(dbUrl)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	taskRepo := repositories.NewTaskRepo(db)
	taskHandler := handlers.NewTaskHandler(taskRepo)

	mws := []middleware.Middleware{
		middleware.Logging,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/tasks", func(w http.ResponseWriter, req *http.Request) {
        switch req.Method {
        case http.MethodGet:
            methodHandler(taskHandler.GetAll, http.MethodGet, mws...)(w, req)
        case http.MethodPost:
            methodHandler(taskHandler.Create, http.MethodPost, mws...)(w, req)
        default:
            http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        }
    })

    mux.HandleFunc("/tasks/", func(w http.ResponseWriter, req *http.Request) {
        switch req.Method {
        case http.MethodGet:
            methodHandler(taskHandler.GetById, http.MethodGet, mws...)(w, req)
        case http.MethodPut:
            methodHandler(taskHandler.Update, http.MethodPut, mws...)(w, req)
        default:
            http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
        }
    })

	err = http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatal(err)
	}
}

func methodHandler(handlerFunc http.HandlerFunc, allowedMethod string, mws ...middleware.Middleware) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method != allowedMethod {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		chain := middleware.ChainMiddleware(handlerFunc, mws...)
		chain.ServeHTTP(w, req)
	}
}
	