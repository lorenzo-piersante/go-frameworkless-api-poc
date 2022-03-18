FROM golang:1.18

RUN apt-get update
RUN apt-get -y upgrade
RUN apt-get install sqlite3

RUN mkdir -p /home/api
WORKDIR /home/api
COPY . /home/api

RUN go install

CMD ["go", "run", "main.go"]