FROM golang:1.24

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Установка file
RUN apt-get update && apt-get install -y file && rm -rf /var/lib/apt/lists/*

# Билдим под amd64
RUN GOOS=linux GOARCH=amd64 go build -o /app/migrate ./cmd/migrate/main.go

# Диагностика архитектуры
RUN file /app/migrate

# Права на исполняемый файл
RUN chmod +x /app/migrate

# Собираем основное приложение (опционально)
RUN GOOS=linux GOARCH=amd64 go build -o /app/main ./cmd/main.go

EXPOSE 8080

CMD ["/app/main"]
