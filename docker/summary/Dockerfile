FROM golang:1.20-alpine as base

LABEL maintainer="EduardoAvila <me@eduaravila.com>"

WORKDIR /summary_server

COPY . .

# SHOW FILES 
RUN ls || cat

RUN go mod download

# DEBUG MODE
FROM base as debug

RUN apk add build-base
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN ls || cat 

FROM base as dev

RUN go install github.com/cosmtrek/air@latest

FROM base as prod

ENV ENV=production
# create dir for the binary

WORKDIR /summary_server

RUN  go build cmd/summary_server/main.go

# execute the binary
CMD [ "./main" ]