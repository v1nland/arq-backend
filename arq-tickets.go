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
        rows.Scan(&tic.Id, &tic.Id_departamentos, &tic.Id_usuarios, &tic.Consulta, &tic.Respuesta, &tic.Finalizado)
        fmt.Println(tic.Id, tic.Id_departamentos, tic.Id_usuarios, tic.Consulta, tic.Respuesta, tic.Finalizado)
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

func (Ti Ticket) FetchTicketsIdcond() (id_cond int) (tickets []Ticket, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select * from tickets where id_departamentos = (select id_departamentos from departamentos where id_condominio=?)", id_cond)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var tic Ticket
        rows.Scan(&tic.Id, &tic.Id_departamentos, &tic.Id_usuarios, &tic.Consulta, &tic.Respuesta, &tic.Finalizado)
        fmt.Println(tic.Id, tic.Id_departamentos, tic.Id_usuarios, tic.Consulta, tic.Respuesta, tic.Finalizado)
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
	tickets, err := ti.FetchTickets(idcondominio)
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

func (Ti Ticket) FetchTicketsIdcondFinal() (id_cond int) (tickets []Ticket, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select * from tickets where id_departamentos = (select id_departamentos from departamentos where id_condominio=?) and finalizado = 1", id_cond)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var tic Ticket
        rows.Scan(&tic.Id, &tic.Id_departamentos, &tic.Id_usuarios, &tic.Consulta, &tic.Respuesta, &tic.Finalizado)
        fmt.Println(tic.Id, tic.Id_departamentos, tic.Id_usuarios, tic.Consulta, tic.Respuesta, tic.Finalizado)
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
	tickets, err := ti.FetchTickets(idcondominio)
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

func (Ti Ticket) UpdateTicketsFinal() (id_cond int) (tickets []Ticket, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("update tickets set finalizado= 1 where id=?", id_tic)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var tic Ticket
        rows.Scan(&tic.Id, &tic.Id_departamentos, &tic.Id_usuarios, &tic.Consulta, &tic.Respuesta, &tic.Finalizado)
        fmt.Println(tic.Id, tic.Id_departamentos, tic.Id_usuarios, tic.Consulta, tic.Respuesta, tic.Finalizado)
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
	tickets, err := ti.FetchTickets(idtic)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": tickets,
		"count": len(tickets),
	})
}
