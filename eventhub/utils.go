package main

import (
    "github.com/satori/go.uuid"
    "net/http"
    "strings"
    "github.com/garyburd/redigo/redis"
)


func GenAccessToken() string {

    accessToken := uuid.NewV4()
    return accessToken.String()
}


// 重构：考虑使用中间件location获取
func getScheme(r *http.Request) string{

    scheme := "http://"
    if r.TLS != nil {
        scheme = "https://"
    }

    return scheme
}


func GenWebhook(site string, token string, r *http.Request) string{

    scheme := getScheme(r)
    webhook := strings.Join([]string{scheme, r.Host, "/webhook/", site, "/", token}, "")
    return webhook
}


func IsValidSiteType(site string) bool {
    _, err := SiteMapper[site]
    return err
}


func GetWebhook(site string, token string) (string, error) {

    redisCli, _ :=  GetRedis()
    return redis.String(redisCli.Do("HGET", site, token))
}


func GetRedis() (redis.Conn, error) {

    return redis.DialURL(RedisUrl)
}


/*
// bad taste
func GetOauthCallback(site string, r *http.Request) string{
    
    v, err = OauthCallbacks[site]
    if err {
        return v
    }
    scheme := getScheme(r)
    callback := strings.Join([]string{scheme, r.Host, "/callback/", site}, "")
    OauthCallbacks[site] = callback
    return callback
}
*/


func GenStateString(site string) string {
    
    return ""
}
