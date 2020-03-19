package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"lsc-ecosys/ota-poc/otaku/api/middlewares"
	"lsc-ecosys/ota-poc/otaku/api/responses"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Initialize connection to db and wire up routes
func (a *App) Initialize(DBHost, DBPort, DBUser, DBName, DBPassword string) {
	var err error
	DBURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DBHost, DBPort, DBUser, DBName, DBPassword)

	fmt.Println(DBURI)
	a.DB, err = gorm.Open("postgres", DBURI)
	if err != nil {
		fmt.Printf("\n Cannot connect to DB %s", DBName)
		log.Fatal("Error: ", err)
	} else {
		fmt.Printf("Successfully connected to DB %s", DBName)
	}

	//a.DB.Debug().AutoMigrate(&models.Artifact{}) // auto migration

	a.Router = mux.NewRouter().StrictSlash(true)
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	a.Router.Use(middlewares.SetContentTypeMiddleware) // setting content-type

	a.Router.HandleFunc("/", home).Methods("GET")
	a.Router.HandleFunc("/artifacts", a.GetAllArtifacts).Methods("GET")

	a.Router.HandleFunc("/rollouts", a.GetAllRollouts).Methods("GET")
	a.Router.HandleFunc("/rollout", a.CreateRollout).Methods("POST")

	a.Router.HandleFunc("/artifact/check", a.CheckUpdateByVersion).Methods("GET").Queries("version", "{version}")
}

func (a *App) RunServer() {
	log.Printf("\nServer starting on port 5000")
	log.Fatal(http.ListenAndServe(":5000", a.Router))
}

func home(w http.ResponseWriter, r *http.Request) { // this is the home route
	responses.JSON(w, http.StatusOK, "Welcome To OTAKU")
}
