FROM golang:1.23.6-bookworm
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o fetch-receipt-processor .
EXPOSE 8080
CMD ["/build/fetch-receipt-processor"]