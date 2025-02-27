@"
🔹 Reto - Monitoreo de Uso de Memoria

Crea un script en PowerShell que:
    Liste los procesos en ejecución.
    Filtre los 10 procesos que más memoria consumen.
    Calcule el total de memoria usada por esos 10 procesos.
    Muestre la información en pantalla y la guarde en un archivo con la fecha en el nombre.
🔍 Comandos a investigar:
    Get-Process
    Sort-Object
    Select-Object
    Measure-Object
    Out-File
"@

$date= Get-Date -Format "yyyyMMdd_HHmm"
$filePath=".\P_memery_$date.txt"

#Out-string más apto con out-file

Write-Host "Seleccionando  los 10 procesos más costosos en  terminos de momeria! "

Get-Process | Sort-Object -Property WorkingSet64 -Descending | Select-Object -First 10 | Out-File -FilePath $filePath
Get-Process | Sort-Object -Property WorkingSet64 -Descending | Select-Object -First 10 | Measure-Object -Sum -Property WorkingSet64 | Out-File -FilePath $filePath -Append

Get-Content $filePath 
