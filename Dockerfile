FROM golang:1.20-bullseye AS debug
WORKDIR /app
COPY . .
RUN go install github.com/cosmtrek/air@v1.42.0
CMD air

FROM golang:1.20-bullseye as builder
WORKDIR /opt/app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN cd cmd && go build -trimpath -ldflags="-w -s" -o "gocraw"

FROM gcr.io/distroless/base-debian11 as dev
COPY --from=builder opt/app/cmd/gocraw /gocraw
CMD ["/gocraw"]
