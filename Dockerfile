FROM golang:1.19.13-alpine3.18 as build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN  GOOS=linux GOARCH=amd64 go build -o bin/libraryservice cmd/*go

############################################### FINAL ##############################################
FROM alpine
ARG FILENAME="library.yaml"
ARG FILEDIR="./config"

WORKDIR /app
COPY --from=build /app/bin/libraryservice /app/bin/libraryservice
COPY $FILEDIR/${FILENAME} /app/bin/config/${FILENAME}

#start command
ENTRYPOINT ["/app/bin/libraryservice"]




