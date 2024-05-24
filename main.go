package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/MatiF100/Throw-Muffin-API/controllers"
	"github.com/MatiF100/Throw-Muffin-API/database"
	"github.com/MatiF100/Throw-Muffin-API/middlewares"
	"github.com/gin-gonic/gin"
	"gopkg.in/fsnotify.v1"

	"github.com/gofor-little/env"

	_ "github.com/lib/pq"
)

type App struct {
	port       string
	local_mode bool
}

func main() {
	var app App = App{}

	setup_env(&app)
	setup_db(&app)
	setup_azure(&app)

	router := initRouter()

	router.Run(":8080")
}

func Ping(context *gin.Context) {
	context.JSON(200, gin.H{"message": "2137pong"})
}

func initRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")
	{
		router.GET("/", Ping)
		auth := api.Group("/auth")
		{
			auth.POST("/login", controllers.GenerateToken)
			auth.POST("/register", controllers.RegisterUser)
			auth.POST("/refresh-token", controllers.RefreshToken)
		}

		secured := api.Group("").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}

	return router
}

func setup_env(app *App) {
	if err := env.Load(".env"); err != nil {
		log.Println("Failed to locate .env file. Using system variables and/or defaults")
		app.local_mode = true
	}
}

func setup_db(app *App) {
	dbString := env.Get("dbString", fmt.Sprintf("host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable"))

	database.Connect(dbString)
	database.Migrate()
}

func setup_azure(app *App) {
	// creates a new file watcher for App_offline.htm
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("ERROR", err)
	}
	defer watcher.Close()

	// watch for App_offline.htm and exit the program if present
	// This allows continuous deployment on App Service as the .exe will not be
	// terminated otherwise
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if strings.HasSuffix(event.Name, "app_offline.htm") {
					fmt.Println("Exiting due to app_offline.htm being present")
					os.Exit(0)
				}
			}
		}
	}()

	// get the current working directory and watch it
	currentDir, err := os.Getwd()
	if err := watcher.Add(currentDir); err != nil {
		fmt.Println("ERROR", err)
	}

	// Azure App Service sets the port as an Environment Variable
	// This can be random, so needs to be loaded at startup
	port := os.Getenv("HTTP_PLATFORM_PORT")

	// default back to 8080 for local dev
	if port == "" {
		port = "8080"
	}

	app.port = port
}
