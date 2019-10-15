package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
)

func (mu Multa) FetchMultas() (multas []Multa, err error) {
    // Opens DB
    db := GetConnection()

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

func GetMultas(c *gin.Context){
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
