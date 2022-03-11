FROM golang:1.17.2-alpine3.13
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o bin/togo-write -v .
EXPOSE 8880
CMD [ "/app/bin/togo-write" ]
