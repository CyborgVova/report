FROM golang:1.21-alpine
ENV GOPATH="/"
ENV DB_HOST="host.docker.internal"
ENV DB_PORT="5432"
ENV DB_USER="user"
ENV DB_PASSWORD="user"
ENV DB_NAME="report"
WORKDIR /report
COPY . .
RUN go build cmd/app/app.go
CMD ["./app"]