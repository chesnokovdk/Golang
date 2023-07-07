package main

import (
	"fmt"
	"net/http"
	"pr1/controllers"
	"pr1/middleware"
	"pr1/models"
	"pr1/views"

	"github.com/gorilla/mux"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "qwe123"
	dbname   = "lenslocked_dev"
)

var (
	homeView    *views.View
	contactView *views.View
	faqView     *views.View
)

var h http.Handler = http.HandlerFunc(err404)

func err404(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>We are terribly sorry</h1>")
	fmt.Fprint(w, "<h1>But this is <b>Error 404</b></h1>")
}

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	services, err := models.NewServices(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer services.Close()
	services.AutoMigrate()

	// us.DestructiveReset()

	r := mux.NewRouter()

	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers(services.User)
	galleriesC := controllers.NewGalleries(services.Gallery, services.Image, r)

	userMw := middleware.User{
		UserService: services.User,
	}
	requireUserMw := middleware.RequireUser{}

	// galleriesC.New is an http.Handler, so we use Apply
	newGallery := requireUserMw.Apply(galleriesC.New)
	// galleriecsC.Create is an http.HandlerFunc, so we use ApplyFn
	createGallery := requireUserMw.ApplyFn(galleriesC.Create)

	r.NotFoundHandler = h
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.Handle("/faq", staticC.Faq).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	r.Handle("/login", usersC.LoginView).Methods("GET")
	r.HandleFunc("/login", usersC.Login).Methods("POST")
	r.HandleFunc("/cookietest", usersC.CookieTest).Methods("GET")

	// Image routes
	imageHandler := http.FileServer(http.Dir("./images/"))
	r.PathPrefix("/images/").Handler(http.StripPrefix("/images/", imageHandler))

	r.HandleFunc("/galleries/{id:[0-9]+}/images", requireUserMw.ApplyFn(galleriesC.ImageUpload)).Methods("POST")

	r.Handle("/galleries/new", newGallery).Methods("GET")
	r.Handle("/galleries", createGallery).Methods("POST")
	r.Handle("/galleries", requireUserMw.ApplyFn(galleriesC.Index)).Methods("GET").Name(controllers.IndexGalleries)
	r.HandleFunc("/galleries/{id:[0-9]+}/edit", requireUserMw.ApplyFn(galleriesC.Edit)).Methods("GET").Name(controllers.EditGallery)
	r.HandleFunc("/galleries/{id:[0-9]+}/update", requireUserMw.ApplyFn(galleriesC.Update)).Methods("POST")
	r.HandleFunc("/galleries/{id:[0-9]+}/delete", requireUserMw.ApplyFn(galleriesC.Delete)).Methods("POST")
	r.HandleFunc("/galleries/{id:[0-9]+}", galleriesC.Show).Methods("GET").Name(controllers.ShowGallery)

	fmt.Println("Starting server on: http://localhost:3000/")
	http.ListenAndServe(":3000", userMw.Apply(r))
}
