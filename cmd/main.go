package main

import (
	"os"
	"post-api/api/routes"
	"post-api/pkg"

	"github.com/gin-gonic/gin"

	_ "github.com/joho/godotenv/autoload"
)

// init() will run first
func init() {
	modeGin := "debug"
	if os.Getenv("GIN_MODE") != "" {
		modeGin = os.Getenv("GIN_MODE")
	}
	gin.SetMode(modeGin)
}

func main() {
	r := gin.Default()
	dbGorm := pkg.Connect()

	routes.PostRoute(r, dbGorm)

	port := ":8080"
	if os.Getenv("APP_PORT") != "" {
		port = os.Getenv("APP_PORT")
	}
	r.Run(port)
}

// docker run --name post-posgres -p 5432:5432 -e POSTGRES_PASSWORD=yourpassword -d postgres
