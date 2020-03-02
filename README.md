[![GeeksCli](https://circleci.com/gh/JavierRodriguez78/GeeksCli/tree/circleci-project-setup.svg?style=svg)](https://circleci.com/gh/JavierRodriguez78/GeeksCli/tree/circleci-project-setup)

## **GeeksCli** es un cliente desarrollado en Go para gestionar tus entornos de programación desde tu terminal.

### Desarrollo
Para poder seguir desarrollando la aplicación deberás tener instalado Go, para ello:

1º Descarga e instala [Go](https://golang.org/dl/)

2º Descarga el proyecto de GitHub.

3º Inicializamos el gestor de Módulos en Go 
```
$> go mod init
```
4º Instalamos las dependencias de los Módulos de nuestro proyecto.
```
$> go install
```
5º Para compilar el proyecto sólo debemos ejecutar
```
$> go build
```
6º Para ejecutar el proyecto
```
$> go run main.go
```


### Estructura del proyecto
![](https://github.com/JavierRodriguez78/GeeksCli/blob/master/assets/Estructura.png)

El proyecto dispone de un directorio denominado **services** donde añadiremos el código de los diveros servicios de tercero que desarrollemos.
Y otra directorio denominado lib donde almacenaremos nuestras **librerias**

El archivo de arranque es **main.go**

### Configuración.
1º Debemos añadir  en la servicio de DigitalOcean en el archivo TakenSource.go dentro de la constante **pat** el acceso Token que generamos en DigitalOcean para poder dar acceso los pasos para obtener dicho token los podéis seguir [aquí](https://www.digitalocean.com/docs/apis-clis/api/create-personal-access-token/)


