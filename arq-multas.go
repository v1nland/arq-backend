package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
)

// // // // TAKE ALL MULTAS // // // //
func (mu MultaDpto) FetchMultas() (multas []MultaDpto, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("SELECT multas.id, multas.grado, multas.id_departamentos, multas.monto, multas.fecha, multas.causa, departamentos.numero as numero_dpto, departamentos.dueno, departamentos.residente, departamentos.correo as correo_dueno, departamentos.correo_residente, departamentos.telefono as telefono_dueno, departamentos.telefono_residente FROM multas join departamentos on multas.id_departamentos = departamentos.id")
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var mul MultaDpto
        rows.Scan(&mul.Id, &mul.Grado, &mul.Id_departamentos, &mul.Monto, &mul.Fecha, &mul.Causa,  &mul.Num_dpto, &mul.Dueno_dpto, &mul.Residente_dpto, &mul.Correo_dueno, &mul.Correo_residente, &mul.Telefono_dueno, &mul.Telefono_residente)
        fmt.Println(mul.Id, mul.Grado, mul.Id_departamentos, mul.Monto, mul.Fecha, mul.Causa, mul.Num_dpto, mul.Dueno_dpto, mul.Residente_dpto, mul.Correo_dueno, mul.Correo_residente, mul.Telefono_dueno, mul.Telefono_residente)
        multas = append(multas, mul)
	}
	defer rows.Close()

	return
}

func GetMultas(c *gin.Context){
    // Results container
    mu := MultaDpto{}
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

// // // // TAKE ALL MULTAS WITH FECHA BETWEEN (fechai, fechaf) // // // //
func (mu Multa) FetchMultasByFecha(fechai string, fechaf string,id_dpto string) (multas []Multa, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select * from multas where (fecha between ? and ?) and id_departamentos = ?", fechai, fechaf, id_dpto)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var mul Multa
        rows.Scan(&mul.Id, &mul.Grado, &mul.Id_departamentos, &mul.Monto, &mul.Fecha, &mul.Causa)
        fmt.Println(mul.Id, mul.Grado, mul.Id_departamentos, mul.Monto, mul.Fecha, mul.Causa)
        multas = append(multas, mul)
	}
	defer rows.Close()

	return
}

func GetMultasByFecha(c *gin.Context){
    // URL parameters
    var fechai = c.Param("fechai")
    var fechaf = c.Param("fechaf")
    var id_dpto = c.Param("id_dpto")

    fmt.Println(fechai);
    fmt.Println(fechaf);
    fmt.Println(id_dpto);

    // Results container
    mu := Multa{}
    // Fetch from database
	multas, err := mu.FetchMultasByFecha(fechai, fechaf, id_dpto)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": multas,
		"count": len(multas),
	})
}

// // // // TAKE COUNT (MULTAS) WITH FECHA BETWEEN (fechai, fechaf) // // // //
func (mu Multa) FetchCountMultasByFecha(fechai string, fechaf string, id_dpto string) (multas []Multa, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select count(*) from multas where (fecha between ? and ?) and id_departamentos = ?", fechai, fechaf, id_dpto)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var mul Multa
        rows.Scan(&mul.Id, &mul.Grado, &mul.Id_departamentos, &mul.Monto, &mul.Fecha, &mul.Causa)
        fmt.Println(mul.Id, mul.Grado, mul.Id_departamentos, mul.Monto, mul.Fecha, mul.Causa)
        multas = append(multas, mul)
	}
	defer rows.Close()

	return
}

func GetCountMultasByFecha(c *gin.Context){
    // URL parameters
    var fechai = c.Param("fechai")
    var fechaf = c.Param("fechaf")
    var id_dpto = c.Param("id_dpto")

    fmt.Println(fechai);
    fmt.Println(fechaf);
    fmt.Println(id_dpto);

    // Results container
    mu := Multa{}
    // Fetch from database
	multas, err := mu.FetchCountMultasByFecha(fechai, fechaf,id_dpto)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"count": len(multas),
	})
}

