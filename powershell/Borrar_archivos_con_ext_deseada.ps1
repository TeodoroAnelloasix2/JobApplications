param(
    [string]$directorio="C:\Users\giaco",
    [string]$ext=".txt"
)


$logs=Get-childitem -Path $directorio -Recurse -Filter "*$ext"
$logCount= $logs.Count
$totalSize=($logs | Measure-Object -Property Length -Sum).sum
$totalSizeMB=[math]::Round($totalSize/1MB,2)

$oldest= $logs | Sort-Object LastWriteTime | Select-Object -First 1
$newest= $logs | Sort-Object LastWriteTime -Descending | Select-Object -First 1

@"

   Crear comentario multilínea 
"@

$text=@"
    Cantidad de archivos encontrados: $logcount
    Tamaño total: $totalSizeMB MB
    Mas reciente: $($newest.Lastwritetime) $newest
    Mas antiguo: $($oldest.Lastwritetime) $oldest

"@
$text | out-file -FilePath "$directorio\resumen.txt"

$oldlogs= $logs | where-object { $_.LastWriteTime -lt (Get-Date).AddDays(-30)}

foreach ($log in $oldlogs){
Get-Content $log | Write-Output 
Remove-Item $log.fullname -Confirm }


Get-Content "$directorio\resumen.txt" | Write-Output

#Get-Content "test.txt" | Write-Output