package main

type Condominio struct{
    Id int `json:"id"`
    Nombre string `json:"nombre" form:"nombre"`
    Ubicacion string `json:"ubicacion" form:"ubicacion"`
}
