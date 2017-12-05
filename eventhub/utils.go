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


func GenWebhook(site string, token string, r *http.Request) string{
    scheme := "http://"
    if r.TLS != nil {
        scheme = "https://"
    }

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
