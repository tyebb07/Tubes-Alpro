package main

import "fmt"

const NMAX = 99

type Layanan struct {
	idLayanan   string
	namaLayanan string
	harga       int
}

type Pasien struct {
	idPasien string
	nama     string
	umur     int
}

type Transaksi struct {
	idTransaksi   string
	tanggal       string
	idPasien      string
	daftarLayanan [10]Layanan
	jmlLayanan    int
	totalBiaya    int
}

type arrPasien [NMAX]Pasien
type arrLayanan [NMAX]Layanan
type arrTransaksi [NMAX]Transaksi

// pencarian
func cariPasien(P arrPasien, nP int, idCari string) int {

/* 	mengembalikan nilai idx sebagai indeks array yang dicari melalui idPasien 
	menggunakan sequential search pada menu cari pasien dan akan mengembalikan 
	nilai idx = -1 jika idPasien yang dicari tidak ditemukan */

	var idx, i int
	idx = -1

	for i = 0; i < nP && idx == -1; i++ {
		if P[i].idPasien == idCari {
			idx = i
		}
	}
	return idx
}

func cariLayanan(L arrLayanan, nL int, idCari string) int {

/* 	mengembalikan nilai idx sebagai indeks array yang dicari melalui idLayanan 
	pada menu cari layanan dan akan mengembalikan nilai idx = -1 jika idLayanan 
	yang dicari tidak ditemukan */
	

	var idx, i int
	idx = -1

	for i = 0; i < nL && idx == -1; i++ {
		if L[i].idLayanan == idCari {
			idx = i
		}
	}
	return idx
}

// sequential search pake nama
func cariPasienByNama(P arrPasien, nP int, namaCari string) int {

/* 	mengembalikan nilai idx sebagai indeks array yang dicari melalui nama pasien 
	pada menu cari pasien dan akan mengembalikan nilai idx = -1 jika nama pasien 
	yang dicari tidak ditemukan */

	var idx, i int
	idx = -1

	for i = 0; i < nP && idx == -1; i++ {
		if P[i].nama == namaCari {
			idx = i
		}
	}
	return idx
}

func urutkanPasienByID(P *arrPasien, nP int) {
	var i, j, min int
	var temp Pasien

	for i = 0; i < nP - 1;i++{
		min = i
		for j = i + 1 ; j < nP; j++{
					if P[j].idPasien < P[min].idPasien {
				min = j
			}
		}
		temp = P[i]
		P[i] = P[min]
		P[min] = temp
	}
}

