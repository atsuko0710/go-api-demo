package main

import (
	"errors"
	"go-api-demo/router"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main()  {
	// gin.SetMode(gin.ReleaseMode)
	g := gin.New()

	middlewares := []gin.HandlerFunc{}

	router.Load(
		g,
		middlewares...
	)

	go func ()  {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Print("The router has been deployed successfully.")
	}()

	log.Printf("Start to listening the incoming requests on http address: %s", ":8080")
	log.Printf(http.ListenAndServe(":8080", g).Error())
}

// 健康检查自检
func pingServer() error {
	for i := 0; i < 2; i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get("http://127.0.0.1:8080" + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		
		// 两次健康自检都不成功则终止进程
		// Sleep for a second to continue the next ping.
		log.Print("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}