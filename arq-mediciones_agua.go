package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
)

//Take all data----------------------------------------------------------------------------------

func (Ma Medicion_agua) FetchMediciones_agua() (mediciones_agua []Medicion_agua, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("SELECT * FROM mediciones_agua")
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var med Medicion_agua
        rows.Scan(&med.Id, &med.Fecha, &med.Litros, &med.id_departamentos)
        fmt.Println(med.Id, med.Fecha, med.Litros, med.id_departamentos)
        conserjes = append(mediciones_agua, med)
	}
	defer rows.Close()

	return
}

func GetMediciones_agua(c *gin.Context){
    // Results container
    ma := Medicion_agua{}
    // Fetch from database
	mediciones_agua, err := ec.FetchMediciones_agua()
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": mediciones_agua,
		"count": len(mediciones_agua),
	})
}

//SELECT MEDICIONES DE AGUA POR ID_DPTO-------------------------------------------------------------

func (Ma Medicion_agua) FetchMediciones_aguaIddpto(id_departamentos int) (mediciones_agua []Medicion_agua, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select * from mediciones_agua where id_departamentos=?", id_departamentos)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var med Medicion_agua
        rows.Scan(&med.Id, &med.Fecha, &med.Litros, &med.id_departamentos)
        fmt.Println(med.Id, med.Fecha, med.Litros, med.id_departamentos)
        conserjes = append(mediciones_agua, med)
	}
	defer rows.Close()

	return
}

func GetMediciones_aguaIddpto(c *gin.Context){
    // URL parameters
    var iddepartamento = c.Param("iddepartamento")

    fmt.Println(iddepartamento);
    // Results container
    ma := Medicion_agua{}
    // Fetch from database
	mediciones_agua, err := ec.FetchMediciones_aguaIddpto(iddepartamento)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": mediciones_agua,
		"count": len(mediciones_agua),
	})
}

//select mediciones por idcondominio--------------------------------------------------------------

func (Ma Medicion_agua) FetchMediciones_aguaIdcond(id_condominio int) (mediciones_agua []Medicion_agua, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select * from mediciones_agua where id_departamentos in (select id from departamentos where id_condominio=?)", id_condominio)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var med Medicion_agua
		var dep Departamento
        rows.Scan(&med.Id, &med.Fecha, &med.Litros, &med.id_departamentos, &dep.id_condominio)
        fmt.Println(med.Id, med.Fecha, med.Litros, med.id_departamentos, dep.id_condominio)
        conserjes = append(mediciones_agua, med)
	}
	defer rows.Close()

	return
}

func GetMediciones_aguaIdcond(c *gin.Context){
    // URL parameters
    var idcond = c.Param("idcond")

    fmt.Println(idcond);
    // Results container
    ma := Medicion_agua{}
    // Fetch from database
	mediciones_agua, err := ec.FetchMediciones_aguaIdcond(idcond)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": mediciones_agua,
		"count": len(mediciones_agua),
	})
}

//select todas las mediciones en un mes, año entregado-----------------------------------------------

func (Ma Medicion_agua) FetchMediciones_aguaFecha(ano_inicio int, mes_inicio int, ano_final int, mes_final int) (mediciones_agua []Medicion_agua, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select * from mediciones_agua where fecha between '?-?-01 00:00:00' and '?-?-01 00:00:00'", ano_inicio, mes_inicio, ano_final, mes_final)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var med Medicion_agua
        rows.Scan(&med.Id, &med.Fecha, &med.Litros, &med.id_departamentos)
        fmt.Println(med.Id, med.Fecha, med.Litros, med.id_departamentos)
        conserjes = append(mediciones_agua, med)
	}
	defer rows.Close()

	return
}

func GetMediciones_aguaFecha(c *gin.Context){
    // URL parameters
    var ano_inicio = c.Param("ano_inicio")
    var mes_inicio = c.Param("mes_inicio")
    var ano_final = c.Param("ano_final")
    var mes_final = c.Param("mes_final")

    fmt.Println(ano_inicio);
    fmt.Println(mes_inicio);
    fmt.Println(ano_final);
    fmt.Println(mes_final);
    // Results container
    ma := Medicion_agua{}
    // Fetch from database
	mediciones_agua, err := ec.FetchMediciones_aguaFecha(ano_inicio, mes_inicio, ano_final, mes_final)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": mediciones_agua,
		"count": len(mediciones_agua),
	})
}
