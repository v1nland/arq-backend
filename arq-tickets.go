package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
)

// // // // TAKE ALL TICKETS // // // //
func (Ti Ticket) FetchTickets() (tickets []Ticket, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("SELECT * FROM tickets")
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var tic Ticket
        rows.Scan(&tic.Id, &tic.Id_departamentos, &tic.Id_usuarios, &tic.Consulta, &tic.Respuesta, &tic.Finalizado, &tic.Asunto, &tic.Fecha)
        fmt.Println(tic.Id, tic.Id_departamentos, tic.Id_usuarios, tic.Consulta, tic.Respuesta, tic.Finalizado, tic.Asunto, tic.Fecha)
        tickets = append(tickets, tic)
	}
	defer rows.Close()

	return
}

func GetTickets(c *gin.Context){
    // Results container
    ti := Ticket{}
    // Fetch from database
	tickets, err := ti.FetchTickets()
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": tickets,
		"count": len(tickets),
	})
}

// // // // TAKE ALL TICKETS WITH ID_COND=id_cond // // // //
func (Ti Ticket) FetchTicketsByCondID(id_cond string) (tickets []Ticket, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select * from tickets where id_departamentos in (select id from departamentos where id_condominio=?)", id_cond)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var tic Ticket
        rows.Scan(&tic.Id, &tic.Id_departamentos, &tic.Id_usuarios, &tic.Consulta, &tic.Respuesta, &tic.Finalizado, &tic.Asunto, &tic.Fecha)
        fmt.Println(tic.Id, tic.Id_departamentos, tic.Id_usuarios, tic.Consulta, tic.Respuesta, tic.Finalizado, tic.Asunto, tic.Fecha)
        tickets = append(tickets, tic)
	}
	defer rows.Close()

	return
}

func GetTicketsByCondID(c *gin.Context){
    // URL parameters
    var id_cond = c.Param("id_cond")

    fmt.Println(id_cond);

    // Results container
    ti := Ticket{}
    // Fetch from database
	tickets, err := ti.FetchTicketsByCondID(id_cond)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": tickets,
		"count": len(tickets),
	})
}

// // // // TAKE ALL TICKETS WITH ID_COND=id_cond AND finalizado=1 // // // //
func (Ti Ticket) FetchTicketsFinalizadosByCondID(id_cond string) (tickets []Ticket, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select * from tickets where id_departamentos in (select id from departamentos where id_condominio =?) and finalizado = 1", id_cond)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var tic Ticket
        rows.Scan(&tic.Id, &tic.Id_departamentos, &tic.Id_usuarios, &tic.Consulta, &tic.Respuesta, &tic.Finalizado, &tic.Asunto, &tic.Fecha)
        fmt.Println(tic.Id, tic.Id_departamentos, tic.Id_usuarios, tic.Consulta, tic.Respuesta, tic.Finalizado, tic.Asunto, tic.Fecha)
        tickets = append(tickets, tic)
	}
	defer rows.Close()

	return
}

func GetTicketsFinalizadosByCondID(c *gin.Context){
    // URL parameters
    var id_cond = c.Param("id_cond")

    fmt.Println(id_cond);

    // Results container
    ti := Ticket{}
    // Fetch from database
	tickets, err := ti.FetchTicketsFinalizadosByCondID(id_cond)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": tickets,
		"count": len(tickets),
	})
}

// // // // END TICKET WITH ID=id // // // //
func (Ti Ticket) FetchEndTicketByID(id string) (tickets []Ticket, err error) {
    // Opens DB
    db := GetConnection()

    // SQL que
	rows, err := db.Query("update tickets set finalizado= 1 where id=?", id)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var tic Ticket
        rows.Scan(&tic.Id, &tic.Id_departamentos, &tic.Id_usuarios, &tic.Consulta, &tic.Respuesta, &tic.Finalizado, &tic.Asunto, &tic.Fecha)
        fmt.Println(tic.Id, tic.Id_departamentos, tic.Id_usuarios, tic.Consulta, tic.Respuesta, tic.Finalizado, tic.Asunto, tic.Fecha)
        tickets = append(tickets, tic)
	}
	defer rows.Close()

	return
}

func GetEndTicketByID(c *gin.Context){
    // URL parameters
    var id = c.Param("id")

    fmt.Println(id);

    // Results container
    ti := Ticket{}
    // Fetch from database
	tickets, err := ti.FetchEndTicketByID(id)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": tickets,
		"count": len(tickets),
	})
}

