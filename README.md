# 💵 Cotação Dólar - API e Cliente em Go

Este projeto consiste em uma aplicação dividida em duas partes: **servidor** e **cliente**. A aplicação permite buscar a cotação do dólar (USD para BRL) de uma API pública, armazenar no banco de dados (SQLite) e exibir a cotação no terminal.

---

## 📦 Estrutura do Projeto

```
.
├── client              # Aplicação cliente que consome a API local
│   └── main.go
├── server              # Servidor HTTP em Go
│   ├── domain
│   │   ├── Entity              # Definições de entidades
│   │   ├── errors              # Tratamento centralizado de erros
│   │   └── response            # Estruturas de resposta da API
│   ├── repositorys             # Interação com banco de dados SQLite
│   ├── rest-clients            # Cliente HTTP que consome a AwesomeAPI
│   └── main.go                 # Inicialização do servidor e rotas
```

---

## 🚀 Como funciona

### 🔧 Servidor (`/server`)

- Inicia um servidor HTTP escutando na porta `8080`
- Rota `/cotacao`:
  - Faz uma requisição para a [AwesomeAPI](https://docs.awesomeapi.com.br/api-de-moedas) para obter a cotação do dólar
  - Persiste o resultado em um banco SQLite (`cotacoes.db`)
  - Retorna apenas o campo `bid` (valor atual do dólar)

### 💻 Cliente (`/client`)

- Faz uma requisição para `http://localhost:8080/cotacao`
- Exibe a cotação do dólar no terminal

---

## 🛠️ Tecnologias Utilizadas

- Go (Golang)
- SQLite (banco de dados local)
- HTTP Client/Server
- AwesomeAPI (cotação USD/BRL)
- Tratamento de erros customizados

---

## ▶️ Como Executar o Projeto

### 1. Clonar o repositório
```bash
git clone https://github.com/gerps2/cotacao-real-dolar.git
cd cotacao-real-dolar
```

### 2. Rodar o servidor
```bash
cd server
go run main.go
```

### 3. Em outro terminal, rodar o cliente
```bash
cd client
go run main.go
```

---

## 📌 Observações

- O servidor possui timeout de 200ms para chamada externa e 10ms para persistência no banco.
- O cliente possui timeout de 100ms para resposta do servidor.
- Ideal para estudos sobre chamadas HTTP, tratamento de erros e persistência com SQLite.

---

## 📝 Licença

Este projeto está sob a licença MIT. Sinta-se livre para usá-lo e modificá-lo.