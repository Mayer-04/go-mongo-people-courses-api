# PRUEBA TÉCNICA – Backend Junior/Mid con Go + MongoDB

>[!IMPORTANT]
> Puedes utilizar documentación oficial de Go, MongoDB, Gin/Fiber/Echo/Chi u otros frameworks permitidos. No usar foros ni IA.

## Contexto

Debes crear una **API REST** usando **Go (golang ≥ 1.20 hasta 1.24)** y **MongoDB**, para gestionar dos entidades:

- **Personas**
- **Cursos**

Relación:

*Cada persona puede estar inscrita en uno o varios cursos (relación 1:N).*

Objetivo:

- Definir los modelos usando structs de Go.

- Conectarte a MongoDB usando el driver oficial.

- Crear CRUD completo para ambas entidades.

- Implementar un endpoint especial con **aggregate** + **$lookup** + **$unwind**.

- Implementar **paginación**, **validaciones**, **middlewares**, y manejo estructurado de errores.

Framework sugerido (no obligatorio):

- Gin, Fiber o Echo

## Requisitos técnicos

**Debes usar:**

- Go 1.20–1.24

- MongoDB (driver oficial)

- Un framework HTTP para Go (Gin, Fiber, Echo…)

- `.env` para la conexión `MONGODB_URI`

- Contextos (`context.Context`) en todos los accesos a la DB

**Se valorará:**

- Uso de interfaces para los repositorios.

- Código modular y limpio.

- Manejo de errores con una estructura común.

- Logs.

- Índices de MongoDB.

## Modelos

Colección: **personas**

Campos mínimos:

| Campo    | Tipo       | Requerido | Descripción                |
| -------- | ---------- | --------- | -------------------------- |
| `nombre` | string     | sí        | Nombre de la persona       |
| `cedula` | string     | sí, única | Identificación             |
| `email`  | string     | opcional  | Email                      |
| `cursos` | []ObjectID | sí        | Cursos donde está inscrito |

## Curso

Colección: **cursos**

| Campo           | Tipo   | Requerido                               |
| --------------- | ------ | --------------------------------------- |
| `nombre`        | string | sí                                      |
| `codigo`        | string | sí, único                               |
| `descripcion`   | string | opcional                                |
| `duracionHoras` | int    | opcional                                |
| `nivel`         | string | opcional (básico, intermedio, avanzado) |

## Endpoints requeridos

### Personas

```bash
GET    /personas            (con paginación: ?page=1&limit=10)
POST   /personas
GET    /personas/:id
PUT    /personas/:id
DELETE /personas/:id   (obligatorio)
```

### Cursos

```bash
GET    /cursos
POST   /cursos
GET    /cursos/:id
PUT    /cursos/:id
DELETE /cursos/:id
```

>[!NOTE]
> Los POST/PUT deben validar con precisión los datos.

## Lógica especial

Implementar:

### GET /personas/cedula/:cedula/detalle-cursos

Debe devolver:

- Datos de la persona

- Lista completa de cursos donde está inscrita

**Requisito obligatorio:**

Debe usarse aggregate con `$lookup` + `$unwind` + `$group`.

Ejemplo de respuesta:

```json
{
  "persona": {
    "nombre": "María Gómez",
    "cedula": "987654321",
    "email": "maria@example.com"
  },
  "cursos": [
    {
      "nombre": "Go Básico",
      "codigo": "GO-101",
      "nivel": "Básico"
    },
    {
      "nombre": "MongoDB Avanzado",
      "codigo": "MDB-301",
      "nivel": "Avanzado"
    }
  ]
}
```

Si la cédula no existe, devolver:

```json
404 { "error": "Persona no encontrada" }
```

## Middleware obligatorio

Debes implementar dos **middlewares**:

### Request Logger

Debe imprimir:

- Método

- URL

- Tiempo de respuesta

- Código de estado

## API-Key Middleware

Toda petición debe incluir header:

```bash
X-API-KEY: <valor>
```

El valor debe validarse desde la variable de entorno:

```bash
API_KEY=miclave123
```

Si es incorrecta:

```bash
403 Forbidden
```

## Validaciones mínimas

### Para crear/actualizar Persona

- nombre obligatorio

- cedula obligatoria y única

- cursos debe contener al menos 1 curso existente

### Para crear/actualizar Curso

- nombre obligatorio

- codigo obligatorio y único

### Para todos los endpoints

- Manejar errores de MongoDB (duplicate key, formatos de ObjectID incorrectos, etc.)

- Devolver mensajes claros

## Estructura sugerida del proyecto

```bash
project/
 ├─ cmd/
 │   └─ server/
 │       └─ main.go
 ├─ internal/
 │   ├─ config/
 │   │   └─ mongo.go
 │   ├─ models/
 │   │   ├─ persona.go
 │   │   └─ curso.go
 │   ├─ repository/
 │   │   ├─ persona_repository.go
 │   │   └─ curso_repository.go
 │   ├─ services/
 │   │   ├─ persona_service.go
 │   │   └─ curso_service.go
 │   ├─ handlers/
 │   │   ├─ persona_handler.go
 │   │   └─ curso_handler.go
 │   └─ middlewares/
 │       ├─ logger.go
 │       └─ apikey.go
 └─ go.mod
```

## Índices obligatorios en MongoDB

Debes crear (automático o manual):

- Índice único para cedula en personas

- Índice único para codigo en cursos

- Índice por cédula para acelerar el endpoint especial

## Entregables

Debes entregar un **repositorio público** con:

- Código fuente

- **README** con:
    -Cómo correr el proyecto

    -Ejemplos de request/response (`POST/PUT`)

    -Variables de entorno necesarias

    -Explicación breve del aggregate usado

## Extras valorados (opcionales)

- Tests unitarios básicos (services o handlers)

- Dockerfile + docker-compose

- Manejo de errores centralizado

- Documentación con Swagger/OpenAPI
