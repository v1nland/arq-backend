package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func (co Departamento) FetchDepartamentos() (departamentos []Departamento, err error) {
    // Opens DB
    db, err := sql.Open("mysql", "b5ef0f5f54c166:fc7bd410@tcp(us-cdbr-iron-east-02.cleardb.net:3306)/heroku_eb1b634cb7a7bd8")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    // SQL query
	rows, err := db.Query("SELECT id, numero, id_condominio, nombre_dueño, nombre_arrendatario FROM departamentos")
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var dpto Departamento
        rows.Scan(&dpto.Id, &dpto.Numero, &dpto.Id_condominio, &dpto.Nombre_dueño, &dpto.Nombre_arrendatario)
        fmt.Println(dpto.Id, dpto.Numero, dpto.Id_condominio, dpto.Nombre_dueño, dpto.Nombre_arrendatario)
        departamentos = append(departamentos, dpto)
	}
	defer rows.Close()

	return
}

func GETCondominios(c *gin.Context){
    // Results container
    co := Departamento{}
    // Fetch from database
	departamentos, err := co.FetchDepartamentos()
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": departamentos,
		"count": len(departamentos),
	})
}
