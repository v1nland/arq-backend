package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func (co Condominio) FetchCondominiosPorID(condominio_id string) (condominios []Condominio, err error) {
    // Opens DB
    db, err := sql.Open("mysql", "b5ef0f5f54c166:fc7bd410@tcp(us-cdbr-iron-east-02.cleardb.net:3306)/heroku_eb1b634cb7a7bd8")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    // SQL query
	rows, err := db.Query("SELECT id, nombre, ubicacion FROM condominios WHERE id = ?", condominio_id)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var cond Condominio
        rows.Scan(&cond.Id, &cond.Nombre, &cond.Ubicacion)
        fmt.Println(cond.Id, cond.Nombre, cond.Ubicacion)
        condominios = append(condominios, cond)
	}
	defer rows.Close()

	return
}

func GETCondominiosPorID(c *gin.Context){
    var id = c.Param("id")

    // Results container
    co := Condominio{}
    // Fetch from database
	condominios, err := co.FetchCondominiosPorID(id)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": condominios,
		"count": len(condominios),
	})
}
