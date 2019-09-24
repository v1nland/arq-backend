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

    // Run API
    router.Run()
}
