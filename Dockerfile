FROM golang:alpine AS build

WORKDIR /build
ADD . /build

RUN go build ./cmd/pastebin-go/

FROM alpine:latest

# Export necessary port
EXPOSE 80
# Add  application
WORKDIR /dist
COPY --from=build /build/cmd/pastebin-go/views /dist/views
COPY --from=build /build/pastebin-go /dist/main

ENTRYPOINT ["/dist/main"]