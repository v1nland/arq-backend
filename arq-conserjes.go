package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
)

//Take all data----------------------------------------------------------------------------------

func (Cs Conserje) FetchConserjes() (conserjes []Conserje, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("SELECT * FROM conserjes")
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var cos Conserje
        rows.Scan(&cos.Id, &cos.Nombre, &cos.Rut, &cos.id_condominio)
        fmt.Println(cos.Id, cos.Nombre, cos.Rut, cos.id_condominio)
        conserjes = append(conserjes, cos)
	}
	defer rows.Close()

	return
}

func GetConserjes(c *gin.Context){
    // Results container
    cs := Conserje{}
    // Fetch from database
	conserjes, err := ec.FetchConserjes()
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": conserjes,
		"count": len(conserjes),
	})
}

//select conserjes por condominio--------------------------------------------------------------------

func (Cs Conserje) FetchConserjesIdcond(id_condominio int) (conserjes []Conserje, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select * from conserjes where id_condominio=?", id_condominio)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var cos Conserje
        rows.Scan(&cos.Id, &cos.Nombre, &cos.Rut, &cos.id_condominio)
        fmt.Println(cos.Id, cos.Nombre, cos.Rut, cos.id_condominio)
        conserjes = append(conserjes, cos)
	}
	defer rows.Close()

	return
}

func GetConserjesIdcond(c *gin.Context){
    // URL parameters
    var idcondominio = c.Param("idcondominio")

    fmt.Println(idcondominio);
    // Results container
    cs := Conserje{}
    // Fetch from database
	conserjes, err := ec.FetchConserjesIdcond(idcondominio)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": conserjes,
		"count": len(conserjes),
	})
}
