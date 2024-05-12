package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CreateDiskPage(c *gin.Context) {
	// Acceder a la sesión
	session := sessions.Default(c)
	email := session.Get("email")
	rol := session.Get("rol")

	if rol != "Administrador" {
		// Si el usuario no está autenticado, redirige a la página de inicio de sesión
		c.Redirect(http.StatusFound, "/login")
		return
	}

	hosts, _ := consultarHosts(email.(string))

	c.HTML(http.StatusOK, "createDisk.html", gin.H{
		"email": email,
		"hosts": hosts,
	})
}

func CreateDisk(c *gin.Context) {
	// Definir la URL del servidor
	serverURL := "http://localhost:8081/json/addDisk"

	// Obtener los datos del formulario
	nameDisk := c.PostForm("nameDisk")
	rutaDisk := c.PostForm("rutaDisk")
	osDisk := c.PostForm("osDisk")
	distriDisk := c.PostForm("distriDisk")
	arquiDiskStr := c.PostForm("arquiDisk")
	arquiDisk, _ := strconv.Atoi(arquiDiskStr)
	idHostDiskStr := c.PostForm("idHostDisk")
	idHostDisk, _ := strconv.Atoi(idHostDiskStr)

	// Crear un objeto Host con los datos del formulario
	disco := Disco{
		Nombre:                         nameDisk,
		Ruta_ubicacion:                 rutaDisk,
		Sistema_operativo:              osDisk,
		Distribucion_sistema_operativo: distriDisk,
		arquitectura:                   arquiDisk,
		Host_id:                        idHostDisk,
	}

	// Serializar el objeto host como JSON
	jsonData, err := json.Marshal(disco)
	if err != nil {
		// Manejar el error, por ejemplo, responder con un error HTTP
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al serializar el objeto Host"})
		return
	}

	// Crea una solicitud HTTP POST con el JSON como cuerpo
	req, err := http.NewRequest("POST", serverURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return
	}

	// Establece el encabezado de tipo de contenido
	req.Header.Set("Content-Type", "application/json")

	// Realiza la solicitud HTTP
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	// Verificar el código de estado de la respuesta
	if resp.StatusCode != http.StatusOK {
		// Manejar el error
		c.JSON(resp.StatusCode, gin.H{"error": "Error en la respuesta del servidor"})
		return
	}

	// Responder con una confirmación o redirección si es necesario
	c.HTML(http.StatusOK, "createDisk.html", nil)
}

func consultarHosts(email string) ([]Host, error) {
	serverURL := "http://localhost:8081/json/consultHost" // Cambia esto por la URL de tu servidor en el puerto 8081

	persona := Persona{Email: email}
	jsonData, err := json.Marshal(persona)
	if err != nil {
		return nil, err
	}

	// Crea una solicitud HTTP POST con el JSON como cuerpo
	req, err := http.NewRequest("POST", serverURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	// Establece el encabezado de tipo de contenido
	req.Header.Set("Content-Type", "application/json")

	// Realiza la solicitud HTTP
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Verifica la respuesta del servidor (resp.StatusCode) aquí si es necesario
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("La solicitud al servidor no fue exitosa")
	}

	// Lee la respuesta del cuerpo de la respuesta HTTP
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var hosts []Host

	// Decodifica los datos de respuesta en la variable machines.
	if err := json.Unmarshal(responseBody, &hosts); err != nil {
		// Maneja el error de decodificación aquí
	}

	return hosts, nil
}
