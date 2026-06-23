# TobaLapor - Product Requirements Document (PRD)

## Project Overview

TobaLapor adalah platform pengaduan masyarakat berbasis mobile yang ditujukan untuk Kabupaten Toba. Sistem ini memungkinkan masyarakat melaporkan permasalahan publik dan memantau progres penanganannya secara transparan.

Fokus utama aplikasi bukan hanya menerima laporan, tetapi menyediakan transparansi terhadap proses penanganan laporan hingga selesai.

Platform terdiri dari:

* Mobile App (Flutter) untuk masyarakat
* Web Admin untuk pengelolaan laporan

---

# Problem Statement

Masyarakat sering tidak mengetahui perkembangan laporan yang telah mereka kirimkan kepada pemerintah daerah.

Permasalahan yang ingin diselesaikan:

* Tidak mengetahui apakah laporan diterima.
* Tidak mengetahui siapa yang menangani laporan.
* Tidak mengetahui progres penanganan laporan.
* Tidak adanya transparansi setelah laporan dikirim.

---

# Goals

* Mempermudah masyarakat melaporkan masalah publik.
* Memberikan transparansi proses penanganan laporan.
* Mempermudah distribusi laporan kepada dinas terkait.
* Menyediakan riwayat progres penanganan laporan.

---

# User Roles

## 1. User (Masyarakat)

Warga Kabupaten Toba yang membuat laporan.

### Permissions

* Register akun
* Login
* Logout
* Mengelola profil
* Membuat laporan
* Upload foto laporan
* Melihat laporan milik sendiri
* Melihat detail laporan
* Melihat status laporan
* Melihat riwayat progres laporan
* Menerima notifikasi perubahan status

### Features

#### Authentication

* Register
* Login
* Logout

#### Profile

* Lihat profil
* Edit profil

#### Report Management

* Buat laporan
* Upload foto
* Tambahkan lokasi
* Tambahkan deskripsi

#### Report Tracking

* Daftar laporan saya
* Detail laporan
* Timeline progres laporan

#### Notifications

* Notifikasi perubahan status
* Notifikasi laporan selesai

---

## 2. Admin Dinas

Petugas dinas yang menangani laporan sesuai dinasnya.

Admin Dinas hanya dapat mengakses laporan yang ditugaskan ke dinasnya.

### Permissions

* Login
* Melihat laporan dinas
* Melihat detail laporan
* Mengubah status laporan
* Menambahkan catatan progres
* Upload bukti penyelesaian

### Features

#### Dashboard

* Total laporan aktif
* Total laporan selesai
* Total laporan menunggu tindakan

#### Report Management

* Daftar laporan dinas
* Detail laporan
* Filter laporan berdasarkan status

#### Report Handling

* Update status
* Tambah progres
* Upload foto bukti
* Tandai laporan selesai

---

## 3. Super Admin

Role tertinggi dalam sistem.

### Permissions

* Mengakses seluruh laporan
* Mengelola dinas
* Mengelola akun admin dinas
* Memverifikasi laporan
* Menentukan dinas tujuan laporan
* Mengelola statistik sistem

### Features

#### Dashboard

* Total laporan
* Total laporan aktif
* Total laporan selesai
* Statistik per dinas

#### Report Management

* Melihat seluruh laporan
* Verifikasi laporan
* Menolak laporan
* Menentukan dinas tujuan
* Reassign laporan ke dinas lain

#### Agency Management

* Tambah dinas
* Edit dinas
* Hapus dinas

#### User Management

* Tambah admin dinas
* Edit admin dinas
* Nonaktifkan admin dinas

---

# Agencies (Dinas)

## Dinas PUPR

Menangani:

* Jalan rusak
* Jembatan rusak
* Drainase tersumbat
* Infrastruktur umum

Contoh laporan:

* Jalan Berlubang di Desa Lumban Gaol

---

## Dinas Lingkungan Hidup

Menangani:

* Sampah menumpuk
* Pencemaran lingkungan
* Pohon tumbang
* Kebersihan fasilitas umum

Contoh laporan:

* Sampah Menumpuk di Pasar Balige

---

## Dinas Perhubungan

Menangani:

* Lampu jalan mati
* Rambu lalu lintas rusak
* Marka jalan pudar

Contoh laporan:

* Lampu Jalan Mati di Simpang Soposurung

---

## Dinas Pendidikan

Menangani:

* Fasilitas sekolah rusak
* Sarana pendidikan umum

Contoh laporan:

* Atap Ruang Kelas Bocor

---

## Dinas Kesehatan

Menangani:

* Fasilitas puskesmas
* Sarana pelayanan kesehatan masyarakat

Contoh laporan:

* Toilet Puskesmas Tidak Dapat Digunakan

---

# Workflow

1. User membuat laporan.
2. Laporan masuk ke Super Admin.
3. Super Admin memverifikasi laporan.
4. Super Admin menentukan dinas tujuan.
5. Admin Dinas menerima laporan.
6. Admin Dinas memperbarui status.
7. User menerima pembaruan status.
8. Laporan selesai.

---

# Report Status

* PENDING_VERIFICATION
* ASSIGNED
* IN_PROGRESS
* COMPLETED
* REJECTED

---

# MVP Scope

## Included

* Authentication
* Role Based Access Control
* Report Management
* Photo Upload
* Location Storage
* Status Tracking
* Progress History
* Notification System
* Admin Dashboard

## Excluded

* Chat Realtime
* AI Classification
* Multi Kabupaten
* WhatsApp Integration
* Public Interactive Map
* Rating System
* Anonymous Reporting
* Machine Learning

---

# Technology Stack

## Mobile

* Flutter
* BLoC

## Backend

* Golang
* Gin Framework
* GORM
* JWT Authentication

## Database

* PostgreSQL

## Storage

* Cloudinary atau MinIO

## Admin Panel

* Laravel Filament

---

# Initial Database Tables

* users
* roles
* agencies
* reports
* report_images
* report_statuses
* report_histories
* notifications

---

# Sample Dashboard Data

Total Reports: 30

Distribution:

* PUPR: 10
* Lingkungan Hidup: 8
* Perhubungan: 5
* Pendidikan: 4
* Kesehatan: 3

Example Statistics:

* Completed: 18
* In Progress: 8
* Pending: 4

---

# Core Value

TobaLapor berfokus pada transparansi penanganan laporan masyarakat Kabupaten Toba melalui sistem pelacakan status dan riwayat progres yang jelas hingga laporan selesai ditangani.
