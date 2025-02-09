param(
    [string]$source="C:\yourpath",
    [string]$dest="C:\yourpath",
    [int]$days=5

)

#Crear carpeta de backup con marca de tiempo

$timestamp=Get-date -Format "ddMMyyyy"
$backupdir="$dest\Backup_$timestamp"

if(!(Test-path $backupdir)){ New-Item -ItemType directory -path $backupdir }

$recentFiles=get-childitem -path $source -recurse | where-object {
    $_.LastWriteTime -ge (Get-Date).AddDays(-$days)
}

$log=@()


foreach ($file in $recentFiles){
    
    $destPath="$backupdir\$($file.name)"
    copy-item -path $file.fullname -Destination $destPath
    #Agerar archivo copiado a log

    $log+="Copiando: $($file.Fullname) ---> $destpath"
}

$log | out-file -filepath "$backupdir\informe_backup.txt"



$log | ForEach-Object { Write-Output $_ }
