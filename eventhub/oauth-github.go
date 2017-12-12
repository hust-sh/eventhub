package main

import (
    "fmt"
	"net/http"
	"golang.org/x/oauth2"
    "log"
    "github.com/gin-gonic/gin"
    "github.com/google/go-github/github"
)


func GithubLoginHandler(c *gin.Context) {
	url := githubOauthConfig.AuthCodeURL(githubStateString)
    log.Printf(url)
    c.Redirect(http.StatusTemporaryRedirect, url)
}


func GithubCallbackHandler(c *gin.Context) {

    state := c.Query("state")
	if state != githubStateString {
		log.Printf("invalid oauth state, expected '%s', got '%s'\n", githubStateString, state)
        c.Redirect(http.StatusTemporaryRedirect, "/siteentry")
		return
	}

	code := c.Query("code")
    log.Printf("code:%s", code)
	token, err := githubOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Printf("oauthConf.Exchange() failed with '%s'\n", err)
        c.Redirect(http.StatusTemporaryRedirect, "/siteentry")
		return
	}
    log.Printf("token: %s", token.AccessToken)

    oauthClient := githubOauthConfig.Client(oauth2.NoContext, token)
    client := github.NewClient(oauthClient)
    user, _, err := client.Users.Get(c, "")
    if err != nil {
        log.Printf("client.User.Get() failed with %s\n", err)
        c.Redirect(http.StatusTemporaryRedirect, "/siteentry")
        return
    }

    repos, repo_err := GetRepositories(c, client)
    if repo_err != nil {
        log.Printf("get repo failed")
        c.Redirect(http.StatusTemporaryRedirect, "/siteentry")
        return
    }

    var response = "Here is the repositories you owed:\n"
    for _, repo := range repos {
        response = fmt.Sprintf("%s %s\n", response, *repo.FullName) 
    }
    
    response = fmt.Sprintf("Hello %s.\n %s", *user.Login, response)
    log.Printf(response)
    c.String(http.StatusOK, response)

}


func GetRepositories(ctx *gin.Context, client *github.Client) ([]*github.Repository, error) {

    opt := &github.RepositoryListOptions{
        ListOptions: github.ListOptions{PerPage:10},
    }
    
    var allRepos []*github.Repository
    for {
        repos, resp, err := client.Repositories.List(ctx, "", opt)
        if err != nil{
            log.Printf("get repos failed, %v", err)
            return allRepos, err
        }

        allRepos = append(allRepos, repos...)
        if resp.NextPage == 0 {
            break
        }
        opt.Page = resp.NextPage
    }

    return allRepos, nil
}


func GetOrganizations(ctx *gin.Context, client *github.Client) ([]*github.Organization, error) {

    opt := &github.OrganizationsListOptions{
        ListOptions: github.ListOptions{PerPage: 10},
    }
    var allOrganis []*github.Organization
    for {
        Organis, resp, err := client.Organizations.ListAll(ctx, opt)
        if err != nil {
            log.Printf("get organization failed. %v", err)
            return allOrganis, err
        }
        allOrganis = append(allOrganis, Organis...)
        if resp.NextPage == 0 {
            break
        }
        opt.Page = resp.NextPage
    }

    return allOrganis, nil
}
