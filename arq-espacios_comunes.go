package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
)

// // // // TAKE ALL ESPACIOSCOMUNES // // // //
func (Ec EspacioComun) FetchEspaciosComunes() (espacios_comunes []EspacioComun, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("SELECT * FROM espacios_comunes")
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var esc EspacioComun
        rows.Scan(&esc.Id, &esc.Nombre, &esc.Id_condominio, &esc.Estado)
        fmt.Println(esc.Id, esc.Nombre, esc.Id_condominio, esc.Estado)
        espacios_comunes = append(espacios_comunes, esc)
	}
	defer rows.Close()

	return
}

func GetEspaciosComunes(c *gin.Context){
    // Results container
    ec := EspacioComun{}
    // Fetch from database
	espacios_comunes, err := ec.FetchEspaciosComunes()
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": espacios_comunes,
		"count": len(espacios_comunes),
	})
}

// // // // TAKE ALL ESPACIOSCOMUNES WITH ID_COND=id_cond // // // //
func (Ec EspacioComun) FetchEspaciosComunesByCondID(id_cond string) (espacios_comunes []EspacioComun, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select * from espacios_comunes where id_condominio =?", id_cond)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var esc EspacioComun
        rows.Scan(&esc.Id, &esc.Nombre, &esc.Id_condominio, &esc.Estado)
        fmt.Println(esc.Id, esc.Nombre, esc.Id_condominio, esc.Estado)
        espacios_comunes = append(espacios_comunes, esc)
	}
	defer rows.Close()

	return
}

func GetEspaciosComunesByCondID(c *gin.Context){
    //URL parameters
    var idcond = c.Param("id_cond")

    fmt.Println(idcond);
    // Results container
    ec := EspacioComun{}
    // Fetch from database
	espacios_comunes, err := ec.FetchEspaciosComunesByCondID(idcond)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": espacios_comunes,
		"count": len(espacios_comunes),
	})
}

// // // // UPDATE ALL ESPACIOSCOMUNES WITH ID_COND=id_cond // // // //
func (Ec EspacioComun) UpdateEspaciosComunesByID(estado string, id_ec string) (espacios_comunes []EspacioComun, err error) {
    // Opens DB
    db := GetConnection()

    // SQL queryGetPorCond/2/
	rows, err := db.Query("update espacios_comunes set estado= ? where id = ?", estado, id_ec)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var esc EspacioComun
        rows.Scan(&esc.Id, &esc.Nombre, &esc.Id_condominio, &esc.Estado)
        fmt.Println(esc.Id, esc.Nombre, esc.Id_condominio, esc.Estado)
        espacios_comunes = append(espacios_comunes, esc)
	}
	defer rows.Close()

	return
}

func GetUpdateEspaciosComunesByID(c *gin.Context){
    //URL parameters
    var estado = c.Param("estado")
    var id_ec = c.Param("id_ec")

    fmt.Println(estado);
    fmt.Println(id_ec);
    // Results container
    ec := EspacioComun{}
    // Fetch from database
	espacios_comunes, err := ec.UpdateEspaciosComunesByID(estado, id_ec)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": espacios_comunes,
		"count": len(espacios_comunes),
	})
}

// // // // INSERT ESPACIOSCOMUNES// // // //
func (Ec EspacioComun) InsertEspaciosComunes(nombre string, idcond string, estado string) (espacios_comunes []EspacioComun, err error) {
    // Opens DB
    db := GetConnection()

    // SQL queryGetPorCond/2/
	rows, err := db.Query("insert into espacios_comunes(nombre, id_condominio, estado) values (?, ?, ?)",nombre, idcond, estado)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var esc EspacioComun
        rows.Scan(&esc.Id, &esc.Nombre, &esc.Id_condominio, &esc.Estado)
        fmt.Println(esc.Id, esc.Nombre, esc.Id_condominio, esc.Estado)
        espacios_comunes = append(espacios_comunes, esc)
	}
	defer rows.Close()

	return
}

func GetInsertEspaciosComunes(c *gin.Context){
    //URL parameters
    var nombre = c.Param("nombre")
    var id_condominio = c.Param("id_condominio")
    var estado = c.Param("estado")

    fmt.Println(nombre);
    fmt.Println(id_condominio);
    fmt.Println(estado);

    // Results container
    ec := EspacioComun{}
    // Fetch from database
	espacios_comunes, err := ec.InsertEspaciosComunes(nombre, id_condominio, estado)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": espacios_comunes,
		"count": len(espacios_comunes),
	})
}
