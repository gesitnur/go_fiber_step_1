Di dalam Go Fiber (dan banyak *framework* web lainnya), **`c`** adalah nama variabel kependekan dari **Context** (biasa ditulis `ctx` atau `c`).

Tipe datanya dalam Fiber v3 adalah `fiber.Ctx`.

Secara sederhana: **`c` atau *Context* adalah sebuah "kantong ajaib" yang membawa semua informasi tentang siklus HTTP yang sedang berlangsung.**

Bayangkan `c` seperti seorang kurir atau pelayan restoran. Dia yang membawa kertas pesanan dari pelanggan (Browser/Client), dan dia juga yang mengantarkan makanan dari dapur (Server) kembali ke meja pelanggan.

---

## Apa Saja Isi di Dalam `c` (Context) Ini?

Karena `c` merangkum seluruh siklus HTTP, Anda bisa menggunakan `c` untuk melakukan dua hal utama: **Membaca data yang masuk (Request)** dan **Mengirim data keluar (Response)**.

### 1. Untuk Membaca Request (Data dari Client ke Server)

Jika ada data yang dikirim oleh *user*, Anda mengambilnya lewat `c`. Contoh:

* **Mengambil Parameter URL:** `c.Params("id")` (Misal untuk membaca ID di URL `/user/123`).
* **Mengambil Query String:** `c.Query("search")` (Misal untuk membaca kata kunci di URL `/search?search=baju`).
* **Mengambil Body/JSON:** `c.BodyParser(&user)` (Untuk membaca data form atau JSON yang dikirim saat *user* berbelanja atau *login*).
* **Mengambil Header:** `c.Get("Authorization")` (Untuk mengecek token keamanan).

### 2. Untuk Mengirim Response (Data dari Server ke Client)

Seperti kode Anda, `c` juga digunakan untuk menentukan apa yang mau dikembalikan ke *user*. Contoh:

* **Mengirim Teks Biasa:** `c.SendString("Hello")` (Merespons dengan teks polos).
* **Mengirim JSON:** `c.JSON(fiber.Map{"status": "sukses"})` (Paling sering dipakai untuk bikin API).
* **Mengirim File/Gambar:** `c.SendFile("./foto.jpg")`.
* **Mengubah Status HTTP:** `c.Status(404)` (Untuk memberi tahu browser kalau halaman tidak ketemu).

---

## Kenapa di Fiber v3 Tidak Pakai Tanda Bintang (`*`) Lagi?

Ini adalah salah satu perubahan terbesar dari Fiber v2 ke Fiber v3.

* **Di Fiber v2:** Ditulis `func(c *fiber.Ctx)`. Tanda `*` artinya *Context* dikirim sebagai *pointer*.
* **Di Fiber v3:** Ditulis `func(c fiber.Ctx)`. Cukup *interface* biasa.

**Alasannya (Detail Teknis):**
Di balik layar, tim pengembang Fiber mengubah `fiber.Ctx` di v3 menjadi sebuah **Interface**, bukan lagi *struct* biasa. Di dalam bahasa Go, *interface* sudah secara otomatis mereferensikan datanya tanpa perlu kita tulis bintang (`*`) manual lagi.

Tujuan perubahan ini adalah untuk **kebersihan kode (clean code)** dan menghindari *error* akibat salah mengelola memori *pointer* saat aplikasi Anda menghandle ribuan *request* secara bersamaan (fitur *concurrency* Go).

## Contoh Penggunaan `c` yang Lebih Kompleks:

Supaya makin kebayang, ini contoh jika `c` digunakan untuk membuat API yang menerima data:

```go
app.Post("/login", func(c fiber.Ctx) error {
    // 1. Mengambil data "Username" yang diketik user di form login
    username := c.FormValue("username") 
    
    // 2. Memberikan respons balik berupa JSON dan status 200 (OK)
    return c.Status(200).JSON(fiber.Map{
        "pesan": "Selamat datang, " + username,
    })
})

```

Jadi, setiap kali Anda membuat fungsi di dalam *endpoint* Fiber, variabel `c` ini **wajib** dimasukkan sebagai parameter pertama karena dia adalah jembatan komunikasi antara program Go Anda dengan dunia luar.