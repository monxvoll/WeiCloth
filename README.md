# WeiCloth

## Frontend Config

### 1. Start the server

```code
ng serve -o
```


## Backend Config

### Docker Services

#### Start only kafka, prostgres or keycloak instance

__Up__

```code
docker compose up service-name -d
```
__Test__

Use tags if the test is another like integration **-tags=integration -v**

```code
go test package-route -v
```
__Utilities__
```code
docker compose down -v image-name
```