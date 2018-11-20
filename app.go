// app.go

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/sbourne20/examgo5/controller"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	//connectionString := fmt.Sprintf("%s:%s@tcp(rm-d9jxms81910c50lvi.mysql.ap-southeast-5.rds.aliyuncs.com:3306)/%s?charset=utf8&parseTime=True", user, password, dbname)
	connectionString := fmt.Sprintf("%s:%s@tcp(192.168.8.32:3306)/%s?charset=utf8&parseTime=True", user, password, dbname)

	var err error

	a.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {

	a.Router.HandleFunc("/testAgent", controller.TestAgents).Methods("GET")
	a.Router.HandleFunc("/users", a.getUsers).Methods("GET")
	a.Router.HandleFunc("/user", a.createUser).Methods("POST")
	a.Router.HandleFunc("/user/{id}", a.updateUser).Methods("PUT")
	a.Router.HandleFunc("/user/{id}", a.deleteUser).Methods("DELETE")
	a.Router.HandleFunc("/getNews/{id}", controller.GetNews).Methods("GET")
}

func (a *App) getUsers(w http.ResponseWriter, r *http.Request) {
	controller.GetUsers(w, r, a.DB)
}

func (a *App) createUser(w http.ResponseWriter, r *http.Request) {
	controller.CreateUser(w, r, a.DB)
}

func (a *App) updateUser(w http.ResponseWriter, r *http.Request) {
	controller.UpdateUser(w, r, a.DB)
}

func (a *App) deleteUser(w http.ResponseWriter, r *http.Request) {
	controller.DeleteUser(w, r, a.DB)
}
