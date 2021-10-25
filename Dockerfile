FROM golang:1.16.9-alpine3.13
RUN mkdir /app
ADD . /app/
WORKDIR /app/cmd
RUN mkdir tmp
RUN go build
CMD ["./cmd"]