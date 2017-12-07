package main

import (
	"golang.org/x/oauth2"
	"os"
	"golang.org/x/oauth2/github"
)

var RedisUrl = "redis://redis:6379"

var OauthCallbacks = map[string]string{}

var (
    githubOauthConfig = &oauth2.Config {
                        RedirectURL:  "http://localhost:3003/callback/github",
                        ClientID:     os.Getenv("GITHUBKEY"), 
                        ClientSecret: os.Getenv("GITHUB_SECRET"),
                        Scopes:       []string{"user", "repo"},
                        Endpoint:     github.Endpoint}
    githubStateString = "mock"
)

