package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
)

//Take all data----------------------------------------------------------------------------------

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
        rows.Scan(&tic.Id, &tic.Id_departamentos, &tic.Id_condominio, &tic.Consulta, &tic.Respuesta, &tic.Finalizado, &tic.Asunto)
        fmt.Println(tic.Id, tic.Id_departamentos, tic.Id_condominio, tic.Consulta, tic.Respuesta, tic.Finalizado, tic.Asunto)
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

//Select con id_condominio-------------------------------------------------------------------------

func (Ti Ticket) FetchTicketsIdcond(id_cond string) (tickets []Ticket, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select * from tickets where id_condominio=?)", id_cond)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var tic Ticket
        rows.Scan(&tic.Id, &tic.Id_departamentos, &tic.Id_condominio, &tic.Consulta, &tic.Respuesta, &tic.Finalizado, &tic.Asunto)
        fmt.Println(tic.Id, tic.Id_departamentos, tic.Id_condominio, tic.Consulta, tic.Respuesta, tic.Finalizado, tic.Asunto)
        tickets = append(tickets, tic)
	}
	defer rows.Close()

	return
}

func GetTicketsIdcond(c *gin.Context){
    // URL parameters
    var idcondominio = c.Param("idcondominio")

    fmt.Println(idcondominio);

    // Results container
    ti := Ticket{}
    // Fetch from database
	tickets, err := ti.FetchTicketsIdcond(idcondominio)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": tickets,
		"count": len(tickets),
	})
}

//Select con id_ticket y finalizados---------------------------------------------------------------

func (Ti Ticket) FetchTicketsIdcondFinal(id_cond string) (tickets []Ticket, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select * from tickets where id_condominio=? and finalizado = 1", id_cond)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var tic Ticket
        rows.Scan(&tic.Id, &tic.Id_departamentos, &tic.Id_condominio, &tic.Consulta, &tic.Respuesta, &tic.Finalizado, &tic.Asunto)
        fmt.Println(tic.Id, tic.Id_departamentos, tic.Id_condominio, tic.Consulta, tic.Respuesta, tic.Finalizado, tic.Asunto)
        tickets = append(tickets, tic)
	}
	defer rows.Close()

	return
}

func GetTicketsIdcondFinal(c *gin.Context){
    // URL parameters
    var idcondominio = c.Param("idcondominio")

    fmt.Println(idcondominio);

    // Results container
    ti := Ticket{}
    // Fetch from database
	tickets, err := ti.FetchTicketsIdcondFinal(idcondominio)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": tickets,
		"count": len(tickets),
	})
}

//Update con id_ticket para poner finalizado=1-----------------------------------------------------

func (Ti Ticket) UpdateTicketsFinal(id_tic string) (tickets []Ticket, err error) {
    // Opens DB
    db := GetConnection()

    // SQL que
	rows, err := db.Query("update tickets set finalizado= 1 where id=?", id_tic)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var tic Ticket
        rows.Scan(&tic.Id, &tic.Id_departamentos, &tic.Id_condominio, &tic.Consulta, &tic.Respuesta, &tic.Finalizado, &tic.Asunto)
        fmt.Println(tic.Id, tic.Id_departamentos, tic.Id_condominio, tic.Consulta, tic.Respuesta, tic.Finalizado, tic.Asunto)
        tickets = append(tickets, tic)
	}
	defer rows.Close()

	return
}

func GetUpdateTicketsFinal(c *gin.Context){
    // URL parameters
    var idtic = c.Param("idtic")

    fmt.Println(idtic);

    // Results container
    ti := Ticket{}
    // Fetch from database
	tickets, err := ti.UpdateTicketsFinal(idtic)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": tickets,
		"count": len(tickets),
	})
}

//Respuesta ticket--------------------------------------------------------------------------------

func (Ti Ticket) ResponderTickets(id_ticket string, respuesta string) (tickets []Ticket, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("update tickets set respuesta = ? where id = ?", respuesta, id_ticket)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var tic Ticket
        rows.Scan(&tic.Id, &tic.Id_departamentos, &tic.Id_condominio, &tic.Consulta, &tic.Respuesta, &tic.Finalizado, &tic.Asunto)
        fmt.Println(tic.Id, tic.Id_departamentos, tic.Id_condominio, tic.Consulta, tic.Respuesta, tic.Finalizado, tic.Asunto)
        tickets = append(tickets, tic)
	}
	defer rows.Close()

	return
}

func GetResponderTickets(c *gin.Context){
    // URL parameters
    var id_ticket = c.Param("id_ticket")
    var respuesta = c.Param("respuesta")

    fmt.Println(id_ticket);
    fmt.Println(respuesta);

    // Results container
    ti := Ticket{}
    // Fetch from database
	tickets, err := ti.ResponderTickets(id_ticket, respuesta)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": tickets,
		"count": len(tickets),
	})
}

//Insertar nuevo ticket--------------------------------------------------------------------------
func (Ti Ticket) InsertarTickets(id_departamentos string, id_condominio string, consulta string, asunto string) (tickets []Ticket, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("insert into tickets (id_departamentos, id_condominio, consulta, respuesta, finalizado, asunto) values (?, ?, ?, null, 0, ?)", id_departamentos, id_condominio, consulta, asunto)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var tic Ticket
        rows.Scan(&tic.Id, &tic.Id_departamentos, &tic.Id_condominio, &tic.Consulta, &tic.Respuesta, &tic.Finalizado, &tic.Asunto)
        fmt.Println(tic.Id, tic.Id_departamentos, tic.Id_condominio, tic.Consulta, tic.Respuesta, tic.Finalizado, tic.Asunto)
        tickets = append(tickets, tic)
	}
	defer rows.Close()

	return
}

func GetInsertarTickets(c *gin.Context){
    // URL parameters
    var iddepartamentos = c.Param("iddepartamentos")
    var idcondominio = c.Param("idcondominio")
    var consulta = c.Param("consulta")
    var asunto = c.Param("asunto")

    fmt.Println(id_departamentos);
    fmt.Println(id_condominio);
    fmt.Println(consulta);
    fmt.Println(asunto);

    // Results container
    ti := Ticket{}
    // Fetch from database
	tickets, err := ti.InsertarTickets(id_departamentos, id_condominio, consulta, asunto)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": tickets,
		"count": len(tickets),
	})
}

