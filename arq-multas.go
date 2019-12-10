package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
)

//Take all data-------------------------------------------------------------------------------------

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

//Select todas las multas en el mes y ano del depto con id_dpto-----------------------------------------

func (mu Multa) FetchMultasFecha(ano_inicio string, mes_inicio string, ano_final string, mes_final string, id_dpto string) (multas []Multa, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select * from multas where fecha between '?-?-01 00:00:00' and '?-?-01 00:00:00' and departamentos_id = ?", ano_inicio, mes_inicio, ano_final, mes_final, id_dpto)
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

func GetMultasFecha(c *gin.Context){
    // URL parameters
    var anoinicio = c.Param("anoinicio")
    var mesinicio = c.Param("mesinicio")
    var anofinal = c.Param("anofinal")
    var mesfinal = c.Param("mesfinal")
    var iddpto = c.Param("iddpto")

    fmt.Println(anoinicio);
    fmt.Println(mesinicio);
    fmt.Println(anofinal);
    fmt.Println(mesfinal);
    fmt.Println(iddpto);

    // Results container
    mu := Multa{}
    // Fetch from database
	multas, err := mu.FetchMultasFecha(anoinicio, mesinicio, anofinal, mesfinal, iddpto)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": multas,
		"count": len(multas),
	})
}

//Count cantidad de multas en el mes entregado, dpto con id----------------------------------------------

func (mu Multa) FetchCountMultasFecha(ano_inicio string, mes_inicio string, ano_final string, mes_final string, id_dpto string) (multas []Multa, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select count(*) from multas where fecha between '?-?-01 00:00:00' and '?-?-01 00:00:00' and departamentos_id = ?", ano_inicio, mes_inicio, ano_final, mes_final, id_dpto)
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

func GetCountMultasFecha(c *gin.Context){
    // URL parameters
    var anoinicio = c.Param("anoinicio")
    var mesinicio = c.Param("mesinicio")
    var anofinal = c.Param("anofinal")
    var mesfinal = c.Param("mesfinal")
    var iddpto = c.Param("iddpto")

    fmt.Println(anoinicio);
    fmt.Println(mesinicio);
    fmt.Println(anofinal);
    fmt.Println(mesfinal);
    fmt.Println(iddpto);

    // Results container
    mu := Multa{}
    // Fetch from database
	multas, err := mu.FetchCountMultasFecha(anoinicio, mesinicio, anofinal, mesfinal,iddpto)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": multas,
		"count": len(multas),
	})
}

