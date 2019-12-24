package main

type Bodega struct{
    Id int `json:"id"`
    Numero int `json:"numero" form:"numero"`
    Id_departamentos int `json:"id_departamentos" form:"id_departamentos"`
    Estado string `json:"estado" form:"estado"`
}

type Condominio struct{
    Id int `json:"id"`
    Codigo string `json:"codigo"`
    Nombre string `json:"nombre" form:"nombre"`
    Ubicacion string `json:"ubicacion" form:"ubicacion"`
}

type Conserje struct{
    Id int `json:"id"`
    Nombre string `json:"nombre" form:"nombre"`
    Rut string `json:"rut" form:"rut"`
    Id_condominio int `json:"id_condominio" form:"id_condominio"`
}

type Departamento struct{
    Id int `json:"id"`
    Numero int `json:"numero" form:"numero"`
    Password string `json:"password" form:"password"`
    Dueno string `json:"dueno" form:"dueno"`
    Residente string `json:"residente" form:"residente"`
    Telefono string `json:"telefono" form:"telefono"`
    Correo string `json:"correo" form:"correo"`
    Id_condominio string `json:"id_condominio" form:"id_condominio"`
    Telefono_residente string `json:"telefono_residente" form:"telefono_residente"`
    Correo_residente string `json:"correo_residente" form:"correo_residente"`
    Prorrateo float64 `json:"prorrateo" form:"prorrateo"`
    Level string `json:"level" form:"level"`
    Token string `json:"token" form:"token"`
}

type DepartamentoAllData struct{
    Id int `json:"id"`
    Numero int `json:"numero" form:"numero"`
    Password string `json:"password" form:"password"`
    Dueno string `json:"dueno" form:"dueno"`
    Residente string `json:"residente" form:"residente"`
    Telefono string `json:"telefono" form:"telefono"`
    Correo string `json:"correo" form:"correo"`
    Id_condominio string `json:"id_condominio" form:"id_condominio"`
    Telefono_residente string `json:"telefono_residente" form:"telefono_residente"`
    Correo_residente string `json:"correo_residente" form:"correo_residente"`
    Prorrateo float64 `json:"prorrateo" form:"prorrateo"`
    Level string `json:"level" form:"level"`
    N_bodega string `json:"n_bodega" form:"n_bodega"`
    N_estacionamiento string `json:"n_estacionamiento" form:"n_estacionamiento"`

}

type EspacioComun struct{
    Id int `json:"id"`
    Nombre string `json:"nombre" form:"nombre"`
    Id_condominio int `json:"id_condominio" form:"id_condominio"`
    Estado string `json:"estado" form:"estado"`
    Descripcion string `json:"descripcion" form:"descripcion"`
}

type EspacioComunCond struct{
    Id int `json:"id"`
    Nombre string `json:"nombre" form:"nombre"`
    Id_condominio int `json:"id_condominio" form:"id_condominio"`
    Estado string `json:"estado" form:"estado"`
    Descripcion string `json:"descripcion" form:"descripcion"`
    Codigo_cond string `json:"codigo_cond" form:"codigo_cond"`
    Nombre_cond string `json:"nombre_cond" form:"nombre_cond"`
}

type Estacionamiento struct{
    Id int `json:"id"`
    Numero int `json:"numero" form:"numero"`
    Id_departamentos int `json:"id_departamentos" form:"id_departamentos"`
    Estado string `json:"estado" form:"estado"`
    Patente_frecuente string `json:"patente_frecuente" form:"patente_frecuente"`
}

type GastoComun struct{
    Id int `json:"id"`
    Monto int `json:"monto" form:"monto"`
    Detalle string `json:"detalle" form:"detalle"`
    Fecha string `json:"fecha" form:"fecha"`
    Id_departamentos int `json:"id_departamentos" form:"id_departamentos"`
}

type SumaGastoComun struct{
    Suma int `json:"suma"`
}

type MedicionAgua struct{
    Id int `json:"id"`
    Fecha string `json:"fecha" form:"fecha"`
    Litros float64 `json:"litros" form:"litros"`
    Id_departamentos int `json:"id_departamentos" form:"id_departamentos"`
}

