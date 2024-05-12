package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func CreateHostPage(c *gin.Context) {
	// Acceder a la sesión
	session := sessions.Default(c)
	rol := session.Get("rol")

	if rol != "Administrador" {
		// Si el usuario no está autenticado, redirige a la página de inicio de sesión
		c.Redirect(http.StatusFound, "/login")
		return
	}

	c.HTML(http.StatusOK, "createHost.html", nil)
}

func CreateHost(c *gin.Context) {
	// Definir la URL del servidor
	serverURL := "http://localhost:8081/json/addHost"

	// Obtener los datos del formulario
	nombreHost := c.PostForm("nameHost")
	ipHost := c.PostForm("ipHost")
	macHost := c.PostForm("macHost")
	adapHost := c.PostForm("adapHost")
	soHost := c.PostForm("soHost")
	hostnameHost := c.PostForm("hostnameHost")
	ramHostStr := c.PostForm("ramHost")
	ramHost, _ := strconv.Atoi(ramHostStr)
	cpuHostStr := c.PostForm("cpuHost")
	cpuHost, _ := strconv.Atoi(cpuHostStr)
	almaceHostStr := c.PostForm("almaceHost")
	almaceHost, _ := strconv.Atoi(almaceHostStr)
	sshHost := c.PostForm("sshHost")

	// Crear un objeto Host con los datos del formulario
	host := Host{
		Nombre:               nombreHost,
		Ip:                   ipHost,
		Mac:                  macHost,
		Adaptador_red:        adapHost,
		Sistema_operativo:    soHost,
		Hostname:             hostnameHost,
		Ram_total:            ramHost,
		Cpu_total:            cpuHost,
		Almacenamiento_total: almaceHost,
		Ruta_llave_ssh_pub:   sshHost,
	}

	fmt.Println(host)

	// Serializar el objeto host como JSON
	jsonData, err := json.Marshal(host)
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
	c.HTML(http.StatusOK, "createHost.html", nil)
}
