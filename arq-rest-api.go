package main
import (
    "time"
	"github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
)

func StartServer() {
    router := gin.Default() // Initialize API

    router.Use(cors.New(cors.Config{
		// AllowOrigins:     []string{"http://192.168.0.14", "http://localhost"},
		AllowMethods:     []string{"GET", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
        // AllowAllOrigins:  true,
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

    // Test routes
    router.GET("/Test/Database", TestDatabaseConnection)
    router.GET("/Test/URL/:id", GETUrlFunction)

    // MySQL routes
    router.GET("/", HomePage) // Set routes

    // Condominios queries
    router.GET("/Condominios", GETCondominios)
    router.GET("/Condominios/:id", GETCondominiosPorID)

    router.GET("/Bodegas", GETBodegas)
    router.GET("/Departamentos", GETDepartamentos)
    router.GET("/Estacionamientos", GETEstacionamientos)
    router.GET("/Multas", GETMultas)

    // Run API
    router.Run()
}
