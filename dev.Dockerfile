FROM golang:1.18 as run
WORKDIR /app
RUN go install github.com/codegangsta/gin@latest
ENV BIN_APP_PORT=8080
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download
ENTRYPOINT ["gin","run","main.go"]