package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig = &oauth2.Config{
	ClientID:     "YOUR_CLIENT_ID",     // Google Cloud Consoleで取得したクライアントID
	ClientSecret: "YOUR_CLIENT_SECRET", // Google Cloud Consoleで取得したクライアントシークレット
	RedirectURL:  "http://localhost:8080/callback",
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile"},
	Endpoint:     google.Endpoint,
}

func loginHandler(w http.Response, r * http.Request){
	url := googleOauthConfig.AuthCodeURL("ramdomastate" , oauth2.AccesTypeTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)

}

