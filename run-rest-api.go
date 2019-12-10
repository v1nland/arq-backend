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
    // router.GET("/Test/Database", TestDatabaseConnection)
    // router.GET("/Test/URL/:id/:hola", TestUrlParameters)
    // router.GET("/Test/JWT", jwttest)

    // MySQL routes
    router.GET("/", HomePage) // Set routes

    // Bodegas queries
    router.GET("/Bodegas/", GetBodegas)
    router.GET("/Bodegas/:iddepartamento", GetBodegasIddpto)

    // Condominios queries
    router.GET("/Condominios/", GetCondominios)
    router.GET("/Condominios/:id/", GetCondominiosPorID)

    // Conserjes queries
    router.GET("/Conserjes/", GetConserjes)
    router.GET("/Conserjes/:idcondominio/", GetConserjesIdcond)

    // Departamentos queries
    router.GET("/Departamentos/", GetDepartamentos)
    router.GET("/Departamentos/GetLogin/:codigo/:numero/:password/", GetDptoLogin)
    router.GET("/Departamentos/PorCondominio/:idcondominio/", GetDptoCondominio)
    router.GET("/Departamentos/Deptoporid/:iddpto/", GetDptoID)
    router.GET("/Departamentos/TodoDeptoporid/:iddpto/", GetTodoDptoID)

    // Espacios comunes queries
    router.GET("/Espacios_comunes/", GetEspacios_comunes)
    router.GET("/Espacios_comunes/GetPorCond/:idcond/", GetEspacios_comunesIdcond)
    router.GET("/Espacios_comunes/UpdateEspacio/:estado/:idespacio/",GetUpdateEspacios_comunesIdcond)

    // Estacionamientos queries
    router.GET("/Estacionamientos/", GetEstacionamientos)
    router.GET("/Estacionamientos/:iddepartamento", GetEstacionamientosIddpto)

    // Usuarios queries
    router.GET("/Usuarios/", GetUsuarios)
    router.GET("/Usuarios/Login/:rut/:password/", GetUserLogin)
    router.GET("/DatosUsuario/Decode/:token/", DecodeToken)

    // Gastos comunes queries
    router.GET("/Gastos_comunes/", GetGastos_comunes)

    // Mediciones agua queries
    router.GET("/Mediciones_agua/", GetMediciones_agua)
    router.GET("/Mediciones_agua/MedicionesDepto/:iddepartamento/", GetMediciones_aguaIddpto)
    router.GET("/Mediciones_agua/MedicionesCond/:idcond", GetMediciones_aguaIdcond)
    router.GET("/Mediciones_agua/MedicionesFecha/:ano_inicio/:mes_inicio/:ano_final/:mes_final", GetMediciones_aguaFecha)

    // Multas queries
    router.GET("/Multas/", GetMultas)
    router.GET("/Multas/ObtenerMultas/:ano_inicio/:mes_inicio/:ano_final/:mes_final/:iddpto", GetMultasFecha)
    router.GET("/Multas/ContarMultas/:ano_inicio/:mes_inicio/:ano_final/:mes_final/:iddpto", GetMultasFecha)

    //Tickets queries
    router.GET("/Tickets/", GetTickets)
    router.GET("/Tickets/TicketsDepto/:iddepartamento/", GetTicketsIdcond)
    router.GET("/Tickets/TicketsCond/:idcondominio/", GetTicketsIdcondFinal)
    router.GET("/Tickets/Finalizar/:idtic/", GetUpdateTicketsFinal)

    // Pagos gastos comunes queries

    // Run API
    router.Run()
}
