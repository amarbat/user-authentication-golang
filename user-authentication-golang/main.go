package main

import (
  "os"

  middleware "user-authentication-golang/middleware"
  routes "user-authentication-golang/routes"

  "github.com/gin-gonic/gin"
  _ "github.com/heroku/x/hmetrics/onload"
)

func main() {
// -- get port number from environment variables
  port                  := os.Getenv("PORT")

  if port == "" {
// fallback
    port                 = "8000"
  }

  router                := gin.New()
  router.Use(
    gin.Logger())
  routes.UserRoutes(
    router)

  router.Use(
    middleware.Authentication())

// TEST API-1
  router.GET("/api-1", func(c *gin.Context) {

    c.JSON(200, gin.H{"success": "Access granted for api-1"})

  })

// TEST API API-2
  router.GET("/api-2", func(c *gin.Context) {
    c.JSON(200, gin.H{"success": "Access granted for api-2"})
  })

  router.Run(
    ":" + port)
}