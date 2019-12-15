package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
)

//Take all data----------------------------------------------------------------------------------

func (Ec Espacio_comun) FetchEspacios_comunes() (espacios_comunes []Espacio_comun, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("SELECT * FROM espacios_comunes")
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var esc Espacio_comun
        rows.Scan(&esc.Id, &esc.Nombre, &esc.Id_condominio, &esc.Estado)
        fmt.Println(esc.Id, esc.Nombre, esc.Id_condominio, esc.Estado)
        espacios_comunes = append(espacios_comunes, esc)
	}
	defer rows.Close()

	return
}

func GetEspacios_comunes(c *gin.Context){
    // Results container
    ec := Espacio_comun{}
    // Fetch from database
	espacios_comunes, err := ec.FetchEspacios_comunes()
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": espacios_comunes,
		"count": len(espacios_comunes),
	})
}

//select espacios por id_condominio------------------------------------------------------------------

func (Ec Espacio_comun) FetchEspacios_comunesIdcond(id_condominios string) (espacios_comunes []Espacio_comun, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select * from espacios_comunes where id_condominio =?", id_condominios)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var esc Espacio_comun
        rows.Scan(&esc.Id, &esc.Nombre, &esc.Id_condominio, &esc.Estado)
        fmt.Println(esc.Id, esc.Nombre, esc.Id_condominio, esc.Estado)
        espacios_comunes = append(espacios_comunes, esc)
	}
	defer rows.Close()

	return
}

func GetEspacios_comunesIdcond(c *gin.Context){
    //URL parameters
    var idcond = c.Param("idcond")

    fmt.Println(idcond);
    // Results container
    ec := Espacio_comun{}
    // Fetch from database
	espacios_comunes, err := ec.FetchEspacios_comunesIdcond(idcond)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": espacios_comunes,
		"count": len(espacios_comunes),
	})
}

//update espacio comun con idcond, actualiza el estado----------------------------------------------

func (Ec Espacio_comun) UpdateEspacios_comunesIdcond(estado string, id_espacio string) (espacios_comunes []Espacio_comun, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("update espacios_comunes set estado= ? where id_espacio = ?", estado, id_espacio)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var esc Espacio_comun
        rows.Scan(&esc.Id, &esc.Nombre, &esc.Id_condominio, &esc.Estado)
        fmt.Println(esc.Id, esc.Nombre, esc.Id_condominio, esc.Estado)
        espacios_comunes = append(espacios_comunes, esc)
	}
	defer rows.Close()

	return
}

func GetUpdateEspacios_comunesIdcond(c *gin.Context){
    //URL parameters
    var estado = c.Param("estado")
    var idcond = c.Param("idcond")

    fmt.Println(estado);
    fmt.Println(idcond);
    // Results container
    ec := Espacio_comun{}
    // Fetch from database
	espacios_comunes, err := ec.UpdateEspacios_comunesIdcond(estado, idcond)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": espacios_comunes,
		"count": len(espacios_comunes),
	})
}
