# 📌 Documentação da API em Go
## 📌 Visão Geral
Esta API foi desenvolvida em Go (Golang) seguindo as melhores práticas de Clean Architecture, SOLID, DDD e TDD. A API inclui:

✅ Autenticação JWT

✅ CRUD completo de usuários

✅ Arquitetura modularizada e escalável

✅ Testes automatizados para TDD

```plaintext
go-api/
├── cmd/                   # Ponto de entrada principal
│   ├── main.go            # Inicia o servidor
│
├── config/                # Configurações gerais (banco, env, etc.)
│   ├── database.go        # Configuração do banco de dados
│
├── internal/              # Código interno seguindo DDD
│   ├── domain/            # 📌 Camada de Domínio (Entidades e Regras de Negócio)
│   │   ├── usuario.go     # Struct de usuário
│   │
│   ├── repository/        # 📌 Camada de Repositórios (Acesso ao banco)
│   │   ├── usuario_repo.go
│   │
│   ├── usecase/           # 📌 Camada de Casos de Uso (Regras de Negócio)
│   │   ├── usuario_usecase.go
│   │
│   ├── handler/           # 📌 Camada de Handlers (Interface com a Web)
│   │   ├── usuario_handler.go
│
├── pkg/                   # 📌 Pacotes compartilháveis
│   ├── jwt/               # Biblioteca de autenticação JWT
│   │   ├── jwt.go
│
├── test/                  # 📌 Testes Unitários e de Integração
│   ├── usuario_test.go
│
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
