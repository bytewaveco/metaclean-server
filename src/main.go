package main

import (
	"flag"
	"fmt"

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

	if *release {
		gin.SetMode(gin.ReleaseMode)
	}

	server := gin.Default()
	connection := fmt.Sprintf("%s:%d", *host, *port)

	// API definition
	server.POST(fmt.Sprintf("%s/m/files", META_S_API), api.UploadFiles)

	// Start the server
	fmt.Printf("Meta server running on http://%s\n", connection)
	server.Run(connection)
}
