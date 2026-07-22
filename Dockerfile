FROM golang:1.25

WORKDIR /application

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY app/go.mod app/go.sum ./
COPY app .

RUN go mod download
RUN CGO_ENABLED=0 go build -o ./bin/app ./cmd/birth_minder/main.go

CMD ["/bin/app"]
