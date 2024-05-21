# Etapa de construção
FROM golang:1.19-alpine AS builder

# Configurar diretório de trabalho
WORKDIR /app


# Copiar o restante dos arquivos do projeto
COPY . .

# Compilar a aplicação
RUN go build -o loadtester .

# Etapa final
FROM alpine:3.15

# Configurar diretório de trabalho
WORKDIR /app

# Copiar o binário compilado da etapa anterior
COPY --from=builder /app/loadtester .

# Comando de entrada
ENTRYPOINT ["./loadtester"]
