Automatización descargar ydistribuir script si necesario.

Ejemplo de uso:

go run main.go --url "http://ejemplo.com/script.ps1" --path ./directorio_deseado

Resultado final estructura de carpetas:
    script_ib3
        |__ script_descargado.sql(ejemplo)
        |__success
        |__error
        |__log_error_go.txt   

Se tienen que implementar: Copiar script descargado en maquina deseada y segun resultado mover el script en exito o error. 
Siempre dejar constancia de los eventos en archivo log.