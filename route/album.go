package route

import (
	"net/http"

	"github.com/gorilla/mux"
)

// CreateAlbum :
func CreateAlbum(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	w.Write([]byte("CreateAlbum title is " + title))
}

// ReadAlbum :
func ReadAlbum(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	w.Write([]byte("ReadAlbum title is " + title))
}

// UpdateAlbum :
func UpdateAlbum(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	w.Write([]byte("UpdateAlbum title is " + title))
}

// DeleteAlbum :
func DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	w.Write([]byte("DeleteAlbum title is " + title))
}

// GetAllAlbum :
func GetAllAlbum(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetAllAlbum !!"))
}