// // // // INSERT MULTAS // // // //
func (mu Multa) InsertarMultas(cod_cond string, num_dpto string) (multas []Multa, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows1, err := db.Query("select * from multas where id_departamentos in (select id from departamentos where id_condominio in (select id from condominios where codigo = ?) and numero = ?)", cod_cond, num_dpto)
	if err != nil {
		return
	}
    // Take all data
	for rows1.Next() {
		var mul Multa
        rows1.Scan(&mul.Id, &mul.Grado, &mul.Id_departamentos, &mul.Monto, &mul.Fecha, &mul.Causa)
        fmt.Println(mul.Id, mul.Grado, mul.Id_departamentos, mul.Monto, mul.Fecha, mul.Causa)
        multas = append(multas, mul)
	}
	defer rows1.Close()

	return
}

func (mu Multa) InsertarMultasgr1(cod_cond string, num_dpto string, monto string, causa string) (multasgr1 []Multa, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows2, err := db.Query("insert into multas(grado, id_departamentos, monto, fecha, causa) values (1, (select id from departamentos where id_condominio in (select id from condominios where codigo = ?) and numero = ?), ?, current_timestamp, ?)", cod_cond, num_dpto, monto, causa)
	if err != nil {
		return
	}

    // Take all data
	for rows2.Next() {
		var mul Multa
        rows2.Scan(&mul.Id, &mul.Grado, &mul.Id_departamentos, &mul.Monto, &mul.Fecha, &mul.Causa)
        fmt.Println(mul.Id, mul.Grado, mul.Id_departamentos, mul.Monto, mul.Fecha, mul.Causa)
        multasgr1 = append(multasgr1, mul)
	}
	defer rows2.Close()

	return
}

func (mu Multa) InsertarMultasgr2(cod_cond string, num_dpto string, monto string, causa string) (multasgr2 []Multa, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows3, err := db.Query("insert into multas(grado, id_departamentos, monto, fecha, causa) values (2, (select id from departamentos where id_condominio in (select id from condominios where codigo = ?) and numero = ?), ?, current_timestamp, ?)", cod_cond, num_dpto, monto, causa)
	if err != nil {
		return
	}

    // Take all data
	for rows3.Next() {
		var mul Multa
        rows3.Scan(&mul.Id, &mul.Grado, &mul.Id_departamentos, &mul.Monto, &mul.Fecha, &mul.Causa)
        fmt.Println(mul.Id, mul.Grado, mul.Id_departamentos, mul.Monto, mul.Fecha, mul.Causa)
        multasgr2 = append(multasgr2, mul)
	}
	defer rows3.Close()

	return
}

func (mu Multa) InsertarMultasgr3(cod_cond string, num_dpto string, monto string, causa string) (multasgr3 []Multa, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows4, err := db.Query("insert into multas(grado, id_departamentos, monto, fecha, causa) values (3, (select id from departamentos where id_condominio in (select id from condominios where codigo = ?) and numero = ?), ?, current_timestamp, ?)", cod_cond, num_dpto, monto, causa)
	if err != nil {
		return
	}

    // Take all data
	for rows4.Next() {
		var mul Multa
        rows4.Scan(&mul.Id, &mul.Grado, &mul.Id_departamentos, &mul.Monto, &mul.Fecha, &mul.Causa)
        fmt.Println(mul.Id, mul.Grado, mul.Id_departamentos, mul.Monto, mul.Fecha, mul.Causa)
	multasgr3 = append(multasgr3, mul)
	}
	defer rows4.Close()
	return
}

