package main

type Bodega struct{
    Id int `json:"id"`
    Numero int `json:"numero" form:"numero"`
    Num_dpto int `json:"num_dpto" form:"num_dpto"`
}

type Condominio struct{
    Id int `json:"id"`
    Codigo string `json:"codigo"`
    Nombre string `json:"nombre" form:"nombre"`
    Ubicacion string `json:"ubicacion" form:"ubicacion"`
}

type Departamento struct{
    Id int `json:"id"`
    Numero int `json:"numero" form:"numero"`
    Password int `json:"password" form:"password"`
    Dueno string `json:"dueno" form:"dueno"`
    Residente string `json:"residente" form:"residente"`
    Telefono string `json:"telefono" form:"telefono"`
    Correo string `json:"correo" form:"correo"`
    Cod_condominio string `json:"cod_condominio" form:"cod_condominio"`
}

type Estacionamiento struct{
    Id int `json:"id"`
    Numero int `json:"numero" form:"numero"`
    Num_dpto int `json:"num_dpto" form:"num_dpto"`
}

type Multa struct{
    Id int `json:"id"`
    Tipo string `json:"tipo" form:"tipo"`
    Num_dpto int `json:"num_dpto" form:"num_dpto"`
    Precio int `json:"precio" form:"precio"`
    Fecha string `json:"fecha" form:"fecha"`
}

type Usuario struct{
    Id int `json:"id"`
    Rut string `json:"rut" form:"rut"`
    Password string `json:"password" form:"password"`
}
