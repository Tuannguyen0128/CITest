# syntax=docker/dockerfile:1
FROM golang:1.16-alpine
WORKDIR app
COPY go.mod ./
COPY go.sum ./
COPY *.go ./
#Download nessesary module
RUN go mod tidy
#Build app to get .exe
RUN go build -o /docker-gs-ping
#run .exe
CMD ["/docker-gs-ping"]
EXPOSE 8080
