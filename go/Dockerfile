FROM golang:1.20.6-alpine3.18


# Set the working directory inside the container
WORKDIR /usr/src/app

# Copy the rest of the application source code
COPY . .

RUN apk add build-base
RUN go clean -modcache
RUN go mod tidy
RUN go build -o main main.go
RUN chmod +x ./main
CMD [ "./main"]