FROM golang:1.22.6

# set environment module
ENV GO111MODULE=on

# set workspace for this app
WORKDIR /app

# copy file go.mod and go.sum to this workdir or /app
COPY go.mod .
COPY go.sum .

# download dependency / library
RUN go mod download

# copy project file to workdir or /app
COPY . .

# build the apps
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .

# set env for httport to 8080
ENV HTTP_PORT=3000

# expose the port
EXPOSE 3000

# run the application
ENTRYPOINT ["/app/dating-api"]