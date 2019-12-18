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


// // // // QUERY FIXED// // // //
// // // // QUERY FIXED// // // //
// // // // QUERY FIXED// // // //
// // // // QUERY FIXED// // // //
// // // // TAKE ALL MEDICIONESAGUA WITH FECHA BETWEEN (ano_mes_inicial, ano_mes_final) // // // //
func (Ma MedicionAgua) FetchMedicionesAguaByFecha(fechai string, fechaf string) (mediciones_agua []MedicionAgua, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select * from mediciones_agua where fecha between ? and ?", fechai, fechaf)
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
    var fechai = c.Param("fechai")
    var fechaf = c.Param("fechaf")

    fmt.Println(fechai);
    fmt.Println(fechaf);

    // Results container
    ma := MedicionAgua{}
    // Fetch from database
	mediciones_agua, err := ma.FetchMedicionesAguaByFecha(fechai, fechaf)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": mediciones_agua,
		"count": len(mediciones_agua),
	})
}

// // // INSERT MEDICION_AGUA // // //
func (Ma MedicionAgua) InsertMedicionesAgua(litros string, num_dpto string, cod_cond string) (mediciones_agua []MedicionAgua, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("insert into mediciones_agua(fecha, litros, id_departamentos) values (current_timestamp, ? , (select id from departamentos where departamentos.numero = ? and departamentos.id_condominio = ?))", litros, num_dpto, cod_cond)
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

func GetInsertMedicionesAgua(c *gin.Context){
    // URL parameters
    var litros = c.Param("litros")
    var num_dpto = c.Param("num_dpto")
    var cod_cond = c.Param("cod_cond")

    fmt.Println(litros);
    fmt.Println(num_dpto);
    fmt.Println(cod_cond);

    // Results container
    ma := MedicionAgua{}
    // Fetch from database
	mediciones_agua, err := ma.InsertMedicionesAgua(litros, num_dpto, cod_cond)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": mediciones_agua,
		"count": len(mediciones_agua),
	})
}
