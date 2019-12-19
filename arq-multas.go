package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
)

// // // // TAKE ALL MULTAS // // // //
func (mu Multa) FetchMultas() (multas []Multa, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("SELECT id, grado, id_departamentos, monto, fecha, causa FROM multas")
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
func (mu Multa) InsertarMultas(id_dpto string) (multas []Multa, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows1, err := db.Query("select * from multas where id_departamentos = ?", id_dpto)
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

func (mu Multa) InsertarMultasgr1(id_dpto string, monto string, causa string) (multasgr1 []Multa, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows2, err := db.Query("insert into multas(grado, id_departamentos, monto, fecha, causa) values (1, ?, ?, current_timestamp, ?)", id_dpto, monto, causa)
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

func (mu Multa) InsertarMultasgr2(id_dpto string, monto string, causa string) (multasgr2 []Multa, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows3, err := db.Query("insert into multas(grado, id_departamentos, monto, fecha, causa) values (2, ?, ?, current_timestamp, ?)", id_dpto, monto, causa)
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

func (mu Multa) InsertarMultasgr3(id_dpto string, monto string, causa string) (multasgr3 []Multa, err error) {
    // Opens DB
    db := GetConnection()
    fmt.Println("Entré a insertar3")

    // SQL query
	rows4, err := db.Query("insert into multas(grado, id_departamentos, monto, fecha, causa) values (3, ?, ?, current_timestamp, ?)", id_dpto, monto, causa)
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
	fmt.Println("Estoy saliendo de insertar3")
	return
}

func GetInsertarMultas(c *gin.Context){
    // URL parameters
    var iddpto = c.Param("iddpto")
    var monto = c.Param("monto")
    var causa = c.Param("causa")

    fmt.Println(iddpto);
    fmt.Println(monto);
    fmt.Println(causa);

    // Results container
    mu := Multa{}

    // Fetch from database
	multas, err := mu.InsertarMultas(iddpto)
	if err != nil {
        panic(err.Error())
	} 
	if len(multas) <=1{
	multasgr1, err := mu.InsertarMultasgr1(iddpto, monto, causa)
	if err != nil {
	panic(err.Error())
	}
	c.JSON(200,gin.H{
		"rows2":multasgr1,
	})
	} else if len(multas) > 5{
	multasgr3, err := mu.InsertarMultasgr3(iddpto, monto, causa)
	if err != nil {
	panic(err.Error())
	}
	c.JSON(200,gin.H{
		"rows4":multasgr3,
	})
	} else{
	multasgr2, err := mu.InsertarMultasgr2(iddpto, monto, causa)
	if err != nil {
	panic(err.Error())
	}
	c.JSON(200,gin.H{
		"rows3":multasgr2,
	})
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows1": multas,
		"count": len(multas),
	})
}
