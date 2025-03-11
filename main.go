package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	oauthConf        *oauth2.Config
	oauthStateString = "randomstring" // セキュリティのためにランダムな文字列を使用
)

	func init() {
		log.Println("Client ID:", os.Getenv("GOOGLE_CLIENT_ID"))
		oauthConf = &oauth2.Config{
			RedirectURL:  "http://localhost:8080/callback",
			ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),  // 環境変数からクライアントIDを取得
			ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
			Endpoint:     google.Endpoint,
		}
	}
	


func main() {
	http.HandleFunc("/", handleMain)
	http.HandleFunc("/login", handleGoogleLogin)
	http.HandleFunc("/callback", handleGoogleCallback)

	fmt.Println("Started running on http://localhost:8080")
	fmt.Println(http.ListenAndServe(":8080", nil))
}

func handleMain(w http.ResponseWriter, _ *http.Request) {
	const htmlIndex = `<html><body><a href="/login">Google Login</a></body></html>`
	_, err := fmt.Fprint(w, htmlIndex)
	if err != nil {
		return
	}

}
func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := oauthConf.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	token, err := oauthConf.Exchange(context.Background(), code)
	if err != nil {
		fmt.Printf("oauthConf.Exchange() failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	client := oauthConf.Client(context.Background(), token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		fmt.Printf("Get: %s\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	userInfo := struct {
		Email string `json:"email"`
	}{}
	err = json.NewDecoder(response.Body).Decode(&userInfo)
	if err != nil {
		return
	}
	_, err = fmt.Fprintf(w, "User Info: %s\n", userInfo.Email)
	if err != nil {
		return
	}
}
