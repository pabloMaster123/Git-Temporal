var ventanaConfiguracionAbierta = false;

    function abrirVentanaEmergente() {
        ventanaConfiguracionAbierta = true;

        var ventanaEmergente = document.getElementById("ventanaEmergente");
        ventanaEmergente.style.display = "block";
    }

    function abrirVentanaEmergenteConfiguracion(Nombre, Sistema_operativo, Memoria, Cpu) {
        ventanaConfiguracionAbierta = true;

        var ventanaEmergente = document.getElementById("VentanaEmergenteConfiguracion");
        ventanaEmergente.style.display = "block";

        // Asigna los valores de los parámetros a los campos del formulario
        document.getElementById("vmnameConfig").value = Nombre;
        document.getElementById("osConfig").value = Sistema_operativo;
        document.getElementById("memoryConfig").value = Memoria;
        document.getElementById("cpuConfig").value = Cpu;
        document.getElementById("nombreMaquina").value = Nombre;
    }

    function abrirVentanaEmergenteEliminacion( nombre ){
        ventanaConfiguracionAbierta = true;

        var VentanaEmergenteEliminacion = document.getElementById("VentanaEmergenteEliminacion");

        document.getElementById("vmnameDelete").value = nombre;

        VentanaEmergenteEliminacion.style.display = "block";
    }


    function abrirVentanaEmergenteInformacion(nombre, sistemaOperativo, distribucion, ip, ram, cpu, estado, hostname) {
        ventanaConfiguracionAbierta = true;

        var ventanaEmergente = document.getElementById("VentanaEmergenteInformacion");
        ventanaEmergente.style.display = "block";

        // Asignar valores a los elementos <span> con sus respectivos ids

        // Nombre de la Máquina
        document.getElementById("nombreSpan").textContent = nombre;

        // Sistema Operativo de la Máquina
        document.getElementById("sistemaOperativoSpan").textContent = "GNU/" + sistemaOperativo;

        // Distribución de la Máquina
        document.getElementById("distribucionSpan").textContent = distribucion;

        // Modelo de CPU
        if ( cpu == 1 ){
            document.getElementById("cpuSpan").textContent = cpu + " Núcleo";
        }else {
            document.getElementById("cpuSpan").textContent = cpu + " Núcleos";

        }

        // Memoria de la Máquina
        document.getElementById("memoriaSpan").textContent = ram + " MB";

        // Estado de la Máquina
        document.getElementById("estadoSpan").textContent = estado;

        document.getElementById("hostnameSpan").textContent = hostname;

        document.getElementById("passwordSpan").textContent = hostname;

        // IP de la Máquina
        if  ( ip != ""){
            document.getElementById("ipSpan").textContent = ip;
            document.getElementById("urlSpan").textContent = "http://"+ip+":4200"
        }else{
            document.getElementById("ipSpan").textContent = "No asignada";
            document.getElementById("urlSpan").textContent = "No asignada"
        }
    }

    function cerrarVentanaEmergente() {
        ventanaConfiguracionAbierta = false;

        var ventanaEmergente = document.getElementById("ventanaEmergente");
        ventanaEmergente.style.display = "none";
    }

    function cerrarVentanaEmergenteConfiguracion() {
        ventanaConfiguracionAbierta = false;

        var ventanaEmergente = document.getElementById("VentanaEmergenteConfiguracion");
        ventanaEmergente.style.display = "none";
    }

    function cerrarVentanaEmergenteInformacion(){
        ventanaConfiguracionAbierta = false;

        var ventanaEmergente = document.getElementById("VentanaEmergenteInformacion");
        ventanaEmergente.style.display = "none";
    }

    function cerrarVentanaEmergenteEliminar(){
        ventanaConfiguracionAbierta = false;

        var VentanaEmergenteEliminacion = document.getElementById("VentanaEmergenteEliminacion");
        VentanaEmergenteEliminacion.style.display = "none";
    }

    function actualizarTabla() {
        
        if (ventanaConfiguracionAbierta ) {
            return;
        }
        $.ajax({
            url: "/api/machines", // Reemplaza con la URL correcta
            method: "GET",
            dataType: "json",
            success: function(data) {
                // Limpia la tabla actual
                $("#machine-table tbody").empty();
                // Itera a través de los datos y agrega filas a la tabla
                data.forEach(function(machine) {
                    switch (machine.Estado) {
                        case "Apagado":
                            backgroundColor = "#e06666ff"; // Rojo
                            $("#machine-table tbody").append(
                            `<tr style="background-color: ${backgroundColor}">
                                <td>${machine.Nombre}</td>
                                <td>${machine.Ip === "" ? "No asignada" : machine.Ip}</td>
                                <td>${machine.Distribucion_sistema_operativo}</td>
                                <td>${machine.Estado}</td>
                                
                                <td class="button-column">
                                    <form method="post" action="/powerMachine" style="display: inline-block; padding: 0; margin: 0; border: none;">
                                    <input type="hidden" name="nombreMaquina" value="${machine.Nombre}">
                                    <button type="submit" class="btn btn-link" style="padding: 0; margin: 0;">
                                        <img style="width: 35px;" src="/static/images/icons/power.png" alt="Botón 1">
                                    </button>
                                    </form>
                                    <button type="button" class="btn btn-link" style="padding: 0; margin: 0;" onclick="abrirVentanaEmergenteConfiguracion('${machine.Nombre}','${machine.Distribucion_sistema_operativo}','${machine.Ram}','${machine.Cpu}')">
                                        <img style="width: 30px;" src="/static/images/icons/config.png" alt="Botón 2">
                                    </button>
                                    <button type="button" class="btn btn-link" style="padding: 0; margin: 0;" onclick="abrirVentanaEmergenteInformacion('${machine.Nombre}','${machine.Sistema_operativo}','${machine.Distribucion_sistema_operativo}', '${machine.Ip}','${machine.Ram}','${machine.Cpu}', '${machine.Estado}', '${machine.Hostname}')">
                                        <img style="width: 35px;" src="/static/images/icons/info.png" alt="Botón 4">
                                    </button>
                                    <button type="button" class="btn btn-link" style="padding: 0; margin: 0;" onclick="abrirVentanaEmergenteEliminacion('${machine.Nombre}')">
                                        <img style="width: 30px;" src="/static/images/icons/delete.png" alt="Botón 3">
                                    </button>
                                </td>
                                </td>
                            </tr>`
                            );
                            break;
                        case "Encendido":
                            backgroundColor = "#93c47dff"; // Verde
                            $("#machine-table tbody").append(
                            `<tr style="background-color: ${backgroundColor}">
                                <td>${machine.Nombre}</td>
                                <td style="position: relative;">
                                    ${machine.Ip === "" ? "No asignada" : machine.Ip}
                                    
                                    <!-- Muestra el botón solo si la IP está asignada -->
                                    ${machine.Ip !== "" ? `
                                    <button type="button" class="btn btn-link" style="padding: 0; margin: 0;" onclick="copiarText('${machine.Ip}')">
                                        <img style="width: 30px;" src="/static/images/icons/copy.png" alt="Botón 5">
                                    </button>
                                    ` : ''}
                                </td>

                                <td>${machine.Distribucion_sistema_operativo}</td>
                                <td>${machine.Estado}</td>
                                
                                <td class="button-column">
                                    <form method="post" action="/powerMachine" style="display: inline-block; padding: 0; margin: 0; border: none;">
                                    <input type="hidden" name="nombreMaquina" value="${machine.Nombre}">
                                    <button type="submit" class="btn btn-link"" style="padding: 0; margin: 0;">
                                        <img style="width: 35px;" src="/static/images/icons/power.png" alt="Botón 1">
                                    </button>
                                    </form>
                                    <button type="button" class="btn btn-link" style="padding: 0; margin: 0;" onclick="abrirVentanaEmergenteConfiguracion('${machine.Nombre}','${machine.Distribucion_sistema_operativo}','${machine.Ram}','${machine.Cpu}')" disabled>
                                        <img style="width: 30px;" src="/static/images/icons/config.png" alt="Botón 2">
                                    </button>
                                    <button type="button" class="btn btn-link" style="padding: 0; margin: 0;" onclick="abrirVentanaEmergenteInformacion('${machine.Nombre}','${machine.Sistema_operativo}','${machine.Distribucion_sistema_operativo}', '${machine.Ip}','${machine.Ram}','${machine.Cpu}', '${machine.Estado}', '${machine.Hostname}', '${machine.Hostname}')">
                                        <img style="width: 35px;" src="/static/images/icons/info.png" alt="Botón 3">
                                    </button>
                                    <form method="post" action="/deleteMachine" style="display: inline-block; padding: 0; margin: 0; border: none;">
                                        <input type="hidden" name="nombreMaquina" value="${machine.Nombre}">
                                        <button type="submit" class="btn btn-link" style="padding: 0; margin: 0;" disabled>
                                            <img style="width: 30px;" src="/static/images/icons/delete.png" alt="Botón 4">
                                        </button>
                                    </form>
                                </td>
                            </tr>`
                            );
                            break;
                        case "Procesando":
                            backgroundColor = "#83DEE3"; // Azul
                            $("#machine-table tbody").append(
                            `<tr style="background-color: ${backgroundColor}">
                                <td>${machine.Nombre}</td>
                                <td>${machine.Ip === "" ? "No asignada" : machine.Ip}</td>
                                <td>${machine.Distribucion_sistema_operativo}</td>
                                <td>${machine.Estado}</td>
                            
                                <td class="button-column">
                                    <form method="post" action="/powerMachine" style="display: inline-block; padding: 0; margin: 0; border: none;">
                                    <input type="hidden" name="nombreMaquina" value="${machine.Nombre}">
                                    <button type="submit" class="btn btn-link" style="padding: 0; margin: 0;" disabled>
                                        <img style="width: 35px;" src="/static/images/icons/power.png" alt="Botón 1">
                                    </button>
                                    </form>
                                    <button type="button" class="btn btn-link" style="padding: 0; margin: 0;" onclick="abrirVentanaEmergenteConfiguracion('${machine.Nombre}','${machine.Sistema_operativo}','${machine.Memoria}','${machine.Cpu}')" disabled>
                                        <img style="width: 30px;" src="/static/images/icons/config.png" alt="Botón 2">
                                    </button>
                                    <button type="button" class="btn btn-link" style="padding: 0; margin: 0;" onclick="abrirVentanaEmergenteInformacion('${machine.Nombre}','${machine.Sistema_operativo}','${machine.Distribucion_sistema_operativo}', '${machine.Ip}','${machine.Ram}','${machine.Cpu}', '${machine.Estado}', '${machine.Hostname}')">
                                        <img style="width: 35px;" src="/static/images/icons/info.png" alt="Botón 4">
                                    </button>
                                    <form method="post" action="/deleteMachine" style="display: inline-block; padding: 0; margin: 0; border: none;">
                                        <input type="hidden" name="nombreMaquina" value="${machine.Nombre}">
                                        <button type="submit" class="btn btn-link" style="padding: 0; margin: 0;" disabled>
                                            <img style="width: 30px;" src="/static/images/icons/delete.png" alt="Botón 3">
                                        </button>
                                    </form>
                                </td>
                            </tr>`
                            );
                            break;
                        default:
                            backgroundColor = ""; // Puedes proporcionar un valor predeterminado si es necesario
                    }
     
                });
            },
            error: function(error) {
                console.error("Error al obtener datos: " + error);
            }
        });
    }

    function copiarText(texto) {
        // Crea un elemento de entrada temporal
        const tempInput = document.createElement("input");
        tempInput.value = texto;
        document.body.appendChild(tempInput);
        tempInput.select();

        // Intenta copiar el texto al portapapeles
        document.execCommand("copy");

        // Elimina el elemento de entrada temporal
        document.body.removeChild(tempInput);
    }

    // Llama a actualizarTabla al cargar la página y periódicamente para mantener los datos actualizados
    actualizarTabla();
    setInterval(actualizarTabla, 500);
