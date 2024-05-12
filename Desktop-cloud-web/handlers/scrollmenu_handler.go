package handlers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Scrollmenu(c *gin.Context) {

	// Acceder a la sesión
	session := sessions.Default(c)
	email := session.Get("email")
	rol := session.Get("rol")

	// Recuperar o inicializar un arreglo de máquinas virtuales en la sesión del usuario
	machines, _ := consultarMaquinas(email.(string))

	c.HTML(http.StatusOK, "scrollmenu.html", gin.H{
		"email":    email,
		"machines": machines,
		"rol":      rol,
	})
}

func ActualizacionesMaquinas(c *gin.Context) {

	// Acceder a la sesión
	session := sessions.Default(c)
	email := session.Get("email")

	// Obtén las máquinas actualizadas (por ejemplo, desde una base de datos)
	machines, err := consultarMaquinas(email.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener actualizaciones de máquinas"})
		return
	}

	// Devuelve las máquinas en formato JSON
	c.JSON(http.StatusOK, machines)
}
