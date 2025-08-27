# Dockerfile

# Stage de construção
FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN apk --no-cache add ca-certificates

# Copia os arquivos de go.mod e go.sum para o cache de dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante do código fonte
COPY . .

# Constrói o executável da aplicação
# CGO_ENABLED=0 desabilita cgo para criar um binário estático
# -o estudo-go define o nome do executável
# ./cmd/main.go é o caminho para o arquivo main
RUN CGO_ENABLED=0 go build -o estudo-go ./cmd/main.go

# Stage final (imagem mais leve)
FROM alpine:latest

WORKDIR /app

# Copia o binário construído da stage de construção
COPY --from=builder /app/estudo-go .

# Copia o certificado SSL padrão do alpine para garantir que o driver mysql consiga se conectar
# (pode ser necessário dependendo da sua versão do MySQL ou driver)
RUN apk --no-cache add ca-certificates

# Expõe a porta que a aplicação Go vai escutar
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["/app/estudo-go"]