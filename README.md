Nama Proyek: PairProject Hacktiv8 theme "Beverage"
Waktu Pengerjaan: 3 Hari
Deskripsi Singkat: Aplikasi Command Line Interface (CLI) untuk manajemen operasional kafe yang dibangun menggunakan bahasa Go dan database MySQL. Sistem ini memfasilitasi manajemen pengguna, pengelolaan produk dan add-ons, proses transaksi pesanan pelanggan, serta pembuatan laporan penjualan toko. 
Poin Pencapaian:
•	Menerapkan pola desain Layered Architecture (terdiri dari Model, Repository, Usecase, dan Handler) untuk memisahkan logika bisnis dari antarmuka interaksi pengguna dan interaksi database, sehingga kode lebih terstruktur dan mudah dipelihara. 
•	Merancang skema relational database menggunakan MySQL yang terdiri dari 9 tabel yang saling berelasi (Users, Products, Categories, Add_ons, Orders, dll.) dengan aturan integritas data seperti Foreign Keys, Unique Constraints, dan referensial aksi ON DELETE CASCADE. 
•	Mengimplementasikan Database Transactions untuk memproses pemesanan (checkout) guna memastikan pencatatan tabel orders, order items, payments, dan pengurangan jumlah stok produk terjadi secara atomik (menggunakan tx.Begin, tx.Commit, dan tx.Rollback). 
•	Membangun sistem Role-Based Access Control (RBAC) dan autentikasi sederhana untuk membedakan hak akses serta menu dasbor antara entitas Admin dan Customer. 
•	Menyusun operasi query SQL yang kompleks (melibatkan JOIN, agregasi SUM/COUNT, dan GROUP BY) untuk menyajikan dasbor analitik bagi Admin, termasuk kalkulasi total pendapatan nyata, 5 produk terlaris, dan pelanggan dengan riwayat pemesanan tertinggi. 
