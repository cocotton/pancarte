package pancarte

import (
	"log"
	"net/http"
	"os"

	"github.com/cocotton/pancarte/pancarte/authentication"
	"github.com/cocotton/pancarte/pancarte/door"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
)

// Pancarte structure is used to manage Pancarte
type Pancarte struct {
	DB     *mgo.Session
	Router *mux.Router
}

// InitDB initializes the connection with the database
func (p *Pancarte) InitDB(dbName string) {
	var err error

	p.DB, err = mgo.Dial(dbName)
	if err != nil {
		log.Fatal(err)
	}
	p.DB.SetMode(mgo.Monotonic, true)
}

// InitRouter initializes the mux router and routes
func (p *Pancarte) InitRouter() {
	p.Router = mux.NewRouter()

	p.Router.HandleFunc("/addDoor", authentication.Validate(func(w http.ResponseWriter, r *http.Request) {
		door.AddDoor(w, r, p.DB)
	})).Methods("POST")
	p.Router.HandleFunc("/getDoor/{doorID}", func(w http.ResponseWriter, r *http.Request) {
		door.GetDoor(w, r, p.DB)
	}).Methods("GET")
	p.Router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		authentication.Login(w, r, p.DB)
	}).Methods("POST")
	p.Router.HandleFunc("/logout", authentication.Logout).Methods("GET")
}

// CheckEnv checks if the required environment variables exist
func (p *Pancarte) CheckEnv() {
	if len(os.Getenv("PANCARTE_SECRET")) == 0 {
		log.Fatal("Environment variable Pancarte_Secret is not set. Exiting.")
	}
}

// Run launches the http server
func (p *Pancarte) Run(port string) {
	http.ListenAndServe(port, handlers.LoggingHandler(os.Stdout, p.Router))
}
