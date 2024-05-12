package handlers

type Persona struct {
	Nombre      string
	Apellido    string
	Email       string
	Contrasenia string
	Rol         string
}

type Host struct {
	Id                             int
	Nombre                         string
	Mac                            string
	Ip                             string
	Hostname                       string
	Ram_total                      int
	Cpu_total                      int
	Almacenamiento_total           int
	Ram_usada                      int
	Cpu_usada                      int
	Almacenamiento_usado           int
	Adaptador_red                  string
	Estado                         string
	Ruta_llave_ssh_pub             string
	Sistema_operativo              string
	Distribucion_sistema_operativo string
}

type Disco struct {
	Id                             int
	Nombre                         string
	Ruta_ubicacion                 string
	Sistema_operativo              string
	Distribucion_sistema_operativo string
	arquitectura                   int
	Host_id                        int
}

type Maquina_virtual struct {
	Uuid                           string
	Nombre                         string
	Ram                            int
	Cpu                            int
	Ip                             string
	Estado                         string
	Hostname                       string
	Persona_email                  string
	Host_id                        int
	Disco_id                       int
	Sistema_operativo              string
	Distribucion_sistema_operativo string
}

type Conetendor struct {
	ConetendorId string
	Imagen       string
	Comando      string
	Creado       string
	Status       string
	Puerto       string
	Nombre       string
}

type Imagen struct {
	Repositorio string
	Tag         string
	ImagenId    string
	Creacion    string
	Tamanio     string
}
