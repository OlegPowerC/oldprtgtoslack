FROM golang:alpine as builder
 
RUN apk update && apk add --no-cache git
RUN apk add --no-cache openssh
RUN apk add build-base
 
CMD mkdir /opt/messnotify
WORKDIR /opt/messnotify
COPY messnotify.go messnotify.go
COPY settings.json settings.json

RUN go build messnotify.go

FROM alpine  
WORKDIR /root/
ENV TZ=Europe/Moscow

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /opt/messnotify/. .
RUN apk --update add tzdata && cp /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
# Expose port 8788 to the outside
EXPOSE 8788

# Command to run the executable
CMD ["./messnotify"]
