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
    router.GET("/Bodegas/:id_dpto/", GetBodegasByDptoID)

    // Condominios queries
    router.GET("/Condominios/", GetCondominios)
    router.GET("/Condominios/byID/:id/", GetCondominiosByID)
    router.GET("/Condominios/byUsuarioID/:id_usr/", GetCondominiosByUserID)

    // Conserjes queries
    router.GET("/Conserjes/", GetConserjes)
    router.GET("/Conserjes/:id_cond/", GetConserjesByCondID)

    // Departamentos queries
    router.GET("/Departamentos/", GetDepartamentos)
    router.GET("/Departamentos/Login/:codigo/:numero/:password/", GetDptoLogin)
    router.GET("/Departamentos/byCondominioID/:id_cond/", GetDptoByCondominioID)
    router.GET("/Departamentos/byID/:id_dpto/", GetDptoByID)
    router.GET("/Departamentos/AllData/byID/:id_dpto/", GetDepartamentosAllDataByID)

    // Espacios comunes queries
    router.GET("/EspaciosComunes/", GetEspaciosComunes)
    router.GET("/EspaciosComunes/byCondominioID/:id_cond/", GetEspaciosComunesByCondID)
    router.GET("/EspaciosComunes/Update/:estado/:id_ec/", GetUpdateEspaciosComunesByID)

    // Estacionamientos queries
    router.GET("/Estacionamientos/", GetEstacionamientos)
    router.GET("/Estacionamientos/:id_dpto/", GetEstacionamientosByDptoID)

    // Gastos comunes queries
    router.GET("/GastosComunes/", GetGastosComunes)

    // needs fix
    router.GET("/GastosComunes/Balance/:id_dpto/", GetBalanceGastosComunes)

    // Mediciones agua queries
    router.GET("/MedicionesAgua/", GetMedicionesAgua)
    router.GET("/MedicionesAgua/byDptoID/:id_dpto/", GetMedicionesAguaByDptoID)
    router.GET("/MedicionesAgua/byCondominioID/:id_cond/", GetMedicionesAguaByCondID)

    // needs fix
    router.GET("/MedicionesAgua/MedicionesFecha/:ano_inicio/:mes_inicio/:ano_final/:mes_final/", GetMedicionesAguaByFecha)

    // Multas queries
    router.GET("/Multas/", GetMultas)

    // needs fix
    router.GET("/Multas/ObtenerMultas/:ano_inicio/:mes_inicio/:ano_final/:mes_final/:id_dpto/", GetMultasByFecha)
    router.GET("/Multas/ContarMultas/:ano_inicio/:mes_inicio/:ano_final/:mes_final/:iddpto/", GetCountMultasByFecha)

    // Pagos gastos comunes queries
    router.GET("/PagosGC/", GetPagos)
    router.GET("/PagosGC/byDptoID/:id_dpto/", GetPagosByDptoID)
    router.GET("/PagosGC/byCondominioID/:id_cond/", GetPagosByCondominioID)

    // needs fix
    router.GET("/PagosGC/PagosCondMes/:id_cond/:anoinicio/:mesinicio/:anofinal/:mesfinal/", GetPagosByFechaAndCondominioID)

    //Tickets queries
    router.GET("/Tickets/", GetTickets)
    router.GET("/Tickets/byCondID/:id_cond/", GetTicketsByCondID)
    router.GET("/Tickets/byDptoID/:id_dpto/", GetTicketsByDptoID)
    router.GET("/Tickets/FinalizadosByCondID/:id_cond/", GetTicketsFinalizadosByCondID)
    router.GET("/Tickets/Finalizar/:id/", GetEndTicketByID)
    router.GET("/Tickets/Responder/:id/:respuesta/", GetResponderTicketByID) // falta que marque el ticket como finalizado una vez respondido
    router.GET("/Tickets/Insertar/:id_dpto/:id_cond/:consulta/:asunto/", GetInsertarTicket)

    // Tokens queries
    router.GET("/DatosUsuario/Decode/:token/", DecodeToken)

    // Usuarios queries
    router.GET("/Usuarios/", GetUsuarios)
    router.GET("/Usuarios/Login/:rut/:password/", GetUserLogin)

    // Run API
    router.Run()
}
