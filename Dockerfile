FROM golang:1.20-alpine3.16 as base64
RUN apk update
WORKDIR /src/desafioKlever
COPY go.mod go.sun ./
RUN go build -o desafioKlever

FROM alpine:3.16 as binary
COPY --from=base /src/desafioKlever/desafioKlever .
COPY --from=base /src/desafioKlever/web ./web
EXPOSE 8000
ENTRYPOINT [ "/desafioKlever" ]