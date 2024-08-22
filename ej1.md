# Proyecto en Golang con Gin

Este proyecto está desarrollado en Golang y utiliza el framework Gin para construir una API sencilla. A continuación, se detallan los pasos seguidos para su configuración y desarrollo.

## Pasos Realizados

1. **Inicialización del Proyecto**
    - Ejecuté `go mod init <nombre_del_proyecto>` para iniciar el módulo Go, ya que elegí Go como el lenguaje para este proyecto.

2. **Instalación de Gin**
    - Para el desarrollo del proyecto, decidí utilizar el framework Gin. Instalé Gin con el siguiente comando:
      ```bash
      go get -u github.com/gin-gonic/gin
      ```

3. **Creación del Archivo Principal**
    - Creé el archivo `main.go`, que servirá como el archivo principal de la aplicación.

4. **Definición del Endpoint**
    - Definí el endpoint principal que manejará las solicitudes. El endpoint se configura para recibir y procesar las solicitudes correspondientes.

5. **Definición del Body Esperado**
    - Definí el formato del cuerpo (body) esperado para las solicitudes. Asegúrate de que el formato del body coincida con lo definido en el código para evitar errores.

6. **Configuración de la Base de Datos**
    - Utilicé un mapa simple en memoria para almacenar los datos de manera temporal durante el desarrollo. Esto es adecuado para pruebas y desarrollo, pero considera usar una base de datos persistente para un entorno de producción.

7. **Manejo de Errores y Respuestas**
    - Implementé el método para manejar los posibles errores que puedan surgir, devolviendo los códigos de estado HTTP correspondientes:
        - **201 Created**: Cuando se crea un ítem correctamente, se devuelve el ítem creado.
        - **400 Bad Request**: Si falta el `id` en la solicitud.
        - **409 Conflict**: Si el estado ya existe. (Este código es debatible; en algunos casos, un código 403 Forbidden o un 400 Bad Request podrían ser más apropiados. Evita usar códigos de error del servidor como 5xx. Alternativamente, se podría devolver un 200 OK con el ítem antiguo si aplica).

## Ejecución

Para ejecutar la aplicación, utiliza el comando:
```bash
go run main.go
