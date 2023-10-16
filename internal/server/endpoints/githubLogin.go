package endpoints

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"net/http"
)

var (
	githubOauthConfig = oauth2.Config{
		ClientID:     "065d047663d40d183c04",
		ClientSecret: "7b7c2239b98e0b66d53e6b2adbfd8722561512f4",
		RedirectURL:  "http://localhost:5173/profile",
		Endpoint:     github.Endpoint,
		Scopes:       []string{"user:email"},
	}
)

func GithubLog() {
	http.HandleFunc("/login", handleLogin)

	http.HandleFunc("/auth/callback", handleCallback)

	http.ListenAndServe(":3000", nil)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	url := githubOauthConfig.AuthCodeURL("", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusFound)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	// Handle the OAuth callback from GitHub
	code := r.URL.Query().Get("code")
	_, err := githubOauthConfig.Exchange(r.Context(), code)
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Use the token to make requests to GitHub API on behalf of the user
	// You can also store the token in the database for future use
	// ...
}
