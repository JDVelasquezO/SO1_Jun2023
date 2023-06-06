FROM golang

WORKDIR /back

COPY . .

RUN go mod download

EXPOSE 8000

CMD ["go", "run", "Main.go"]
