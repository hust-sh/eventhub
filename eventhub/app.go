package main

import (
    "log"
    "net/http"
    "os"
    "os/signal"
    "time"
    "context"
)

func main() {

    router := setRouter()
    
    srv := &http.Server{
        Addr: ":3003",
        Handler: router,
    }

    go func() {
        if err := srv.ListenAndServe(); err != nil {
            log.Printf("listen: %s\n", err)
        }
    }()

    // Wait for interrupt signal to gracefully shutdown the server
    // with a timeout of 5 seconds
    quit := make(chan os.Signal)
    signal.Notify(quit, os.Interrupt)
    <-quit

    ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancle()
    if err := srv.Shutdown(ctx); err != nil {
        log.Fatal("Server Shutdown:", err)
    }

    log.Println("Server exiting")
}
