package app

import (
  "fmt"
  "log"
  "net/http"

  "github.com/api/app/handler"
  "github.com/api/app/model"
  "github.com/api/config"
  "github.com/gorilla/mux"
  "github.com/jinzhu/gorm"
)

// App has router and db instances
type App struct {
  Router *mux.Router
  DB     *gorm.DB
}

// App initialize with predefined configuration
func (a *App) Initialize(config *config.Config) {
  dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
    config.DB.Username,
    config.DB.Password,
    config.DB.Name,
    config.DB.Charset)

  db, err := gorm.Open(config.DB.Dialect, dbURI)
  if err != nil {
    log.Fatal("Could not connect database")
  }

  a.DB = model.DBMigrate(db)
  a.Router = mux.NewRouter()
  a.setRouters()
}

// Set all required routers
func (a *App) setRouters() {
  // Routing for handling the projects
  a.Get("/famouspersons", a.GetAllFamousPersons)
  a.Post("/famouspersons", a.CreateFamousPerson)
  a.Get("/famouspersons/{name}", a.GetFamousPerson)
  a.Put("/famouspersons/{name}", a.UpdateFamousPerson)
  a.Delete("/famouspersons/{name}", a.DeleteFamousPerson)
}

// Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
  a.Router.HandleFunc(path, f).Methods("GET")
}

// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
  a.Router.HandleFunc(path, f).Methods("POST")
}

// Wrap the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
  a.Router.HandleFunc(path, f).Methods("PUT")
}

// Wrap the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
  a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Handlers to manage FamousPerson Data
func (a *App) GetAllFamousPersons(w http.ResponseWriter, r *http.Request) {
  handler.GetAllFamousPersons(a.DB, w, r)
}

func (a *App) CreateFamousPerson(w http.ResponseWriter, r *http.Request) {
  handler.CreateFamousPerson(a.DB, w, r)
}

func (a *App) GetFamousPerson(w http.ResponseWriter, r *http.Request) {
  handler.GetFamousPerson(a.DB, w, r)
}

func (a *App) UpdateFamousPerson(w http.ResponseWriter, r *http.Request) {
  handler.UpdateFamousPerson(a.DB, w, r)
}

func (a *App) DeleteFamousPerson(w http.ResponseWriter, r *http.Request) {
  handler.DeleteFamousPerson(a.DB, w, r)
}

// Run the app on it's router
func (a *App) Run(host string) {
  log.Fatal(http.ListenAndServe(host, a.Router))
}
