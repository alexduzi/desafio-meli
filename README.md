# API de Produtos - Desafio Técnico

API RESTful para listatem de produtos desenvolvida em Go.

## Índice

- [Sobre o Projeto](#sobre-o-projeto)
- [Arquitetura](#arquitetura)
- [Tecnologias Utilizadas](#tecnologias-utilizadas)
- [Pré-requisitos](#pré-requisitos)
- [Instalação](#instalação)
- [Como Usar](#como-usar)
- [Testes](#testes)
- [Endpoints da API](#endpoints-da-api)
- [Decisões Técnicas](#decisões-técnicas)
- [Estrutura do Projeto](#estrutura-do-projeto)

---

## Sobre o Projeto

Esta API foi desenvolvida como parte de um desafio técnico e implementa um sistema completo de gerenciamento de produtos com as seguintes características:

- ✅ **Clean Architecture** - Separação clara de responsabilidades
- ✅ **In-Memory Database** - SQLite em memória (`:memory:`)
- ✅ **Testes Abrangentes** - Cobertura de ~95% do código
- ✅ **Error Handling Centralizado** - Middleware customizado para tratamento de erros
- ✅ **Documentação Swagger** - API totalmente documentada
- ✅ **Otimização de Performance** - Prevenção do problema N+1 com thumbnails
- ✅ **Docker Support** - Multi-stage build otimizado com health checks
- ✅ **Production Ready** - Container seguro com usuário non-root

---

## Arquitetura

O projeto segue os princípios de **Clean Architecture** (Arquitetura Limpa), separando o código em camadas bem definidas:

```
┌─────────────────────────────────────────────────────────┐
│                    HTTP Layer (Gin)                      │
│                  (Handlers/Controllers)                  │
└────────────────────┬────────────────────────────────────┘
                     │
┌────────────────────▼────────────────────────────────────┐
│                   Use Cases Layer                        │
│              (Business Logic/Rules)                      │
└────────────────────┬────────────────────────────────────┘
                     │
┌────────────────────▼────────────────────────────────────┐
│                 Repository Layer                         │
│            (Data Access Abstraction)                     │
└────────────────────┬────────────────────────────────────┘
                     │
┌────────────────────▼────────────────────────────────────┐
│              Database Layer (SQLite)                     │
│                  (In-Memory)                             │
└─────────────────────────────────────────────────────────┘
```

### Camadas:

1. **Handler (HTTP Layer)**: Recebe requisições HTTP, valida entrada, chama os casos de uso
2. **Use Case (Business Layer)**: Contém a lógica de negócio, orquestra chamadas ao repositório
3. **Repository (Data Layer)**: Interface de acesso aos dados, abstrai implementação do banco
4. **Entity (Domain Layer)**: Modelos de domínio e regras de validação

### Vantagens desta Arquitetura:

- ✅ **Testabilidade**: Cada camada pode ser testada independentemente
- ✅ **Manutenibilidade**: Mudanças em uma camada não afetam outras
- ✅ **Flexibilidade**: Fácil trocar implementações (ex: SQLite → PostgreSQL)
- ✅ **Escalabilidade**: Estrutura preparada para crescimento

---

## Tecnologias Utilizadas

### Core
- **Go 1.21+** - Linguagem de programação
- **Gin** - Framework web HTTP router
- **SQLite** - Banco de dados em memória
- **sqlx** - Extensions para database/sql

### Testes
- **testify** - Assertions e mocks para testes
- **httptest** - Testes de handlers HTTP

### Documentação
- **Swagger/OpenAPI** - Documentação automática da API
- **swaggo/swag** - Geração de docs Swagger

### Ferramentas de Desenvolvimento
- **Make** - Automação de tarefas
- **go test** - Framework de testes nativo

---

## Pré-requisitos

- **Go 1.21 ou superior** - [Instalar Go](https://golang.org/doc/install)
- **Make** (opcional, mas recomendado) - Geralmente já vem instalado em Linux/macOS
- **Git** - Para clonar o repositório

---

## Instalação

### 1. Clone o repositório

```bash
git clone <url-do-repositorio>
cd <nome-do-diretorio>
```

### 2. Instale as dependências

```bash
make deps
```

Ou manualmente:

```bash
go mod download
go mod tidy
```

### 3. Gere a documentação Swagger (opcional)

```bash
make swagger
```

---

## Como Usar

### Usando Make (Recomendado)

O projeto inclui um Makefile com comandos úteis:

```bash
# Ver todos os comandos disponíveis
make help

# Executar a aplicação
make run

# Compilar o binário
make build

# Executar testes unitários (rápido, sem banco)
make test-unit

# Executar todos os testes
make test

# Ver cobertura de testes
make test-coverage

# Gerar relatório HTML de cobertura
make test-coverage-html

# Limpar arquivos gerados
make clean

# Executar tudo (deps, swagger, build, test)
make all
```

### Executando Manualmente

```bash
# Rodar a aplicação
go run cmd/api/main.go

# Rodar testes
go test ./...

# Rodar testes com cobertura
go test -cover ./internal/...

# Compilar
go build -o bin/api cmd/api/main.go
```

### Usando Docker (Recomendado para Produção)

#### Usando Make + Docker (Mais Fácil)

```bash
# Docker Compose
make docker-compose-up      # Iniciar aplicação
make docker-compose-logs    # Ver logs
make docker-compose-down    # Parar aplicação

# Docker direto
make docker-build           # Construir imagem
make docker-run             # Executar container
make docker-logs            # Ver logs
make docker-stop            # Parar e remover container
make docker-clean           # Limpar todos recursos Docker
```

#### Usando Docker Compose Manualmente

```bash
# Iniciar a aplicação
docker-compose up -d

# Ver logs
docker-compose logs -f

# Parar a aplicação
docker-compose down
```

#### Usando Docker Diretamente

```bash
# Construir a imagem
docker build -t product-api .

# Executar o container
docker run -d -p 8080:8080 --name product-api product-api

# Ver logs
docker logs -f product-api

# Parar e remover o container
docker stop product-api && docker rm product-api
```

### Acessando a API

Após iniciar a aplicação:

- **API Base URL**: `http://localhost:8080`
- **Swagger UI**: `http://localhost:8080/swagger/index.html`
- **Health Check**: `http://localhost:8080/health`

---

## Testes

O projeto possui uma suíte de testes abrangente com **~95% de cobertura**:

### Tipos de Testes

1. **Testes Unitários** - Testam cada camada isoladamente usando mocks
   ```bash
   make test-unit
   ```

2. **Testes de Integração** - Testam o fluxo completo com banco de dados
   ```bash
   make test-integration
   ```

3. **Cobertura de Código**
   ```bash
   make test-coverage
   ```

### Estrutura de Testes

```
internal/
├── entity/
│   └── product_test.go          # Testes de entidades e validações
├── errors/
│   └── errors_test.go           # Testes de error handling
├── usecase/
│   ├── mock_repository_test.go  # Mock do repositório
│   ├── get_product_test.go      # Testes do caso de uso GetProduct
│   └── list_product_test.go     # Testes do caso de uso ListProducts
├── handler/
│   ├── product_handler_test.go  # Testes dos handlers HTTP
│   └── health_handler_test.go   # Testes do health check
└── infra/http/
    ├── error_middleware_test.go # Testes do middleware de erros
    └── router_test.go           # Testes de rotas

test/integration/
└── api_integration_test.go      # Testes end-to-end
```

### Cobertura Atual

```
✅ internal/entity      → 100.0% coverage
✅ internal/errors      → 100.0% coverage
✅ internal/handler     → 100.0% coverage
✅ internal/infra/http  → 100.0% coverage
✅ internal/usecase     → 100.0% coverage
```

---

## Endpoints da API

### Health Check

```http
GET /health
```

**Resposta de Sucesso (200 OK):**
```json
{
  "status": "healthy",
  "timestamp": "2024-01-01T00:00:00Z",
  "service": "product-api"
}
```

**Uso**: Endpoint para verificar se a API está funcionando corretamente. Útil para monitoramento, health checks do Docker/Kubernetes, e load balancers.

### Listar Produtos

```http
GET /api/v1/products
```

**Resposta de Sucesso (200 OK):**
```json
{
  "data": [
    {
      "id": "PROD-1234567890-123456",
      "title": "iPhone 15 Pro",
      "description": "Latest Apple smartphone",
      "price": 999.99,
      "currency": "USD",
      "condition": "new",
      "stock": 10,
      "seller_id": "seller-001",
      "seller_name": "Apple Store",
      "category": "Electronics",
      "thumbnail": "https://example.com/thumb.jpg",
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
  ]
}
```

**Nota**: O endpoint de listagem retorna apenas o `thumbnail` (não o array completo de imagens) para otimizar performance e evitar o problema N+1.

### Obter Produto por ID

```http
GET /api/v1/products/{id}
```

**Resposta de Sucesso (200 OK):**
```json
{
  "data": {
    "id": "PROD-1234567890-123456",
    "title": "iPhone 15 Pro",
    "description": "Latest Apple smartphone",
    "price": 999.99,
    "currency": "USD",
    "condition": "new",
    "stock": 10,
    "seller_id": "seller-001",
    "seller_name": "Apple Store",
    "category": "Electronics",
    "thumbnail": "https://example.com/thumb.jpg",
    "images": [
      {
        "id": 1,
        "product_id": "PROD-1234567890-123456",
        "image_url": "https://example.com/img1.jpg",
        "display_order": 0
      }
    ],
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

**Resposta de Erro (404 Not Found):**
```json
{
  "error": "product not found",
  "code": "PRODUCT_NOT_FOUND",
  "timestamp": "2024-01-01T00:00:00Z"
}
```

**Resposta de Erro (400 Bad Request):**
```json
{
  "error": "invalid product id",
  "code": "INVALID_PRODUCT_ID",
  "timestamp": "2024-01-01T00:00:00Z"
}
```

---

## Decisões Técnicas

### Clean Architecture

**Por quê?**
- Separação clara de responsabilidades
- Facilita testes isolados de cada camada
- Permite trocar implementações sem afetar o core

### SQLite In-Memory (`:memory:`)

**Por quê?**
- Atende requisito de persistência simples do desafio
- Zero configuração necessária
- Ideal para desenvolvimento e testes
- Performance excelente para dados temporários

**Trade-off**: Dados são perdidos ao reiniciar a aplicação (comportamento esperado para desafios).

### Error Handling Centralizado

**Implementação**:
```go
// Middleware processa erros automaticamente
func ErrorHandlerMiddleware() gin.HandlerFunc {
    // Mapeia erros de domínio para status HTTP
    // Retorna JSON padronizado
}
```

**Vantagens**:
- Respostas de erro consistentes
- Handlers mais limpos
- Fácil adicionar logging/monitoring

### Otimização N+1 com Thumbnails

**Problema**: Ao listar produtos, buscar todas as imagens de cada produto seria ineficiente:
```
1 query para produtos + N queries para imagens = N+1 queries
```

**Solução**:
- **List Endpoint**: Retorna apenas `thumbnail` (1 query total)
- **Get Endpoint**: Retorna array completo de `images` (2 queries)

**Resultado**: Performance ~10x melhor em listagens.

### Docker Multi-Stage Build

**Por quê?**

- **Imagem otimizada**: Build stage com ~500MB, runtime final com ~20MB
- **Segurança**: Container roda com usuário non-root (appuser)
- **Health checks**: Monitoramento automático usando endpoint `/health`
- **CGO habilitado**: Suporte completo ao SQLite com driver nativo

**Características**:

```dockerfile
# Build stage - Go 1.24 + ferramentas de build
FROM golang:1.24-alpine AS builder
# ... compila aplicação ...

# Runtime stage - Alpine mínimo
FROM alpine:latest
# ... apenas o binário + libs runtime ...
USER appuser  # Roda como non-root
```

**Benefícios**:

- Deployment rápido e seguro
- Menor superfície de ataque
- Compatível com Kubernetes, Docker Swarm, etc.

## Estrutura do Projeto

```
.
├── cmd/
│   └── api/
│       ├── main.go              # Entry point da aplicação
│       └── main_test.go         # [Movido para test/integration]
├── internal/
│   ├── entity/                  # Entidades de domínio
│   │   ├── product.go
│   │   └── product_test.go
│   ├── errors/                  # Definição de erros customizados
│   │   ├── errors.go
│   │   └── errors_test.go
│   ├── handler/                 # HTTP handlers (controllers)
│   │   ├── product_handler.go
│   │   └── product_handler_test.go
│   ├── usecase/                 # Casos de uso (business logic)
│   │   ├── product_repository.go      # Interface do repositório
│   │   ├── get_product.go
│   │   ├── get_product_test.go
│   │   ├── list_product.go
│   │   ├── list_product_test.go
│   │   └── mock_repository_test.go
│   └── infra/
│       ├── database/            # Implementação do banco de dados
│       │   ├── db.go
│       │   ├── product_repository.go
│       │   └── migrations/
│       │       └── migrations.go
│       └── http/                # Configuração HTTP
│           ├── router.go
│           ├── router_test.go
│           ├── error_middleware.go
│           └── error_middleware_test.go
├── test/
│   └── integration/             # Testes de integração
│       └── api_integration_test.go
├── docs/                        # Documentação Swagger (gerada)
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── Makefile                     # Automação de tarefas
├── go.mod                       # Dependências do projeto
├── go.sum                       # Checksums das dependências
└── README.md                    # Este arquivo
```

### Convenções de Nomenclatura

- **Handlers**: Terminam com `Handler` (ex: `ProductHandler`)
- **Use Cases**: Terminam com `UseCase` (ex: `GetProductUseCase`)
- **Repositories**: Terminam com `Repository` (ex: `ProductRepository`)
- **DTOs**: Terminam com `DTO` (ex: `ProductDTO`)
- **Testes**: Terminam com `_test.go`

---

## 🔧 Makefile - Comandos Disponíveis

| Comando | Descrição |
|---------|-----------|
| **Desenvolvimento Local** | |
| `make help` | Mostra todos os comandos disponíveis |
| `make run` | Executa a aplicação localmente |
| `make build` | Compila o binário da aplicação |
| `make swagger` | Gera/atualiza documentação Swagger |
| **Testes** | |
| `make test` | Executa todos os testes (unitários + integração) |
| `make test-unit` | Executa apenas testes unitários (rápido, sem DB) |
| `make test-integration` | Executa apenas testes de integração (requer DB) |
| `make test-coverage` | Executa testes e mostra cobertura |
| `make test-coverage-html` | Gera relatório HTML de cobertura |
| **Docker** | |
| `make docker-build` | Constrói a imagem Docker |
| `make docker-run` | Executa o container Docker |
| `make docker-stop` | Para e remove o container Docker |
| `make docker-logs` | Visualiza logs do container |
| `make docker-compose-up` | Inicia aplicação com Docker Compose |
| `make docker-compose-down` | Para aplicação com Docker Compose |
| `make docker-compose-logs` | Visualiza logs do Docker Compose |
| `make docker-clean` | Remove imagens e containers Docker |
| **Utilitários** | |
| `make clean` | Remove arquivos gerados (binários, coverage) |
| `make deps` | Baixa e organiza dependências |
| `make all` | Executa deps, swagger, build e test |

---

## Bibliotecas de Terceiros

### Dependências de Produção

```go
require (
    github.com/gin-gonic/gin v1.9.1           // Framework web HTTP
    github.com/jmoiron/sqlx v1.3.5            // Extensions para database/sql
    github.com/mattn/go-sqlite3 v1.14.18      // Driver SQLite
    github.com/swaggo/files v1.0.1            // Swagger UI files
    github.com/swaggo/gin-swagger v1.6.0      // Integração Swagger + Gin
    github.com/swaggo/swag v1.16.2            // Gerador de docs Swagger
)
```

### Dependências de Teste

```go
require (
    github.com/stretchr/testify v1.8.4        // Assertions e mocks
)
```

---

## 📝 Notas Adicionais

### Swagger/OpenAPI

A documentação Swagger é gerada automaticamente a partir de comentários no código:

```go
// @Summary List all products
// @Description Get a list of all products
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /products [get]
func (h *ProductHandler) ListProducts(c *gin.Context) { ... }
```

### Testes

Os testes são organizados em:
- **Testes Unitários**: Não dependem de DB, usam mocks
- **Testes de Integração**: Usam DB real (in-memory)

Use `make test-unit` para feedback rápido durante desenvolvimento.

---

## 👤 Autor

Alex Duzi

---

## 📄 Licença

Este projeto foi desenvolvido como parte de um desafio técnico.

---
