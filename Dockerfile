FROM golang:alpine AS build

WORKDIR /build
ADD . /build

RUN go build ./cmd/bingo/

FROM alpine:latest

# Export necessary port
EXPOSE 80
# Add  application
WORKDIR /dist
COPY --from=build /build/cmd/bingo/views /dist/views
COPY --from=build /build/bingo /dist/main

ENTRYPOINT ["/dist/main"]