package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GestionContenedores(c *gin.Context) {
	// Renderiza la plantilla HTML

	// Acceder a la sesión
	session := sessions.Default(c)
	email := session.Get("email")

	if email == nil {
		// Si el usuario no está autenticado, redirige a la página de inicio de sesión
		c.Redirect(http.StatusFound, "/login")
		return
	}

	// Recuperar o inicializar un arreglo de máquinas virtuales en la sesión del usuario
	machines, _ := MaquinasActualesC(email.(string))

	c.HTML(http.StatusOK, "gestionContenedores.html", gin.H{
		"email":    email,
		"machines": machines,
	})
}

func CrearContenedor(c *gin.Context) {

	serverURL := "http://localhost:8081/json/crearContenedor"

	// Acceder a la sesión
	session := sessions.Default(c)
	email := session.Get("email")

	maquinaVirtual := c.PostForm("maquinaVirtual")
	nombreImagen := c.PostForm("nombreImagen")
	comando := "docker run"

	// Dividir la cadena en IP y hostname
	partes := strings.Split(maquinaVirtual, " - ")
	if len(partes) != 2 {
		// Manejar un error si el formato no es el esperado
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato de máquina virtual incorrecto"})
		return
	}

	ip := partes[0]
	hostname := partes[1]

	payload := map[string]interface{}{
		"imagen":   nombreImagen,
		"comando":  comando,
		"ip":       ip,
		"hostname": hostname,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
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

	var respuesta map[string]string

	err = json.NewDecoder(resp.Body).Decode(&respuesta)
	if err != nil {
		log.Println("Error al decodificar el body de la respuesta")
		return
	}

	mensaje := respuesta["mensaje"]

	// Recuperar o inicializar un arreglo de máquinas virtuales en la sesión del usuario
	machines, _ := MaquinasActualesC(email.(string))

	// Renderizar la plantilla HTML con los datos recibidos, incluyendo el mensaje
	c.HTML(http.StatusOK, "gestionContenedores.html", gin.H{
		"mensaje":  mensaje, // Pasar el mensaje al contexto de renderizado
		"email":    email,
		"machines": machines,
	})

}

func MaquinasActualesC(email string) ([]Maquina_virtual, error) {
	serverURL := "http://localhost:8081/json/consultMachine" // Cambia esto por la URL de tu servidor en el puerto 8081

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

	var machines []Maquina_virtual

	// Decodifica los datos de respuesta en la variable machines.
	if err := json.Unmarshal(responseBody, &machines); err != nil {
		// Maneja el error de decodificación aquí
	}

	return machines, nil
}

func obtenerContenedores(maquinaVirtual string) ([]Conetendor, error) {
	serverURL := "http://localhost:8081/json/ContenedoresVM" // Cambia esto por la URL de tu servidor en el puerto 8081

	partes := strings.Split(maquinaVirtual, " - ")

	ip := partes[0]
	hostname := partes[1]

	payload := map[string]interface{}{
		"ip":       ip,
		"hostname": hostname,
	}

	jsonData, err := json.Marshal(payload)
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
	var contenedores []Conetendor

	// Decodifica los datos de respuesta en la variable machines.
	if err := json.Unmarshal(responseBody, &contenedores); err != nil {
		// Maneja el error de decodificación aquí
	}

	return contenedores, nil

}

func GetContendores(c *gin.Context) {

	// Acceder a la sesión para obtener el email del usuario
	maquinaVirtual := c.PostForm("buscarMV")

	log.Println("Maquina Virtual:", maquinaVirtual)

	// Obtener los datos de las máquinas utilizando el email del usuario
	contenedores, err := obtenerContenedores(maquinaVirtual)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Devolver los datos en formato JSON
	c.JSON(http.StatusOK, contenedores)
}
