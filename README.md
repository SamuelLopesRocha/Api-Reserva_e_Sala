# 🏫 API de Reservas

## 📌 Descrição

Esta API é responsável pela criação e listagem de **reservas** e **salas**.  
Ela valida a existência da sala antes de permitir a criação de uma reserva.  
Além disso, ao criar uma turma, é possível atribuir uma sala específica.

Os dados são armazenados localmente utilizando **SQLite**.  
A aplicação foi desenvolvida com **Go (Gin Framework)**, segue o padrão **MVC** e está totalmente **conteinerizada com Docker**, facilitando sua integração em um ambiente de **microsserviços**.

---

## 🚀 Como Executar com Docker

### Pré-requisitos

- Docker e Docker Compose instalados
- Ter a API de Turmas ativa

### Passos

```bash
git clone https://github.com/MarceloHenrique1711/Reserva-de-salas-Api.git
cd Reserva-de-salas-Api
docker-compose up
🌐 Integração com Microsserviços
A API de Reservas faz parte de um ecossistema baseado em microsserviços. Atualmente, ela se comunica com o seguinte serviço externo:

🔗 Serviço de Turmas
Descrição: Permite a criação de turmas com possibilidade de associar uma sala.

Integração: A API de Turmas realiza uma requisição GET para verificar salas disponíveis.

✅ Endpoint consultado:
GET http://localhost:6000/salas
Se a sala for válida, a turma pode ser criada com uma sala atribuída. Caso contrário, a turma será criada, mas sem sala.

📬 Endpoints
🔹 GET /sala/
Cria uma nova sala, com turma podendo ser atribuida a uma sala.

Exemplo de corpo (JSON):
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
Resposta:
{
  "mensagem": "Turma criada com sucesso"
}

🔹 GET /reserva/
Lista todas as reservas realizadas.

Exemplo de resposta (JSON):
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
  Resposta:
{
  "mensagem": "Turma criada com sucesso"
}


🛠️ Tecnologias Utilizadas
Go 1.24

Gin (Framework)

SQLite

Docker / Docker Compose

Padrão MVC

🏗️ Arquitetura
A aplicação segue o padrão MVC (Model-View-Controller).

Estrutura de Diretórios
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
