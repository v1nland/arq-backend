package main
import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
)

// // // // TAKE ALL DEPARTAMENTOS // // // //
func (de DepartamentoAllData) FetchDepartamentos() (departamentos []DepartamentoAllData, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("SELECT departamentos.id, departamentos.numero, departamentos.password, departamentos.dueno, departamentos.residente, departamentos.telefono, departamentos.correo, departamentos.id_condominio, departamentos.telefono_residente, departamentos.correo_residente, departamentos.prorrateo, coalesce(estacionamientos.numero, -1) as n_estacionamientos, coalesce(bodegas.numero, -1) as n_bodegas FROM departamentos left outer join estacionamientos on departamentos.id = estacionamientos.id_departamentos left outer join bodegas on estacionamientos.id_departamentos = bodegas.id_departamentos")
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var dpto DepartamentoAllData
        rows.Scan(&dpto.Id, &dpto.Numero, &dpto.Password, &dpto.Dueno, &dpto.Residente, &dpto.Telefono, &dpto.Correo, &dpto.Id_condominio, &dpto.Telefono_residente, &dpto.Correo_residente, &dpto.Prorrateo, &dpto.N_bodega, &dpto.N_estacionamiento)
        dpto.Level = "user"
        fmt.Println(dpto.Id, dpto.Numero, dpto.Password, dpto.Dueno, dpto.Residente, dpto.Telefono, dpto.Correo, dpto.Id_condominio, dpto.Telefono_residente, dpto.Correo_residente, dpto.Prorrateo, dpto.N_bodega, dpto.N_estacionamiento)
        departamentos = append(departamentos, dpto)
	}
	defer rows.Close()

	return
}

func GetDepartamentos(c *gin.Context){
    // Results container
    de := DepartamentoAllData{}
    // Fetch from database
	departamentos, err := de.FetchDepartamentos()
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": departamentos,
		"count": len(departamentos),
	})
}

// // // // LOGIN DEPARTAMENTOS // // // //
func (de Departamento) FetchDptoLogin(cond_code string, dpto_num string, dpto_pass string) (departamentos []Departamento, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select * from departamentos where (id_condominio = (select id from condominios where codigo = ?)) and numero= ? and password = ?", cond_code, dpto_num, dpto_pass)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var dpto Departamento
        rows.Scan(&dpto.Id, &dpto.Numero, &dpto.Password, &dpto.Dueno, &dpto.Residente, &dpto.Telefono, &dpto.Correo, &dpto.Id_condominio, &dpto.Telefono_residente, &dpto.Correo_residente, &dpto.Prorrateo)

        token := jwt.New(jwt.SigningMethodHS256)
        claims := token.Claims.(jwt.MapClaims)
        claims["id"] = dpto.Id
        claims["numero"] = dpto.Numero
        claims["password"] = dpto.Password
        claims["dueno"] = dpto.Dueno
        claims["residente"] = dpto.Residente
        claims["telefono"] = dpto.Telefono
        claims["correo"] = dpto.Correo
        claims["id_cond"] = dpto.Id_condominio
        claims["telefono_residente"] = dpto.Telefono_residente
        claims["correo_residente"] = dpto.Correo_residente
        claims["prorrateo"] = dpto.Prorrateo
        claims["level"] = "user"
        dpto_token, err := token.SignedString( SecretKey() )
        dpto.Token = dpto_token
        dpto.Level = "user"

        _ = err

        fmt.Println(dpto.Id, dpto.Numero, dpto.Password, dpto.Dueno, dpto.Residente, dpto.Telefono, dpto.Correo, dpto.Id_condominio, dpto.Telefono_residente, dpto.Correo_residente, dpto.Prorrateo, dpto.Token)
        departamentos = append(departamentos, dpto)
	}
	defer rows.Close()

	return
}

