package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
)

// // // // TAKE ALL CONDOMINIOS // // // //
func (co Condominio) FetchCondominios() (condominios []Condominio, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("SELECT id, codigo, nombre, ubicacion FROM condominios")
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var cond Condominio
        rows.Scan(&cond.Id, &cond.Codigo, &cond.Nombre, &cond.Ubicacion)
        fmt.Println(cond.Id, cond.Codigo, cond.Nombre, cond.Ubicacion)
        condominios = append(condominios, cond)
	}
	defer rows.Close()

	return
}

func GetCondominios(c *gin.Context){
    // Results container
    co := Condominio{}
    // Fetch from database
	condominios, err := co.FetchCondominios()
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": condominios,
		"count": len(condominios),
	})
}

// // // // TAKE ALL CONDOMINIOS WITH ID=id_condominios // // // //
func (co Condominio) FetchCondominiosByID(id_condominios string) (condominios []Condominio, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("SELECT id, codigo, nombre, ubicacion FROM condominios WHERE id = ?", id_condominios)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var cond Condominio
        rows.Scan(&cond.Id, &cond.Codigo, &cond.Nombre, &cond.Ubicacion)
        fmt.Println(cond.Id, cond.Codigo, cond.Nombre, cond.Ubicacion)
        condominios = append(condominios, cond)
	}
	defer rows.Close()

	return
}

func GetCondominiosByID(c *gin.Context){
    // Get URL parameters
    var id = c.Param("id")

    // Results container
    co := Condominio{}

    // Fetch from database
	condominios, err := co.FetchCondominiosByID(id)
	if err != nil {
        panic(err.Error())
	}

    // Render via GET method
	c.JSON(200, gin.H{
		"rows": condominios,
		"count": len(condominios),
	})
}

// // // // TAKE ALL BODEGAS WITH USUARIO_ID=id_departamentos // // // //
func (co Condominio) FetchCondominiosByUserID(id_usr string) (condominios []Condominio, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("SELECT id, codigo, nombre, ubicacion FROM condominios WHERE id in (select id_condominio from usuarios_condominios where id_usuarios = ?)", id_usr)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var cond Condominio
        rows.Scan(&cond.Id, &cond.Codigo, &cond.Nombre, &cond.Ubicacion)
        fmt.Println(cond.Id, cond.Codigo, cond.Nombre, cond.Ubicacion)
        condominios = append(condominios, cond)
	}
	defer rows.Close()

	return
}

func GetCondominiosByUserID(c *gin.Context){
    // Get URL parameters
    var id_usr = c.Param("id_usr")

    // Results container
    co := Condominio{}

    // Fetch from database
	condominios, err := co.FetchCondominiosByUserID(id_usr)
	if err != nil {
        panic(err.Error())
	}

    // Render via GET method
	c.JSON(200, gin.H{
		"rows": condominios,
		"count": len(condominios),
	})
}
