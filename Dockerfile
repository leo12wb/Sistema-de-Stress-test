# Imagem base
FROM golang:latest

# Define o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copia o código fonte para o diretório de trabalho
COPY . .

# Compila o aplicativo
RUN go build -o load-test .

# Comando padrão para executar o aplicativo
CMD ["./load-test"]