func GetDptoLogin(c *gin.Context){
    // URL parameters
    var codigo = c.Param("codigo")
    var numero = c.Param("numero")
    var password = c.Param("password")

    fmt.Println(codigo);
    fmt.Println(numero);
    fmt.Println(password);

    // Results container
    de := Departamento{}
    // Fetch from database
	departamentos, err := de.FetchDptoLogin(codigo, numero, password)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": departamentos,
		"count": len(departamentos),
	})
}

// // // // SELECT DEPARTAMENTOS BY ID_COND=id_cond // // // //
func (de Departamento) FetchDptoByCondominioID(id_cond string) (departamentos []Departamento, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select * from departamentos where id_condominio = ?", id_cond)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var dpto Departamento
        rows.Scan(&dpto.Id, &dpto.Numero, &dpto.Password, &dpto.Dueno, &dpto.Residente, &dpto.Telefono, &dpto.Correo, &dpto.Id_condominio, &dpto.Telefono_residente, &dpto.Correo_residente, &dpto.Prorrateo)
        fmt.Println(dpto.Id, dpto.Numero, dpto.Password, dpto.Dueno, dpto.Residente, dpto.Telefono, dpto.Correo, dpto.Id_condominio, dpto.Telefono_residente, dpto.Correo_residente, dpto.Prorrateo)
        departamentos = append(departamentos, dpto)
	}
	defer rows.Close()

	return
}

func GetDptoByCondominioID(c *gin.Context){
    // URL parameters
    var id_cond = c.Param("id_cond")

    fmt.Println(id_cond);

    // Results container
    de := Departamento{}
    // Fetch from database
	departamentos, err := de.FetchDptoByCondominioID(id_cond)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": departamentos,
		"count": len(departamentos),
	})
}

// // // // SELECT DEPARTAMENTOS BY ID=id_dpto // // // //
func (de Departamento) FetchDptoByID(id_dpto string) (departamentos []Departamento, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("select * from departamentos where id= ?", id_dpto)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var dpto Departamento
        rows.Scan(&dpto.Id, &dpto.Numero, &dpto.Password, &dpto.Dueno, &dpto.Residente, &dpto.Telefono, &dpto.Correo, &dpto.Id_condominio, &dpto.Telefono_residente, &dpto.Correo_residente, &dpto.Prorrateo)
        fmt.Println(dpto.Id, dpto.Numero, dpto.Password, dpto.Dueno, dpto.Residente, dpto.Telefono, dpto.Correo, dpto.Id_condominio, dpto.Telefono_residente, dpto.Correo_residente, dpto.Prorrateo)
        departamentos = append(departamentos, dpto)
	}
	defer rows.Close()

	return
}

func GetDptoByID(c *gin.Context){
    // URL parameters
    var id_dpto = c.Param("id_dpto")

    fmt.Println(id_dpto);

    // Results container
    de := Departamento{}
    // Fetch from database
	departamentos, err := de.FetchDptoByID(id_dpto)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": departamentos,
		"count": len(departamentos),
	})
}

// // // // SELECT DEPARTAMENTOS BY ID=id_dpto AND bodegas AND estacionamientos // // // //
func (de DepartamentoAllData) FetchDepartamentosAllDataByID(id_dpto string) (departamentos []DepartamentoAllData, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("SELECT departamentos.id, departamentos.numero, departamentos.password, departamentos.dueno, departamentos.residente, departamentos.telefono, departamentos.correo, departamentos.id_condominio, departamentos.telefono_residente, departamentos.correo_residente, departamentos.prorrateo, coalesce(estacionamientos.numero, -1) as n_estacionamientos, coalesce(bodegas.numero, -1) as n_bodegas FROM departamentos left outer join estacionamientos on departamentos.id = estacionamientos.id_departamentos left outer join bodegas on estacionamientos.id_departamentos = bodegas.id_departamentos where departamentos.id = ?", id_dpto)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var dpto DepartamentoAllData
        rows.Scan(&dpto.Id, &dpto.Numero, &dpto.Password, &dpto.Dueno, &dpto.Residente, &dpto.Telefono, &dpto.Correo, &dpto.Id_condominio, &dpto.Telefono_residente, &dpto.Correo_residente, &dpto.Prorrateo, &dpto.N_bodega, &dpto.N_estacionamiento)
        dpto.Level = "user"
        fmt.Println(dpto.Id, dpto.Numero, dpto.Password, dpto.Dueno, dpto.Residente, dpto.Telefono, dpto.Correo, dpto.Id_condominio, dpto.Telefono_residente, dpto.Correo_residente, dpto.Prorrateo, dpto.N_bodega, dpto.N_estacionamiento)
        departamentos = append(departamentos, dpto)
	}
	defer rows.Close()

	return
}

