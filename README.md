🚀 Go User System

Sistema completo de autenticação e gerenciamento de usuários, desenvolvido com Go (API REST) e ASP.NET Core MVC (Frontend), utilizando JWT, PostgreSQL e Docker.

🎯 Objetivo: demonstrar domínio real de backend, frontend, autenticação, containers e boas práticas

🧩 Arquitetura do Projeto

O sistema é dividido em três containers independentes:

┌─────────────┐     ┌─────────────┐     ┌──────────────┐
│  Frontend   │ --> │   API Go    │ --> │  PostgreSQL  │
│ ASP.NET MVC │     │   Gin/JWT   │     │   Database   │
└─────────────┘     └─────────────┘     └──────────────┘

📦 Containers

Frontend: ASP.NET Core MVC (C#)

Backend: Go + Gin (API REST)

Database: PostgreSQL

Orquestração: Docker Compose

🛠️ Tecnologias Utilizadas
Backend (API)

Go (Golang)

Gin Framework

JWT (JSON Web Token)

PostgreSQL

GORM

Docker

Frontend

ASP.NET Core MVC (.NET 8)

Razor Views

Session + Middleware de autenticação

Docker

🔐 Autenticação (JWT)

O sistema utiliza JWT stateless para autenticação:

Usuário se cadastra (/auth/register)

Usuário faz login (/auth/login)

API retorna um JWT assinado

Token é enviado no header:

Authorization: Bearer <token>


Middleware valida o token e extrai o user_id

Endpoints protegidos só funcionam com token válido

📡 Endpoints da API
🔓 Públicos
Método	Rota	Descrição
POST	/auth/register	Cadastro de usuário
POST	/auth/login	Login e geração JWT
GET	/health	Health check
🔐 Protegidos (JWT)
Método	Rota	Descrição
GET	/me	Dados do usuário logado
GET	/users	Lista usuários
GET	/users/:id	Busca usuário por ID
PUT	/users/:id	Atualiza usuário
DELETE	/users/:id	Remove usuário
🧑‍💻 Testes via Terminal (curl)
📌 Cadastro de usuário
curl -X POST http://localhost:8080/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Geovanni Souza",
    "email": "geo@email.com",
    "password": "123456"
  }'

🔑 Login (gera token)
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "geo@email.com",
    "password": "123456"
  }'

👤 Buscar perfil do usuário logado
curl http://localhost:8080/me \
  -H "Authorization: Bearer SEU_TOKEN_AQUI"

🖥️ Frontend (ASP.NET MVC)

O frontend consome a API e implementa:

Tela de login

Controle de sessão

Middleware próprio para proteger rotas

Tela de perfil com dados do usuário autenticado

📍 Acesso:

http://localhost:5103

🐳 Docker & Execução
▶️ Subir todo o sistema
docker compose up --build

🔄 Rebuild limpo
docker compose down -v
docker compose build --no-cache
docker compose up

🧠 Boas Práticas Aplicadas

Separação de responsabilidades (handlers, services, repositories)

JWT stateless

Middleware de autenticação

Docker multi-stage build

Containers independentes

Código organizado e legível

Pensado para produção

🎯 O que este projeto demonstra

✔ Backend real em Go
✔ API REST profissional
✔ Autenticação JWT
✔ Integração com banco de dados
✔ Frontend funcional
✔ Docker na prática
✔ Organização de projeto
✔ Pensamento arquitetural

📌 Próximos passos (roadmap)

 Validação avançada

 Refresh token

 Testes automatizados

 Deploy em cloud (Railway / VPS)

 CI/CD

👨‍💻 Autor

Geovanni Souza
Backend / Full Stack Developer

📎 GitHub: https://github.com/Geovannisouza23