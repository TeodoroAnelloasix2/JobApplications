 @"
 🔹 Reto - Monitoreo de Procesos Activos
Crea un script en PowerShell que:
    Liste los procesos en ejecución.
    Filtre los 5 procesos que consumen más CPU.
    Guarde el resultado en un archivo de texto con la fecha en el nombre.
🔍 Comandos a investigar:
    Get-Process
    Sort-Object
    Select-Object
    Out-File
"@

Get-Process | Sort-Object -Property CPU -Descending | Select-Object  -First 5 | Format-Table -AutoSize | Out-File -FilePath .\Process.txt

Write-Host "Leyendo los procesos.... "
Write-Host "Volcando 5 más dispendiosos en Process.txt"

$p=Get-Content -LiteralPath .\Process.txt
$p | Write-Output