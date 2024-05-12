package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ProfilePage(c *gin.Context) {
	// Acceder a la sesión
	session := sessions.Default(c)
	email := session.Get("email")
	nombre := session.Get("nombre")
	apellido := session.Get("apellido")
	rol := session.Get("rol")

	if email == nil {
		// Si el usuario no está autenticado, redirige a la página de inicio de sesión
		c.Redirect(http.StatusFound, "/loginPage")
		return
	}

	fmt.Println("Valor del correo electrónico:", email)

	c.HTML(http.StatusOK, "profile.html", gin.H{
		"email":    email,
		"nombre":   nombre,
		"apellido": apellido,
		"rol":      rol,
	})
}
