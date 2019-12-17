package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
)

// // // // TAKE ALL GASTOSCOMUNES // // // //
func (Gc GastoComun) FetchGastosComunes() (gastos_comunes []GastoComun, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("SELECT * FROM gastos_comunes")
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var gac GastoComun
        rows.Scan(&gac.Id, &gac.Monto, &gac.Detalle, &gac.Fecha, &gac.Id_departamentos)
        fmt.Println(gac.Id, gac.Monto, gac.Detalle, gac.Fecha, gac.Id_departamentos)
        gastos_comunes = append(gastos_comunes, gac)
	}
	defer rows.Close()

	return
}

func GetGastosComunes(c *gin.Context){
    // Results container
    gc := GastoComun{}
    // Fetch from database
	gastos_comunes, err := gc.FetchGastosComunes()
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": gastos_comunes,
		"count": len(gastos_comunes),
	})
}

// // // // GASTOSCOMUNES BALANCE // // // //
func (Gc GastoComun) BalanceGastosComunes(id_dpto string) (gastos_comunes []GastoComun, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("SELECT departamentos.numero, pagos_gastos_comunes.fecha, departamentos.dueno, gastos_comunes.monto as 'montogc', pagos_gastos_comunes.monto as 'pagosgc' from departamentos join pagos_gastos_comunes on departamentos.id = pagos_gastos_comunes.id_departamentos join gastos_comunes on pagos_gastos_comunes.id_departamentos = gastos_comunes.id_departamentos where departamentos.id = ?", id_dpto)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var gac GastoComun
        rows.Scan(&gac.Id, &gac.Monto, &gac.Detalle, &gac.Fecha, &gac.Id_departamentos)
        fmt.Println(gac.Id, gac.Monto, gac.Detalle, gac.Fecha, gac.Id_departamentos)
        gastos_comunes = append(gastos_comunes, gac)
	}
	defer rows.Close()

	return
}

func GetBalanceGastosComunes(c *gin.Context){
    //URL parameters
    var id_dpto = c.Param("id_dpto")

    fmt.Println(id_dpto);
    // Results container
    gc := GastoComun{}
    // Fetch from database
	gastos_comunes, err := gc.BalanceGastosComunes(id_dpto)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": gastos_comunes,
		"count": len(gastos_comunes),
	})
}
