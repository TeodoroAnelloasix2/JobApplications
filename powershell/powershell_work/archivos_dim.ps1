@"
🔹 Reto - Gestión de Archivos y Tamaños

Crea un script en PowerShell que:

    Busque todos los archivos dentro de una carpeta específica y sus subcarpetas.
    Ordene los archivos por tamaño (de mayor a menor).
    Muestre en pantalla los 10 archivos más grandes.
    Calcule el tamaño total de esos 10 archivos.
    Guarde la información en un archivo con la fecha en el nombre.

🔍 Comandos a investigar:

    Get-ChildItem
    Sort-Object
    Select-Object
    Measure-Object
    Out-File
"@


$w_path=Read-host "Inserta una ruta del sistema para extraer los 10 archivos más pesados"

if ($w_path -eq "."){
    $w_path=Get-Location
    
}

write-host "Analisis en curso de-->  $w_path"
$arx=Get-ChildItem -Path $w_path -File -Recurse | Sort-Object -Property Length -Descending | Select-Object -First 10

Write-Host "Los 10 archivos más grandes: "

$arx | Format-Table name,Length -AutoSize

$dim_tot=($archivos| Measure-Object -Property Length -Sum ).Sum

$mas_pesado=$arx | Select-Object -First 1


$date = Get-Date -Format "yyyyMMdd_HHmm"
$file_path = ".\Archivos_Grandes_$date.txt"


"Archivos más grandes en $w_path" | Out-File -FilePath $file_path

"El más grande es el: $($mas_pesado.Name) con $($mas_pesado.Length)" |Out-File -FilePath $file_path -Append

"El tamaño total es de: $dim_tot" | Out-File -FilePath $file_path -Append

$arx | Format-Table Name,Length -AutoSize | Out-File $file_path -Append


Write-Host "Información guardada en: $file_Path"