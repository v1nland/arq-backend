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
    router.GET("/Condominios/PorID/:id/", GetCondominiosPorID)
    router.GET("/Condominios/PorUsuarioID:idusuario/", GetCondominiosPorIDusuario)

    // Conserjes queries
    router.GET("/Conserjes/", GetConserjes)
    router.GET("/Conserjes/:idcondominio/", GetConserjesIdcond)

    // Departamentos queries
    router.GET("/Departamentos/", GetDepartamentos)
    router.GET("/Departamentos/GetLogin/:codigo/:numero/:password/", GetDptoLogin)
    router.GET("/Departamentos/PorCondominio/:idcondominio/", GetDptoCondominio)
    router.GET("/Departamentos/Deptoporid/:iddpto/", GetDptoID)
    router.GET("/Departamentos/DeptoEstBod/:iddpto/", GetDepartamentosEstBod)

    // Espacios comunes queries
    router.GET("/EspaciosComunes/", GetEspacios_comunes)
    router.GET("/EspaciosComunes/GetPorCond/:idcond/", GetEspacios_comunesIdcond)
    router.GET("/EspaciosComunes/UpdateEspacio/:estado/:idespacio/",GetUpdateEspacios_comunesIdcond)

    // Estacionamientos queries
    router.GET("/Estacionamientos/", GetEstacionamientos)
    router.GET("/Estacionamientos/:iddepartamento", GetEstacionamientosIddpto)

    // Usuarios queries
    router.GET("/Usuarios/", GetUsuarios)
    router.GET("/Usuarios/Login/:rut/:password/", GetUserLogin)
    router.GET("/DatosUsuario/Decode/:token/", DecodeToken)

    // Gastos comunes queries
    router.GET("/GastosComunes/", GetGastos_comunes)
    router.GET("/GastosComunes/:iddpto", GetBalanceGastos_comunes)

    // Mediciones agua queries
    router.GET("/MedicionesAgua/", GetMediciones_agua)
    router.GET("/MedicionesAgua/MedicionesDepto/:iddepartamento/", GetMediciones_aguaIddpto)
    router.GET("/MedicionesAgua/MedicionesCond/:idcond", GetMediciones_aguaIdcond)
    router.GET("/MedicionesAgua/MedicionesFecha/:ano_inicio/:mes_inicio/:ano_final/:mes_final", GetMediciones_aguaFecha)

    // Multas queries
    router.GET("/Multas/", GetMultas)
    router.GET("/Multas/ObtenerMultas/:ano_inicio/:mes_inicio/:ano_final/:mes_final/:iddpto", GetMultasFecha)
    router.GET("/Multas/ContarMultas/:ano_inicio/:mes_inicio/:ano_final/:mes_final/:iddpto", GetCountMultasFecha)

    //Tickets queries
    router.GET("/Tickets/", GetTickets)
    router.GET("/Tickets/TicketsDepto/:iddepartamento/", GetTicketsIdcond)
    router.GET("/Tickets/TicketsCond/:idcondominio/", GetTicketsIdcondFinal)
    router.GET("/Tickets/Finalizar/:idtic/", GetUpdateTicketsFinal)
    router.GET("/Tickets/Responder/:idtic/:respuesta", GetResponderTickets)
    router.GET("/Tickets/Insertar/:iddepartamentos/:idcondominio/:consulta/:asunto/", GetInsertarTickets)

    // Pagos gastos comunes queries
    router.GET("/Pagos_gastos_comunes/", GetPagos)
    router.GET("/Pagos_gastos_comunes/PagosDpto/:iddepartamento/", GetPagosIddpto)
    router.GET("/Pagos_gastos_comunes/PagosCond/:idcondominio/", GetPagosIdcondominio)
    router.GET("/Pagos_gastos_comunes/PagosCondMes/:idcondominio/:anoinicio/:mesinicio/:anofinal/:mesfinal/", GetPagosIdcondominiomes)

    // Run API
    router.Run()
}
