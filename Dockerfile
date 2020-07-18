FROM golang:1.14

# Todo: Heeaaaallthhhhhcheeeeeck
# HEALTHCHECK CMD curl -f http://localhost:8080/health || exit 1;

ADD . /go/src/github.com/psolru/terrastate-http/
WORKDIR /go/src/github.com/psolru/terrastate-http/
RUN go get \
 && go build -o /go/bin/terrastate-http .

ENTRYPOINT terrastate-http
EXPOSE 8080
