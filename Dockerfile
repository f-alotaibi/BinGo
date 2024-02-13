FROM golang:alpine AS build

WORKDIR /build
ADD . /build
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN templ generate .
RUN go build .

FROM alpine:latest

# Export necessary port
EXPOSE 80
# Add  application
WORKDIR /dist
COPY --from=build /build/bingo /dist/main

ENTRYPOINT ["/dist/main"]