# library-go

Cara untuk menjalankan aplikasi

Langkah pertama yang harus dilakukan adalah clone repository ini. 

```
git clone https://github.com/faizalardhi16/library-go.git
```


1. Download dan Install GO pada url berikut : https://go.dev/doc/install
2. Sesuaikan dengan OS yang digunakan untuk menjalankan GO
3. Download dan Install XAMPP untuk menjalankan MySQL database di https://www.apachefriends.org/download.html
4. Sesuaikan dengan OS yang digunakan untuk running XAMPP
5. Import database joybox.sql pada repository ini
6. Run xampp

Setelah menjalankan semua langkah diatas, buka folder yang sudah di clone tersebut dan jalankan command dibawah ini :

```
go run main.go
```

Terdapat beberapa endpoint yang bisa digunakan

# USER

Buat sebuah user baru, user bisa dibuat melalui Postman. Dimana, digunakan url http://localhost:5041/api/v1/register

Dengan body seperti dibawah ini 

```
{
    "email":"natalie@gmail.com",
    "first_name":"Natalie",
    "last_name": "Sasha",
    "password": "123123"
}
```

Setelah register, dapat dicoba untuk login ke akun yang sudah didaftarkan tersebut di http://localhost:5041/api/v1/session

```
{
    "email":"natalie@gmail.com",
    "password": "123123"
}
```

Token yang didapatkan ketika register maupun login dapat digunakan untuk melakukan transaksi


# TRANSACTION

Untuk melakukan test pada API transaction, kita membutuhkan User yang terdaftar dan buku yang sudah diupload. Oleh karena itu, dengan import DB yang sudah disediakan akan mempercepat prosesnya.

Untuk melakukan booking pada buku, dibuat terlebih dahulu transaksinya, dimana 1 transaksi berelasi dengan 1 id buku.

Transaksi dapat dibuat dengan menggunakan url http://localhost:5041/api/v1/transaction

```
{
    "book_id": "3f83675d-2111-4d5e-ab06-0c07a1cadc4b"
}
```

Jangan lupa untuk menambahkan authorization Bearer Token di tab Headernya

Token didapatkan ketika login / mendaftar!


