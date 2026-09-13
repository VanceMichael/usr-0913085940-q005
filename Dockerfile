FROM golang:1.22 AS build
WORKDIR /src
COPY go.mod ./
COPY cmd ./cmd
RUN CGO_ENABLED=0 go build -o /out/server ./cmd/server
FROM gcr.io/distroless/static-debian12
COPY --from=build /out/server /server
EXPOSE 8080
ENTRYPOINT ["/server"]
