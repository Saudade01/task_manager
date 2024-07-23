# Golang image'ını temel alarak başlıyoruz
FROM golang:1.20-alpine

# Çalışma dizinini ayarla
WORKDIR /app

# Modül dosyalarını kopyala
COPY go.mod ./
COPY go.sum ./

# Bağımlılıkları yükle
RUN go mod download

# Proje dosyalarını kopyala
COPY . .

# Uygulamayı derle
RUN go build -o /task_manager

# MySQL bağlantı bilgilerini ortam değişkeni olarak ayarla
ENV MYSQL_USER=myuser
ENV MYSQL_PASSWORD=Orhun8722
ENV MYSQL_DATABASE=task_manager
ENV MYSQL_HOST=127.0.0.1
ENV MYSQL_PORT=3306

# Uygulamayı başlat
CMD ["/task_manager"]
