FROM golang:latest as builder
 
RUN mkdir -p /app
 
WORKDIR /app
 
COPY ./ /app
 
RUN CGO_ENABLED=0 GOOS=linux go build -v -o server

# Use the official Alpine image for a lean production container.
# https://hub.docker.com/_/alpine
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine:3
RUN apk add --no-cache ca-certificates

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/server /server

#Copy the file hack for counter
COPY --from=builder /app/data /data
 
CMD ["./server"]
