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

func loginHandler(w http.Response, r *http.Request) {
	url := googleOauthConfig.AuthCodeURL("ramdomastate", oauth2.AccesTypeTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
func callbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "No Code in the URL", http.StatusBadRequest)
		return
	}

	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
//静的ファイルの提供
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets",http.StripPrefix())
}
