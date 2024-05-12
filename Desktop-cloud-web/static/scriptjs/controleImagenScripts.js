function actualizarTabla() {
    $.ajax({
        url: "/api/images",
        method: "GET", // Cambiado a método GET
        dataType: "json",
        success: function(data) {
            try {
                // Limpia la tabla actual
                $("#imagen-table tbody").empty();
                // Itera a través de los datos y agrega filas a la tabla
                data.forEach(function(imagenes) {
                    var backgroundColor = "#93c47dff"; // Verde
                    $("#imagen-table tbody").append(
                        `<tr style="background-color: ${backgroundColor}">
                            <td>${imagenes.Repositorio}</td>
                            <td>${imagenes.Tag}</td>
                            <td>${imagenes.Tamanio}</td>
                        </tr>`
                    );
                });
            } catch (error) {
                console.error("Error al procesar los datos JSON: " + error);
            }
        },
        error: function(xhr, status, error) {
            console.error("Error al obtener datos de la tabla: " + error);
        }
    });
}

$(document).ready(function() {
    // Capturar el evento de clic del botón de búsqueda
    $("#buscarBtn").click(function(event) {
        // Evitar que el formulario se envíe automáticamente
        event.preventDefault();
        
        // Llamar a la función actualizarTabla()
        actualizarTabla();
    });
});