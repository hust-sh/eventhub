/*
Oauth2 Protocol Flow:

     +--------+                               +---------------+
     |        |--(A)- Authorization Request ->|   Resource    |
     |        |                               |     Owner     |
     |        |<-(B)-- Authorization Grant ---|               |
     |        |                               +---------------+
     |        |
     |        |                               +---------------+
     |        |--(C)-- Authorization Grant -->| Authorization |
     | Client |                               |     Server    |
     |        |<-(D)----- Access Token -------|               |
     |        |                               +---------------+
     |        |
     |        |                               +---------------+
     |        |--(E)----- Access Token ------>|    Resource   |
     |        |                               |     Server    |
     |        |<-(F)--- Protected Resource ---|               |
     +--------+                               +---------------+

Oauth的过程实际上是应用程序(Client)通过用户(Resource Owner)授权(Grant)获得访问指定服务器(Resource Server)资源的过程。
授权过程：
	A. 应用程序(Client)询问用户是否愿意给用户的开放指定的权限（开放的权限内容有下文的Scope指定）。(此时url会跳转到授权服务的授权页面。)
    B. 若用户选择愿意，授权服务器(Authorization Server)将会重定向到Client指定的一个url callback中，并携带一个临时的code。
    C/D. 客户端通过code向授权中心换取Access Token
    E. Client通过Access Token可访问资源服务器(Resource Server)的指定内容(由Scope决定）
    F. 资源服务器返回Client 相应的资源数据


Google 授权过程：
    1. Google要求先要在Google Cloud Platform(https://console.developers.google.com/)中创建自己的web app，并获得OAuth client ID和client secret。
    2. google授权中心授权之后需重定向到客户端程序，并携带code。因此客户端程序需准备好回调url，如下文的http://localhost:3000/GoogleCallback
    3. 客户端程序解析回调url，并获得code。再通过code换取Access Token
    4. 最后客户端程序便可以使用访问对应的资源了

关键词:

- oauth2.Config 源码注释

    type Config struct {
        // ClientID is the application's ID.
        ClientID string
    
        // ClientSecret is the application's secret.
        ClientSecret string
    
        // Endpoint contains the resource server's token endpoint
        // URLs. These are constants specific to each server and are
        // often available via site-specific packages, such as
        // google.Endpoint or github.Endpoint.
        Endpoint Endpoint
    
        // RedirectURL is the URL to redirect users going through
        // the OAuth flow, after the resource owner's URLs.
        RedirectURL string
    
        // Scope specifies optional requested permissions.
        Scopes []string
    }

- Endpoint
    google的Endpoint定义如下:

    // google.Endpoint is Google's OAuth 2.0 endpoint.
    var Endpoint = oauth2.Endpoint{
        AuthURL:  "https://accounts.google.com/o/oauth2/auth",
        TokenURL: "https://accounts.google.com/o/oauth2/token",
    }

- Scopes
  Scopes指明当次授权请求期望访问哪些资源(可指定多条)
  如https://www.googleapis.com/auth/gmail.metadata代表访问gmail的元数据
  如何知道google的哪些资源可访问，可参考google API或者前往oauthplayground探索(https://developers.google.com/oauthplayground/)

- ClientID, ClientSecret
  需再某个google应用中(如开发这的google app)中创建凭据中获取


Inspired by https://jacobmartins.com/2016/02/29/getting-started-with-oauth2-in-go/ 
*/

package main

import (
    "net/http"
    "log"
	"golang.org/x/oauth2"
    "github.com/gin-gonic/gin"
)

func GoogleLoginHandler(c *gin.Context) {
    url := googleOauthConfig.AuthCodeURL(googleStateString)
    log.Printf(url)
    c.Redirect(http.StatusTemporaryRedirect, url)
}


func GoogleCallbackHandler(c *gin.Context) {
    state := c.Query("state")
    if state != googleStateString {
        log.Printf("invalid oauth state, expectd %s, but %s\n", googleStateString , state)
        c.Redirect(http.StatusTemporaryRedirect, "/siteentry")
        return
    }

    code := c.Query("code")
    log.Printf("code:%s", code)
    token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
    if err != nil {
        log.Printf("oauthConfig.Exchange() failed with %s\n", err)
        c.Redirect(http.StatusTemporaryRedirect, "/siteentry")
        return
    }

    log.Printf("token:%s\n", token.AccessToken)

    c.String(http.StatusOK, "Oauth Success")
}
