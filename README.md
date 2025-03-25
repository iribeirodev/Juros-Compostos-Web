# Juros Compostos Web  

Este é um projeto em **Go** utilizando o **Gin** para simular cálculos de **juros compostos** e **parcelas com amortização**.

## 📌 Tecnologias  

- [Go](https://go.dev/)  
- [Gin Gonic](https://github.com/gin-gonic/gin)  

## ⚙️ Instalação  

1. Clone o repositório:  
   ```sh
   git clone https://github.com/iribeirodev/juros-compostos-web.git
   cd juros-compostos-web


2. Instale as dependências
	```sh
	go mod tidy
	```

3. Inicie o servidor

   ```sh
   go run cmd/main.go
   ```

O servidor rodará na porta 8080.

Como Testar
As requisições podem ser testadas utilizando o arquivo teste.http. Caso esteja usando o VSCode, basta abrir o arquivo e clicar em "Send Request".

Simulação de Juros Compostos
Endpoint:
````
POST http://localhost:8080/juros-compostos
Content-Type: application/json

{
  "valor_inicial": 1000,
  "taxa": 5,
  "periodos": 12
}
````

Simulação de Parcelas (Sistema de Amortização)
Endpoint:

```
POST http://localhost:8080/simulacao-parcelas
Content-Type: application/json

{
  "valor_produto": 10000,
  "parcelas": 12,
  "taxa": 8
}
```

Verificar Saúde da API
Endpoint:

```
GET http://localhost:8080/ping
```

📂 Estrutura do Projeto

juros-compostos-web/
│── cmd/
│   ├── main.go        # Arquivo principal para iniciar o servidor
│── internal/
│   ├── handlers/      # Handlers das rotas
│   ├── models/        # Modelos de entrada e saída de dados
│   ├── services/      # Lógica de negócio
│── routes/
│   ├── routes.go      # Definição das rotas
│── teste.http         # Arquivo de teste das requisições
│── go.mod             # Dependências do projeto
│── README.md          # Documentação do projeto

📜 Licença
Este projeto é open-source e pode ser usado livremente.

