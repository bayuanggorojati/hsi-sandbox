# Tugas Pekan 8 - Clean Architecture dengan Fiber dan GORM

## Deskripsi tugas

Tugas ini bertujuan untuk membangun aplikasi *backend* untuk peminjaman buku dengan *Fiber* dan *GORM* menggunakan pendekatan *Clean Architecture*. Secara umum, detil deskripsi tugas sudah dijelaskan di dokumen yang Tugas Pekan 8. Saya akan fokus pada pendekatan, asumsi dan implementasi yang saya lakukan pada tugas ini.

## Pendekatan implementasi

Pada dasarnya, implementasi *clean architecture* dengan *Fiber* pada fungsi-fungsi terkait dengan penyimpanan buku sudah dijelaskan di sesi praktikum. Oleh karena itu, untuk implementasi fungsi-fungsi terkait dengan transaksi peminjaman tinggal mengikuti pola yang sama dengan yang sudah dilakukan pada sesi praktikum. Akan tetapi, ada beberapa hal yang cukup *tricky* dan ini juga yang menjadikan alasan saya menulis README ini, yaitu yang terkait dengan **relasi** antara tipe data peminjaman dengan buku, dan yang terkait dengan tipe data **Date** dan penulisannya di *JSON Body* pada *HTTP POST Request*, yang akan dijelaskan lebih lanjut pada sub-bab berikut ini.

### Relasi tipe data buku dan peminjam

Pertama, saya asumsi bahwa transaksi peminjaman memiliki relasi *one-to-many* dengan tipe data buku. Alasan saya adalah seseorang bisa meminjam lebih dari satu buku pada satu kali transaksi peminjaman. Untuk implementasinya, saya merujuk pada informasi dari dokumentasi *GORM* [di sini] (https://gorm.io/docs/has_many.html). Untuk lebih detilnya, implementasi kedua tipe data (buku dan peminjam) adalah sebagai berikut:

```
type Peminjam struct {
	ID                  int       
	Nama                string    
	TanggalPeminjaman   time.Time 
	TanggalPengembalian time.Time 
	BukuPinjaman        []Buku
}

type Buku struct {
	ID         int
	Judul      string
	Penulis    string
	Tahun      int
	PeminjamID int
}
```

Pada kode di atas, **PeminjamID** merupakan *foreign key* dan seseorang dapat meminjam lebih dari satu buku dalam satu transaksi peminjaman yang diimplementasikan dengan *Array* **Buku**.

### Tipe data *Date* dan penulisannya di **JSON**

Pertama, seperti yang sudah terlihat pada potongan kode di atas, Tanggal Peminjaman dan Pengembalian diimplementasikan dengan tipe data **time.Time**. Namun, ada hal lain yang perlu ditambahkan pada struct *Peminjaman* tersebut seperti yang dijelaskan pada [dokumentasi Fiber] (https://docs.gofiber.io/api/ctx/#bodyparser) > For example, if you want to parse a JSON body with a field called Pass, you would use a struct field of `json:"pass"`. Bentuk final tipe data Peminjam setelah menambahkan *struct field json* adalah sebagai berikut:

```
type Peminjam struct {
	ID                  int       `json:"ID"`
	Nama                string    `json:"Nama"`
	TanggalPeminjaman   time.Time `json:"TanggalPeminjaman"`
	TanggalPengembalian time.Time `json:"TanggalPengembalian"`
	BukuPinjaman        []Buku
}
```

Kedua, format penulisan tipe data *date/time* pada *JSON Body* juga sangat penting, karena jika format penulisan tidak sesuai standar, maka field tersebut tidak dapat di-*parse* dengan benar ke struct *Peminjam*. Standar atau konsensus format penulisan *date/time* pada *JSON* adalah **ISO 8601** (baca [artikel ini] (https://docs.jsonata.org/date-time)). Jadi, contoh penulisannya adalah sebagai berikut: `"2025-09-18T10:00:00.000Z"`.

## Daftar API

### Get All Books

Method: **GET**

Endpoint: `http://<IP>:3000/api/bukus`

Contoh: `http://localhost:3000/api/bukus`

### Get Book by ID

Method: **GET**

Endpoint: `http://<IP>:3000/api/bukus/:id`

Contoh: `http://localhost:3000/api/bukus/1`

### CREATE a Book

Method: **POST**

Endpoint: `http://<IP>:3000/api/bukus`

Contoh: `http://localhost:3000/api/bukus`

Contoh JSON Body: 
```
{
    "ID": 1,
    "Judul": "Intro to Fiber Go",
    "Penulis": "Gilang Golang",
    "Tahun": 2020
}
```

### Get All Rental Transactions

Method: **GET**

Endpoint: `http://<IP>:3000/api/peminjams`

Contoh: `http://localhost:3000/api/peminjams`

### Get Rental Transaction by ID

Method: **GET**

Endpoint: `http://<IP>:3000/api/peminjams/:id`

Contoh: `http://localhost:3000/api/peminjams/1`

### CREATE a Rental Transaction

Method: **POST**

Endpoint: `http://<IP>:3000/api/peminjams`

Contoh: `http://localhost:3000/api/peminjams`

Contoh JSON Body: 
```
{
    "ID": 1,
    "Nama": "Golang Fiber",
    "TanggalPeminjaman": "2025-09-18T10:00:00.000Z",
    "TanggalPengembalian": "2025-09-25T10:00:00.000Z",
    "BukuPinjaman": [
        {
            "ID": 1,
            "Judul": "Intro to Fiber Go",
            "Penulis": "Gilang Golang",
            "Tahun": 2020
        },
        {
            "ID": 5,
            "Judul": "Intro to Golang",
            "Penulis": "Gilang Golang",
            "Tahun": 2019
        }
    ]
}
```