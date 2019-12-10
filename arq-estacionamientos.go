package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
)

//Take all data-------------------------------------------------------------------------------------

func (est Estacionamiento) FetchEstacionamientos() (estacionamientos []Estacionamiento, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("SELECT id, numero, id_departamentos, estado, patente_frec FROM estacionamientos")
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var esta Estacionamiento
        rows.Scan(&esta.Id, &esta.Numero, &esta.Id_departamentos, &esta.Estado, &esta.Patente_frecuente)
        fmt.Println(esta.Id, esta.Numero, esta.Id_departamentos, esta.Estado, esta.Patente_frecuente)
        estacionamientos = append(estacionamientos, esta)
	}
	defer rows.Close()

	return
}

func GetEstacionamientos(c *gin.Context){
    // Results container
    est := Estacionamiento{}
    // Fetch from database
	estacionamientos, err := est.FetchEstacionamientos()
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

func (est Estacionamiento) FetchEstacionamientosIddpto(id_departamentos string) (estacionamientos []Estacionamiento, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select * from estacionamientos where id_departamentos=?", id_departamentos)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var esta Estacionamiento
        rows.Scan(&esta.Id, &esta.Numero, &esta.Id_departamentos, &esta.Estado, &esta.Patente_frecuente)
        fmt.Println(esta.Id, esta.Numero, esta.Id_departamentos, esta.Estado, esta.Patente_frecuente)
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
