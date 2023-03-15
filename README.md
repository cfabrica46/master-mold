# cookiecutter-app-image

Este Cookiecutter incluye todo lo necesario para trabajar en su proyecto de API. El cookiecutter le pedirá algunos input y creara un directorio con un proyecto (**build-image**) que incluyen el pipeline listos para realizar:

* Versionamiento semantico
* Publicación de la imagen en harbor y revisión de seguridad.

## Prerequisitos:

* [pre-commit](https://pre-commit.com/)
* [cookiecutter](https://cookiecutter.readthedocs.io/en/2.0.2/)

## Instalación:

```
$ cookiecutter https://gitlab.falabella.tech/fif/integracion/forthehorde/templates/cookiecutters/images/cookiecutter-api-go-image.git --checkout v0.1.23
```

Los input:

* api_name [api_test]: Nombre de la aplicación 
* runner_tag [integracion-gitlab-dcc-qa]: Tag en donde se ejecutará
* harbor_project [integracion-api-dev]: Nombre del proyecto harbor a utilizar
* image_description [Breve descripcion de la api]: Descripcion de que hace la api,

La salida del comando creara un directorio con el nombre del **api_name**

## Estructura de directorio

Dentro del directorio se el siguiente árbol de directorio:

```
$ cd api-test/
$ tree -a
.
└── build-image
    ├── .gitignore
    ├── .gitlab-ci.yml
    ├── .version
    ├── README.md
    └── ias
        ├── .dockerignore
        └── Dockerfile

5 directories, 22 files
```

El directorio:

* **build-image** = Directorio que contiene toda la lógica de programación de su API, y un pipeline (```.gitlab-ci.yml```) que se encargar del versionamiento semántico, creación de la imagen mediante el Dockerfile y la publicación de la imagen a harbor en el proyecto que se definió en la pregunta del cookiecutter **harbor_project**. **No modificar el pipeline**.

## Versionamiento semantico 

Build-image tiene un archivo oculto **.version** con la versión 0.1.0, que debe ir editando para que el pipeline de semversión genere los tags (para el repositorio, y la versión de la imagen) de su API.

Mas información de versionamiento semantico [acá](https://gitlab.falabella.tech/fif/arquitectura/devops-and-cloud/gitlab-templates/semversion-tagging-pipeline).
