# WeiCloth

### 1. Levantar el Backend Core (Go)
En la **primera terminal** (ubicado en la carpeta donde están los archivos `.go`), ejecuta:
```bash
go run main.go handlers.go models.go
```

### 2. Levantar el Frontend
En una **segunda terminal**, levantar el servidor web local ( dentro d la carpeta ):
```bash
http-server -p 3000
```

### 3. Exponer a Internet
En una **tercera terminal**, ejecutar este comando para crear el túnel público:
```bash
cloudflared tunnel --url http://localhost:3000
```
