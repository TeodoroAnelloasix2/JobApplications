@"
🔹 Reto - Detección de Procesos No Respondientes

Crea un script en PowerShell que:
    Liste todos los procesos en ejecución.
    Filtre aquellos que no están respondiendo.
    Guarde la lista en un archivo con la fecha en el nombre.
    Muestre en pantalla cuántos procesos no responden.
🔍 Comandos a investigar:
    Get-Process
    Where-Object
    Out-File
    Measure-Object
"@
Write-Host 
Write-Host

$procesos_not_responding=Get-Process | Where-Object {$_.Responding -eq $false}

$tot=$procesos_not_responding | Measure-Object

if ($tot.Count -ge 1 ){

    Write-Output $tot 

    Write-Host "Procesos que no están respondiendo.... "
    $procesos_not_responding  |  Format-Table -AutoSize -Property Id, Name

}
else{
    Write-Output " No se han detectado procesos que no están respondiendo correctamente! "
}