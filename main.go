package main

import "fmt"

const NMAX = 999

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

func cariPasien(P arrPasien, nP int, idCari string) int {
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
	var idx, i int
	idx = -1

	for i = 0; i < nL && idx == -1; i++ {
		if L[i].idLayanan == idCari {
			idx = i
		}
	}
	return idx
}

// sequential search pakai Nama
func cariPasienByNama(P arrPasien, nP int, namaCari string) int {
	var idx, i int
	idx = -1

	for i = 0; i < nP && idx == -1; i++ {
		if P[i].nama == namaCari {
			idx = i
		}
	}
	return idx
}

// binary search berdasarkan ID 
func cariPasienByID(P arrPasien, nP int, idCari string) int {
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

func menuPasien(P *arrPasien, nP *int) {
	var pilih int
	pilih = -1

	for pilih != 0 {
		fmt.Println("\n==== MENU KELOLA PASIEN ====")
		fmt.Println("1.Tambah pasien")
		fmt.Println("2.Tampilkan pasien")
		fmt.Println("3.Ubah pasien")
		fmt.Println("4.Hapus pasien")
		fmt.Println("5.Cari pasien")
		fmt.Println("0.Kembali ke menu utama")
		fmt.Print("Pilih : ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			tambahPasien(P, nP)
		} else if pilih == 2 {
			tampilkanPasien(*P, *nP)
		} else if pilih == 3 {
			ubahPasien(P, *nP)
		} else if pilih == 4 {
			hapusPasien(P, nP)
		} else if pilih == 5 {
			menuCariPasien(*P, *nP)
		} else if pilih != 0 {
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func tambahPasien(P *arrPasien, nP *int) {
	if *nP < NMAX {
		fmt.Println("\n--- Tambah Pasien ---")
		fmt.Print("Masukkan ID Pasien : ")
		fmt.Scan(&P[*nP].idPasien)
		fmt.Print("Masukkan Nama      : ")
		fmt.Scan(&P[*nP].nama)
		fmt.Print("Masukkan Umur      : ")
		fmt.Scan(&P[*nP].umur)

		*nP++
		fmt.Println("Data berhasil ditambahkan")
	} else {
		fmt.Println("Data sudah penuh, tidak bisa menambah pasien")
	}
}

func tampilkanPasien(P arrPasien, nP int) {
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
	var pilih, idx int
	var stringCari string

	fmt.Println("\n--- Menu Pencarian Pasien ---")
	fmt.Println("1. Berdasarkan nama (Sequential Search)")
	fmt.Println("2. Berdasarkan ID (Binary Search)")
	fmt.Print("Pilih: ")
	fmt.Scan(&pilih)

	if pilih == 1 {
		fmt.Print("Masukkan Nama Pasien yang dicari: ")
		fmt.Scan(&stringCari)
		idx = cariPasienByNama(P, nP, stringCari)
	} else if pilih == 2 {
		fmt.Print("Masukkan ID Pasien yang dicari: ")
		fmt.Scan(&stringCari)
		idx = cariPasienByID(P, nP, stringCari)
	} else {
		fmt.Println("Pilihan tidak valid")
		idx = -1
	}

	if pilih == 1 || pilih == 2 {
		if idx != -1 {
			fmt.Println("\nData ditemukan!")
			fmt.Println("ID Pasien :", P[idx].idPasien)
			fmt.Println("Nama      :", P[idx].nama)
			fmt.Println("Umur      :", P[idx].umur)
		} else {
			fmt.Println("\nData tidak ditemukan")
		}
	}
}

func menuLayanan(L *arrLayanan, nL *int) {
	var pilih int
	pilih = -1

	for pilih != 0 {
		fmt.Println("\n=== MENU KELOLA LAYANAN ===")
		fmt.Println("1.Tambah layanan")
		fmt.Println("2.Tampilkan layanan")
		fmt.Println("0.Kembali ke menu utama")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			tambahLayanan(L, nL)
		} else if pilih == 2 {
			tampilkanLayanan(*L, *nL)
		} else if pilih != 0 {
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func tambahLayanan(L *arrLayanan, nL *int) {
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

func menuTransaksi(T *arrTransaksi, nT *int, P arrPasien, nP int, L arrLayanan, nL int) {
	var pilih int
	pilih = -1

	for pilih != 0 {
		fmt.Println("\n=== MENU TRANSAKSI & SORTING ===")
		fmt.Println("1. Catat kunjungan baru")
		fmt.Println("2. Tampilkan riwayat transaksi")
		fmt.Println("3. Urutkan transaksi dari biaya")
		fmt.Println("4. Urutkan transaksi dari tanggal")
		fmt.Println("0. Kembali ke Menu Utama")
		fmt.Print("Pilih: ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			tambahKunjungan(T, nT, P, nP, L, nL)
		} else if pilih == 2 {
			tampilkanTransaksi(*T, *nT)
		} else if pilih == 3 {
			urutkanTransaksiBiayaDesc(T, *nT)
			tampilkanTransaksi(*T, *nT)
		} else if pilih == 4 {
			urutkanTransaksiTanggalAsc(T, *nT)
			tampilkanTransaksi(*T, *nT)
		} else if pilih != 0 {
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func tambahKunjungan(T *arrTransaksi, nT *int, P arrPasien, nP int, L arrLayanan, nL int) {
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
						fmt.Println("-> Layanan masuk ke nota!")
					} else {
						fmt.Println("-> ID Layanan tidak valid!")
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
	var i int
	fmt.Println("\n--- Riwayat Transaksi ---")

	if nT == 0 {
		fmt.Println("Belum ada riwayat transaksi.")
	} else {
		for i = 0; i < nT; i++ {
			fmt.Println("----------------------------")
			fmt.Printf("Tanggal: %s | ID Trans: %s | Pasien: %s | Total Biaya: Rp%d\n", T[i].tanggal, T[i].idTransaksi, T[i].idPasien, T[i].totalBiaya)
		}
		fmt.Println("----------------------------")
	}
}

func urutkanTransaksiBiayaDesc(T *arrTransaksi, nT int) {
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
	fmt.Println(">> Data berhasil diurutkan berdasarkan biaya secara aescending")
}

// insertion sort
func urutkanTransaksiTanggalAsc(T *arrTransaksi, nT int) {
	var i, j int
	var key Transaksi

	for i = 1; i < nT; i++ {
		key = T[i]
		j = i - 1

		for j >= 0 && T[j].tanggal > key.tanggal {
			T[j+1] = T[j]
			j--
		}

		T[j+1] = key
	}
	fmt.Println(">> Data berhasil diurutkan berdasarkan tanggal secara ascending")
}

func statistik(T arrTransaksi, nT int, L arrLayanan, nL int) {
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

func menuUtama() {
	fmt.Println("\n===== SIM-KLIK MENU UTAMA =====")
	fmt.Println("1. Kelola pasien")
	fmt.Println("2. Kelola layanan")
	fmt.Println("3. Transaksi & Pengurutan")
	fmt.Println("4. Statistik klinik")
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
			menuPasien(&P, &nP)
		} else if pilih == 2 {
			menuLayanan(&L, &nL)
		} else if pilih == 3 {
			menuTransaksi(&T, &nT, P, nP, L, nL)
		} else if pilih == 4 {
			statistik(T, nT, L, nL)
		} else if pilih == 0 {
			fmt.Println("Terima kasih sudah menggunakan SIM-KLIK")
		} else {
			fmt.Println("Menu tidak  valid")
		}
	}
}