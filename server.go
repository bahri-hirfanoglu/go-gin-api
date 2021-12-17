package main

import (
	"gin-framework-api/middlewares"
	"gin-framework-api/routes"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func LogOutPath(file string) {
	f, _ := os.Create(file)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	USER := os.Getenv("USER_NANE")
	PASS := os.Getenv("PASS_WORD")
	FILE := os.Getenv("FILE_NAME")
	PORT := os.Getenv("SERVER_PORT")

	LogOutPath(FILE)
	server := gin.New()

	server.Use(gin.Recovery(), gin.Logger(), middlewares.BasicAuth(USER, PASS))
	routes.Routes(server)
	log.Fatal(server.Run(":" + PORT))
}
