# {{cookiecutter.api_name}}

{{cookiecutter.image_description}}

## Diagrama de flujo

<Agregar aqui un diagrama de secuencia o de flujo describiendo la logica de la API>

## Como compilar

Para compilar simplemente se debe hacer `make build` lo que generara un binario en `build/bin/api-test`

## Correr tests y coverage

Para correr los tests se debe hacer `make test` y para obtener el coverage `make coverage`

## Configuración

| Variables       |  Command && Shortcut   |                Descripción                |
| --------------- | :--------------------: | :---------------------------------------: |
| PORT            |      --port=, -p       | Puerto que escucha el servicio http, por defecto `8080`         |
| LOGGING_LEVEL   | --logging_level=,-l    | Nivel de logger, por defecto `info`                             |
| TRACING_ENABLED | --tracing_enabled=, -t | Indica si tracing se encuentra habilitado, por defecto `false`  |
| METRICS_ENABLED | --metrics_enabled=, -m | Indica si metrics se encuentra habilitado, por defecto `true`   |
| DD_TRACING_ENABLED  | --dd_tracing_enabled=,  | Especifica si se debe configurar tracing para datadog, por defecto `false`  |

## Ambientes

| Entorno |                 Url                  |
| ------- | :----------------------------------: |
| Dev     | http://integracion-k8s-dev.fif.tech/ |
| QA      | http://integracion-k8s-qa.fif.tech/  |
| PROD    | http://integracion-k8s.fif.tech/ |

## Documentación
### https://confluence.falabella.tech/x/CjCTDg
```json
{
    "name": "{{cookiecutter.api_name}}", //Nombre de la API
    "capability": "", //Categoría de la API (Customer,CreditCard,Account,etc)
    "countries": [], //Paises donde esta desplegada la API
    "kafka": {}, // Si tiene uso de kafka
    "paths": [
        "/"
    ], //Path que tendra disponible la API
    "repository": "https://gitlab.falabella.tech/fif/integracion/forthehorde/{{cookiecutter.api_name}}", //Repositorio de la API
    "status": "", //Status de la API (active,deprecated,etc)
    "uhuras": [
        {
            "idService" : "", //Nombre del servicio (ul-go-..., etc) 
            "country" : "", //Pais del servicio que estamos invocando
            "operation" : "", //Operacion del servicio que necesitamos
            "version" : "" //Version del Servicio y Operacion 
        }
    ], //Servicios Uhuras que usa la API
    "apis": [], // Apis que usa la API
    "tags": [], 
    "version": "VERSION_ENV", //Version de la API que se seteará al generar el tag el pipeline
    "teamDev": "integraciones-mdw",
    "DocReference": "https://confluence.falabella.tech/x/CjCTDg"
}
```
## Anexos
### Flujo de despliegues
* https://confluence.falabella.tech/x/Si1LDQ

### Cookiecutter
* https://gitlab.falabella.tech/fif/integracion/forthehorde/templates/cookiecutters/images/cookiecutter-api-go-image

