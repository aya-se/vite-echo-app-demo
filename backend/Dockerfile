FROM golang:1.20-alpine
ENV TZ /usr/share/zoneinfo/Asia/Tokyo
WORKDIR /backend
RUN go install github.com/cosmtrek/air@latest
COPY . .
RUN go mod download
CMD ["air"]
EXPOSE 8080