package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoginPage(c *gin.Context) {
	// Acceder a la sesión
	session := sessions.Default(c)
	email := session.Get("email")

	if email != nil {
		// Si el usuario no está autenticado, redirige a la página de inicio de sesión
		c.Redirect(http.StatusFound, "/mainPage")
		return
	}

	c.HTML(http.StatusOK, "login.html", nil)
}

func Login(c *gin.Context) {
	// Obtener los datos del formulario
	email := c.PostForm("email")
	password := c.PostForm("password")

	// Crear una estructura Account y convertirla a JSON
	persona := Persona{Email: email, Contrasenia: password}
	jsonData, err := json.Marshal(persona)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	usuario, er := sendJSONToServer(jsonData)
	fmt.Println(usuario)

	if er == nil {
		session := sessions.Default(c)
		session.Set("email", email) // Almacena el email del usuario en la sesión
		session.Set("nombre", usuario.Nombre)
		session.Set("apellido", usuario.Apellido)
		session.Set("rol", usuario.Rol)
		session.Save()
		// Redirigir al usuario a la página principal
		c.Redirect(302, "/mainPage")
	} else {
		c.Redirect(http.StatusFound, "/login")
	}
}

func LoginTemp(c *gin.Context) {
	session := sessions.Default(c)
	serverURL := "http://localhost:8081/json/createGuestMachine" // Cambia esto por la URL de tu servidor en el puerto 8081

	clientIP := c.ClientIP()
	distribucion := c.PostForm("osCreate")

	//Crea un mapa con la dirección IP del cliente
	data := map[string]string{
		"ip":           clientIP,
		"distribucion": distribucion,
	}

	// Convierte el mapa a JSON
	jsonBody, err := json.Marshal(data)
	if err != nil {
		// Maneja el error si la conversión falla
		fmt.Println("Error al convertir a JSON:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
		return
	}

	// Crea una solicitud HTTP con el cuerpo JSON
	req, err := http.NewRequest("POST", serverURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		// Maneja el error si la creación de la solicitud falla
		fmt.Println("Error al crear la solicitud HTTP:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
		return
	}
	req.Header.Set("Content-Type", "application/json")

	// Realiza la solicitud
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error al realizar la solicitud HTTP:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		// Lee el cuerpo de la respuesta
		responseBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error al leer el cuerpo de la respuesta:", err)
			return
		}

		// Convierte el cuerpo de la respuesta en un mapa
		var data map[string]string
		if err := json.Unmarshal(responseBody, &data); err != nil {
			fmt.Println("Error al decodificar el JSON:", err)
			return
		}

		// Accede a los datos del mapa
		mensaje := data["mensaje"]
		fmt.Println("Mensaje recibido:", mensaje)

		session.Set("email", mensaje)
		session.Set("nombre", "Usuario")
		session.Set("apellido", "Temporal")
		session.Set("rol", "Invitado")
		session.Save()

		c.Redirect(http.StatusSeeOther, "/controlMachine")
	} else {

	}

}

func sendJSONToServer(jsonData []byte) (Persona, error) {
	serverURL := "http://localhost:8081/json/login" // Cambia esto por la URL de tu servidor en el puerto 8081
	var usuario Persona

	// Crea una solicitud HTTP POST con el JSON como cuerpo
	req, err := http.NewRequest("POST", serverURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return usuario, err
	}

	// Establece el encabezado de tipo de contenido
	req.Header.Set("Content-Type", "application/json")

	// Realiza la solicitud HTTP
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return usuario, err
	}
	defer resp.Body.Close()

	// Verifica la respuesta del servidor (resp.StatusCode) aquí si es necesario
	if resp.StatusCode != http.StatusOK {
		return usuario, errors.New("Error en la respuesta del servidor")
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return usuario, err
	}
	var resultado map[string]interface{}

	if err := json.Unmarshal(responseBody, &resultado); err != nil {
		fmt.Println("Error al deserializar")
		return usuario, err
	}
	specsMap, _ := resultado["usuario"].(map[string]interface{})
	specsJSON, err := json.Marshal(specsMap)
	if err != nil {
		fmt.Println("Error al serializar el usuario:", err)
		return usuario, err
	}

	err = json.Unmarshal(specsJSON, &usuario)
	if err != nil {
		fmt.Println("Error al deserializar el usuario:", err)
		return usuario, err
	}

	return usuario, nil
}

func GuestLoginSend(c *gin.Context) {
	// Acceder a la sesión
	session := sessions.Default(c)
	email := session.Get("email")

	if email == nil {
		// Si el usuario no está autenticado, redirige a la página de inicio de sesión
		c.Redirect(http.StatusFound, "/login")
		return
	}

	userEmail := email.(string)

	// Obtener los datos del formulario
	vmname := c.PostForm("vmnameCreate")
	if vmname == "" {
		// Si el nombre de la máquina virtual está vacío, mostrar un mensaje de error en el HTML
		c.HTML(http.StatusOK, "controlMachine.html", gin.H{
			"ErrorMessage": "El nombre de la máquina virtual no puede estar vacío.",
		})
		return
	}
	ditOs := c.PostForm("osCreate")
	memoryStr := c.PostForm("memoryCreate")
	memory, err := strconv.Atoi(memoryStr)
	cpuStr := c.PostForm("cpuCreate")
	cpu, _ := strconv.Atoi(cpuStr)
	os := "Linux"

	// Crear una estructura Account y convertirla a JSON
	maquina_virtual := Maquina_virtual{Nombre: vmname, Sistema_operativo: os, Distribucion_sistema_operativo: ditOs, Ram: memory, Cpu: cpu, Persona_email: userEmail}
	clientIP := c.ClientIP()

	payload := map[string]interface{}{
		"specifications": maquina_virtual,
		"clientIP":       clientIP,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if sendJSONMachineToServer(jsonData) {
		// Registro exitoso, muestra un mensaje de éxito en el HTML
		c.HTML(http.StatusOK, "controlMachine.html", gin.H{
			"SuccessMessage": "Solicitud para crear màquina virtual enviada con èxito.",
		})
	} else {
		// Registro erróneo, muestra un mensaje de error en el HTML
		c.HTML(http.StatusOK, "controlMachine.html", gin.H{
			"ErrorMessage": "Error al enviar la solicitud para crear màquina virtual. Intente de nuevo",
		})
	}
}
