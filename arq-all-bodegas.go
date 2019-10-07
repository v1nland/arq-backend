package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func (co Bodega) FetchBodegas() (bodegas []Bodega, err error) {
    // Opens DB
    db, err := sql.Open("mysql", "b5ef0f5f54c166:fc7bd410@tcp(us-cdbr-iron-east-02.cleardb.net:3306)/heroku_eb1b634cb7a7bd8")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()

    // SQL query
	rows, err := db.Query("SELECT id, numero, num_dpto FROM bodegas")
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var bod Bodega
        rows.Scan(&bod.Id, &bod.Numero, &bod.Num_dpto)
        fmt.Println(bod.Id, bod.Numero, bod.Num_dpto)
        bodegas = append(bodegas, bod)
	}
	defer rows.Close()

	return
}

func GETCondominios(c *gin.Context){
    // Results container
    co := Bodega{}
    // Fetch from database
	bodegas, err := co.FetchBodegas()
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": bodegas,
		"count": len(bodegas),
	})
}
