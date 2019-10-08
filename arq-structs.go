package main

type Condominio struct{
    Id int `json:"id"`
    Nombre string `json:"nombre" form:"nombre"`
    Ubicacion string `json:"ubicacion" form:"ubicacion"`
}

type Bodega struct{
    Id int `json:"id"`
    Numero int `json:"numero" form:"numero"`
    Num_dpto int `json:"num_dpto" form:"num_dpto"`
}

type Departamento struct{
    Id int `json:"id"`
    Numero int `json:"numero" form:"numero"`
    Id_condominio int `json:"id_condominio" form:"id_condominio"`
    Nombre_dueño string `json:"nombre_dueño" form:"nombre_dueño"`
    Nombre_arrendatario string `json:"nombre_arrendatario" form:"nombre_arrendatario"`
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
