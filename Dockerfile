FROM golang:1.25

RUN apt-get update && apt-get install gettext-base

WORKDIR /application

COPY app/go.mod app/go.sum ./
COPY app .

COPY config.dist.yml .

RUN go mod download
RUN CGO_ENABLED=0 go build -o /bin/app ./cmd/birth_minder/main.go

CMD ["/bin/bash", "-c", "cat config.yml && /bin/app"]