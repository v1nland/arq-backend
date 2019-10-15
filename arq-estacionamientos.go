package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
)

func (est Estacionamiento) FetchEstacionamientos() (estacionamientos []Estacionamiento, err error) {
    // Opens DB
    db := GetConnection()

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
