FROM golang:1.19

WORKDIR /go/delivery
COPY server.go go.mod ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/server

FROM ghcr.io/cloudquery/cloudquery:3.4.0

EXPOSE 8080
COPY --from=0 /bin/server /bin/server
ENTRYPOINT []
CMD ["/bin/server"]