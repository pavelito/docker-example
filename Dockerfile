FROM golang:latest
 
RUN mkdir -p /app
 
WORKDIR /app
 
COPY . /app
 
RUN go build ./main.go

RUN go get -u github.com/cosmtrek/air
 
CMD ["./main"]