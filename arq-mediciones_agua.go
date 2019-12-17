package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
)

// // // // TAKE ALL MEDICIONESAGUA // // // //
func (Ma MedicionAgua) FetchMedicionesAgua() (mediciones_agua []MedicionAgua, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("SELECT * FROM mediciones_agua")
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var med MedicionAgua
        rows.Scan(&med.Id, &med.Fecha, &med.Litros, &med.Id_departamentos)
        fmt.Println(med.Id, med.Fecha, med.Litros, med.Id_departamentos)
        mediciones_agua = append(mediciones_agua, med)
	}
	defer rows.Close()

	return
}

func GetMedicionesAgua(c *gin.Context){
    // Results container
    ma := MedicionAgua{}
    // Fetch from database
	mediciones_agua, err := ma.FetchMedicionesAgua()
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": mediciones_agua,
		"count": len(mediciones_agua),
	})
}

// // // // TAKE ALL MEDICIONESAGUA WITH ID_DPTO=id_dpto // // // //
func (Ma MedicionAgua) FetchMedicionesAguaByDptoID(id_dpto string) (mediciones_agua []MedicionAgua, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select * from mediciones_agua where id_departamentos=?", id_dpto)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var med MedicionAgua
        rows.Scan(&med.Id, &med.Fecha, &med.Litros, &med.Id_departamentos)
        fmt.Println(med.Id, med.Fecha, med.Litros, med.Id_departamentos)
        mediciones_agua = append(mediciones_agua, med)
	}
	defer rows.Close()

	return
}

func GetMedicionesAguaByDptoID(c *gin.Context){
    // URL parameters
    var id_dpto = c.Param("id_dpto")

    fmt.Println(id_dpto);
    // Results container
    ma := MedicionAgua{}
    // Fetch from database
	mediciones_agua, err := ma.FetchMedicionesAguaByDptoID(id_dpto)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": mediciones_agua,
		"count": len(mediciones_agua),
	})
}

// // // // TAKE ALL MEDICIONESAGUA WITH ID_COND=id_cond // // // //
func (Ma MedicionAgua) FetchMedicionesAguaByCondID(id_cond string) (mediciones_agua []MedicionAgua, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select * from mediciones_agua where id_departamentos in (select id from departamentos where id_condominio=?)", id_cond)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var med MedicionAgua
        rows.Scan(&med.Id, &med.Fecha, &med.Litros, &med.Id_departamentos)
        fmt.Println(med.Id, med.Fecha, med.Litros, med.Id_departamentos)
        mediciones_agua = append(mediciones_agua, med)
	}
	defer rows.Close()

	return
}

func GetMedicionesAguaByCondID(c *gin.Context){
    // URL parameters
    var id_cond = c.Param("id_cond")

    fmt.Println(id_cond);
    // Results container
    ma := MedicionAgua{}
    // Fetch from database
	mediciones_agua, err := ma.FetchMedicionesAguaByCondID(id_cond)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": mediciones_agua,
		"count": len(mediciones_agua),
	})
}


// // // // QUERIES NEEDING FIX // // // //
// // // // QUERIES NEEDING FIX // // // //
// // // // QUERIES NEEDING FIX // // // //
// // // // QUERIES NEEDING FIX // // // //
// // // // QUERIES NEEDING FIX // // // //
// // // // QUERIES NEEDING FIX // // // //
// // // // QUERIES NEEDING FIX // // // // 
// // // // TAKE ALL MEDICIONESAGUA WITH FECHA BETWEEN (ano_mes_inicial, ano_mes_final) // // // //
func (Ma MedicionAgua) FetchMedicionesAguaByFecha(ano_inicio string, mes_inicio string, ano_final string, mes_final string) (mediciones_agua []MedicionAgua, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select * from mediciones_agua where fecha between '?-?-01 00:00:00' and '?-?-01 00:00:00'", ano_inicio, mes_inicio, ano_final, mes_final)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var med MedicionAgua
        rows.Scan(&med.Id, &med.Fecha, &med.Litros, &med.Id_departamentos)
        fmt.Println(med.Id, med.Fecha, med.Litros, med.Id_departamentos)
        mediciones_agua = append(mediciones_agua, med)
	}
	defer rows.Close()

	return
}

func GetMedicionesAguaByFecha(c *gin.Context){
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
    ma := MedicionAgua{}
    // Fetch from database
	mediciones_agua, err := ma.FetchMedicionesAguaByFecha(ano_inicio, mes_inicio, ano_final, mes_final)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": mediciones_agua,
		"count": len(mediciones_agua),
	})
}
