package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
)

// // // // TAKE ALL PAGOSGASTOSCOMUNES // // // //
func (pag PagosGastosComunes) FetchPagos() (pagos_gastos_comunes []PagosGastosComunes, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("SELECT id, monto, fecha, id_departamentos FROM pagos_gastos_comunes")
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var pago PagosGastosComunes
        rows.Scan(&pago.Id, &pago.Monto, &pago.Fecha, &pago.Id_departamento)
        fmt.Println(pago.Id, pago.Monto, pago.Fecha, pago.Id_departamento)
        pagos_gastos_comunes = append(pagos_gastos_comunes, pago)
	}
	defer rows.Close()

	return
}

func GetPagos(c *gin.Context){
    // Results container
    pag := PagosGastosComunes{}
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

// // // // TAKE ALL PAGOSGASTOSCOMUNES WITH ID_DPTO=id_dpto // // // //
func (pag PagosGastosComunes) FetchPagosByDptoID(id_dpto string) (pagos_gastos_comunes []PagosGastosComunes, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select * from pagos_gastos_comunes where id_departamentos=?", id_dpto)
	if err != nil {
        return
    }

    // Take all data
	for rows.Next() {
		var pago PagosGastosComunes
        rows.Scan(&pago.Id, &pago.Monto, &pago.Fecha, &pago.Id_departamento)
        fmt.Println(pago.Id, pago.Monto, pago.Fecha, pago.Id_departamento)
        pagos_gastos_comunes = append(pagos_gastos_comunes, pago)
	}
	defer rows.Close()

	return
}

func GetPagosByDptoID(c *gin.Context){
    // URL parameters
    var id_dpto = c.Param("id_dpto")

    fmt.Println(id_dpto);
    // Results container
    pag := PagosGastosComunes{}
    // Fetch from database
	pagos_gastos_comunes, err := pag.FetchPagosByDptoID(id_dpto)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": pagos_gastos_comunes,
		"count": len(pagos_gastos_comunes),
	})
}

// // // // TAKE ALL PAGOSGASTOSCOMUNES WITH ID_COND=id_cond // // // //
func (pag PagosGastosComunes) FetchPagosByCondominioID(id_cond string) (pagos_gastos_comunes []PagosGastosComunes, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select * from pagos_gastos_comunes where id_departamentos in (select id from departamentos where id_condominio = ?)", id_cond)
	if err != nil {
        return
    }

    // Take all data
	for rows.Next() {
		var pago PagosGastosComunes
        rows.Scan(&pago.Id, &pago.Monto, &pago.Fecha, &pago.Id_departamento)
        fmt.Println(pago.Id, pago.Monto, pago.Fecha, pago.Id_departamento)
        pagos_gastos_comunes = append(pagos_gastos_comunes, pago)
	}
	defer rows.Close()

	return
}

func GetPagosByCondominioID(c *gin.Context){
    // URL parameters
    var id_cond = c.Param("id_cond")

    fmt.Println(id_cond);
    // Results container
    pag := PagosGastosComunes{}
    // Fetch from database
	pagos_gastos_comunes, err := pag.FetchPagosByCondominioID(id_cond)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": pagos_gastos_comunes,
		"count": len(pagos_gastos_comunes),
	})
}

// // // // TAKE ALL MEDICIONESAGUA WITH ID_DPTO=id_dpto AND FECHA BETWEEN (ano_mes_inicial, ano_mes_final) // // // //
func (pag PagosGastosComunes) FetchPagosByFechaAndCondominioID(id_cond string, anoinicio string, mesinicio string, anofinal string, mesfinal string) (pagos_gastos_comunes []PagosGastosComunes, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select * from pagos_gastos_comunes where id_departamentos in (select id from departamentos where id_condominio = ?) and fecha between '?-?-01 00:00:00' and '?-?-01 00:00:00", id_cond, anoinicio, mesinicio, anofinal, mesfinal)
	if err != nil {
        return
    }

    // Take all data
	for rows.Next() {
		var pago PagosGastosComunes
        rows.Scan(&pago.Id, &pago.Monto, &pago.Fecha, &pago.Id_departamento)
        fmt.Println(pago.Id, pago.Monto, pago.Fecha, pago.Id_departamento)
        pagos_gastos_comunes = append(pagos_gastos_comunes, pago)
	}
	defer rows.Close()

	return
}

func GetPagosByFechaAndCondominioID(c *gin.Context){
    // URL parameters
    var id_cond = c.Param("id_cond")
    var anoinicio = c.Param("anoinicio")
    var mesinicio = c.Param("mesinicio")
    var anofinal = c.Param("anofinal")
    var mesfinal = c.Param("mesfinal")

    fmt.Println(id_cond);
    fmt.Println(anoinicio);
    fmt.Println(mesinicio);
    fmt.Println(anofinal);
    fmt.Println(mesfinal);
    // Results container
    pag := PagosGastosComunes{}
    // Fetch from database
	pagos_gastos_comunes, err := pag.FetchPagosByFechaAndCondominioID(id_cond, anoinicio, mesinicio, anofinal, mesfinal)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": pagos_gastos_comunes,
		"count": len(pagos_gastos_comunes),
	})
}
