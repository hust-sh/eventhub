package main

import (
	"os"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
)

var RedisUrl = "redis://redis:6379"

var OauthCallbacks = map[string]string{}

const htmlIndex = `<html><body>
<a href="/githublogin">Log in with Github</a></br>
<a href="/googlelogin">Log in with Google</a>
</body></html>
`

var (
    githubOauthConfig = &oauth2.Config {
                        RedirectURL:  "http://localhost:3003/callback/github",
                        ClientID:     os.Getenv("GITHUBKEY"), 
                        ClientSecret: os.Getenv("GITHUB_SECRET"),
                        Scopes:       []string{"user", "repo"},
                        Endpoint:     github.Endpoint}
    githubStateString = "githubmock"   // todo: should be random
    googleOauthConfig = &oauth2.Config {
                        RedirectURL:   "http://localhost:3003/callback/google",
                        ClientID:      os.Getenv("GOOGLEKEY"),
                        ClientSecret:  os.Getenv("GOOGLE_SECRET"),
                        Scopes:        []string{"https://www.googleapis.com/auth/gmail.metadata"},
                        Endpoint:      google.Endpoint,
    }
    googleStateString = "googlemock"   //todo: should be random
)

