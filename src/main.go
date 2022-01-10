package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"meta-server/api"
)

const META_S_VERSION = 1

var META_S_API = fmt.Sprintf("api/v%d", META_S_VERSION)

func main() {
	release := flag.Bool("release", false, "Run the server in release mode.")
	host := flag.String("h", "localhost", "Server hostname.")
	port := flag.Int("p", 3333, "Server port.")
	flag.Parse()

	cleanupError := os.RemoveAll("/tmp/MetaClean")

	if cleanupError != nil {
		panic(cleanupError)
	}

	if *release {
		gin.SetMode(gin.ReleaseMode)
	}

	server := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3333", "https://metaclean.pro"}
	corsConfig.AllowMethods = []string{"GET", "POST"}
	server.Use(cors.New(corsConfig))
	connection := fmt.Sprintf("%s:%d", *host, *port)

	// API definition
	server.GET(fmt.Sprintf("%s/status", META_S_API), api.Status)
	server.GET(fmt.Sprintf("%s/m/files", META_S_API), api.DownloadFiles)
	server.POST(fmt.Sprintf("%s/m/files", META_S_API), api.UploadFiles)
	// For AWS container checks, send Okay
	server.GET("/", api.Status)

	// Start the server
	fmt.Printf("Meta server running on http://%s\n", connection)
	server.Run(connection)
}
