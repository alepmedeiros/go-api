# 📌 Documentação da API em Go
## 📌 Visão Geral
Esta API foi desenvolvida em Go (Golang) seguindo as melhores práticas de Clean Architecture, SOLID, DDD e TDD. A API inclui:

✅ Autenticação JWT

✅ CRUD completo de usuários

✅ Arquitetura modularizada e escalável

✅ Testes automatizados para TDD

```plaintext
go-api/
│
├── auth/                # authenticações com JWT
│   ├── auth.go        
│
├── controlle/              # Camada que expoe endpoints
│   ├── auth_controller.go
│   │── teste_controller.go
│   │── usuario_controller.go
│   
│── middleware/        # 📌 Bibliteca de validação dos endpoints
│   |── middleware.go
│
│── models/           # 📌 Camada dos models
│   ├── usuario.go
│   
|── resources/          # 📌 Camada de configuração com banco de dados
|   ├── postgre_conectar.go
|── routes/            # camada das rotas
|   ├── authentication_routes.go
|   ├── teste_routes.go
|   ├── usuario_routes.go
|   
│── main.go            # Inicia o servidor
├── go.mod                 # Gerenciamento de dependências
```

## 🚀 Configuração e Instalação
#### 1️⃣ Clonando o repositório
```sh
git clone https://github.com/seu-usuario/go-api.git
cd go-api
```

#### 2️⃣ Instalando dependências
```sh
go mod tidy
```

## 📌 Rotas da API
#### 🔐 Autenticação
##### 🔹 Login (POST /auth/ogin)
```json
{
  "email": "admin@email.com",
  "senha": "123456"
}
```
#### 📌 Retorno: Token JWT
##### 👤 Usuários
###### 🔹 Criar usuário (POST /usuarios)
```json
{
  "nome": "Carlos",
  "email": "carlos@email.com",
  "senha": "123456"
}
```
🔹 Listar usuários (GET /usuarios)

🔹 Buscar usuário (GET /usuarios/:id)

🔹 Atualizar usuário (PUT /usuarios/:id)

🔹 Deletar usuário (DELETE /usuarios/:id)

### ✅ Testes Automatizados (TDD)

```sh
go test ./...
```
## 📌 Tecnologias Utilizadas

✅ Go (Golang) 1.20+

✅ Gin (framework web)

✅ GORM (ORM para Go)

✅ PostgreSQL

✅ JWT para autenticação

✅ Testes com `testing`

## 📜 Licença
Este projeto é open-source e está disponível sob a licença MIT.
