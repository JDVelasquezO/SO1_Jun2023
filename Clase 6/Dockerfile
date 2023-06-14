FROM golang

WORKDIR /back

COPY . .

RUN go mod download

ENV DB_HOST localhost

EXPOSE 8000

CMD ["go", "run", "Main.go"]
