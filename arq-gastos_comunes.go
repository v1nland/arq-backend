package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
)

//Take all data----------------------------------------------------------------------------------

func (Gc Gasto_comun) FetchGastos_comunes() (gastos_comunes []Gasto_comun, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("SELECT * FROM gastos_comunes")
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var gac Gasto_comun
        rows.Scan(&gac.Id, &gac.Monto, &gac.Detalle, &gac.Fecha, &gac.Id_departamentos)
        fmt.Println(gac.Id, gac.Monto, gac.Detalle, gac.Fecha, gac.Id_departamentos)
        gastos_comunes = append(gastos_comunes, gac)
	}
	defer rows.Close()

	return
}

func GetGastos_comunes(c *gin.Context){
    // Results container
    gc := Gasto_comun{}
    // Fetch from database
	gastos_comunes, err := gc.FetchGastos_comunes()
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": gastos_comunes,
		"count": len(gastos_comunes),
	})
}

