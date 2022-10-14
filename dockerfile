FROM golang:1.18

##buat folder APP
RUN mkdir /app

##set direktori utama
WORKDIR /app

##copy seluruh file ke app
COPY ./ /app

##buat executeable
RUN go build -o main .

##jalankan executeable
CMD ["./main"]
