FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN rm .env && mv .env.compose .env
RUN go build -v -o /usr/local/bin/onecv ./main.go

CMD ["onecv"]
