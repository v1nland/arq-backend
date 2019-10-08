package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func (est Estacionamiento) FetchEstacionamientos() (estacionamientos []Estacionamiento, err error) {
    // Opens DB
    db, err := sql.Open("mysql", "b5ef0f5f54c166:fc7bd410@tcp(us-cdbr-iron-east-02.cleardb.net:3306)/heroku_eb1b634cb7a7bd8")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    // SQL query
	rows, err := db.Query("SELECT id, numero, num_dpto FROM estacionamientos")
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var esta Estacionamiento
        rows.Scan(&esta.Id, &esta.Numero, &esta.Num_dpto)
        fmt.Println(esta.Id, esta.Numero, esta.Num_dpto)
        estacionamientos = append(estacionamientos, esta)
	}
	defer rows.Close()

	return
}

func GETCondominios(c *gin.Context){
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
