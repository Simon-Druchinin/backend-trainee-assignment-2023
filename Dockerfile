FROM golang:1.21.0

RUN go version
ENV GOPATH=/

COPY . .


# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# make wait-for-postgres.sh executable
RUN chmod a+x docker/*.sh

RUN go mod download
RUN go build -o app ./cmd/main.go
