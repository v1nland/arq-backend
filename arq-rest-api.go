package main
import (
    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
)

func StartServer() {
    router := gin.Default() // Initialize API

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
