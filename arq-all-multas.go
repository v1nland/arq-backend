package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func (mu Multa) FetchMultas() (multas []Multa, err error) {
    // Opens DB
    db, err := sql.Open("mysql", "b5ef0f5f54c166:fc7bd410@tcp(us-cdbr-iron-east-02.cleardb.net:3306)/heroku_eb1b634cb7a7bd8")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    // SQL query
	rows, err := db.Query("SELECT id, tipo, num_dpto, precio, fecha FROM multas")
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var mul Multa
        rows.Scan(&mul.Id, &mul.Tipo, &mul.Num_dpto, &mul.Precio, &mul.Fecha)
        fmt.Println(mul.Id, mul.Tipo, mul.Num_dpto, mul.Precio, mul.Fecha)
        multas = append(multas, mul)
	}
	defer rows.Close()

	return
}

func GETCondominios(c *gin.Context){
    // Results container
    mu := Multa{}
    // Fetch from database
	multas, err := mu.FetchMultas()
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": multas,
		"count": len(multas),
	})
}