type MedicionAguaDpto struct{
    Id int `json:"id"`
    Fecha string `json:"fecha" form:"fecha"`
    Litros float64 `json:"litros" form:"litros"`
    Id_departamentos int `json:"id_departamentos" form:"id_departamentos"`
    Num_dpto string `json:"num_dpto" form:"num_dpto"`
    Dueno_dpto string `json:"dueno_dpto" form:"dueno_dpto"`
    Residente_dpto string `json:"residente_dpto" form:"residente_dpto"`
    Correo_dueno string `json:"correo_dueno" form:"correo_dueno"`
    Correo_residente string `json:"correo_residente" form:"correo_residente"`
    Telefono_dueno string `json:"telefono_dueno" form:"telefono_dueno"`
    Telefono_residente string `json:"telefono_residente" form:"telefono_residente"`

}

type SumaMedicionAgua struct{
    Suma float64 `json:"suma"`
}


type Multa struct{
    Id int `json:"id"`
    Grado int `json:"grado" form:"grado"`
    Id_departamentos int `json:"id_departamentos" form:"id_departamentos"`
    Monto int `json:"monto" form:"monto"`
    Fecha string `json:"fecha" form:"fecha"`
    Causa string `json:"causa" form:"causa"`

}

type MultaDpto struct{
    Id int `json:"id"`
    Grado int `json:"grado" form:"grado"`
    Id_departamentos int `json:"id_departamentos" form:"id_departamentos"`
    Monto int `json:"monto" form:"monto"`
    Fecha string `json:"fecha" form:"fecha"`
    Causa string `json:"causa" form:"causa"`
    Num_dpto string `json:"num_dpto" form:"num_dpto"`
    Dueno_dpto string `json:"dueno_dpto" form:"dueno_dpto"`
    Residente_dpto string `json:"residente_dpto" form:"residente_dpto"`
    Correo_dueno string `json:"correo_dueno" form:"correo_dueno"`
    Correo_residente string `json:"correo_residente" form:"correo_residente"`
    Telefono_dueno string `json:"telefono_dueno" form:"telefono_dueno"`
    Telefono_residente string `json:"telefono_residente" form:"telefono_residente"`
}

type PagosGastosComunes struct{
    Id int `json:"id"`
    Monto int `json:"monto" form:"monto"`
    Fecha string `json:"fecha" form:"fecha"`
    Id_departamento int `json:"id_departamentos" form:"id_departamentos"`

}

type PagosGastosComunesDpto struct{
    Id int `json:"id"`
    Monto int `json:"monto" form:"monto"`
    Fecha string `json:"fecha" form:"fecha"`
    Id_departamento int `json:"id_departamentos" form:"id_departamentos"`
    Num_dpto string `json:"num_dpto" form:"num_dpto"`
    Dueno_dpto string `json:"dueno_dpto" form:"dueno_dpto"`
    Residente_dpto string `json:"residente_dpto" form:"residente_dpto"`
    Correo_dueno string `json:"correo_dueno" form:"correo_dueno"`
    Correo_residente string `json:"correo_residente" form:"correo_residente"`
    Telefono_dueno string `json:"telefono_dueno" form:"telefono_dueno"`
    Telefono_residente string `json:"telefono_residente" form:"telefono_residente"`
}

type SumaPagosGastosComunes struct{
    Suma int `json:"suma"`
}

type Ticket struct{
    Id int `json:"id"`
    Id_departamentos int `json:"id_departamentos" form:"id_departamentos"`
    Id_usuarios int `json:"id_usuarios" form:"id_usuarios"`
    Consulta string `json:"consulta" form:"consulta"`
    Respuesta string `json:"respuesta" form:"respuesta"`
    Finalizado int `json:"finalizado" form:"finalizado"`
    Asunto string `json:"asunto" form:"asunto"`
    Fecha string `json:"fecha" form:"fecha"`

}

type Usuario struct{
    Id int `json:"id"`
    Rut string `json:"rut" form:"rut"`
    Password string `json:"password" form:"password"`
    Nombre string `json:"nombre" form:"nombre"`
    Level string `json:"level" form:"level"`
    Token string `json:"token" form:"token"`
}

type Usuarios_condominios struct{
    Id int `json:"id"`
    Id_usuarios int `json:"id_usuarios"`
    Id_condominio int `json:"id_condominio"`
}
