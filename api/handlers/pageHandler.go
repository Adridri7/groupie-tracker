package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Page struct {
	Title string
	Error string

	Data1 []Artist
	Data2 DatesLocation

	Data5 Artist
	Data6 map[string][]string
}

var VarArtists []Artist
var VarDateLocation DatesLocation

const webTitle string = "Groupie-Tracker"

func RenderTemplate(w http.ResponseWriter, tmpl string, page Page) {
	t, err := template.ParseFiles("./web/templates/" + tmpl + ".html")

	if err != nil {
		ErrorPage(w, http.StatusBadRequest, Page{Title: webTitle, Error: "Oops, are you looking for a ghost..?"})
		return
	}

	if err := t.Execute(w, page); err != nil {
		ErrorPage(w, http.StatusBadRequest, Page{Title: webTitle, Error: "Oops, are you looking for a ghost..?"})
		return
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	//VarArtists = VarArtists[0:17]
	if VarArtists == nil {
		ErrorPage(w, http.StatusInternalServerError, Page{Title: webTitle, Error: "Oops, looks like the groupie API is empty..."})
		return
	}
	if r.URL.Path != "/" {
		idStr := string(r.URL.Path[1:])
		id, err := strconv.Atoi(idStr)
		id -= 1
		if err != nil || id < 0 || id > len(VarArtists)-1 {
			ErrorPage(w, http.StatusNotFound, Page{Title: webTitle, Error: "Oops, the page you were looking for could not be found... :)"})
			return
		} else {
			RenderTemplate(w, "Artist", Page{Title: webTitle, Data5: VarArtists[id], Data6: VarDateLocation.Index[id].DatesLocations})
			return
		}
	}

	switch r.Method {
	case "GET":
		RenderTemplate(w, "Home", Page{Title: webTitle, Data1: VarArtists})
		log.Printf("HTTP Response Code : %v", (http.StatusOK))
	default:
		ErrorPage(w, http.StatusMethodNotAllowed, Page{Title: webTitle, Error: "Oops, looks like you're not allowed to do this.."})
	}
}

func About(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about" {
		ErrorPage(w, http.StatusNotFound, Page{Title: webTitle, Error: "Oops, the page you were looking for could not be found... :)"})
		return
	}

	switch r.Method {
	case "GET":
		RenderTemplate(w, "About", Page{Title: webTitle})
		log.Printf("HTTP Response Code : %v", (http.StatusOK))
	default:
		ErrorPage(w, http.StatusMethodNotAllowed, Page{Title: webTitle, Error: "Oops, looks like you're not allowed to do this.."})
	}
}

func ErrorPage(w http.ResponseWriter, errorCode int, page Page) {
	RenderTemplate(w, "Error", page)
	log.Printf("HTTP Response Code : %v", errorCode)
}
