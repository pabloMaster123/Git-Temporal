package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// DatosCatalogo representa los datos para el catálogo de máquinas virtuales
type DatosDashboard struct {
	Total_maquinas_creadas    int
	Total_maquinas_encendidas int
	Total_usuarios            int
	Total_estudiantes         int
	Total_invitados           int
	Total_RAM                 int
	Total_RAM_usada           int
	Total_CPU                 int
	Total_CPU_usada           int
}

func DashboardHandler(c *gin.Context) {

	// Acceder a la sesión
	session := sessions.Default(c)
	rol := session.Get("rol")
	if rol != "Administrador" {
		// Si el usuario no está autenticado, redirige a la página de inicio de sesión
		c.Redirect(http.StatusFound, "/login")
		return
	}

	// Calcula los datos para el catálogo (esto es solo un ejemplo, debes obtener estos datos de tu lógica)
	datosDashboard, _ := consultarMetricas()

	c.HTML(http.StatusOK, "dashboard.html", gin.H{
		"email":          "email",
		"machines":       nil,
		"machinesChange": nil,
		"datosDashboard": datosDashboard,
	})
}

func consultarMetricas() (DatosDashboard, error) {
	var metricas DatosDashboard

	resp, err := http.Get("http://localhost:8081/json/consultMetrics")
	if err != nil {
		return metricas, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return metricas, err
	}

	err = json.NewDecoder(resp.Body).Decode(&metricas)
	if err != nil {
		return metricas, err
	}

	return metricas, nil
}
