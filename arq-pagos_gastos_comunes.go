package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
)

//Take all data-------------------------------------------------------------------------------------

func (pag Pagos_gastos_comunes) FetchPagos() (pagos_gastos_comunes []Pagos_gastos_comunes, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("SELECT id, monto, fecha, id_departamentos FROM pagos_gastos_comunes")
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var pago Pagos_gastos_comunes
        rows.Scan(&pago.Id, &pago.Monto, &pago.Fecha, &pago.Id_departamento)
        fmt.Println(pago.Id, pago.Monto, pago.Fecha, pago.Id_departamento)
        pagos_gastos_comunes = append(pagos_gastos_comunes, pago)
	}
	defer rows.Close()

	return
}

func GetPagos(c *gin.Context){
    // Results container
    pag := Pagos_gastos_comunes{}
    // Fetch from database
	pagos_gastos_comunes, err := pag.FetchPagos()
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": pagos_gastos_comunes,
		"count": len(pagos_gastos_comunes),
	})
}


//Select con id_departamentos------------------------------------------------------------------------

func (pag Pagos_gastos_comunes) FetchPagosIddpto(id_dpto string) (pagos_gastos_comunes []Pagos_gastos_comunes, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select * from pagos_gastos_comunes where id_departamentos=?", id_dpto)
	if err != nil {
        return
    }

    // Take all data
	for rows.Next() {
		var pago Pagos_gastos_comunes
        rows.Scan(&pago.Id, &pago.Monto, &pago.Fecha, &pago.Id_departamento)
        fmt.Println(pago.Id, pago.Monto, pago.Fecha, pago.Id_departamento)
        pagos_gastos_comunes = append(pagos_gastos_comunes, pago)
	}
	defer rows.Close()

	return
}

func GetPagosIddpto(c *gin.Context){
    // URL parameters
    var iddepartamento = c.Param("iddepartamento")

    fmt.Println(iddepartamento);
    // Results container
    pag := Pagos_gastos_comunes{}
    // Fetch from database
	pagos_gastos_comunes, err := pag.FetchPagosIddpto(iddepartamento)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": pagos_gastos_comunes,
		"count": len(pagos_gastos_comunes),
	})
}

//select con id condominio---------------------------------------------------------------------------

func (pag Pagos_gastos_comunes) FetchPagosIdcondominio(id_condominio string) (pagos_gastos_comunes []Pagos_gastos_comunes, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select * from pagos_gastos_comunes where id_departamentos in (select id from departamentos where id_condominio = ?)", id_condominio)
	if err != nil {
        return
    }

    // Take all data
	for rows.Next() {
		var pago Pagos_gastos_comunes
        rows.Scan(&pago.Id, &pago.Monto, &pago.Fecha, &pago.Id_departamento)
        fmt.Println(pago.Id, pago.Monto, pago.Fecha, pago.Id_departamento)
        pagos_gastos_comunes = append(pagos_gastos_comunes, pago)
	}
	defer rows.Close()

	return
}

func GetPagosIdcondominio(c *gin.Context){
    // URL parameters
    var idcondominio = c.Param("idcondominio")

    fmt.Println(idcondominio);
    // Results container
    pag := Pagos_gastos_comunes{}
    // Fetch from database
	pagos_gastos_comunes, err := pag.FetchPagosIdcondominio(idcondominio)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": pagos_gastos_comunes,
		"count": len(pagos_gastos_comunes),
	})
}

//Select con id_condominio y mes------------------------------------------------------------------

func (pag Pagos_gastos_comunes) FetchPagosIdcondominiomes(id_condominio string, anoinicio string, mesinicio string, anofinal string, mesfinal string) (pagos_gastos_comunes []Pagos_gastos_comunes, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select * from pagos_gastos_comunes where id_departamentos in (select id from departamentos where id_condominio = ?) and fecha between '?-?-01 00:00:00' and '?-?-01 00:00:00", id_condominio, anoinicio, mesinicio, anofinal, mesfinal)
	if err != nil {
        return
    }

    // Take all data
	for rows.Next() {
		var pago Pagos_gastos_comunes
        rows.Scan(&pago.Id, &pago.Monto, &pago.Fecha, &pago.Id_departamento)
        fmt.Println(pago.Id, pago.Monto, pago.Fecha, pago.Id_departamento)
        pagos_gastos_comunes = append(pagos_gastos_comunes, pago)
	}
	defer rows.Close()

	return
}

func GetPagosIdcondominiomes(c *gin.Context){
    // URL parameters
    var idcondominio = c.Param("idcondominio")
    var anoinicio = c.Param("anoinicio")
    var mesinicio = c.Param("mesinicio")
    var anofinal = c.Param("anofinal")
    var mesfinal = c.Param("mesfinal")

    fmt.Println(idcondominio);
    fmt.Println(anoinicio);
    fmt.Println(mesinicio);
    fmt.Println(anofinal);
    fmt.Println(mesfinal);
    // Results container
    pag := Pagos_gastos_comunes{}
    // Fetch from database
	pagos_gastos_comunes, err := pag.FetchPagosIdcondominiomes(idcondominio, anoinicio, mesinicio, anofinal, mesfinal)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": pagos_gastos_comunes,
		"count": len(pagos_gastos_comunes),
	})
}


