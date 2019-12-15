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
}

type Espacio_comun struct{
    Id int `json:"id"`
    Nombre string `json:"nombre" form:"nombre"`
    Id_condominio int `json:"id_condominio" form:"id_condominio"`
    Estado string `json:"estado" form:"estado"`
}

type Estacionamiento struct{
    Id int `json:"id"`
    Numero int `json:"numero" form:"numero"`
    Id_departamentos int `json:"id_departamentos" form:"id_departamentos"`
    Estado string `json:"estado" form:"estado"`
    Patente_frecuente string `json:"patente_frecuente" form:"patente_frecuente"`
}

type Gasto_comun struct{
    Id int `json:"id"`
    Monto int `json:"monto" form:"monto"`
    Detalle string `json:"detalle" form:"detalle"`
    Fecha string `json:"fecha" form:"fecha"`
    Id_departamentos int `json:"id_departamentos" form:"id_departamentos"`
}

type Medicion_agua struct{
    Id int `json:"id"`
    Fecha string `json:"fecha" form:"fecha"`
    Litros float64 `json:"monto" form:"monto"`
    Id_departamentos int `json:"id_departamentos" form:"id_departamentos"`
}

type Multa struct{
    Id int `json:"id"`
    Grado int `json:"grado" form:"grado"`
    Id_departamentos int `json:"id_departamentos" form:"id_departamentos"`
    Monto int `json:"monto" form:"monto"`
    Fecha string `json:"fecha" form:"fecha"`
    Causa string `json:"causa" form:"causa"`
}

type Pagos_gastos_comunes struct{
    Id int `json:"id"`
    Monto int `json:"monto" form:"monto"`
    Fecha string `json:"fecha" form:"fecha"`
    Id_departamento int `json:"id_departamentos" form:"id_departamentos"`
}

type Ticket struct{
    Id int `json:"id"`
    Id_departamentos int `json:"id_departamentos" form:"id_departamentos"`
    Id_usuarios int `json:"id_usuarios" form:"id_usuarios"`
    Consulta string `json:"consulta" form:"consulta"`
    Respuesta string `json:"respuesta" form:"respuesta"`
    Finalizado int `json:"finalizado" form:"finalizado"`
}

type Usuario struct{
    Id int `json:"id"`
    Rut string `json:"rut" form:"rut"`
    Password string `json:"password" form:"password"`
    Level string `json:"level" form:"level"`
    Token string `json:"token" form:"token"`
}

type Usuarios_condominios struct{
    Id int `json:"id"`
    Id_usuarios int `json:"id_usuarios"`
    Id_condominio int `json:"id_condominio"`
}
