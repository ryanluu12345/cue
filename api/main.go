package main

import (
	"fmt"
	"net/http"
	"golang.org/x/oauth2"
)

func main() {
	http.HandleFunc("/user/auth", handleAuth)
	http.ListenAndServe(":8080", nil)
}

func handleAuth(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
			getOAuthToken()
		case "DELETE":
			rmOAuthToken()
	}
}

var conf = &oauth2.Config{
	ClientID:     "YOUR_CLIENT_ID",
	ClientSecret: "YOUR_CLIENT_SECRET",
	Scopes:       []string{"SCOPE1", "SCOPE2"},
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://provider.com/o/oauth2/auth",
		TokenURL: "https://provider.com/o/oauth2/token",
	},
}

func getOAuthToken() {

}

func rmOAuthToken() {

}
