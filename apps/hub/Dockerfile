FROM golang:alpine as builder

ENV GOARCH=amd64 
ENV GOOS=linux 
ENV CGO_ENABLED=0 

WORKDIR /src
COPY . .

RUN go build -mod vendor -o server .
RUN chmod +x server

FROM scratch
WORKDIR /usr/local/bin/

COPY --from=builder /src/server /usr/local/bin/server
COPY --from=builder /src/conf /usr/local/bin/conf

EXPOSE 8080/tcp
ENTRYPOINT ["/usr/local/bin/server"]