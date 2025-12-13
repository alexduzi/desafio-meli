# API de Produtos - Desafio Técnico

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![Test Coverage](https://img.shields.io/badge/coverage-95%25-brightgreen)
![Docker](https://img.shields.io/badge/Docker-Supported-2496ED?style=flat&logo=docker)

API RESTful para listagem de produto desenvolvida em Go com Clean Architecture.

---

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

O projeto segue os princípios de **Clean Architecture** (Arquitetura Limpa), separando o código em camadas bem definidas.

Para visualizar os diagramas completos da arquitetura, consulte: [Architecture Diagrams](docs/architecture.md)

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
- **Docker** - Containerização

---

## Pré-requisitos

- **Go 1.21 ou superior** - [Instalar Go](https://golang.org/doc/install)
- **Make** (opcional, mas recomendado) - Geralmente já vem instalado em Linux/macOS
- **Docker** (opcional) - [Instalar Docker](https://docs.docker.com/get-docker/)
- **Git** - Para clonar o repositório

---

## Instalação

### 1. Clone o repositório

```bash
git clone <url-do-repositorio>
cd <nome-do-diretorio>
```

### 2. Configure o ambiente e instale dependências

```bash
# Opção 1: Usando Make (recomendado)
make setup

# Opção 2: Manualmente
cp .env.example .env
go mod download
go mod tidy
```

### 3. Gere a documentação Swagger (opcional)

```bash
make swagger
```

---

## Como Usar

### 🚀 Quick Start

```bash
# Setup inicial (primeira vez)
make setup

# Executar a aplicação
make run

# Ou usando Docker Compose
make docker-compose-up
```

### Usando Make (Recomendado)

O projeto inclui um Makefile com comandos úteis:

```bash
# Ver todos os comandos disponíveis
make help

# Setup inicial do projeto
make setup

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

# Executar tudo (setup, swagger, build, test)
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

### Usando Docker

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

### 📝 Exemplos de Uso com curl

```bash
# Health check
curl http://localhost:8080/health

# Listar todos os produtos
curl http://localhost:8080/api/v1/products

# Obter produto específico
curl http://localhost:8080/api/v1/products/MLB001

# Com formatação JSON (requer jq)
curl http://localhost:8080/api/v1/products | jq
```

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
   make test-coverage        # Console output
   make test-coverage-html   # HTML report
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

---

### Listar Produtos

```http
GET /api/v1/products
```

**Resposta de Sucesso (200 OK):**
```json
{
  "data": [
    {
      "id": "MLB001",
      "title": "iPhone 15 Pro Max 256GB - Titanium Blue",
      "description": "Latest Apple flagship smartphone...",
      "price": 1299.99,
      "currency": "USD",
      "condition": "new",
      "stock": 45,
      "seller_id": "SELLER001",
      "seller_name": "TechWorld Store",
      "category": "Electronics > Smartphones",
      "thumbnail": "https://images.unsplash.com/photo-1696446702230...",
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
  ]
}
```

**Nota**: O endpoint de listagem retorna apenas o `thumbnail` (não o array completo de imagens) para otimizar performance e evitar o problema N+1.

---

### Obter Produto por ID

```http
GET /api/v1/products/{id}
```

**Parâmetros:**
- `id` (path) - ID do produto (ex: MLB001)

**Resposta de Sucesso (200 OK):**
```json
{
  "data": {
    "id": "MLB001",
    "title": "iPhone 15 Pro Max 256GB - Titanium Blue",
    "description": "Latest Apple flagship smartphone...",
    "price": 1299.99,
    "currency": "USD",
    "condition": "new",
    "stock": 45,
    "seller_id": "SELLER001",
    "seller_name": "TechWorld Store",
    "category": "Electronics > Smartphones",
    "thumbnail": "https://images.unsplash.com/photo-1696446702230...",
    "images": [
      {
        "id": 1,
        "product_id": "MLB001",
        "image_url": "https://images.unsplash.com/photo-1696446702230...",
        "display_order": 0
      },
      {
        "id": 2,
        "product_id": "MLB001",
        "image_url": "https://images.unsplash.com/photo-1695048133142...",
        "display_order": 1
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

### 1. Clean Architecture com Inversão de Dependência

**Decisão**: Separar a aplicação em camadas distintas com inversão de dependência.

**Justificativa**:
- **Testabilidade**: Cada camada pode ser testada isoladamente com mocks
- **Flexibilidade**: Fácil trocar implementações (ex: SQLite → PostgreSQL)
- **Manutenibilidade**: Mudanças em uma camada não afetam outras

**Trade-offs**:
- Mais código boilerplate inicialmente
- Curva de aprendizado maior
- **Benefício**: Manutenibilidade e testabilidade no longo prazo compensam a complexidade inicial

---

### 2. SQLite In-Memory Database

**Decisão**: Usar SQLite com configuração `:memory:` ao invés de persistência em arquivo.

**Justificativa**:
- Atende ao requisito do desafio de "simular persistência de dados"
- Zero configuração necessária - funciona imediatamente em qualquer sistema
- Perfeito para desenvolvimento e testes
- Excelente performance para dados temporários

**Trade-offs**:
- Dados são perdidos ao reiniciar a aplicação (comportamento esperado para este desafio)
- Não adequado para produção (limitação reconhecida)
- **Benefício**: Simplicidade e portabilidade para um desafio técnico

---

### 3. Otimização N+1 com Thumbnails

**Problema Identificado**:
Ao listar produtos, buscar todas as imagens de cada produto criaria N+1 queries:
```
1 query para produtos + N queries para imagens = Problema de performance
```

**Solução Implementada**:
- **List endpoint**: Retorna apenas `thumbnail` (1 query total)
- **Detail endpoint**: Retorna array completo de `images` (2 queries)

**SQL Otimizado para Listagem**:
```sql
SELECT p.*, 
       (SELECT image_url FROM product_images 
        WHERE product_id = p.id 
        ORDER BY display_order ASC 
        LIMIT 1) as thumbnail
FROM products p
```

**Impacto**:
- ~10x melhor performance em operações de listagem
- Menor tamanho de payload
- Melhor experiência do usuário

---

### 4. Error Handling Centralizado

**Decisão**: Implementar middleware de tratamento de erros ao invés de tratar erros em cada handler.

**Benefícios**:
- Formato de resposta de erro consistente em todos os endpoints
- Código de handlers mais limpo (apenas retornam erros)
- Ponto único para logging/monitoring
- Fácil estender com serviços de rastreamento de erros

**Implementação**:
```go
// Handler apenas retorna o erro
func (h *Handler) GetProduct(c *gin.Context) {
    result, err := h.useCase.Execute(input)
    if err != nil {
        _ = c.Error(err)  // Middleware cuida do resto
        return
    }
    c.JSON(200, result)
}
```

---

### 5. Docker Multi-Stage Build

**Decisão**: Usar multi-stage build com práticas de segurança.

**Características**:
- Build stage: ~500MB (compilador Go + ferramentas)
- Runtime stage: ~20MB (Alpine + binário apenas)
- Execução com usuário non-root
- Integração com health check endpoint

**Benefícios de Segurança**:
- Superfície de ataque mínima
- Sem ferramentas desnecessárias na imagem de produção
- Segue best practices do Docker

---

## Estrutura do Projeto

```
.
├── cmd/
│   └── api/
│       └── main.go              # Entry point da aplicação
├── internal/
│   ├── entity/                  # Entidades de domínio
│   │   ├── product.go
│   │   └── product_test.go
│   ├── errors/                  # Definição de erros customizados
│   │   ├── errors.go
│   │   └── errors_test.go
│   ├── handler/                 # HTTP handlers (controllers)
│   │   ├── product_handler.go
│   │   ├── product_handler_test.go
│   │   ├── health_handler.go
│   │   └── health_handler_test.go
│   ├── usecase/                 # Casos de uso (business logic)
│   │   ├── product_repository.go
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
│       │       ├── 001_schema.sql
│       │       ├── 002_seed.sql
│       │       └── migrations.go
│       └── http/                # Configuração HTTP
│           ├── router.go
│           ├── router_test.go
│           ├── error_middleware.go
│           └── error_middleware_test.go
├── test/
│   └── integration/             # Testes de integração
│       └── api_integration_test.go
├── docs/                        # Documentação
│   ├── architecture.md          # Diagramas de arquitetura
│   ├── docs.go                  # Swagger gerado
│   ├── swagger.json
│   └── swagger.yaml
├── .env.example                 # Exemplo de variáveis de ambiente
├── .env                         # Configurações locais (git ignored)
├── Dockerfile                   # Multi-stage build
├── docker-compose.yml           # Orquestração Docker
├── Makefile                     # Automação de tarefas
├── go.mod                       # Dependências do projeto
├── go.sum                       # Checksums das dependências
└── README.md                    # Este arquivo
```

---

## 🔧 Makefile - Comandos Disponíveis

| Comando | Descrição |
|---------|-----------|
| **Setup** | |
| `make setup` | Setup inicial do projeto (copia .env, instala deps) |
| **Desenvolvimento Local** | |
| `make run` | Executa a aplicação localmente |
| `make build` | Compila o binário da aplicação |
| `make swagger` | Gera/atualiza documentação Swagger |
| **Testes** | |
| `make test` | Executa todos os testes (unitários + integração) |
| `make test-unit` | Executa apenas testes unitários (rápido, sem DB) |
| `make test-integration` | Executa apenas testes de integração |
| `make test-coverage` | Executa testes e mostra cobertura |
| `make test-coverage-html` | Gera relatório HTML de cobertura |
| **Docker** | |
| `make docker-build` | Constrói a imagem Docker |
| `make docker-run` | Executa o container Docker |
| `make docker-stop` | Para e remove o container |
| `make docker-logs` | Visualiza logs do container |
| `make docker-compose-up` | Inicia aplicação com Docker Compose |
| `make docker-compose-down` | Para aplicação Docker Compose |
| `make docker-clean` | Remove imagens e containers |
| **Utilitários** | |
| `make clean` | Remove arquivos gerados |
| `make deps` | Baixa e organiza dependências |
| `make all` | Executa setup, swagger, build e test |

---

## 📚 Documentação Adicional

- **[Diagramas de Arquitetura](docs/architecture.html)** - Visualização interativa da arquitetura
  - **Como visualizar:** Abra o arquivo `docs/architecture.html` em qualquer navegador
  - Também disponível em Markdown: [docs/architecture.md](docs/architecture.md)
- **[Swagger UI](http://localhost:8080/swagger/index.html)** - Documentação interativa da API (quando o servidor está rodando)

---

## 👤 Autor

Alex Duzi - duzihd@gmail.com

---

## 📄 Licença

Este projeto foi desenvolvido como parte de um desafio técnico.

---