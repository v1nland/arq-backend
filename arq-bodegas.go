package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
)

func (bo Bodega) FetchBodegas() (bodegas []Bodega, err error) {
    // Opens DB
    db := GetConnection()

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

func GetBodegas(c *gin.Context){
    // Results container
    bo := Bodega{}
    // Fetch from database
	bodegas, err := bo.FetchBodegas()
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": bodegas,
		"count": len(bodegas),
	})
}
