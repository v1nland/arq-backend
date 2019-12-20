package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
)

// // // // TAKE ALL PAGOSGASTOSCOMUNES // // // //
func (pag PagosGastosComunesDpto) FetchPagos() (pagos_gastos_comunes []PagosGastosComunesDpto, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("SELECT pagos_gastos_comunes.id, pagos_gastos_comunes.monto, pagos_gastos_comunes.fecha, pagos_gastos_comunes.id_departamentos, departamentos.numero as numero_dpto, departamentos.dueno, departamentos.residente, departamentos.correo as correo_dueno, departamentos.correo_residente, departamentos.telefono as telefono_dueno, departamentos.telefono_residente FROM pagos_gastos_comunes join departamentos on pagos_gastos_comunes.id_departamentos = departamentos.id")
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var pago PagosGastosComunesDpto
        rows.Scan(&pago.Id, &pago.Monto, &pago.Fecha, &pago.Id_departamento, &pago.Num_dpto, &pago.Dueno_dpto, &pago.Residente_dpto, &pago.Correo_dueno, &pago.Correo_residente, &pago.Telefono_dueno, &pago.Telefono_residente)
        fmt.Println(pago.Id, pago.Monto, pago.Fecha, pago.Id_departamento, pago.Num_dpto, pago.Dueno_dpto, pago.Residente_dpto, pago.Correo_dueno, pago.Correo_residente, pago.Telefono_dueno, pago.Telefono_residente)
        pagos_gastos_comunes = append(pagos_gastos_comunes, pago)
	}
	defer rows.Close()

	return
}

func GetPagos(c *gin.Context){
    // Results container
    pag := PagosGastosComunesDpto{}
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

// // // // TAKE ALL PAGOSGASTOSCOMUNES WITH ID_DPTO=id_dpto AND FECHA BETWEEN (fechai, fechaf) // // // //
func (pag PagosGastosComunes) FetchPagosByFechaAndCondominioID(id_cond string, fechai string, fechaf string) (pagos_gastos_comunes []PagosGastosComunes, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select * from pagos_gastos_comunes where id_departamentos in (select id from departamentos where id_condominio = ?) and fecha between ? and ?", id_cond, fechai, fechaf)
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
    var fechai = c.Param("fechai")
    var fechaf = c.Param("fechaf")

    fmt.Println(id_cond);
    fmt.Println(fechai);
    fmt.Println(fechaf);

    // Results container
    pag := PagosGastosComunes{}
    // Fetch from database
	pagos_gastos_comunes, err := pag.FetchPagosByFechaAndCondominioID(id_cond, fechai, fechaf)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": pagos_gastos_comunes,
		"count": len(pagos_gastos_comunes),
	})
}

// // // // INSERT PAGOS_GASTOS_COMUNES // // // //
func (pag PagosGastosComunes) InsertPagos(monto string, num_dpto string, cod_cond string) (pagos_gastos_comunes []PagosGastosComunes, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("insert into pagos_gastos_comunes(monto, fecha, id_departamentos) values (?, current_timestamp, (select id from departamentos where numero = ? and id_condominio = (select id from condominios where codigo = ?)))", monto, num_dpto, cod_cond)
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

func GetInsertPagos(c *gin.Context){
    // URL parameters
    var monto = c.Param("monto")
    var num_dpto = c.Param("num_dpto")
    var cod_cond = c.Param("cod_cond")

    fmt.Println(monto);
    fmt.Println(num_dpto);
    fmt.Println(cod_cond);

    // Results container
    pag := PagosGastosComunes{}
    // Fetch from database
	pagos_gastos_comunes, err := pag.InsertPagos(monto, num_dpto, cod_cond)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": pagos_gastos_comunes,
		"count": len(pagos_gastos_comunes),
	})
}


// // // // UPDATE PAGOS_GASTOS_COMUNES // // // //
func (pag PagosGastosComunes) UpdatePagos(monto string, fecha string, id_dpto string, idpago string) (pagos_gastos_comunes []PagosGastosComunes, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("update pagos_gastos_comunes set monto = ?, fecha = ?, id_departamentos = ? where id = ?", monto, fecha, id_dpto, idpago)
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

func GetUpdatePagos(c *gin.Context){
    // URL parameters
    var monto = c.Param("monto")
    var fecha = c.Param("fecha")
    var id_dpto = c.Param("id_dpto")
    var idpago = c.Param("idpago")

    fmt.Println(monto);
    fmt.Println(fecha);
    fmt.Println(id_dpto);
    fmt.Println(idpago);

    // Results container
    pag := PagosGastosComunes{}
    // Fetch from database
	pagos_gastos_comunes, err := pag.UpdatePagos(monto, fecha, id_dpto, idpago)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": pagos_gastos_comunes,
		"count": len(pagos_gastos_comunes),
	})
}
