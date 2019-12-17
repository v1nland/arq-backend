package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
)

// // // // TAKE ALL BODEGAS // // // //
func (bo Bodega) FetchBodegas() (bodegas []Bodega, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("SELECT id, numero, id_departamentos, estado FROM bodegas")
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var bod Bodega
        rows.Scan(&bod.Id, &bod.Numero, &bod.Id_departamentos, &bod.Estado)
        fmt.Println(bod.Id, bod.Numero, bod.Id_departamentos, bod.Estado)
        bodegas = append(bodegas, bod)
	}
	defer rows.Close()

	return
}

func GetBodegas(c *gin.Context){
    // Results container
    bo := Bodega{}
    // Fetch from database
	bodegas, err := bo.FetchBodegas()
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": bodegas,
		"count": len(bodegas),
	})
}

// // // // TAKE ALL BODEGAS WITH DPTO_ID=id_departamentos // // // //
func (bo Bodega) FetchBodegasByDptoID(id_dpto string) (bodegas []Bodega, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select * from bodegas where id_departamentos=?", id_dpto)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var bod Bodega
        rows.Scan(&bod.Id, &bod.Numero, &bod.Id_departamentos, &bod.Estado)
        fmt.Println(bod.Id, bod.Numero, bod.Id_departamentos, bod.Estado)
        bodegas = append(bodegas, bod)
	}
	defer rows.Close()

	return
}

func GetBodegasByDptoID(c *gin.Context){
    // URL parameters
    var id_dpto = c.Param("id_dpto")

    fmt.Println(id_dpto);
    // Results container
    bo := Bodega{}
    // Fetch from database
	bodegas, err := bo.FetchBodegasByDptoID(id_dpto)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": bodegas,
		"count": len(bodegas),
	})
}
