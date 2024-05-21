# Etapa de construção
FROM golang:1.19-alpine AS builder

# Configurar diretório de trabalho
WORKDIR /app

# Copiar o arquivo main.go para o contêiner
COPY main.go .

# Compilar a aplicação
RUN go build -o loadtester main.go

# Etapa final
FROM alpine:3.15

# Configurar diretório de trabalho
WORKDIR /app

# Copiar o binário compilado da etapa de construção
COPY --from=builder /app/loadtester .

# Comando de entrada padrão
ENTRYPOINT ["./loadtester"]
