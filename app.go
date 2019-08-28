package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rest_api_example/models"
	"log"
	"net/http"
	"strconv"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (app *App) Initialize(user, password, dbName string) {
	connectionString := fmt.Sprintf("%s:%s@/%s", user, password, dbName)

	var err error
	app.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	app.Router = mux.NewRouter()
	app.initializeRoutes()
}

func (app *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, app.Router))
}

func (app *App) initializeRoutes() {
	app.Router.HandleFunc("/users", app.getUsers).Methods("GET")
	app.Router.HandleFunc("/users", app.createUser).Methods("POST")
	app.Router.HandleFunc("/users/{id:[0-9]+}", app.getUser).Methods("GET")
	app.Router.HandleFunc("/users/{id:[0-9]+}", app.updateUser).Methods("PUT")
	app.Router.HandleFunc("/users/{id:[0-9]+}", app.deleteUser).Methods("DELETE")
}

func (app *App) getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responseWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	u := models.User{ID: id}
	if err := u.GetUser(app.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			responseWithError(w, http.StatusNotFound, "User not found")
		default:
			responseWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	responseWithJSON(w, http.StatusOK, u)
}

func (app *App) getUsers(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count < 0 || count > 20 {
		count = 20
	}

	if start < 0 {
		start = 1
	}

	user := models.User{}

	users, err := user.GetUsers(app.DB, start, count)
	if err != nil {
		responseWithJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseWithJSON(w, http.StatusOK, users)
}

func (app *App) createUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		responseWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := user.CreateUser(app.DB); err != nil {
		responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseWithJSON(w, http.StatusCreated, user)
}

func (app *App) updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responseWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user := models.User{ID: id}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		responseWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := user.UpdateUser(app.DB); err != nil {
		responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseWithJSON(w, http.StatusOK, user)
}

func (app *App) deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		responseWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user := models.User{ID: id}
	if err := user.DeleteUser(app.DB); err != nil {
		responseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func responseWithError(w http.ResponseWriter, code int, message string) {
	responseWithJSON(w, code, map[string]string{"error": message})
}

func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
