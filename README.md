# 🛒 Toko Online Backend - Go & SQLite

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)
![Gin](https://img.shields.io/badge/Gin-Web%20Framework-00ADD8?style=for-the-badge&logo=gin)
![GORM](https://img.shields.io/badge/GORM-ORM-00ADD8?style=for-the-badge&logo=go)
![SQLite](https://img.shields.io/badge/SQLite-Database-003B57?style=for-the-badge&logo=sqlite)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)

RESTful API backend untuk sistem e-commerce yang dibangun dengan **Go (Golang)**, **Gin Web Framework**, **GORM**, dan **SQLite**. Project ini menyediakan endpoint lengkap untuk manajemen produk dan sistem transaksi/checkout dengan fitur database transaction untuk menjaga integritas data.

---

## ✨ Fitur Utama

### 📦 CRUD Produk
- ✅ **Create** - Tambah produk baru dengan validasi data
- ✅ **Read** - Tampilkan semua produk atau produk spesifik berdasarkan ID
- ✅ **Update** - Edit data produk yang sudah ada
- ✅ **Delete** - Hapus produk dari database

### 🛍️ Sistem Transaksi/Checkout
- ✅ **Order Management** - Buat pesanan dengan multiple items
- ✅ **Stock Validation** - Cek ketersediaan stok sebelum checkout
- ✅ **Transaction Rollback** - Jika stok tidak mencukupi atau terjadi error, seluruh transaksi dibatalkan
- ✅ **Auto Stock Deduction** - Stok produk otomatis berkurang setelah order berhasil
- ✅ **Order History** - Simpan detail lengkap setiap transaksi

### 🔧 Fitur Teknis
- ✅ **Auto Migration** - Database schema otomatis dibuat saat server dijalankan
- ✅ **CORS Middleware** - Support request dari frontend berbeda origin/port
- ✅ **Clean Architecture** - Struktur folder modular dan terorganisir
- ✅ **Error Handling** - Response error yang jelas dan informatif

---

## 🚀 Teknologi yang Digunakan

| Teknologi | Versi | Deskripsi |
| :--- | :--- | :--- |
| **Go (Golang)** | 1.21+ | Bahasa pemrograman utama |
| **Gin Web Framework** | v1.10.0 | HTTP web framework untuk routing dan middleware |
| **GORM** | v1.25.12 | ORM (Object-Relational Mapping) untuk database operations |
| **SQLite** | v1.5.6 | Database engine (file-based, tanpa instalasi server) |
| **gin-contrib/cors** | v1.7.2 | Middleware CORS untuk cross-origin requests |

---

## 📁 Struktur Folder Project

```
toko-online-backend/
├── config/
│   └── db.go                 # Konfigurasi database & auto-migration
├── controllers/
│   ├── product.go            # Logika CRUD produk
│   └── order.go              # Logika transaksi/checkout
├── models/
│   ├── product.go            # Schema/struct Product
│   └── order.go              # Schema/struct Order & OrderItem
├── routes/
│   ├── product_routes.go     # Endpoint produk
│   └── order_routes.go       # Endpoint order
├── go.mod                    # Go module dependencies
├── go.sum                    # Checksum dependencies
├── main.go                   # Entry point aplikasi
├── toko_online.db            # SQLite database (auto-generated)
└── README.md                 # Dokumentasi project
```

---

## 📦 Instalasi & Menjalankan Project

### Prasyarat
- **Go** versi 1.21 atau lebih tinggi terinstall di sistem Anda
  - Download: [https://golang.org/dl/](https://golang.org/dl/)

### Langkah-langkah Instalasi

1. **Clone atau navigasi ke project directory**
   ```bash
   cd d:\portofolio-project\toko-online-backend
   ```

2. **Download dependencies**
   ```bash
   go mod tidy
   ```

3. **Jalankan server**
   ```bash
   go run main.go
   ```

4. **Server akan berjalan di port 8080**
   ```
   Starting server...
   Database connected and migrated successfully
   Server running on port 8080
   ```

5. **Database file `toko_online.db` akan otomatis terbentuk**
   - File ini berisi tabel `products`, `orders`, dan `order_items`

---

## 📡 Dokumentasi API Endpoint

### Base URL
```
http://localhost:8080
```

---

### 📦 Produk Endpoints

#### 1. Create Product - Tambah Produk Baru

**Endpoint:** `POST /products`

**Request Body:**
```json
{
  "name": "Laptop Gaming ASUS ROG",
  "category": "Elektronik",
  "description": "Laptop gaming dengan spesifikasi tinggi untuk gaming dan editing",
  "price": 25000000,
  "stock": 10,
  "image": "https://example.com/images/laptop-gaming.jpg"
}
```

**Response (201 Created):**
```json
{
  "message": "Product created successfully",
  "data": {
    "id": 1,
    "name": "Laptop Gaming ASUS ROG",
    "category": "Elektronik",
    "description": "Laptop gaming dengan spesifikasi tinggi untuk gaming dan editing",
    "price": 25000000,
    "stock": 10,
    "image": "https://example.com/images/laptop-gaming.jpg",
    "created_at": "2026-07-16T10:18:00Z",
    "updated_at": "2026-07-16T10:18:00Z"
  }
}
```

---

#### 2. Get All Products - Tampilkan Semua Produk

**Endpoint:** `GET /products`

**Response (200 OK):**
```json
{
  "message": "Products retrieved successfully",
  "data": [
    {
      "id": 1,
      "name": "Laptop Gaming ASUS ROG",
      "category": "Elektronik",
      "description": "Laptop gaming dengan spesifikasi tinggi",
      "price": 25000000,
      "stock": 10,
      "image": "https://example.com/images/laptop.jpg",
      "created_at": "2026-07-16T10:18:00Z",
      "updated_at": "2026-07-16T10:18:00Z"
    },
    {
      "id": 2,
      "name": "Sepatu Nike Air Jordan",
      "category": "Fashion",
      "description": "Sepatu sneakers original dengan desain klasik",
      "price": 3500000,
      "stock": 25,
      "image": "https://example.com/images/sepatu-jordan.jpg",
      "created_at": "2026-07-16T10:20:00Z",
      "updated_at": "2026-07-16T10:20:00Z"
    }
  ]
}
```

---

#### 3. Get Product by ID - Tampilkan Produk Spesifik

**Endpoint:** `GET /products/:id`

**Example:** `GET /products/1`

**Response (200 OK):**
```json
{
  "message": "Product retrieved successfully",
  "data": {
    "id": 1,
    "name": "Laptop Gaming ASUS ROG",
    "category": "Elektronik",
    "description": "Laptop gaming dengan spesifikasi tinggi",
    "price": 25000000,
    "stock": 10,
    "image": "https://example.com/images/laptop.jpg",
    "created_at": "2026-07-16T10:18:00Z",
    "updated_at": "2026-07-16T10:18:00Z"
  }
}
```

**Response (404 Not Found):**
```json
{
  "error": "Product not found"
}
```

---

#### 4. Update Product - Edit Produk

**Endpoint:** `PUT /products/:id`

**Example:** `PUT /products/1`

**Request Body:**
```json
{
  "name": "Laptop Gaming ASUS ROG Updated",
  "category": "Elektronik",
  "description": "Laptop gaming dengan spesifikasi tinggi - Updated description",
  "price": 28000000,
  "stock": 8,
  "image": "https://example.com/images/laptop-updated.jpg"
}
```

**Response (200 OK):**
```json
{
  "message": "Product updated successfully",
  "data": {
    "id": 1,
    "name": "Laptop Gaming ASUS ROG Updated",
    "category": "Elektronik",
    "description": "Laptop gaming dengan spesifikasi tinggi - Updated description",
    "price": 28000000,
    "stock": 8,
    "image": "https://example.com/images/laptop-updated.jpg",
    "created_at": "2026-07-16T10:18:00Z",
    "updated_at": "2026-07-16T10:30:00Z"
  }
}
```

---

#### 5. Delete Product - Hapus Produk

**Endpoint:** `DELETE /products/:id`

**Example:** `DELETE /products/1`

**Response (200 OK):**
```json
{
  "message": "Product deleted successfully"
}
```

**Response (404 Not Found):**
```json
{
  "error": "Product not found"
}
```

---

### 🛍️ Order Endpoints

#### 6. Create Order - Checkout Pesanan

**Endpoint:** `POST /orders`

**Request Body:**
```json
{
  "customer_name": "John Doe",
  "items": [
    {
      "product_id": 1,
      "quantity": 2
    },
    {
      "product_id": 2,
      "quantity": 1
    }
  ]
}
```

**Response (201 Created):**
```json
{
  "message": "Order created successfully",
  "data": {
    "id": 1,
    "customer_name": "John Doe",
    "total_price": 53500000,
    "order_items": [
      {
        "id": 1,
        "order_id": 1,
        "product_id": 1,
        "product": {
          "id": 1,
          "name": "Laptop Gaming ASUS ROG",
          "category": "Elektronik",
          "description": "Laptop gaming dengan spesifikasi tinggi",
          "price": 25000000,
          "stock": 8,
          "image": "https://example.com/images/laptop.jpg",
          "created_at": "2026-07-16T10:18:00Z",
          "updated_at": "2026-07-16T10:30:00Z"
        },
        "quantity": 2,
        "price": 25000000,
        "created_at": "2026-07-16T11:05:00Z",
        "updated_at": "2026-07-16T11:05:00Z"
      },
      {
        "id": 2,
        "order_id": 1,
        "product_id": 2,
        "product": {
          "id": 2,
          "name": "Sepatu Nike Air Jordan",
          "category": "Fashion",
          "description": "Sepatu sneakers original dengan desain klasik",
          "price": 3500000,
          "stock": 24,
          "image": "https://example.com/images/sepatu-jordan.jpg",
          "created_at": "2026-07-16T10:20:00Z",
          "updated_at": "2026-07-16T10:20:00Z"
        },
        "quantity": 1,
        "price": 3500000,
        "created_at": "2026-07-16T11:05:00Z",
        "updated_at": "2026-07-16T11:05:00Z"
      }
    ],
    "created_at": "2026-07-16T11:05:00Z",
    "updated_at": "2026-07-16T11:05:00Z"
  }
}
```

**Response (400 Bad Request - Stok Tidak Mencukupi):**
```json
{
  "error": "Stok produk Laptop Gaming ASUS ROG tidak mencukupi"
}
```

**Response (404 Not Found - Produk Tidak Ditemukan):**
```json
{
  "error": "Product not found"
}
```

---

### 🧪 Testing Endpoint

**Endpoint:** `GET /ping`

**Response (200 OK):**
```json
{
  "message": "pong"
}
```

---

## 🔒 Keamanan & Best Practices

- ✅ **Transaction Management** - Menggunakan GORM transaction untuk menjaga integritas data
- ✅ **Input Validation** - Validasi data input menggunakan Gin binding
- ✅ **Error Handling** - Response error yang jelas dan informatif
- ✅ **CORS Configuration** - Support cross-origin requests dengan konfigurasi yang aman
- ✅ **SQL Injection Prevention** - GORM ORM secara otomatis mencegah SQL injection

---

## 📝 Catatan Penting

1. **Database File** - File `toko_online.db` akan otomatis dibuat saat pertama kali server dijalankan
2. **Stock Management** - Stok produk akan otomatis berkurang setelah order berhasil dibuat
3. **Transaction Rollback** - Jika ada error selama proses checkout, seluruh transaksi akan dibatalkan
4. **Port Configuration** - Server berjalan di port 8080 secara default
5. **CORS** - API dapat diakses dari frontend berbeda origin/port

---

## 🤝 Kontribusi

Project ini dikembangkan sebagai showcase portofolio untuk mendemonstrasikan kemampuan dalam membangun RESTful API dengan Go.

---

## 👤 Developer Identity

| Detail | Informasi |
| :--- | :--- |
| **Nama** | Daffa |
| **Peran** | Backend / Full-Stack Developer |
| **Fokus Teknologi** | Go (Golang), React.js, JavaScript, SQLite |
| **Kontak & Media Sosial** | [GitHub](https://github.com/daffa) | [LinkedIn](https://linkedin.com/in/daffa) |

---

## 📄 License

Project ini dilisensikan under MIT License.

---

## 🎯 Future Enhancements

- [ ] Authentication & Authorization (JWT)
- [ ] Pagination untuk list produk
- [ ] Search & Filter produk
- [ ] Order History endpoint
- [ ] Payment Gateway Integration
- [ ] Admin Dashboard API
- [ ] Rate Limiting
- [ ] Docker Containerization
- [ ] Unit Testing
- [ ] API Documentation (Swagger)

---

**Dibuat dengan ❤️ menggunakan Go & Gin**
