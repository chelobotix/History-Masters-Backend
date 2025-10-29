# 🐳 Docker - History Masters Backend

Este documento contiene las instrucciones para ejecutar la aplicación usando Docker.

## 📋 Requisitos Previos

- Docker instalado (versión 20.10 o superior)
- Docker Compose instalado (versión 1.29 o superior)
- Make (opcional, pero recomendado)

## 🚀 Inicio Rápido

### Opción 1: Usando Docker Compose (Recomendado)

Esta opción levanta tanto la API como la base de datos PostgreSQL automáticamente:

```bash
# 1. Crear archivo .env con tus variables de entorno
cp .env.example .env  # Edita este archivo con tus configuraciones

# 2. Levantar todos los servicios
docker-compose up -d

# 3. Ver los logs
docker-compose logs -f

# 4. Detener los servicios
docker-compose down
```

### Opción 2: Usando Make (Más Simple)

Si tienes Make instalado, puedes usar comandos más simples:

```bash
# Ver todos los comandos disponibles
make help

# Levantar el entorno de desarrollo
make dev-up

# Ver los logs
make dev-logs

# Reconstruir y levantar
make dev-rebuild

# Detener todo
make dev-down

# Limpiar completamente (incluyendo datos de la BD)
make dev-clean
```

### Opción 3: Solo Docker (API únicamente)

Si ya tienes PostgreSQL corriendo localmente:

```bash
# 1. Construir la imagen
docker build -t history-masters-api .

# 2. Ejecutar el contenedor
docker run -d \
  --name history_masters_api \
  -p 8080:8080 \
  --env-file .env \
  history-masters-api

# 3. Ver los logs
docker logs -f history_masters_api

# 4. Detener el contenedor
docker stop history_masters_api
docker rm history_masters_api
```

## 🔧 Configuración

### Variables de Entorno

Crea un archivo `.env` en la raíz del proyecto con las siguientes variables:

```env
# Configuración del Servidor
SERVER_PORT=:8080

# Configuración de la Base de Datos
DB_HOST=postgres          # Usar 'postgres' para docker-compose, 'localhost' para local
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=history_masters
DB_SSLMODE=disable

# Configuración de Log
LOG_LEVEL=info
```

## 📦 Estructura del Dockerfile

El Dockerfile utiliza una construcción multi-stage:

1. **Etapa Builder**: Compila la aplicación Go en un contenedor Alpine
2. **Etapa Runtime**: Copia solo el binario compilado a una imagen Alpine limpia

Esto resulta en una imagen final muy pequeña y segura.

## 🛠️ Comandos Útiles

### Docker Compose

```bash
# Levantar servicios en background
docker-compose up -d

# Levantar servicios en foreground (ver logs)
docker-compose up

# Ver logs de un servicio específico
docker-compose logs -f api
docker-compose logs -f postgres

# Reconstruir las imágenes
docker-compose build

# Reconstruir y levantar
docker-compose up -d --build

# Detener servicios
docker-compose stop

# Detener y eliminar contenedores
docker-compose down

# Detener y eliminar contenedores + volúmenes (¡CUIDADO: borra la BD!)
docker-compose down -v

# Ver el estado de los servicios
docker-compose ps

# Ejecutar comandos en un contenedor
docker-compose exec api /bin/sh
docker-compose exec postgres psql -U postgres -d history_masters
```

### Docker

```bash
# Listar imágenes
docker images

# Listar contenedores en ejecución
docker ps

# Listar todos los contenedores (incluyendo detenidos)
docker ps -a

# Ver logs de un contenedor
docker logs -f <container_name>

# Ejecutar shell en un contenedor
docker exec -it <container_name> /bin/sh

# Detener un contenedor
docker stop <container_name>

# Eliminar un contenedor
docker rm <container_name>

# Eliminar una imagen
docker rmi <image_name>
```

## 🔍 Troubleshooting

### El contenedor no se inicia

```bash
# Ver los logs del contenedor
docker-compose logs api

# Ver todos los contenedores (incluyendo los que fallaron)
docker ps -a

# Verificar que PostgreSQL esté corriendo
docker-compose ps
```

### Error de conexión a la base de datos

1. Verifica que PostgreSQL esté corriendo: `docker-compose ps`
2. Verifica las variables de entorno en `.env`
3. Asegúrate de usar `DB_HOST=postgres` cuando uses docker-compose

### La aplicación no responde

```bash
# Verificar que el puerto esté expuesto correctamente
docker port history_masters_api

# Verificar los logs
docker logs -f history_masters_api

# Probar la conexión
curl http://localhost:8080/health
```

### Limpiar todo y empezar de nuevo

```bash
# Detener y eliminar todo (contenedores y volúmenes)
docker-compose down -v

# Eliminar imágenes huérfanas y caché
docker system prune -a

# Levantar de nuevo
docker-compose up -d --build
```

## 📊 Monitoreo

### Ver uso de recursos

```bash
# Ver estadísticas en tiempo real
docker stats

# Ver solo los contenedores de este proyecto
docker stats history_masters_api history_masters_db
```

## 🔐 Seguridad

- El contenedor corre con un usuario no-root (`appuser`)
- No se incluyen archivos `.env` en la imagen (ver `.dockerignore`)
- Se usa una imagen base Alpine por su tamaño reducido y seguridad
- El binario se compila de forma estática sin CGO

## 📝 Notas Adicionales

- El puerto por defecto es `8080`, puedes cambiarlo en el archivo `.env`
- Los datos de PostgreSQL persisten en un volumen Docker
- Para desarrollo local, es recomendable usar `docker-compose`
- Para producción, considera usar orquestadores como Kubernetes

