# Simulador de saque em caixas eletrônicos

## Requisitos

- Go (versão go1.18.1 ou superior)

## Instalação
1. Crie as pastas do Go caso você ainda não tenha:
   
   ```
   mkdir go && cd go && mkdir pkg && mkdir bin && mkdir src
   
   ```
   
2. Clone o repositório do projeto dentro da pasta src:

   ```
   cd src && git clone https://github.com/rbs-batista/atm.git
   
   ```

3. Acesse o diretório do projeto:

   ```
   cd atm
   ```

3. Instale as dependências do projeto:

   ```
   go get
   ```
## Como Rodar

1. Para executar o projeto, utilize o seguinte comando:

   ```
   go run main.go
   ```

2. O servidor será iniciado e estará acessível no seguinte endereço:

   ```
   http://localhost:8000
   ```

## Como Testar

1. Para executar os testes do projeto, utilize o seguinte comando:

   ```
   cd Tests && go test
   ```
