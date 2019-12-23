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
    router.GET("/Departamentos/UpdateDpto/:password/:dueno/:residente/:telefono/:correo/:id_condominio/:telefono_residente/:correo_residente/:prorrateo/:iddpto/", GetUpdateDpto)

    // Espacios comunes queries
    router.GET("/EspaciosComunes/", GetEspaciosComunes)
    router.GET("/EspaciosComunes/byCondominioID/:id_cond/", GetEspaciosComunesByCondID)
    router.GET("/EspaciosComunes/Update/:nombre/:estado/:descripcion/:id_ec/", GetUpdateEspaciosComunesByID)
    router.GET("/EspaciosComunes/Insertar/:nombre/:id_condominio/:estado/:descripcion/", GetInsertEspaciosComunes)
    router.GET("/EspaciosComunes/byID/:idesp/", GetEspaciosComunesByID) //new

    // Estacionamientos queries
    router.GET("/Estacionamientos/", GetEstacionamientos)
    router.GET("/Estacionamientos/:id_dpto/", GetEstacionamientosByDptoID)

    // Gastos comunes queries
    router.GET("/GastosComunes/", GetGastosComunes)
    router.GET("/GastosComunes/Insertar/:monto/:detalle/:num_dpto/:cod_cond/", GetInsertGastosComunes)
    router.GET("/GastosComunes/Update/:monto/:detalle/:fecha/:id_dpto/:id_gc/", GetUpdateGastosComunes)

    // needs fix
    router.GET("/GastosComunes/Balance/:id_dpto/", GetBalanceGastosComunes)

    // Mediciones agua queries
    router.GET("/MedicionesAgua/", GetMedicionesAgua)
    router.GET("/MedicionesAgua/byDptoID/:id_dpto/", GetMedicionesAguaByDptoID)
    router.GET("/MedicionesAgua/byCondominioID/:id_cond/", GetMedicionesAguaByCondID)
    router.GET("/MedicionesAgua/Insertar/:litros/:num_dpto/:cod_cond/", GetInsertMedicionesAgua)
    router.GET("/MedicionesAgua/MedicionesFecha/:fechai/:fechaf/", GetMedicionesAguaByFecha)
    router.GET("/MedicionesAgua/byID/:idmed/", GetMedicionesAguaByID) //new

    // Multas queries
    router.GET("/Multas/", GetMultas)
    router.GET("/Multas/Insertar/:codcond/:num_dpto/:monto/:causa/", GetInsertarMultas)
    router.GET("/Multas/ObtenerMultas/:fechai/:fechaf/:id_dpto/", GetMultasByFecha)
    router.GET("/Multas/ContarMultas/:fechai/:fechaf/:iddpto/", GetCountMultasByFecha)
    router.GET("/Multas/UpdateMultas/:grado/:id_dpto/:monto/:fecha/:causa/:idmul/", GetUpdateMultas)
    router.GET("/Multas/byID/:idmulta/", GetMultasByID)

    // Pagos gastos comunes queries
    router.GET("/PagosGC/", GetPagos)
    router.GET("/PagosGC/byDptoID/:id_dpto/", GetPagosByDptoID)
    router.GET("/PagosGC/byCondominioID/:id_cond/", GetPagosByCondominioID)
    router.GET("/PagosGC/PagosCondMes/:fechai/:fechaf/", GetPagosByFechaAndCondominioID)
    router.GET("/PagosGC/Insertar/:monto/:num_dpto/:cod_cond/", GetInsertPagos)
    router.GET("/PagosGC/Update/:monto/:fecha/:id_dpto/:idpago/", GetUpdatePagos)
    router.GET("/PagosGC/ByID/:pago_id/", GetPagosByID)

    //Tickets queries
    router.GET("/Tickets/", GetTickets)
    router.GET("/Tickets/byCondID/:id_cond/", GetTicketsByCondID)
    router.GET("/Tickets/byDptoID/:id_dpto/", GetTicketsByDptoID)
    router.GET("/Tickets/FinalizadosByCondID/:id_cond/", GetTicketsFinalizadosByCondID)
    router.GET("/Tickets/Finalizar/:id/", GetEndTicketByID)
    router.GET("/Tickets/Responder/:id/:respuesta/", GetResponderTicketByID)
    router.GET("/Tickets/Insertar/:id_dpto/:id_cond/:consulta/:asunto/", GetInsertarTicket)
    router.GET("/Tickets/CountPending/", GetCountPendingTickets)

    // Tokens queries
    router.GET("/DatosUsuario/Decode/:token/", DecodeToken)

    // Usuarios queries
    router.GET("/Usuarios/", GetUsuarios)
    router.GET("/Usuarios/Login/:rut/:password/", GetUserLogin)
    router.GET("/Usuarios/Update/:rut/:password/:nombre/:iduser", GetUpdateUsuarios)

    // Run API
    router.Run()
}
