package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
)

// // // // TAKE ALL ESPACIOSCOMUNES // // // //
func (Ec EspacioComunCond) FetchEspaciosComunes() (espacios_comunes []EspacioComunCond, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("SELECT espacios_comunes.id, espacios_comunes.nombre, espacios_comunes.id_condominio, espacios_comunes.estado, espacios_comunes.descripcion, condominios.codigo, condominios.nombre FROM espacios_comunes join condominios on espacios_comunes.id_condominio = condominios.id")
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var esc EspacioComunCond
        rows.Scan(&esc.Id, &esc.Nombre, &esc.Id_condominio, &esc.Estado, &esc.Descripcion, &esc.Codigo_cond, &esc.Nombre_cond)
        fmt.Println(esc.Id, esc.Nombre, esc.Id_condominio, esc.Estado, esc.Descripcion, esc.Codigo_cond, esc.Nombre_cond)
        espacios_comunes = append(espacios_comunes, esc)
	}
	defer rows.Close()

	return
}

func GetEspaciosComunes(c *gin.Context){
    // Results container
    ec := EspacioComunCond{}
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
        rows.Scan(&esc.Id, &esc.Nombre, &esc.Id_condominio, &esc.Estado, &esc.Descripcion)
        fmt.Println(esc.Id, esc.Nombre, esc.Id_condominio, esc.Estado, esc.Descripcion)
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
func (Ec EspacioComun) UpdateEspaciosComunesByID(nombre string, estado string, descripcion string, id_ec string) (espacios_comunes []EspacioComun, err error) {
    // Opens DB
    db := GetConnection()

    // SQL queryGetPorCond/2/
	rows, err := db.Query("update espacios_comunes set nombre =?, estado= ?, descripcion = ? where id = ?", nombre, estado, descripcion, id_ec)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var esc EspacioComun
        rows.Scan(&esc.Id, &esc.Nombre, &esc.Id_condominio, &esc.Estado, &esc.Descripcion)
        fmt.Println(esc.Id, esc.Nombre, esc.Id_condominio, esc.Estado, esc.Descripcion)
        espacios_comunes = append(espacios_comunes, esc)
	}
	defer rows.Close()

	return
}

func GetUpdateEspaciosComunesByID(c *gin.Context){
    //URL parameters
    var nombre = c.Param("nombre")
    var estado = c.Param("estado")
    var id_ec = c.Param("id_ec")
    var descripcion = c.Param("descripcion")

    fmt.Println(nombre);
    fmt.Println(estado);
    fmt.Println(id_ec);
    fmt.Println(descripcion);
    // Results container
    ec := EspacioComun{}
    // Fetch from database
	espacios_comunes, err := ec.UpdateEspaciosComunesByID(nombre, estado, descripcion, id_ec)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": espacios_comunes,
		"count": len(espacios_comunes),
	})
}

// // // // INSERT ESPACIOSCOMUNES BY ID// // // //
func (Ec EspacioComun) InsertEspaciosComunes(nombre string, idcond string, estado string, descripcion string) (espacios_comunes []EspacioComun, err error) {
    // Opens DB
    db := GetConnection()

    // SQL queryGetPorCond/2/
	rows, err := db.Query("insert into espacios_comunes(nombre, id_condominio, estado, descripcion) values (?, ?, ?, ?)",nombre, idcond, estado, descripcion)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var esc EspacioComun
        rows.Scan(&esc.Id, &esc.Nombre, &esc.Id_condominio, &esc.Estado, &esc.Descripcion)
        fmt.Println(esc.Id, esc.Nombre, esc.Id_condominio, esc.Estado, esc.Descripcion)
        espacios_comunes = append(espacios_comunes, esc)
	}
	defer rows.Close()

	return
}

func GetInsertEspaciosComunes(c *gin.Context){
    //URL parameters
    var nombre = c.Param("nombre")
    var estado = c.Param("estado")
    var id_condominio = c.Param("id_condominio")
    var descripcion = c.Param("descripcion")

    fmt.Println(nombre);
    fmt.Println(id_condominio);
    fmt.Println(estado);
    fmt.Println(descripcion);

    // Results container
    ec := EspacioComun{}
    // Fetch from database
	espacios_comunes, err := ec.InsertEspaciosComunes(nombre, id_condominio, estado, descripcion)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": espacios_comunes,
		"count": len(espacios_comunes),
	})
}

// // // // TAKE ALL ESPACIOSCOMUNES POR ID // // // //
func (Ec EspacioComun) FetchEspaciosComunesByID(idesp string) (espacios_comunes []EspacioComun, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("SELECT * FROM espacios_comunes where id = ?", idesp)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var esc EspacioComun
        rows.Scan(&esc.Id, &esc.Nombre, &esc.Id_condominio, &esc.Estado, &esc.Descripcion)
        fmt.Println(esc.Id, esc.Nombre, esc.Id_condominio, esc.Estado, esc.Descripcion)
        espacios_comunes = append(espacios_comunes, esc)
	}
	defer rows.Close()

	return
}

func GetEspaciosComunesByID(c *gin.Context){
    //URL parameters
    var idesp = c.Param("idesp")

    fmt.Println(idesp);
    // Results container
    ec := EspacioComun{}
    // Fetch from database
	espacios_comunes, err := ec.FetchEspaciosComunesByID(idesp)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": espacios_comunes,
		"count": len(espacios_comunes),
	})
}
