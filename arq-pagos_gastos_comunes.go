package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
)

//Take all data-------------------------------------------------------------------------------------

func (pag Pagos_gastos_comunes) FetchPagos() (pagos_gastos_comunes []Pago_gastos_comunes, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("SELECT id, monto, fecha, id_departamentos FROM pagos_gastos_comunes")
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var pago Pagos_gastos_comunes
        rows.Scan(&pago.Id, &pago.Monto, &pago.fecha, &pago.Id_departamentos)
        fmt.Println(pago.Id, pago.Monto, pago.fecha, pago.Id_departamentos)
        pagos_gastos_comunes = append(pagos_gastos_comunes, pago)
	}
	defer rows.Close()

	return
}

func GetPagos(c *gin.Context){
    // Results container
    pag := Pago_gastos_comunes{}
    // Fetch from database
	pagos_gastos_comunes, err := pag.FetchPagos()
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": pagos_gastos_comunes,
		"count": len(pagos_gastos_comunes),
	})
}


//Select con id_departamentos------------------------------------------------------------------------

func (pag Pagos_gastos_comunes) FetchPagosIddpto() (id_departamentos int) (pagos_gastos_comunes []Pago_gastos_comunes, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select * from pagos_gastos_comunes where id_departamentos=?", id_dpto)
	if err != nil {

    // Take all data
	for rows.Next() {
		var pago Pagos_gastos_comunes
        rows.Scan(&pago.Id, &pago.Monto, &pago.fecha, &pago.Id_departamentos)
        fmt.Println(pago.Id, pago.Monto, pago.fecha, pago.Id_departamentos)
        pagos_gastos_comunes = append(pagos_gastos_comunes, pago)
	}
	defer rows.Close()

	return
}

func GetPagosIddpto(c *gin.Context){
    // URL parameters
    var iddepartamento = c.Param("iddepartamento")

    fmt.Println(iddepartamento);
    // Results container
    pag := Pago_gastos_comunes{}
    // Fetch from database
	pagos_gastos_comunes, err := pag.FetchPagosIddpto(iddepartamento)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": pagos_gastos_comunes,
		"count": len(pagos_gastos_comunes),
	})
}		return
	}

