function actualizarTabla() {
    $.ajax({
        url: "/api/images",
        method: "POST",
        dataType: "json",
        success: function(data) {
            // Limpia la tabla actual
            $("#imagen-table tbody").empty();
            // Itera a través de los datos y agrega filas a la tabla
            data.forEach(function(imagen) {
                backgroundColor = "#93c47dff"; // Verde
                $("#imagen-table tbody").append(
                    `<tr style="background-color: ${backgroundColor}">
                        <td>${imagen.Repositorio}</td>
                        <td>${imagen.Tag}</td>
                        <td>${imagen.Tamanio}</td>
                    </tr>`
                );
            });
        },
        error: function(error) {
            console.error("Error al obtener datos de la tabla: " + error);
        }
    });
}

actualizarTabla();