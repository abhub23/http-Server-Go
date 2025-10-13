FROM golang:1.25.2-alpine3.22 AS base

WORKDIR /src

COPY . .
# No .exe because its linux (alpine) and not windows
RUN go build -o server

FROM alpine:3.22.2

WORKDIR /src

COPY --from=base ./src/index.html .

COPY --from=base ./src/server .

EXPOSE 8080

CMD [ "./server" ]