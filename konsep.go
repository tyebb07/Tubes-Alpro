package main

import "fmt"

// =====================================================================
// MASTER PLAN KONSEP SIM-KLIK (KLINIK KECANTIKAN)
// =====================================================================
// ATURAN WAJIB TUBES :
// 1. Pakai Golang dengan algoritma terstruktur
// 2. Wajib modular (fungsi/prosedur) dengan parameter, hindari variabel global di dalam subprogram
// 3. Pakai Array statis dan Struct yang dideklarasikan secara global
// 4. DILARANG pakai 'break' atau 'return' di dalam perulangan (looping)
// =====================================================================

// ---------------------------------------------------------------------
// 0. KONSEP STRUKTUR DATA (WADAH)
// ---------------------------------------------------------------------
// - Butuh konstanta NMAX (misal: 999) agar batas array mudah diubah dari satu tempat
// - Struct Layanan: idLayanan, namaLayanan, harga
// - Struct Pasien: idPasien, nama, umur, noHP
// - Struct Transaksi: idTransaksi, tanggal, idPasien, daftarLayanan (Array khusus maksimal 10), jmlLayanan, totalBiaya
// - Tipe Array: arrPasien, arrLayanan, arrTransaksi untuk membungkus Struct menjadi laci penyimpanan statis sebesar NMAX

// ---------------------------------------------------------------------
// FITUR A: CRUD DATA PASIEN & LAYANAN
// ---------------------------------------------------------------------
// - Tujuan: Admin bisa Menambah (Create), Menampilkan (Read), Mengubah (Update), dan Menghapus (Delete) data
// - Konsep Tambah: Parameter array dan counter jumlah data wajib dikirim menggunakan pointer (*) agar data di wadah aslinya ikut bertambah
// - Konsep Tampil: Cukup looping dari index 0 sampai batas jumlah data saat ini, lalu cetak isinya satu per satu (tidak butuh pointer karena hanya membaca)
// - Konsep Hapus: Karena array statis, penghapusan dilakukan dengan menggeser elemen array yang ada di belakang maju ke depan untuk menutupi index yang dihapus, lalu nilai counter dikurangi (n--)

// ---------------------------------------------------------------------
// FITUR B: PENCATATAN KUNJUNGAN (TRANSAKSI)
// ---------------------------------------------------------------------
// - Tujuan: Kasir mencatat tanggal kedatangan, identitas pasien, dan layanan yang diambil[cite: 17].
// - Alur Logika:
//   1. Input ID Pasien, lalu lakukan pencarian (searching) untuk mengecek validitas pasien[cite: 19, 167]. Jika tidak ketemu (idx = -1), tolak transaksi
//   2. Jika valid, buat perulangan untuk memasukkan layanan satu per satu
//   3. Akali larangan 'break' dengan menggunakan variabel boolean (misal: lanjut = true) sebagai syarat looping
//   4. Jika kasir mengetik kata "STOP", ubah lanjut = false agar perulangan berhenti secara natural
//   5. Hitung total biaya secara otomatis dengan menambahkan harga setiap layanan yang terinput

// ---------------------------------------------------------------------
// FITUR C: PENCARIAN DATA (SEARCHING)
// ---------------------------------------------------------------------
// - Sequential Search (Cari Pasien By Nama):
//   > Logika: Mengecek array satu per satu dari urutan paling depan
//   > Cocok untuk nama karena datanya diinput acak (tidak terurut abjad)
//   > Berhenti looping pakai kondisi tambahan "&& idx == -1" agar tidak melanggar larangan 'break'
// - Binary Search (Cari Pasien By ID):
//   > Syarat mutlak: Array WAJIB diurutkan (sorting) terlebih dahulu berdasarkan ID Pasien
//   > Logika: Langsung melompat ke posisi tengah, lalu membuang setengah area pencarian (kiri atau kanan) hingga data ditemukan

// ---------------------------------------------------------------------
// FITUR D: PENGURUTAN DATA TRANSAKSI (SORTING) -> [BELUM DI-CODE]
// ---------------------------------------------------------------------
// - Tujuan: Menampilkan riwayat transaksi yang sudah terurut berdasarkan Total Biaya atau Tanggal
// - Selection Sort (Berdasarkan Total Biaya):
//   > Bekerja dengan mencari nilai minimum atau maksimum dari sisa deretan array yang belum terurut
//   > Setelah ditemukan, nilainya ditukar (swap) ke posisi paling depan
// - Insertion Sort (Berdasarkan Tanggal):
//   > Algoritma ini bekerja seperti saat kita menyusun kartu remi di tangan
//   > Elemen data ditarik dan disisipkan mundur ke posisi yang tepat di dalam area array yang sudah terurut

// ---------------------------------------------------------------------
// FITUR E: STATISTIK KLINIK (REPORTING) -> [BELUM DI-CODE]
// ---------------------------------------------------------------------
// - Tujuan: Mengetahui jumlah kunjungan harian dan layanan paling laris (juara 1)
// - Statistik Harian:
//   > Minta input tanggal yang ingin dicari, lalu lakukan perulangan pada seluruh array Transaksi
//   > Hitung ada berapa transaksi yang memiliki tanggal yang sama
// - Layanan Paling Laris (Max Finding):
//   > Buat satu array counter sementara yang panjangnya sama dengan jumlah layanan
//   > Looping seluruh riwayat kunjungan pasien, jika layanan tersebut diambil, tambahkan angka pada counter layanan tersebut
//   > Terakhir, lakukan algoritma pencarian nilai terbesar (Maximum) pada array counter tersebut untuk mencetak nama layanan juaranya

// ---------------------------------------------------------------------
// MENU UTAMA (MAIN FUNCTION)
// ---------------------------------------------------------------------
// - Menjadi ruangan kontrol utama tempat variabel array global dideklarasikan secara nyata dan dikirim ke subprogram
// - Menggunakan looping dan switch-case untuk membuat menu interaktif
// - Perulangan menu ditahan menggunakan variabel kondisi (misal: jalan = true). Selama nilainya true, layar akan terus memunculkan menu
// - Begitu user memilih menu Keluar, ubah jalan = false agar program berhenti

func main(){
	var x int
	fmt.Scan(&x)
	fmt.Println(x)
}