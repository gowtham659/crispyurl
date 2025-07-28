FROM golang:latest

WORKDIR /server

#copy all project files
COPY . /server/

#downloand dependencies mentioned in go.mod
RUN go mod download

#build
RUN go build -o app

#excute
CMD ["./app"]

FROM container-registry.oracle.com/database/enterprise:latest


