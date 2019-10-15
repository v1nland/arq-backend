package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
)

func (de Departamento) FetchDepartamentos() (departamentos []Departamento, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("SELECT id, numero, id_condominio, dueño, residente, telefono, correo FROM departamentos")
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var dpto Departamento
        rows.Scan(&dpto.Id, &dpto.Numero, &dpto.Id_condominio, &dpto.Nombre_dueno, &dpto.Nombre_arrendatario)
        fmt.Println(dpto.Id, dpto.Numero, dpto.Id_condominio, dpto.Nombre_dueno, dpto.Nombre_arrendatario)
        departamentos = append(departamentos, dpto)
	}
	defer rows.Close()

	return
}

func GetDepartamentos(c *gin.Context){
    // Results container
    de := Departamento{}
    // Fetch from database
	departamentos, err := de.FetchDepartamentos()
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": departamentos,
		"count": len(departamentos),
	})
}