func GetDepartamentosAllDataByID(c *gin.Context){
    // URL parameters
    var id_dpto = c.Param("id_dpto")

    fmt.Println(id_dpto);
    // Results container
    de := DepartamentoAllData{}
    // Fetch from database
	departamentos, err := de.FetchDepartamentosAllDataByID(id_dpto)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": departamentos,
		"count": len(departamentos),
	})
}

// // // UPDATE DEPARTAMENTO CON ID_DPTO // // //
func (de Departamento) UpdateDpto(password string, dueno string, residente string, telefono string, correo string, id_condominio string, telefono_residente string, correo_residente string, prorrateo string, id_dpto string) (departamentos []Departamento, err error) {
    // Opens DB
    db := GetConnection()

    // SQL query
	rows, err := db.Query("update departamentos set password = ?, dueno = ?, residente = ?, telefono = ?, correo = ?, id_condominio= ?, telefono_residente = ?, correo_residente = ?, prorrateo = ? where id = ?", password, dueno, residente, telefono, correo, id_condominio, telefono_residente, correo_residente, prorrateo, id_dpto)
	if err != nil {
		return
	}

    // Take all data
	for rows.Next() {
		var dpto Departamento
        rows.Scan(&dpto.Id, &dpto.Numero, &dpto.Password, &dpto.Dueno, &dpto.Residente, &dpto.Telefono, &dpto.Correo, &dpto.Id_condominio, &dpto.Telefono_residente, &dpto.Correo_residente, &dpto.Prorrateo)
        dpto.Level = "user"
        fmt.Println(dpto.Id, dpto.Numero, dpto.Password, dpto.Dueno, dpto.Residente, dpto.Telefono, dpto.Correo, dpto.Id_condominio, dpto.Telefono_residente, dpto.Correo_residente, dpto.Prorrateo)
        departamentos = append(departamentos, dpto)
	}
	defer rows.Close()

	return
}

func GetUpdateDpto(c *gin.Context){
    // URL parameters
    var password = c.Param("password")
    var dueno = c.Param("dueno")
    var residente = c.Param("residente")
    var telefono = c.Param("telefono")
    var correo = c.Param("correo")
    var id_condominio = c.Param("id_condominio")
    var telefono_residente = c.Param("telefono_residente")
    var correo_residente = c.Param("correo_residente")
    var prorrateo = c.Param("prorrateo")
    var iddpto = c.Param("iddpto")

    fmt.Println(password);
    fmt.Println(dueno);
    fmt.Println(residente);
    fmt.Println(telefono);
    fmt.Println(correo);
    fmt.Println(id_condominio);
    fmt.Println(telefono_residente);
    fmt.Println(correo_residente);
    fmt.Println(prorrateo);
    fmt.Println(iddpto);


    // Results container
    de := Departamento{}
    // Fetch from database
	departamentos, err := de.UpdateDpto(password, dueno, residente, telefono, correo, id_condominio, telefono_residente, correo_residente, prorrateo, iddpto)
	if err != nil {
        panic(err.Error())
	}

    // Show via GET method
	c.JSON(200, gin.H{
		"rows": departamentos,
		"count": len(departamentos),
	})
}
