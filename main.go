package main

import "fmt"

const NMAX int = 999

type Layanan struct {
	idLayanan   string
	namaLayanan string
	harga       int
}

type Pasien struct {
	idPasien string
	nama     string
	umur     int
	noHP     string
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

func main() {
	var P arrPasien
	var L arrLayanan
	var T arrTransaksi
	var nP, nL, nT int
	var op string

	nP = 0
	nL = 0
	nT = 0
	op = ""

	menuUtama(&P, &nP, &L, &nL, &T, &nT, &op)
}

func menuUtama(P *arrPasien, nP *int, L *arrLayanan, nL *int, T *arrTransaksi, nT *int, op *string) {
	for *op != "9" {
		fmt.Println("\n--- MENU UTAMA ---")
		fmt.Println("1. Tambah Pasien")
		fmt.Println("2. Lihat Pasien")
		fmt.Println("3. Ubah Pasien")   
		fmt.Println("4. Hapus Pasien") 
		fmt.Println("5. Tambah Layanan")
		fmt.Println("6. Lihat Layanan")
		fmt.Println("7. Catat Kunjungan")
		fmt.Println("8. Cari Pasien")
		fmt.Println("9. Keluar")
		fmt.Print("Pilihan: ")
		fmt.Scan(op)

		for *op != "1" && *op != "2" && *op != "3" && *op != "4" && *op != "5" && *op != "6" && *op != "7" && *op != "8" && *op != "9" {
			fmt.Print("Pilihan salah, masukkan lagi: ")
			fmt.Scan(op)
		}

		if *op == "1" {
			tambahPasien(P, nP)
		} else if *op == "2" {
			lihatPasien(*P, *nP)
		} else if *op == "3" {
			ubahPasien(P, *nP)
		} else if *op == "4" {
			hapusPasien(P, nP)
		} else if *op == "5" {
			tambahLayanan(L, nL)
		} else if *op == "6" {
			lihatLayanan(*L, *nL)
		} else if *op == "7" {
			tambahKunjungan(T, nT, *P, *nP, *L, *nL)
		} else if *op == "8" {
			menuCari(*P, *nP)
		}
	}
}

func cariPasien(P arrPasien, nP int, idCari string) int {
	var idx, i int
	idx = -1
	i = 0
	for i < nP && idx == -1 {
		if P[i].idPasien == idCari {
			idx = i
		}
		i = i + 1
	}
	return idx
}

func cariLayanan(L arrLayanan, nL int, idCari string) int {
	var idx, i int
	idx = -1
	i = 0
	for i < nL && idx == -1 {
		if L[i].idLayanan == idCari {
			idx = i
		}
		i = i + 1
	}
	return idx
}

// sequential search pake nama
func cariPasienByNama(P arrPasien, nP int, namaCari string) int {
	var idx, i int
	idx = -1
	i = 0
	for i < nP && idx == -1 {
		if P[i].nama == namaCari {
			idx = i
		}
		i = i + 1
	}
	return idx
}

// binary search pake ID
func cariPasienByID(P arrPasien, nP int, idCari string) int {
	var idx, left, right, mid int
	idx = -1
	left = 0
	right = nP - 1
	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if P[mid].idPasien == idCari {
			idx = mid
		} else if P[mid].idPasien > idCari {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return idx
}

func tambahPasien(P *arrPasien, nP *int) {
	var id, nama, noHP string
	var umur int

	if *nP < NMAX {
		fmt.Print("Masukkan ID Pasien: ")
		fmt.Scan(&id)
		fmt.Print("Masukkan Nama Pasien: ")
		fmt.Scan(&nama)
		fmt.Print("Masukkan Umur Pasien: ")
		fmt.Scan(&umur)
		fmt.Print("Masukkan No HP Pasien: ")
		fmt.Scan(&noHP)

		P[*nP].idPasien = id
		P[*nP].nama = nama
		P[*nP].umur = umur
		P[*nP].noHP = noHP
		*nP = *nP + 1
		fmt.Println("Pasien berhasil ditambahkan")
	} else {
		fmt.Println("Kapasitas pasien penuh")
	}
}

func lihatPasien(P arrPasien, nP int) {
	var i int
	if nP == 0 {
		fmt.Println("Data pasien kosong")
	} else {
		i = 0
		for i < nP {
			fmt.Printf("ID: %s \nNama: %s \nUmur: %d \nNo HP: %s\n", P[i].idPasien, P[i].nama, P[i].umur, P[i].noHP)
			i = i + 1
		}
	}
}

func ubahPasien(P *arrPasien, nP int) {
	var idCari, namaBaru, noHPBaru string
	var umurBaru, idx int

	fmt.Println("\n--- UBAH DATA PASIEN ---")
	fmt.Print("Masukkan ID Pasien yang akan diubah: ")
	fmt.Scan(&idCari)

	idx = cariPasien(*P, nP, idCari)

	if idx != -1 {
		fmt.Printf("Data ditemukan! (Nama Lama: %s)\n", P[idx].nama)
		
		fmt.Print("Masukkan Nama baru  : ")
		fmt.Scan(&namaBaru)
		fmt.Print("Masukkan Umur baru  : ")
		fmt.Scan(&umurBaru)
		fmt.Print("Masukkan No HP baru : ")
		fmt.Scan(&noHPBaru)

		P[idx].nama = namaBaru
		P[idx].umur = umurBaru
		P[idx].noHP = noHPBaru
		fmt.Println("Berhasil: Data pasien sukses diubah!")
	} else {
		fmt.Println("Gagal: ID Pasien tidak ditemukan")
	}
}

// disini pakek sistem geser
func hapusPasien(P *arrPasien, nP *int) {
	var idCari string
	var idx, i int

	fmt.Println("\n--- HAPUS DATA PASIEN ---")
	fmt.Print("Masukkan ID Pasien yang akan dihapus: ")
	fmt.Scan(&idCari)

	idx = cariPasien(*P, *nP, idCari)

	if idx != -1 {
		i = idx
		for i < *nP-1 {
			P[i] = P[i+1]
			i = i + 1
		}
		*nP = *nP - 1
		fmt.Println(">> Berhasil: Data pasien sukses dihapus!")
	} else {
		fmt.Println(">> Gagal: ID Pasien tidak ditemukan.")
	}
}

func tambahLayanan(L *arrLayanan, nL *int) {
	var id, nama string
	var harga int

	if *nL < NMAX {
		fmt.Print("Masukkan ID Layanan: ")
		fmt.Scan(&id)
		fmt.Print("Masukkan Nama Layanan: ")
		fmt.Scan(&nama)
		fmt.Print("Masukkan Harga Layanan: ")
		fmt.Scan(&harga)

		L[*nL].idLayanan = id
		L[*nL].namaLayanan = nama
		L[*nL].harga = harga
		*nL = *nL + 1
		fmt.Println("Layanan berhasil ditambahkan")
	} else {
		fmt.Println("Kapasitas layanan penuh")
	}
}

func lihatLayanan(L arrLayanan, nL int) {
	var i int
	if nL == 0 {
		fmt.Println("Data layanan kosong.")
	} else {
		i = 0
		for i < nL {
			fmt.Printf("ID: %s, Nama: %s, Harga: %d\n", L[i].idLayanan, L[i].namaLayanan, L[i].harga)
			i = i + 1
		}
	}
}

func tambahKunjungan(T *arrTransaksi, nT *int, P arrPasien, nP int, L arrLayanan, nL int) {
	var idTrans, tgl, idPasien, idLay string
	var idxPasien, idxLay, idxLayananSekarang int
	var lanjut bool

	if *nT < NMAX {
		fmt.Print("Masukkan ID Transaksi: ")
		fmt.Scan(&idTrans)
		fmt.Print("Masukkan Tanggal: ")
		fmt.Scan(&tgl)
		fmt.Print("Masukkan ID Pasien: ")
		fmt.Scan(&idPasien)

		idxPasien = cariPasien(P, nP, idPasien)

		if idxPasien == -1 {
			fmt.Println("ID Pasien tidak ditemukan.")
		} else {
			T[*nT].idTransaksi = idTrans
			T[*nT].tanggal = tgl
			T[*nT].idPasien = idPasien
			T[*nT].jmlLayanan = 0
			T[*nT].totalBiaya = 0

			lanjut = true
			for T[*nT].jmlLayanan < 10 && lanjut {
				fmt.Print("Masukkan ID Layanan (Ketik STOP jika selesai): ")
				fmt.Scan(&idLay)

				if idLay == "STOP" {
					lanjut = false
				} else {
					idxLay = cariLayanan(L, nL, idLay)
					if idxLay != -1 {
						idxLayananSekarang = T[*nT].jmlLayanan
						T[*nT].daftarLayanan[idxLayananSekarang] = L[idxLay]
						T[*nT].totalBiaya = T[*nT].totalBiaya + L[idxLay].harga
						T[*nT].jmlLayanan = T[*nT].jmlLayanan + 1
						fmt.Println("Layanan ditambahkan.")
					} else {
						fmt.Println("ID Layanan tidak ada.")
					}
				}
			}
			*nT = *nT + 1
			fmt.Println("Transaksi selesai dicatat.")
		}
	} else {
		fmt.Println("Kapasitas transaksi penuh.")
	}
}

func menuCari(P arrPasien, nP int) {
	var pil, namaCari, idCari string
	var hasil int

	fmt.Println("\n--- MENU CARI ---")
	fmt.Println("1. Cari Berdasarkan Nama")
	fmt.Println("2. Cari Berdasarkan ID")
	fmt.Print("Pilihan: ")
	fmt.Scan(&pil)

	for pil != "1" && pil != "2" {
		fmt.Print("Pilihan salah, masukkan lagi: ")
		fmt.Scan(&pil)
	}

	if pil == "1" {
		fmt.Print("Masukkan Nama Pasien yang dicari: ")
		fmt.Scan(&namaCari)
		hasil = cariPasienByNama(P, nP, namaCari)
		if hasil != -1 {
			fmt.Printf("Data Ditemukan! ID: %s\nNama: %s\nUmur: %d\nNo HP: %s\n", P[hasil].idPasien, P[hasil].nama, P[hasil].umur, P[hasil].noHP)
		} else {
			fmt.Println("Data Pasien tidak ditemukan")
		}
	} else if pil == "2" {
		fmt.Print("Masukkan ID Pasien yang dicari: ")
		fmt.Scan(&idCari)
		hasil = cariPasienByID(P, nP, idCari)
		if hasil != -1 {
			fmt.Printf("Data Ditemukan! ID: %s\nNama: %s\nUmur: %d\nNo HP: %s\n", P[hasil].idPasien, P[hasil].nama, P[hasil].umur, P[hasil].noHP)
		} else {
			fmt.Println("Data Pasien tidak ditemukan")
		}
	}
}