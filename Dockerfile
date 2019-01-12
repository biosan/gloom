FROM golang:1.10.7-alpine3.8 AS builder

ENV SRC_DIR=/go/src/github.com/biosan/gloom

# Install git
RUN apk add git
# Install HTTP mux (gorilla/mux)
RUN go get "github.com/gorilla/mux"
# Copy source code
ADD . $SRC_DIR
# Build gloomAPI
RUN cd $SRC_DIR/api; go build -o api; cp api /.


FROM alpine:3.8

WORKDIR /api
# Copy compiled binary
COPY --from=builder /api ./api

EXPOSE 8888

ENTRYPOINT ["./api"]