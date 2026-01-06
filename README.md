# periferia_it_social_network


proyecto fullstack con backend desarrollado en go y frontend en react.


## stacks

- go + Gin + GROM
- PostgreSQL
- Docker / Docker Compose
- JWT Auth
- React + vite


## Como levantar el proyecto

```bash

docker-compose up --build
Backend: http://localhost:8080
Frontend: http://localhost:5173

```


## Credenciales de prueba

```bash
Email: juan@test.com
Password: 123456

```



## Funcionalidades


- Login con JWT
- Middleware de autenticación
- Listar posts
- Crear post con auth
- Like a post con auth
- Logout solo para el frontend
- Seeder automático (usuarios + posts)

## Endpoints principales

- POST /login
- GET /posts
- POST /posts
- POST /posts/:id/like