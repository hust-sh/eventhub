FROM golang:1.9.2

WORKDIR /go/src/app
COPY . .

RUN go get github.com/pilu/fresh  # live reload
RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."
RUN chmod +x init.sh

CMD ["./init.sh"] # ["app"]
