FROM golang:latest

WORKDIR /app
COPY . .

RUN apt-get update && apt-get install -y build-essential
RUN go mod download

CMD ["go run ./app/main.go"]
