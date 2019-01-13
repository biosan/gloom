FROM golang:1.10.7-alpine3.8 AS builder

ENV REPO='github.com/biosan/gloom'

WORKDIR /go/src/${REPO}

# Install git
RUN apk --no-cache add git
# Install HTTP mux (gorilla/mux)
RUN go get "github.com/gorilla/mux"

# Copy source code
COPY . .
# Build gloomAPI
RUN go build -o /api ${REPO}/api


FROM alpine:3.8

WORKDIR /api
# Copy compiled binary
COPY --from=builder /api .

ENV GLOOMAPI_PORT=80
EXPOSE ${GLOOMAPI_PORT}

ENTRYPOINT ["./api"]