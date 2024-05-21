# LoadTester

LoadTester é uma aplicação de linha de comando (CLI) escrita em Go para realizar testes de carga em um serviço web. Ela permite que você especifique a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas. Ao final do teste, um relatório com informações detalhadas é gerado.

## Funcionalidades

- Realização de testes de carga em uma URL especificada.
- Suporte para definir o número total de requests e o nível de concorrência.
- Geração de relatório com tempo total de execução, quantidade de requests realizados, quantidade de requests com status HTTP 200 e distribuição de outros códigos de status HTTP.

## Pré-requisitos

- Go 1.19 ou superior.
- Docker (opcional, para execução via contêiner).

## Instalação

### Clonando o Repositório

```sh
git clone 
```
go mod tidy
go build -o stresstest

docker build -t stresstest .

./Sistema-de-Stress-test --url=http://google.com --requests=1000 --concurrency=10

docker run --rm stresstest --url=http://google.com --requests=1000 --concurrency=10

Exemplo de saída:

Tempo total: 1m23.456s
Total de requests: 1000
Requests com status 200: 950
Distribuição dos códigos de status HTTP:
200: 950
404: 30
500: 20