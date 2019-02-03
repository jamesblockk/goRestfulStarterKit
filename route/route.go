package route

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the list of routes of our API
type Routes []Route

var routes = Routes{
	Route{
		"Hello",
		"GET",
		"/",
		Hello,
	},
	Route{
		"CreateAlbum",
		"POST",
		"/album/{title}",
		CreateAlbum,
	},
	Route{
		"ReadAlbum",
		"Get",
		"/album/{title}",
		ReadAlbum,
	},
	Route{
		"UpdateAlbum",
		"PUT",
		"/album/{title}",
		UpdateAlbum,
	},
	Route{
		"DeleteAlbum",
		"DELETE",
		"/album/{title}",
		DeleteAlbum,
	},
	Route{
		"GetAllAlbum",
		"GET",
		"/album/",
		GetAllAlbum,
	},
}

// Hello :
func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World: "))
}

// New setup the router
func New() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		log.Println(route.Name)
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
