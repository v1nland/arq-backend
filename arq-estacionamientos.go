package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
)

//Take all data-------------------------------------------------------------------------------------

func (est Estacionamientos) FetchEstacionamientos() (estacionamientos []Estacionamiento, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("SELECT id, numero, id_departamentos, estado, patente_frec FROM estacionamientos")
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var esta Estacionamientos
        rows.Scan(&esta.Id, &esta.Numero, &esta.id_departamentos, &esta.estado, &esta.patente_frec)
        fmt.Println(esta.Id, esta.Numero, esta.id_departamentos, esta.estado, esta.patente_frec)
        estacionamientos = append(estacionamientos, esta)
	}
	defer rows.Close()

	return
}

func GetEstacionamientos(c *gin.Context){
    // Results container
    est := Estacionamiento{}
    // Fetch from database
	estacionamientos, err := pag.FetchEstacionamientos()
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": estacionamientos,
		"count": len(estacionamientos),
	})
}


//Select con id_departamentos------------------------------------------------------------------------

func (est Estacionamientos) FetchEstacionamientosIddpto(id_departamentos int) (estacionamientos []Estacionamiento, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select * from estacionamientos where id_departamentos=?", id_dpto)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var esta Estacionamientos
        rows.Scan(&esta.Id, &esta.Numero, &esta.id_departamentos, &esta.estado, &esta.patente_frec)
        fmt.Println(esta.Id, esta.Numero, esta.id_departamentos, esta.estado, esta.patente_frec)
        estacionamientos = append(estacionamientos, esta)
	}
	defer rows.Close()

	return
}

func GetEstacionamientosIddpto(c *gin.Context){
    // URL parameters
    var iddepartamento = c.Param("iddepartamento")

    fmt.Println(iddepartamento);
    // Results container
    est := Estacionamiento{}
    // Fetch from database
	estacionamientos, err := est.FetchEstacionamientosIddpto(iddepartamento)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": estacionamientos,
		"count": len(estacionamientos),
	})
}
