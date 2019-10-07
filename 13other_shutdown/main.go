package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		time.Sleep(10 * time.Second)
		c.String(200, "%s", "zZZZZZZ")
	})

	//r.Run("localhost:8080")
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s \n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shut down")
	ctx, cancle := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancle()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("shut down", err)
	}
	log.Println("server exiting")
}