// binary search pake ID 
func cariPasienByID(P arrPasien, nP int, idCari string) int {
	
/* 	mengembalikan nilai idx sebagai indeks array yang dicari melalui idPasien 
	menggunakan binary search pada menu cari pasien dan akan mengembalikan 
	nilai idx = -1 jika idPasien yang dicari tidak ditemukan */
	
	var kiri, kanan, idx, tengah int

	kiri = 0
	kanan = nP - 1
	idx = -1

	for kiri <= kanan && idx == -1 {
		tengah = (kiri + kanan) / 2

		if P[tengah].idPasien == idCari {
			idx = tengah
		} else if P[tengah].idPasien < idCari {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return idx
}

// pasien
func tambahPasien(P *arrPasien, nP *int) {
/* 	I.S. Terdefinisi array pasien P dan jumlah data nP. */

/*	F.S. Jika array belum penuh, data pasien baru ditambahkan ke P, 
	nilai nP bertambah 1, dan menampilkan pesan "Data berhasil 
	ditambahkan". Jika array penuh, data tidak berubah dan menampilkan 
	pesan "Data sudah penuh, tidak bisa menambah pasien". */
	
	var idBaru string

		if *nP < NMAX {
			fmt.Println("\n--- Tambah Pasien ---")
			fmt.Print("Masukkan ID pasien : ")
			fmt.Scan(&idBaru) 

			if cariPasien(*P, *nP, idBaru) != -1 {
				fmt.Println("Gagal ID pasien sudah ada.. gunakan ID lain")
			} else {
				P[*nP].idPasien = idBaru
				
				fmt.Print("Masukkan Nama      : ")
				fmt.Scan(&P[*nP].nama)
				fmt.Print("Masukkan Umur      : ")
				fmt.Scan(&P[*nP].umur)

				*nP++
				fmt.Println("Data berhasil ditambahkan")
			}
		} else {
			fmt.Println("Data sudah penuh, tidak bisa menambah pasien")
		}
			urutkanPasienByID(P, *nP)
	}

	func tampilkanPasien(P arrPasien, nP int) {

/*	I.S. Terdefinisi array pasien P dan jumlah data nP. */

/* 	F.S. Menampilkan data pasien jika nP > 0 atau menampilkan 
	"Belum ada data pasien" jika nP == 0 */

	var i int
	fmt.Println("\n--- Daftar Pasien ---")

	if nP == 0 {
		fmt.Println("Belum ada data pasien")
	} else {
		for i = 0; i < nP; i++ {
			fmt.Println("----------------------------")
			fmt.Println("No   :", i+1)
			fmt.Println("ID   :", P[i].idPasien)
			fmt.Println("Nama :", P[i].nama)
			fmt.Println("Umur :", P[i].umur)
		}
		fmt.Println("----------------------------")
	}
}

func ubahPasien(P *arrPasien, nP int) {

/* 	I.S. Terdefinisi array pasien P dan jumlah data nP. */

/*	F.S. Jika ID pasien ditemukan, nama dan umur pasien diubah sesuai input baru.
	Jika data kosong atau ID tidak ditemukan, data pasien tidak berubah. */

	var idCari string
	var i int

	fmt.Println("\n--- Ubah Pasien ---")
	if nP == 0 {
		fmt.Println("Data kosong, tidak ada yang bisa diubah.")
	} else {
		fmt.Print("Masukkan ID Pasien yang ingin diubah: ")
		fmt.Scan(&idCari)

		i = cariPasien(*P, nP, idCari)

		if i != -1 {
			fmt.Println("\nData lama:")
			fmt.Println("Nama :", P[i].nama)
			fmt.Println("Umur :", P[i].umur)

			fmt.Print("Nama baru : ")
			fmt.Scan(&P[i].nama)
			fmt.Print("Umur baru : ")
			fmt.Scan(&P[i].umur)

			fmt.Println("Data berhasil diubah")
		} else {
			fmt.Println("ID tidak ditemukan")
		}
	}
}

func hapusPasien(P *arrPasien, nP *int) {

/* 	I.S. Array pasien P dan jumlah data nP terdefinisi. */

/*	F.S. Jika ID pasien ditemukan, data pasien tersebut dihapus,
    elemen setelahnya digeser ke kiri, dan nP berkurang 1.
	Jika data kosong atau ID tidak ditemukan, data pasien tidak berubah. */

	var idCari string
	var i, idx int

	fmt.Println("\n--- Hapus Pasien ---")
	if *nP == 0 {
		fmt.Println("Data kosong, tidak ada yang bisa dihapus")
	} else {
		fmt.Print("Masukkan ID Pasien yang ingin dihapus: ")
		fmt.Scan(&idCari)

		idx = cariPasien(*P, *nP, idCari)

		if idx != -1 {
			for i = idx; i < *nP-1; i++ {
				P[i] = P[i+1]
			}
			*nP--
			fmt.Println("Data berhasil dihapus")
		} else {
			fmt.Println("ID tidak valid atau tidak ditemukan")
		}
	}
}

func menuCariPasien(P arrPasien, nP int) {

/*	I.S. Terdefinisi array P dan jumlah data nP. */

/*	F.S. Mencari data pasien menggunakan nama atau Id pasien,
	dan menampilkan data pasien bila data yang dicari ditemukan
	atau menampilkan pesan "data tidak ditemukan jika data yang 
	dicari tidak ada" */

	var pilih, idx int
	var cari string

	fmt.Println("\n--- Menu Pencarian Pasien ---")
	fmt.Println("1. Berdasarkan nama (Sequential Search)")
	fmt.Println("2. Berdasarkan ID (Binary Search)")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilih)

	if pilih == 1 {
		fmt.Print("Masukkan Nama Pasien yang dicari: ")
		fmt.Scan(&cari)
		idx = cariPasienByNama(P, nP, cari)
	} else if pilih == 2 {
		fmt.Print("Masukkan ID Pasien yang dicari: ")
		fmt.Scan(&cari)
		idx = cariPasienByID(P, nP, cari)
	} else {
		fmt.Println("Pilihan tidak valid")
		idx = -1
	}

	if pilih == 1 || pilih == 2 {
		if idx != -1 {
			fmt.Println("\nData ditemukan")
			fmt.Println("ID Pasien :", P[idx].idPasien)
			fmt.Println("Nama      :", P[idx].nama)
			fmt.Println("Umur      :", P[idx].umur)
		} else {
			fmt.Println("\nData tidak ditemukan")
		}
	}
}

// layanan
func tambahLayanan(L *arrLayanan, nL *int) {

/* 	I.S. Terdefinisi array layanan L dan jumlah data nL. */

/* 	F.S. Jika array belum penuh, satu data layanan baru dapat 
	ditambahkan ke L, nilai nL bertambah 1, dan pesan berhasil, 
	ditampilkan. Jika array penuh, data layanan tidak berubah. */

	if *nL < NMAX {
		fmt.Println("\n--- Tambah Layanan ---")
		fmt.Print("Masukkan ID layanan : ")
		fmt.Scan(&L[*nL].idLayanan)
		fmt.Print("Masukkan Nama       : ")
		fmt.Scan(&L[*nL].namaLayanan)
		fmt.Print("Masukkan harga      : ")
		fmt.Scan(&L[*nL].harga)

		*nL++
		fmt.Println("Data layanan berhasil ditambahkan")
	} else {
		fmt.Println("Data penuh tidak bisa menambah layanan")
	}
}

func tampilkanLayanan(L arrLayanan, nL int) {

/* 	I.S. Terdefinisi array layanan L dan jumlah data nL. */

/* 	F.S. Seluruh data layanan ditampilkan ke layar. Jika 
	nL = 0, ditampilkan pesan bahwa data layanan belum ada. */

	var i int
	fmt.Println("\n---- Daftar Layanan ----")

	if nL == 0 {
		fmt.Println("Belum ada data layanan")
	} else {
		for i = 0; i < nL; i++ {
			fmt.Println("----------------------------")
			fmt.Println("No     :", i+1)
			fmt.Println("ID     :", L[i].idLayanan)
			fmt.Println("Nama   :", L[i].namaLayanan)
			fmt.Println("Harga  :", L[i].harga)
		}
		fmt.Println("----------------------------")
	}
}

// transaksi sama sorting
func tambahKunjungan(T *arrTransaksi, nT *int, P arrPasien, nP int, L arrLayanan, nL int) {

/* 	I.S. Terdefinisi array transaksi T, jumlah transaksi nT, data pasien P, dan data layanan L. */

/* 	F.S. Jika data transaksi belum penuh dan pasien terdaftar, satu transaksi baru dicatat ke T, 
	layanan yang valid dimasukkan, total biaya dihitung, dan nT bertambah 1. Jika pasien tidak 
	terdaftar atau transaksi penuh, data transaksi tidak bertambah. */

	var idTrans, tgl, idPasien, idLay string
	var idxPasien, idxLay int
	var lanjut bool

	if *nT < NMAX {
		fmt.Println("\n----- Catat Transaksi ------")
		fmt.Print("ID Transaksi : ")
		fmt.Scan(&idTrans)
		fmt.Print("Tanggal      : ")
		fmt.Scan(&tgl)
		fmt.Print("ID Pasien    : ")
		fmt.Scan(&idPasien)

		idxPasien = cariPasien(P, nP, idPasien)

		if idxPasien != -1 {
			T[*nT].idTransaksi = idTrans
			T[*nT].tanggal = tgl
			T[*nT].idPasien = idPasien
			T[*nT].jmlLayanan = 0
			T[*nT].totalBiaya = 0

			lanjut = true
			for T[*nT].jmlLayanan < 10 && lanjut {
				fmt.Print("Masukkan ID Layanan (ketik STOP selesai): ")
				fmt.Scan(&idLay)

				if idLay == "STOP" {
					lanjut = false
				} else {
					idxLay = cariLayanan(L, nL, idLay)
					if idxLay != -1 {
						T[*nT].daftarLayanan[T[*nT].jmlLayanan] = L[idxLay]
						T[*nT].totalBiaya = T[*nT].totalBiaya + L[idxLay].harga
						T[*nT].jmlLayanan++
						fmt.Println("Layanan masuk ke nota")
					} else {
						fmt.Println("ID Layanan tidak valid")
					}
				}
			}
			*nT++
			fmt.Println("Transaksi berhasil dicatat!")
		} else {
			fmt.Println("Pasien tidak terdaftar!")
		}
	} else {
		fmt.Println("Data transaksi penuh!")
	}
}

func tampilkanTransaksi(T arrTransaksi, nT int) {

/* 	I.S. Array transaksi T dan jumlah data nT terdefinisi. */

/*	F.S. Seluruh riwayat transaksi ditampilkan ke layar. 
	Jika nT = 0, ditampilkan pesan bahwa belum ada riwayat transaksi. */

	var i int
	fmt.Println("\n--- Riwayat Transaksi ---")

	if nT == 0 {
		fmt.Println("Belum ada riwayat transaksi")
	} else {
		for i = 0; i < nT; i++ {
			fmt.Println("----------------------------")
			fmt.Printf("Tanggal: %s | ID Trans: %s | Pasien: %s | Total Biaya: Rp%d\n", T[i].tanggal, T[i].idTransaksi, T[i].idPasien, T[i].totalBiaya)
		}
		fmt.Println("----------------------------")
	}
}

// sorting
// selection sort
func selectionSortBiayaAsc(T *arrTransaksi, nT int) {

/* 	I.S. Terdefinisi array transaksi T dan jumlah data nT dalam kondisi acak. */

/* 	F.S. Data transaksi pada T terurut menaik berdasarkan totalBiaya menggunakan selection sort. */

	var i, j, min int
	var temp Transaksi

	for i = 0; i < nT-1; i++ {
		min = i
		for j = i + 1; j < nT; j++ {
			if T[j].totalBiaya < T[min].totalBiaya {
				min = j
			}
		}
		temp = T[i]
		T[i] = T[min]
		T[min] = temp
	}
	fmt.Println("Data berhasil diurutkan berdasarkan biaya secara ascending")
}

func selectionSortBiayaDesc(T *arrTransaksi, nT int) {

/*	I.S. Terdefinisi array transaksi T dan jumlah data nT  dalam kondisi acak. */

/* 	F.S. Data transaksi pada T terurut menurun berdasarkan totalBiaya menggunakan selection sort. */

	var i, j, max int
	var temp Transaksi

	for i = 0; i < nT-1; i++ {
		max = i
		for j = i + 1; j < nT; j++ {
			if T[j].totalBiaya > T[max].totalBiaya {
				max = j
			}
		}
		temp = T[i]
		T[i] = T[max]
		T[max] = temp
	}
	fmt.Println("Data berhasil diurutkan berdasarkan Biaya secara Descending")
}

func ubahFormatTanggal(tgl string) string {

/* 	I.S. : String tgl terdefinisi dan berisi tanggal transaksi 
	dengan format "DD-MM-YYYY" (Contoh: "24-04-2006") */

/*	F.S. : Mengembalikan string baru dengan format yang sudah 
	dipotong dan dibalik menjadi "YYYY-MM-DD" (Contoh: "2006-04-24") 
	ini buat data tanggal dapat dibandingkan secara presisi mulai dari tahun, bulan, hingga hari */
	
	var tahun, bulan, hari string
	tahun = tgl[6:10] 
	bulan = tgl[3:5]  
	hari = tgl[0:2]   
	
	return tahun + "-" + bulan + "-" + hari
}

// insertion sort
func insertionSortTanggalAsc(T *arrTransaksi, nT int) {

/* 	I.S. Terdefinisi array transaksi T dan jumlah data nT dalam kondisi acak. */

/* 	F.S. Data transaksi pada T terurut menaik berdasarkan tanggal menggunakan insertion sort. */

	var i, j int
	var key Transaksi

	for i = 1; i < nT; i++ {
		key = T[i]
		j = i - 1

		for j >= 0 && ubahFormatTanggal(T[j].tanggal) > ubahFormatTanggal(key.tanggal) {
			T[j+1] = T[j]
			j--
		}

		T[j+1] = key
	}
	fmt.Println("Data berhasil diurutkan berdasarkan Tanggal secara ascending")
}

func insertionSortTanggalDesc(T *arrTransaksi, nT int) {

/* 	I.S. Terdefinisi array transaksi T dan jumlah data nT dalam kondisi acak. */

/* 	F.S. Data transaksi pada T terurut menurun berdasarkan tanggal menggunakan insertion sort. */

var i, j int
	var key Transaksi

	for i = 1; i < nT; i++ {
		key = T[i]
		j = i - 1

		for j >= 0 && ubahFormatTanggal(T[j].tanggal) < ubahFormatTanggal(key.tanggal) {
			T[j+1] = T[j]
			j--
		}

		T[j+1] = key
	}
	fmt.Println("Data berhasil diurutkan berdasarkan Tanggal (Lengkap) secara Descending")
}

// statistik
func statistik(T arrTransaksi, nT int, L arrLayanan, nL int) {

/* 	I.S. Terdefinisi array transaksi T, jumlah transaksi nT, array layanan L,
    dan jumlah layanan nL. */
	
/* 	F.S. Ditampilkan jumlah transaksi pada tanggal tertentu
    dan layanan yang paling sering dipesan, jika ada transaksi. */

	var tglCari string
	var i, j, max, idxMax, count, idxL int
	var arrCount [NMAX]int

	fmt.Println("\n=== STATISTIK KLINIK ===")
	fmt.Print("Masukkan tanggal kunjungan (contoh: 21-05-2026): ")
	fmt.Scan(&tglCari)

	count = 0
	for i = 0; i < nT; i++ {
		if T[i].tanggal == tglCari {
			count++
		}
	}
	fmt.Println("Total transaksi pada tanggal", tglCari, ":", count)

	if nT > 0 {
		for i = 0; i < nT; i++ {
			for j = 0; j < T[i].jmlLayanan; j++ {
				idxL = cariLayanan(L, nL, T[i].daftarLayanan[j].idLayanan)
				if idxL != -1 {
					arrCount[idxL]++
				}
			}
		}

		max = 0
		idxMax = -1
		for i = 0; i < nL; i++ {
			if arrCount[i] > max {
				max = arrCount[i]
				idxMax = i
			}
		}

		if idxMax != -1 {
			fmt.Println("Layanan paling laris:", L[idxMax].namaLayanan, "(", max, "kali dipesan )")
		}
	}
}

// menu utama
func menuUtama() {
	fmt.Println("\n===== SIM-KLIK MENU UTAMA =====")
	fmt.Println("1. Tambah pasien")
	fmt.Println("2. Tampilkan pasien")
	fmt.Println("3. Ubah pasien")
	fmt.Println("4. Hapus pasien")
	fmt.Println("5. Cari pasien")
	fmt.Println("6. Tambah layanan")
	fmt.Println("7. Tampilkan layanan")
	fmt.Println("8. Catat kunjungan")
	fmt.Println("9. Tampilkan riwayat transaksi")
	fmt.Println("10. Urutkan transaksi dari biaya ascending")
	fmt.Println("11. Urutkan transaksi dari biaya descending")
	fmt.Println("12. Urutkan transaksi dari tanggal ascending")
	fmt.Println("13. Urutkan transaksi dari tanggal descending")
	fmt.Println("14. Statistik klinik")
	fmt.Println("0. Keluar")
	fmt.Print("Pilih menu: ")
}

func main() {
	var P arrPasien
	var L arrLayanan
	var T arrTransaksi
	var nP, nL, nT int
	var pilih int

	nP = 0
	nL = 0
	nT = 0
	pilih = -1

	for pilih != 0 {
		menuUtama()
		fmt.Scan(&pilih)

		if pilih == 1 {
				tambahPasien(&P, &nP)
			} else if pilih == 2 {
				tampilkanPasien(P, nP)
			} else if pilih == 3 {
				ubahPasien(&P, nP)
			} else if pilih == 4 {
				hapusPasien(&P, &nP)
			} else if pilih == 5 {
				menuCariPasien(P, nP)
			} else if pilih == 6 {
				tambahLayanan(&L, &nL)
			} else if pilih == 7 {
				tampilkanLayanan(L, nL)
			} else if pilih == 8 {
				tambahKunjungan(&T, &nT, P, nP, L, nL)
			} else if pilih == 9 {
				tampilkanTransaksi(T, nT)
			} else if pilih == 10 {
				selectionSortBiayaAsc(&T, nT)
				tampilkanTransaksi(T, nT)
			} else if pilih == 11 {
				selectionSortBiayaDesc(&T, nT)
				tampilkanTransaksi(T, nT)
			} else if pilih == 12 {
				insertionSortTanggalAsc(&T, nT)
				tampilkanTransaksi(T, nT)
			} else if pilih == 13 {
				insertionSortTanggalDesc(&T, nT)
				tampilkanTransaksi(T, nT)
			} else if pilih == 14 {
				statistik(T, nT, L, nL)
			} else if pilih == 0 {
				fmt.Println("==== Terima kasih sudah menggunakan SIM-KLIK ====")
			} else {
				fmt.Println("Menu tidak valid")
			}
	}
}