package main
import (
    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func StartServer() {
    // Initialize API
    router := gin.Default()

    // Use CORS
    router.Use(CORSMiddleware())

    // Test routes
    router.GET("/Test/Database", TestDatabaseConnection)
    router.GET("/Test/JWT", jwttest)
    router.GET("/Test/JWT/:token", VerifyToken)
    router.GET("/Test/URL/:id/:hola", TestUrlParameters)

    // MySQL routes
    router.GET("/", HomePage) // Set routes

    // Condominios queries
    router.GET("/Condominios/", GetCondominios)
    router.GET("/Condominios/:id", GetCondominiosPorID)

    router.GET("/Usuarios/", GetUsuarios)
    router.GET("/Usuarios/:rut/:password", GetUserLogin)

    router.GET("/Bodegas/", GetBodegas)

    router.GET("/Departamentos/", GetDepartamentos)
    router.GET("/Departamentos/:codigo/:numero/:password", GetDptoLogin)

    router.GET("/Estacionamientos/", GetEstacionamientos)
    router.GET("/Multas/", GetMultas)

    // Run API
    router.Run()
}
