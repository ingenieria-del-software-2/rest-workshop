FROM golang:1.22.4-bookworm
RUN mkdir /app
WORKDIR /app
COPY . .
RUN go build -o rest .
RUN chmod +x /app/rest
EXPOSE 8080
CMD ["/app/rest"]