// // // // RESPONSE TICKET WITH ID=id // // // //
func (Ti Ticket) FetchResponderTicketByID(id string, respuesta string) (tickets []Ticket, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("update tickets set respuesta = ?, finalizado = 1 where id = ?", respuesta, id)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var tic Ticket
        rows.Scan(&tic.Id, &tic.Id_departamentos, &tic.Id_usuarios, &tic.Consulta, &tic.Respuesta, &tic.Finalizado, &tic.Asunto, &tic.Fecha)
        fmt.Println(tic.Id, tic.Id_departamentos, tic.Id_usuarios, tic.Consulta, tic.Respuesta, tic.Finalizado, tic.Asunto, tic.Fecha)
        tickets = append(tickets, tic)
	}
	defer rows.Close()

	return
}

func GetResponderTicketByID(c *gin.Context){
    // URL parameters
    var id = c.Param("id")
    var respuesta = c.Param("respuesta")

    fmt.Println(id);
    fmt.Println(respuesta);

    // Results container
    ti := Ticket{}
    // Fetch from database
	tickets, err := ti.FetchResponderTicketByID(id, respuesta)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": tickets,
		"count": len(tickets),
	})
}

// // // // INSERT NEW TICKET // // // //
func (Ti Ticket) FetchInsertarTicket(id_dpto string, id_cond string, consulta string, asunto string) (tickets []Ticket, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("insert into tickets (id_departamentos, id_usuarios, consulta, respuesta, finalizado, asunto, fecha) values (?, ?, ?, 0, 0, ?, current_timestamp)", id_dpto, id_cond, consulta, asunto)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var tic Ticket
        rows.Scan(&tic.Id, &tic.Id_departamentos, &tic.Id_usuarios, &tic.Consulta, &tic.Respuesta, &tic.Finalizado, &tic.Asunto, &tic.Fecha)
        fmt.Println(tic.Id, tic.Id_departamentos, tic.Id_usuarios, tic.Consulta, tic.Respuesta, tic.Finalizado, tic.Asunto, tic.Fecha)
        tickets = append(tickets, tic)
	}
	defer rows.Close()

	return
}

func GetInsertarTicket(c *gin.Context){
    // URL parameters
    var id_dpto = c.Param("id_dpto")
    var id_cond = c.Param("id_cond")
    var consulta = c.Param("consulta")
    var asunto = c.Param("asunto")

    fmt.Println(id_dpto);
    fmt.Println(id_cond);
    fmt.Println(consulta);
    fmt.Println(asunto);

    // Results container
    ti := Ticket{}
    // Fetch from database
	tickets, err := ti.FetchInsertarTicket(id_dpto, id_cond, consulta, asunto)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": tickets,
		"count": len(tickets),
	})
}

// // // TAKE ALL TICKETS WITH ID_DPTO=id_dpto // // //

func (Ti Ticket) FetchTicketsByDptoID(id_dpto string) (tickets []Ticket, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select * from tickets where id_departamentos = ?", id_dpto)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var tic Ticket
        rows.Scan(&tic.Id, &tic.Id_departamentos, &tic.Id_usuarios, &tic.Consulta, &tic.Respuesta, &tic.Finalizado, &tic.Asunto, &tic.Fecha)
        fmt.Println(tic.Id, tic.Id_departamentos, tic.Id_usuarios, tic.Consulta, tic.Respuesta, tic.Finalizado, tic.Asunto, tic.Fecha)
        tickets = append(tickets, tic)
	}
	defer rows.Close()

	return
}

func GetTicketsByDptoID(c *gin.Context){
    // URL parameters
    var id_dpto = c.Param("id_dpto")

    fmt.Println(id_dpto);

    // Results container
    ti := Ticket{}
    // Fetch from database
	tickets, err := ti.FetchTicketsByCondID(id_dpto)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": tickets,
		"count": len(tickets),
	})
}

// // // COUNT PENDING TICKETS // // //

func (Ti Ticket) CountPendingTickets() (tickets []Ticket, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select * from tickets where finalizado = 0")
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var tic Ticket
        rows.Scan(&tic.Id, &tic.Id_departamentos, &tic.Id_usuarios, &tic.Consulta, &tic.Respuesta, &tic.Finalizado, &tic.Asunto, &tic.Fecha)
        fmt.Println(tic.Id, tic.Id_departamentos, tic.Id_usuarios, tic.Consulta, tic.Respuesta, tic.Finalizado, tic.Asunto, tic.Fecha)
        tickets = append(tickets, tic)
	}
	defer rows.Close()

	return
}

func GetCountPendingTickets(c *gin.Context){
    // URL parameters

    // Results container
    ti := Ticket{}
    // Fetch from database
	tickets, err := ti.CountPendingTickets()
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"count": len(tickets),
	})
}
