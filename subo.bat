@echo off
echo "Iniciando despliegue a AWS Lambda"

:: Agregar y subir cambios a Git
git add .
git commit -m "last Commit"
git push
IF %ERRORLEVEL% NEQ 0 (
    echo "Error al ejecutar git push. Saliendo..."
    exit /b %ERRORLEVEL%
)

:: Configurar entorno Go para Lambda (Linux y arquitectura amd64)
set GOOS=linux
set GOARCH=amd64

:: Compilar el código Go para Lambda
echo "Compilando el binario para AWS Lambda..."
go build -o bootstrap main.go
IF %ERRORLEVEL% NEQ 0 (
    echo "Error en la compilación de Go. Saliendo..."
    exit /b %ERRORLEVEL%
)

:: Eliminar el archivo ZIP anterior
if exist bootstrap.zip (
    echo "Eliminando bootstrap.zip anterior..."
    del bootstrap.zip
)

:: Crear archivo ZIP
echo "Creando archivo ZIP para Lambda..."
tar.exe -a -cf bootstrap.zip bootstrap
IF %ERRORLEVEL% NEQ 0 (
    echo "Error al crear el archivo ZIP. Saliendo..."
    exit /b %ERRORLEVEL%
)

echo "Archivo ZIP creado con éxito: bootstrap.zip"

:: Aquí puedes agregar la subida a AWS Lambda si lo necesitas
:: Ejemplo (con AWS CLI instalado):
:: aws lambda update-function-code --function-name <tu-lambda-function> --zip-file fileb://bootstrap.zip

echo "Despliegue completado"
pause
