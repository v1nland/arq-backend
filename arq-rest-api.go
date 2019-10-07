package main
import (
    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
)

func StartServer() {
    router := gin.Default() // Initialize API

    // Test routes
    router.GET("/Test/Database", TestDatabaseConnection)

    // MySQL routes
    router.GET("/", HomePage) // Set routes
    router.GET("/Condominios", GETCondominios)
    router.GET("/Bodegas", GETBodegas)
    router.GET("/Departamentos", GETDepartamentos)
    router.GET("/Estacionamientos", GETEstacionamientos)
    router.GET("/Multas", GETMultas)

    // Run API
    router.Run()
}
