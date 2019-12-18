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

// // // // QUERIES NEEDING FIX // // // //
// // // // QUERIES NEEDING FIX // // // //
// // // // QUERIES NEEDING FIX // // // //
// // // // QUERIES NEEDING FIX // // // //
// // // // QUERIES NEEDING FIX // // // //
// // // // QUERIES NEEDING FIX // // // //
// // // // QUERIES NEEDING FIX // // // // 
// // // // TAKE ALL MULTAS WITH FECHA BETWEEN (ano_mes_inicial, ano_mes_final) // // // //
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

// // // // TAKE COUNT (MULTAS) WITH FECHA BETWEEN (ano_mes_inicial, ano_mes_final) // // // //
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
