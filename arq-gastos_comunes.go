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

//balances de gastos-----------------------------------------------------------------------------

func (Gc Gasto_comun) BalanceGastos_comunes(id_dpto string) (gastos_comunes []Gasto_comun, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("SELECT departamentos.numero, pagos_gastos_comunes.fecha, departamentos.dueno, gastos_comunes.monto as 'montogc', pagos_gastos_comunes.monto as 'pagosgc' from departamentos join pagos_gastos_comunes on departamentos.id = pagos_gastos_comunes.id_departamentos join gastos_comunes on pagos_gastos_comunes.id_departamentos = gastos_comunes.id_departamentos where departamentos.id = ?", id_dpto)
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

func GetBalanceGastos_comunes(c *gin.Context){
    //URL parameters
    var iddpto = c.Param("iddpto")

    fmt.Println(iddpto);
    // Results container
    gc := Gasto_comun{}
    // Fetch from database
	gastos_comunes, err := gc.BalanceGastos_comunes(iddpto)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": gastos_comunes,
		"count": len(gastos_comunes),
	})
}
