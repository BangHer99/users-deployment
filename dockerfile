FROM golang:1.18

##buat folder APP
RUN mkdir /app

##set direktori utama
WORKDIR /app

##copy seluruh file ke app
ADD . .

##buat executeable
RUN go build -o users-deployment .

##jalankan executeable
CMD ["./be12/users-deployment"]
