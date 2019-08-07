FROM golang:latest

COPY . /app
WORKDIR /app/src

RUN go get github.com/labstack/echo
RUN go get github.com/pilu/fresh
CMD ["fresh"]
