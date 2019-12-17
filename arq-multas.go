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
func (mu Multa) FetchMultasByFecha(ano_inicio string, mes_inicio string, ano_final string, mes_final string, id_dpto string) (multas []Multa, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select * from multas where fecha between '?-?-01 00:00:00' and '?-?-01 00:00:00' and id_departamentos = ?", ano_inicio, mes_inicio, ano_final, mes_final, id_dpto)
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
    var anoinicio = c.Param("anoinicio")
    var mesinicio = c.Param("mesinicio")
    var anofinal = c.Param("anofinal")
    var mesfinal = c.Param("mesfinal")
    var id_dpto = c.Param("id_dpto")

    fmt.Println(anoinicio);
    fmt.Println(mesinicio);
    fmt.Println(anofinal);
    fmt.Println(mesfinal);
    fmt.Println(id_dpto);

    // Results container
    mu := Multa{}
    // Fetch from database
	multas, err := mu.FetchMultasByFecha(anoinicio, mesinicio, anofinal, mesfinal, id_dpto)
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
func (mu Multa) FetchCountMultasByFecha(ano_inicio string, mes_inicio string, ano_final string, mes_final string, id_dpto string) (multas []Multa, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select count(*) from multas where fecha between '?-?-01 00:00:00' and '?-?-01 00:00:00' and id_departamentos = ?", ano_inicio, mes_inicio, ano_final, mes_final, id_dpto)
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
    var anoinicio = c.Param("anoinicio")
    var mesinicio = c.Param("mesinicio")
    var anofinal = c.Param("anofinal")
    var mesfinal = c.Param("mesfinal")
    var id_dpto = c.Param("id_dpto")

    fmt.Println(anoinicio);
    fmt.Println(mesinicio);
    fmt.Println(anofinal);
    fmt.Println(mesfinal);
    fmt.Println(id_dpto);

    // Results container
    mu := Multa{}
    // Fetch from database
	multas, err := mu.FetchCountMultasByFecha(anoinicio, mesinicio, anofinal, mesfinal,id_dpto)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": multas,
		"count": len(multas),
	})
}
