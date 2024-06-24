package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// ------artistsList------//
type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

// --------typeList--------//
type DatesLocation struct {
	Index []DateLocation `json:"index"`
}

// --------typeBase--------//
type DateLocation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func GetFromJSON(link string, result interface{}) {
	resp, err := http.Get(link)
	if err != nil {
		CheckError(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		CheckError(err)
	}

	err = json.Unmarshal(body, result)
	if err != nil {
		CheckError(err)
	}

}

func CheckError(err error) {
	if err != nil {
		log.Printf("HTTP Response Code : %v", (http.StatusInternalServerError))
		fmt.Println(err)
		return
	}
}
