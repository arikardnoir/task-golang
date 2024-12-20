# task-golang

Este projeto é uma **aplicação web construída em Go**, projetada para ser facilmente implantada utilizando **Docker** ou **Docker Compose**. O objetivo é fornecer um backend simples e escalável para aplicações web, que pode ser executado de forma rápida e isolada em containers.

### Bibliotecas Usadas

- **GORM**: ORM utilizado para interações com o banco de dados SQLite.
- **SQLite**: Banco de dados leve utilizado para persistência de dados.
- **Docker**: Usado para contêinerizar a aplicação e garantir que ela seja executada de maneira isolada e consistente em qualquer ambiente.
- **JWT (JSON Web Token)**: Usado para autenticação e autorização segura na aplicação.
- **Testify**: Biblioteca usada para facilitar a criação de testes automatizados, com suporte a asserções e mocks.

---

## Como Clonar o Repositório

Para começar, clone o repositório para sua máquina local:

      git clone https://github.com/arikardnoir/task-golang
      cd task-golang
      
### Como Executar o Projeto

Existem duas maneiras de executar este projeto: utilizando **Docker** diretamente ou **Docker Compose**.

#### 1. **Executando o Projeto com Docker Simples**

Se você já tem o **Docker** instalado, siga os passos abaixo:

1. **Construir a Imagem Docker**:  
   Primeiro, construa a imagem Docker com o comando:

   ```bash
   docker build -t task-golang .
2. **Após a construção da imagem, execute o container:**
   
   ```bash
   docker run -p 8080:8080 task-golang

#### 2. **Executando o Projeto com Docker Compose**

1. **Instalar o Docker Compose**:  
  Caso não tenha o Docker Compose instalado, você pode seguir as instruções na [[Documentação Oficial]([https://img.shields.io/badge/Switch%20to%20English-blue](https://docs.docker.com/engine/install/debian/))] para instalá-lo.

2. **Construir e Iniciar os Serviços**
   Execute o seguinte comando para construir e iniciar os containers definidos no docker-compose.yml:

   ```bash
   docker-compose up --build

3. **Remover os Serviços**

   ```bash
   docker-compose down

**Acessar a Aplicação:**
Após a execução dos comandos acima, a aplicação estará disponível em:

http://localhost:8080

#### 3. **Como Rodar os Testes**
O projeto inclui testes automatizados para garantir que as funcionalidades principais estão funcionando corretamente. Você pode rodar os testes localmente ou dentro do container Docker.

   #### 1. Rodar os Testes Localmente
   Se você tiver o Go configurado na sua máquina, rode os testes com o seguinte comando:
   
      go test ./..

   #### 2. Rodar os Testes com Docker
   Se preferir rodar os testes dentro do container Docker, siga os passos abaixo:
   1 - Execute o container interativamente:
   
      docker-compose run app sh
    
   2 - Dentro do container, rode os testes:

      go test ./...
