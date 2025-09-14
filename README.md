# 🏫 API de Reservas

API desenvolvida em **Go** com o framework **Gin**, responsável pela criação e listagem de **salas** e **reservas**. Valida a existência de uma sala antes de permitir a criação de uma reserva e permite associar salas a turmas via integração com a **API School System**.

## 📌 Descrição Geral

Esta API:

- Gerencia salas e reservas.
- Valida a existência de salas ao criar reservas.
- Permite criar turmas já associadas a uma sala (via outro serviço).
- Armazena dados localmente com **SQLite**.
- Segue o padrão arquitetural **MVC**.
- Está totalmente conteinerizada com **Docker**.

## 🚀 Como Executar com Docker

### Pré-requisitos

- Docker e Docker Compose instalados
- Ter a **API de Turmas** ativa

### Passos para Execução

```bash
git clone https://github.com/MarceloHenrique1711/Reserva-de-salas-Api.git
cd Reserva-de-salas-Api

# Criar rede docker compartilhada entre APIs
docker network create projeto-apis

# Iniciar os containers
docker-compose up
```

## 🌐 Integração com Microsserviços

Esta API integra-se com o seguinte serviço externo:

### 🔗 Serviço de Turmas

- **Descrição**: Criação de turmas com possibilidade de associar uma sala.
- **Integração**: A API de Turmas faz requisições GET para a API de Reservas.
- **Endpoint Consultado**:
  ```http
  GET http://api_sala:6000/salas
  ```

- Se a sala for válida, a turma será criada com sala atribuída.
- Se inválida, a turma será criada sem sala.

## 📬 Endpoints

### 🔹 `GET /sala/`

Cria uma nova sala (pode incluir dados da turma associada).

**Exemplo de JSON enviado:**

```json
{
  "ativo": true,
  "recursos": "Projetor",
  "sala_id": 1,
  "turma": {
    "ativo": true,
    "descricao": "Turma A",
    "professor_id": 1,
    "sala_id": 1,
    "turma_id": 10
  }
}
```

**Resposta esperada:**

```json
{
  "mensagem": "Turma criada com sucesso"
}
```

---

### 🔹 `GET /reserva/`

Lista todas as reservas cadastradas.

**Exemplo de resposta:**

```json
{
  "reserva_id": 1,
  "data_reserva": "2031-10-14",
  "descricao": "Reunião",
  "sala_id": 1,
  "sala": {
    "sala_id": 1,
    "recursos": "Computador",
    "ativo": true
  }
}
```

## 🛠️ Tecnologias Utilizadas

- Go 1.24
- Gin (Framework Web)
- SQLite (Banco de Dados)
- Docker & Docker Compose
- Arquitetura MVC

## 🏗️ Estrutura do Projeto

```
.
├── config/ 
│   └── config.go   
├── controller/ 
│   ├── reserva_controller.go
│   └── sala_controller.go           
├── models/ 
│   ├── reserva_model.go 
│   └── sala_model.go 
├── repository/ 
│   ├── reserva_repository.go 
│   └── sala_repository.go
├── route/ 
│   ├── reserva_route.go 
│   └── sala_route.go     
├── docker/      
│   ├── Dockerfile
│   ├── docker-compose.yml
│   ├── .dockerignore
│   └── .gitignore                  
├── app.go
├── banco.db 
├── go.mod
└── go.sum
```
