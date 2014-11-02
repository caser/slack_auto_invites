// package main

package slack_auto_invites

import (
	"appengine"
	"appengine/urlfetch"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Configuration struct {
	BaseUrl string
	Token   string
}

func importConfiguration() (string, string) {
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration.BaseUrl, configuration.Token
}

func SetSlackToken(req *http.Request, token string) {
	q := req.URL.Query()
	q.Set("token", token)
	req.URL.RawQuery = q.Encode()
}

func SetFormValues(req *http.Request, fname string, lname string, email string) {
	q := req.URL.Query()
	q.Set("first_name", fname)
	q.Set("last_name", lname)
	q.Set("email", email)
	q.Set("set_active", "true")
	q.Set("_attempts", "1")
	req.URL.RawQuery = q.Encode()
}

func SendInvite(r *http.Request, fname string, lname string, email string) string {
	// client := &http.Client{}

	c := appengine.NewContext(r)
	client := urlfetch.Client(c)

	baseUrl, token := importConfiguration()

	req, _ := http.NewRequest("POST", baseUrl, nil)

	SetSlackToken(req, token)
	SetFormValues(req, fname, lname, email)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request error is: ", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Reading of response error is: ", err)
	}

	return string(body)
	// TODO - add error checking for body of response
	// success: {"ok":true}
	// failure: {"ok":false,"error":"already_in_team"}
}

func inviteHandler(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("fname") != "" && r.FormValue("lname") != "" && r.FormValue("email") != "" {
		fname := r.FormValue("fname")
		lname := r.FormValue("lname")
		email := r.FormValue("email")

		slack_resp := SendInvite(r, fname, lname, email)
		fmt.Fprintf(w, slack_resp)
	} else {
		fmt.Fprintf(w, "No form values given. Please supply first name, last name, and email.")
	}
}

func init() {
	// func main() {
	http.HandleFunc("/", inviteHandler)
}
