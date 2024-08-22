# Estructura del Proyecto y Pasos Realizados
## Pasos Realizados

1. **Creación de Folders**
    - Definí los siguientes folders para estructurar el proyecto de manera clara:
        - **router**: Encargado del enrutamiento entre los endpoints y sus respectivos controllers. Funciona como un "pasamanos" que simplifica el código.
        - **controller**: Recibe el contexto de la solicitud y lo envía al servicio correspondiente para su procesamiento.
        - **service**: Actúa como el intermediario entre el controller y la capa de base de datos.
        - **database**: Gestiona la escritura y lectura de los datos en la base de datos.
        - **models**: Contiene los modelos de datos utilizados en el proyecto. Aunque no es un uso estándar (generalmente cada capa maneja sus propios modelos), en este caso se centralizó aquí por conveniencia.

    - Por costumbre, los folders mencionados se colocan dentro de un folder `src`, mientras que el archivo `main.go` se encuentra fuera. Esta organización puede variar según el lenguaje y no hay una elección claramente superior.

2. **Definición de Estructuras**
    - Creé un struct `Controller` encargado de manejar las solicitudes recibidas desde el router, así como structs para `Service` y `Database`.

3. **Modelos de Datos**
    - Agregué los modelos necesarios para representar los items dentro de la capa `models`.

4. **Estructura Jerárquica**
    - Definí una estructura simple para el controller, que utiliza un struct de service, el cual a su vez utiliza un struct de database. Al ver que la base de datos puede ser reutilizada, definí su struct utilizando generics (opcional en este caso).

5. **Encapsulamiento**
    - Implementé métodos de creación (constructores) para cada struct, manteniendo los atributos privados para respetar el encapsulamiento.

6. **Separación del Router**
    - Externalicé la configuración del router desde el archivo `main.go` para mantener el código más limpio y modular.

7. **Implementación del Código**
    - Implementé el código correspondiente en cada parte del proyecto, respetando la estructura definida.

8. **Integración del Router**
    - Finalmente, añadí el código necesario en el router para conectar los endpoints con sus controllers.

## Ejecución

Para ejecutar la aplicación, asegúrate de tener Go instalado y utiliza el siguiente comando:
```bash
go run main.go