func GetInsertarMultas(c *gin.Context){
    // URL parameters
    var codcond = c.Param("codcond")
    var num_dpto = c.Param("num_dpto")
    var monto = c.Param("monto")
    var causa = c.Param("causa")

    fmt.Println(codcond);
    fmt.Println(num_dpto);
    fmt.Println(monto);
    fmt.Println(causa);

    // Results container
    mu := Multa{}

    // Fetch from database
	multas, err := mu.InsertarMultas(codcond, num_dpto)
	if err != nil {
        panic(err.Error())
	}
	if len(multas) <= 1{
        multasgr1, err := mu.InsertarMultasgr1(codcond, num_dpto, monto, causa)
	    if err != nil {
            panic(err.Error())
        }

        c.JSON(200,gin.H{
            "rows": multasgr1,
            "count": len(multasgr1),
        })
	} else if len(multas) > 5 {
        multasgr3, err := mu.InsertarMultasgr3(codcond, num_dpto, monto, causa)
        if err != nil {
            panic(err.Error())
        }

        c.JSON(200,gin.H{
            "rows": multasgr3,
            "count": len(multasgr3),
        })
	} else {
        multasgr2, err := mu.InsertarMultasgr2(codcond, num_dpto, monto, causa)
        if err != nil {
            panic(err.Error())
        }

        c.JSON(200,gin.H{
            "rows": multasgr2,
            "count": len(multasgr2),
        })
	}
}


// // // // UPDATE MULTAS // // // //
func (mu Multa) UpdateMultas(grado string, id_dpto string, monto string, fecha string, causa string, idmul string) (multas []Multa, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("update multas set grado = ?, id_departamentos = ?, monto = ?, fecha = ?, causa = ? where id = ?", grado, id_dpto, monto, fecha, causa, idmul)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var mul Multa
        rows.Scan(&mul.Id, &mul.Grado, &mul.Id_departamentos, &mul.Monto, &mul.Fecha, &mul.Causa)
        fmt.Println(mul.Id, mul.Grado, mul.Id_departamentos, mul.Monto, mul.Fecha, mul.Causa)
        multas = append(multas, mul)
	}
	defer rows.Close()

	return
}

func GetUpdateMultas(c *gin.Context){
    // URL parameters
    var grado = c.Param("grado")
    var id_dpto = c.Param("id_dpto")
    var monto = c.Param("monto")
    var fecha = c.Param("fecha")
    var causa = c.Param("causa")
    var idmul = c.Param("idmul")

    fmt.Println(grado);
    fmt.Println(id_dpto);
    fmt.Println(monto);
    fmt.Println(fecha);
    fmt.Println(causa);
    fmt.Println(idmul);

    // Results container
    mu := Multa{}
    // Fetch from database
	multas, err := mu.UpdateMultas(grado, id_dpto, monto, fecha, causa, idmul)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": multas,
		"count": len(multas),
	})
}


// // // // TAKE ALL MULTAS BY ID // // // //
func (mu Multa) FetchMultasByID(idmulta string) (multas []Multa, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select * from multas where id = ?", idmulta)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var mul Multa
        rows.Scan(&mul.Id, &mul.Grado, &mul.Id_departamentos, &mul.Monto, &mul.Fecha, &mul.Causa)
        fmt.Println(mul.Id, mul.Grado, mul.Id_departamentos, mul.Monto, mul.Fecha, mul.Causa)
        multas = append(multas, mul)
	}
	defer rows.Close()

	return
}

func GetMultasByID(c *gin.Context){
    // URL parameters
    var idmulta = c.Param("idmulta")

    fmt.Println(idmulta);

    // Results container
    mu := Multa{}
    // Fetch from database
	multas, err := mu.FetchMultasByID(idmulta)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": multas,
		"count": len(multas),
	})
}


// // // // CONTAR MULTAS // // // //
func (mu CuentaMultas) CuentaMultasFecha(fechai string, fechaf string) (multas []CuentaMultas, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select count(*) from multas where fecha between ? and ?", fechai, fechaf)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var mul CuentaMultas
        rows.Scan(&mul.Cuenta)
        fmt.Println(mul.Cuenta)
        multas = append(multas, mul)
	}
	defer rows.Close()

	return
}

func GetCuentaMultasFecha(c *gin.Context){
    // URL parameters
    var fechai = c.Param("fechai")
    var fechaf = c.Param("fechaf")

    fmt.Println(fechai);
    fmt.Println(fechaf);

    // Results container
    mu := CuentaMultas{}
    // Fetch from database
	multas, err := mu.CuentaMultasFecha(fechai, fechaf)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": multas,
		"count": len(multas),
	})
}
