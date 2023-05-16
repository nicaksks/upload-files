package backend

import (
	"cdn/backend/routes"
	"cdn/backend/utils"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	port = utils.Port()
)

func Router() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.LoadHTMLGlob("./frontend/pages/*.html")
	router.Static("/static", "./frontend")

	//Upload
	router.Static("/videos", "./files/videos")
	router.Static("/images", "./files/images")
	router.Static("/gifs", "./files/gifs")
	router.Static("/audios", "./files/audios")
	router.Static("/misc", "./files/misc")

	router.GET("/", routes.Index)
	router.POST("/files", routes.Files)
	return router
}

func Start() {
	server := &http.Server{
		Addr:    port,
		Handler: Router(),
	}
	fmt.Printf("Server online! using port %s\n", strings.Replace(port, ":", "", 1))
	if err := server.ListenAndServe(); err != nil {
		log.Printf("Error: %s", err)
	}
}
