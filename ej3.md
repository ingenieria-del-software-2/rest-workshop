# Ejercicio 3: Test de Integración

En este ejercicio, nos enfocamos en realizar un test de integración para la API. El objetivo es probar la funcionalidad del router sin tener conocimiento del código subyacente. A continuación, se detallan los pasos realizados y los aspectos clave del ejercicio.

## Pasos Realizados

1. **Definición del Test de Integración**
    - Para iniciar el test de integración, creé un archivo de test llamado `main_test.go`.
    - En este archivo, inicialicé el router. En lugar de ejecutar la aplicación con `run`, utilizo el método `ServeHTTP` para simular una serie de solicitudes HTTP, asegurando que el router maneje correctamente las solicitudes sin exponer detalles internos del código.

2. **Simulación de Solicitudes POST**
    - Dado que en este punto no contamos con un endpoint `GET`, lo que quiero probar es la capacidad del sistema para manejar múltiples solicitudes POST.
    - Realizo dos solicitudes POST consecutivas:
        - La primera debe escribir correctamente en la base de datos.
        - La segunda debe detectar que el item ya existe y manejar la situación apropiadamente.

3. **Independencia del Test**
    - Para asegurar que el test sea verdaderamente independiente, no utilizo los modelos definidos en las capas subyacentes. En su lugar, creo las entidades manualmente, simulando el proceso que realizaría un usuario normal al llenar un JSON. *(Puntos extra si puedes adivinar por qué se hace así)*.

4. **Ejecución y Verificación**
    - Ejecuto los tests y verifico que los resultados sean los esperados.
    - Adicionalmente, reviso que los headers sean correctos en caso de ser necesario.

5. **Nota sobre el Orden de Ejecución**
    - Como nota al pie, este ejercicio de integración podría haberse realizado antes que el ejercicio 2, ya que los cambios de código podrían haber sido validados con estos tests. Sin embargo, por decisión pedagógica, se hizo en este orden para la clase.

## Ejecución de los Tests

Para ejecutar los tests, utiliza el siguiente comando:
```bash
go test
