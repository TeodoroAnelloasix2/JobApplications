# 🌟 Imagen base
**FROM** ubuntu  # From indica la base de la imagen desde donde partimos

# 📦 Instalación de paquetes y creación de un archivo version
**RUN** apt-get update && apt-get install -y python3 git vim iputils-ping \  # Run ejecuta comandos
    && echo "1.0.0" >> ./version     

# 🚀 Comando por defecto al iniciar el contenedor
**CMD** ["echo", "-e", "Welcome to this container"]  # cmd es el comando por defecto del contenedor ["ejecutable","param1","param2"...]

# 🎯 Configura un entrypoint para concatenar argumentos
**ENTRYPOINT** ["df"]  # ENTRYPOINT puede concatenar argumentos al comando

# 🚀 Ejecución del contenedor con df -h
# docker run -it --rm helloworld:1.0.1 -h

# Resultado esperado:

# Filesystem      Size  Used Avail Use% Mounted on
# overlay         457G   30G  405G   7% /
# tmpfs            64M     0   64M   0% /dev
# shm              64M     0   64M   0% /dev/shm
# /dev/nvme0n1p2  457G   30G  405G   7% /etc/hosts
# tmpfs           2.9G     0  2.9G   0% /proc/asound
# tmpfs           2.9G     0  2.9G   0% /proc/acpi
# tmpfs           2.9G     0  2.9G   0% /proc/scsi
# tmpfs           2.9G     0  2.9G   0% /sys/firmware
# tmpfs           2.9G     0  2.9G   0% /sys/devices/virtual/powercap

# 📁 Establece el directorio de trabajo
**WORKDIR** /personal  # WORKDIR establece el directorio de trabajo

# 📥 Copia archivos al contenedor
**COPY** ./program.py .  # Copia archivoHost /ruta/contenedor
# Si tengo un index.html en local y quiero copiarlo:
# COPY /local/index.html /ruta/contenedor

# 🐍 Imagen con Python y contenido adicional
**FROM** ubuntu:22.04

**RUN** apt-get update && apt install -y python3 && mkdir /personal

**WORKDIR** /personal

**COPY** ./program.py .

**CMD** ["python3", "program.py"]

# 📤 Añade más contenido (archivos y directorios)
**ADD** ../docsHst/d* /personal/docsContenedor
**ADD** docs.tar .  # ADD es similar a COPY pero también permite traer cosas desde internet



# 🌍 Docker - Uso de ENV, ENTRYPOINT y CMD

## 📌 Definir variables de entorno con `ENV`
Permite establecer variables de entorno que estarán disponibles dentro del contenedor.
Establecer valor con -e  -> docker run -it --rm -e texto="Ya he cambiado el texto" imagetest2:latest

```dockerfile
ENV usuario1="El Italiano" usuario2="Gopher" usuario3=python3
ENV texto="Esto es un valor por defecto"


