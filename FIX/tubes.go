package main

import "fmt"

// ====================================================================
// SARAH: Struct, Searching, dan Manajemen Mood

type CatatanMood struct {
	Tanggal   string
	SkorEmosi int
	Deskripsi string
}

type Tugas struct {
	Nama      string
	Prioritas int
	Durasi    int
	Status    string
}

const NMax = 100

var daftarMood [NMax]CatatanMood
var daftarTugas [NMax]Tugas

// --- 1. SEARCHING ---

// Cari indeks mood berdasarkan tanggal
func seqSearchMood(n int, target string) int {
	ketemu := -1
	i := 0
	for i < n && ketemu == -1 {
		if daftarMood[i].Tanggal == target {
			ketemu = i
		}
		i++
	}
	return ketemu
}

// Cari indeks tugas berdasarkan nama
func seqSearchTugas(n int, target string) int {
	ketemu := -1
	i := 0
	for i < n && ketemu == -1 {
		if daftarTugas[i].Nama == target {
			ketemu = i
		}
		i++
	}
	return ketemu
}

// Cari tugas berdasarkan prioritas (Binary Search)
func binSearchTugasPrioritas(n int, target int) int {
	low := 0
	high := n - 1
	ketemu := -1

	for low <= high && ketemu == -1 {
		mid := (low + high) / 2
		if daftarTugas[mid].Prioritas == target {
			ketemu = mid
		} else if daftarTugas[mid].Prioritas < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return ketemu
}

// --- 2. MANAJEMEN MOOD ---

// Tambah data mood
func tambahMood(n *int) {
	if *n >= NMax {
		fmt.Println("Kapasitas memori mood sudah penuh!")
		return
	}
	var tanggal, deskripsi string
	var skor int

	fmt.Print("Masukkan tanggal (DD/MM/YYYY): ")
	fmt.Scan(&tanggal)
	fmt.Print("Skor emosi (1-10): ")
	fmt.Scan(&skor)
	fmt.Print("Deskripsi perasaan (Gunakan_Garis_Bawah): ")
	fmt.Scan(&deskripsi)

	daftarMood[*n] = CatatanMood{Tanggal: tanggal, SkorEmosi: skor, Deskripsi: deskripsi}
	*n++
	fmt.Println("Mood berhasil ditambahkan!")
}

// Tampilkan semua mood
func tampilMood(n int) {
	fmt.Println("\n=== DAFTAR CATATAN SUASANA HATI ===")
	if n == 0 {
		fmt.Println("Belum ada data.")
		return
	}
	for i := 0; i < n; i++ {
		fmt.Printf("%d. [%s] Skor: %d | Perasaan: %s\n", i+1, daftarMood[i].Tanggal, daftarMood[i].SkorEmosi, daftarMood[i].Deskripsi)
	}
}

// Ubah data mood
func ubahMood(n int) {
	var target string
	fmt.Print("Masukkan Tanggal Mood yang ingin diubah (DD/MM/YYYY): ")
	fmt.Scan(&target)

	idx := seqSearchMood(n, target)
	if idx != -1 {
		var skor int
		var deskripsi string
		fmt.Print("Skor emosi baru (1-10): ")
		fmt.Scan(&skor)
		fmt.Print("Deskripsi perasaan baru (Gunakan_Garis_Bawah): ")
		fmt.Scan(&deskripsi)

		daftarMood[idx].SkorEmosi = skor
		daftarMood[idx].Deskripsi = deskripsi
		fmt.Println("Data mood berhasil diubah!")
	} else {
		fmt.Println("Data mood tidak ditemukan.")
	}
}

// Hapus data mood
func hapusMood(n *int) {
	var target string
	fmt.Print("Masukkan Tanggal Mood yang ingin dihapus (DD/MM/YYYY): ")
	fmt.Scan(&target)

	idx := seqSearchMood(*n, target)
	if idx != -1 {
		for i := idx; i < *n-1; i++ {
			daftarMood[i] = daftarMood[i+1]
		}
		*n--
		fmt.Println("Data mood berhasil dihapus!")
	} else {
		fmt.Println("Data mood tidak ditemukan.")
	}
}

// ====================================================================
// PANDU: Manajemen Tugas dan Sorting

// --- 3. MANAJEMEN TUGAS ---

// Tambah data tugas
func tambahTugas(n *int) {
	if *n >= NMax {
		fmt.Println("Kapasitas memori tugas sudah penuh!")
		return
	}
	var nama string
	var prioritas, durasi int

	fmt.Print("Nama tugas (Gunakan_Garis_Bawah): ")
	fmt.Scan(&nama)
	fmt.Println("Prioritas: 1=Tinggi, 2=Sedang, 3=Rendah")
	fmt.Print("Pilih prioritas (1/2/3): ")
	fmt.Scan(&prioritas)
	fmt.Print("Durasi pengerjaan (menit): ")
	fmt.Scan(&durasi)

	daftarTugas[*n] = Tugas{Nama: nama, Prioritas: prioritas, Durasi: durasi, Status: "Belum"}
	*n++
	fmt.Println("Tugas berhasil ditambahkan!")
}

// Tampilkan semua tugas
func tampilTugas(n int) {
	fmt.Println("\n=== DAFTAR TUGAS HARIAN ===")
	if n == 0 {
		fmt.Println("Belum ada data tugas.")
		return
	}
	for i := 0; i < n; i++ {
		simbol := "[ ]"
		if daftarTugas[i].Status == "Selesai" {
			simbol = "[v]"
		}
		fmt.Printf("%d. %s %s | Prioritas: %d | Durasi: %d menit\n", i+1, simbol, daftarTugas[i].Nama, daftarTugas[i].Prioritas, daftarTugas[i].Durasi)
	}
}

// Tandai tugas selesai
func ubahStatusTugas(n int) {
	var target string
	fmt.Print("Masukkan Nama Tugas yang sudah selesai: ")
	fmt.Scan(&target)

	idx := seqSearchTugas(n, target)
	if idx != -1 {
		daftarTugas[idx].Status = "Selesai"
		fmt.Println("Status berhasil diubah menjadi Selesai!")
	} else {
		fmt.Println("Tugas tidak ditemukan.")
	}
}

// Hapus data tugas
func hapusTugas(n *int) {
	var target string
	fmt.Print("Masukkan Nama Tugas yang ingin dihapus: ")
	fmt.Scan(&target)

	idx := seqSearchTugas(*n, target)
	if idx != -1 {
		for i := idx; i < *n-1; i++ {
			daftarTugas[i] = daftarTugas[i+1]
		}
		*n--
		fmt.Println("Tugas berhasil dihapus!")
	} else {
		fmt.Println("Tugas tidak ditemukan.")
	}
}

// --- 4. SORTING ---

// Urutkan tugas berdasarkan prioritas (Selection Sort)
func selectionSortPrioritas(n int, asc bool) {
	for i := 0; i < n-1; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			if asc {
				if daftarTugas[j].Prioritas < daftarTugas[idx].Prioritas {
					idx = j
				}
			} else {
				if daftarTugas[j].Prioritas > daftarTugas[idx].Prioritas {
					idx = j
				}
			}
		}
		temp := daftarTugas[i]
		daftarTugas[i] = daftarTugas[idx]
		daftarTugas[idx] = temp
	}
	if asc {
		fmt.Println("\n[Selesai] Diurutkan berdasarkan Prioritas (Ascending).")
	} else {
		fmt.Println("\n[Selesai] Diurutkan berdasarkan Prioritas (Descending).")
	}
}

// Urutkan tugas berdasarkan durasi (Insertion Sort)
func insertionSortDurasi(n int, asc bool) {
	for i := 1; i < n; i++ {
		key := daftarTugas[i]
		j := i - 1

		for j >= 0 && ((asc && daftarTugas[j].Durasi > key.Durasi) || (!asc && daftarTugas[j].Durasi < key.Durasi)) {
			daftarTugas[j+1] = daftarTugas[j]
			j--
		}
		daftarTugas[j+1] = key
	}
	if asc {
		fmt.Println("\n[Selesai] Diurutkan berdasarkan Durasi (Ascending).")
	} else {
		fmt.Println("\n[Selesai] Diurutkan berdasarkan Durasi (Descending).")
	}
}

// ====================================================================
// DAMAI: Statistik dan Program Utama (Main)

// --- 5. STATISTIK ---

// Statistik mood mingguan
func statistikMoodMingguan(n int) {
	fmt.Println("\n=== STATISTIK TREN MOOD MINGGUAN ===")
	if n == 0 {
		fmt.Println("Belum ada data mood.")
		return
	}
	total := 0
	for i := 0; i < n; i++ {
		fmt.Printf("%d. %s | Skor: %d | ", i+1, daftarMood[i].Tanggal, daftarMood[i].SkorEmosi)
		for j := 0; j < daftarMood[i].SkorEmosi; j++ {
			fmt.Print("*")
		}
		fmt.Println()
		total += daftarMood[i].SkorEmosi
	}

	rataRata := float64(total) / float64(n)
	fmt.Printf("\nRata-rata skor emosi: %.2f\n", rataRata)
	if rataRata >= 8 {
		fmt.Println("Kesimpulan: Mood sangat baik.")
	} else if rataRata >= 6 {
		fmt.Println("Kesimpulan: Mood cukup baik.")
	} else if rataRata >= 4 {
		fmt.Println("Kesimpulan: Mood kurang stabil.")
	} else {
		fmt.Println("Kesimpulan: Mood perlu diperhatikan.")
	}
}

// Statistik penyelesaian tugas
func statistikPenyelesaianTugas(n int) {
	fmt.Println("\n=== STATISTIK PENYELESAIAN TUGAS ===")
	if n == 0 {
		fmt.Println("Belum ada data tugas.")
		return
	}
	selesai, belum := 0, 0
	for i := 0; i < n; i++ {
		if daftarTugas[i].Status == "Selesai" {
			selesai++
		} else {
			belum++
		}
	}
	persenSelesai := float64(selesai) / float64(n) * 100
	persenBelum := float64(belum) / float64(n) * 100

	fmt.Printf("Total tugas           : %d\n", n)
	fmt.Printf("Tugas selesai         : %d (%.2f%%)\n", selesai, persenSelesai)
	fmt.Printf("Tugas belum selesai   : %d (%.2f%%)\n", belum, persenBelum)
}

// --- 6. MAIN PROGRAM ---
func main() {
	var pilihan, urutPilihan int
	var nMood, nTugas int
	berjalan := true

	for berjalan {
		fmt.Println("\n====================================================")
		fmt.Println("                 APLIKASI MINDFLOW")
		fmt.Println("====================================================")
		fmt.Println("1.  Tambah Catatan Mood")
		fmt.Println("2.  Tampilkan Catatan Mood")
		fmt.Println("3.  Ubah Catatan Mood")
		fmt.Println("4.  Hapus Catatan Mood")
		fmt.Println("5.  Tambah Tugas Harian")
		fmt.Println("6.  Tampilkan Tugas Harian")
		fmt.Println("7.  Tandai Tugas Selesai")
		fmt.Println("8.  Hapus Tugas Harian")
		fmt.Println("9.  Cari Tugas (Berdasarkan Prioritas)")
		fmt.Println("10. Urut Tugas (Berdasarkan Prioritas)")
		fmt.Println("11. Urut Tugas (Berdasarkan Durasi)")
		fmt.Println("12. Statistik Tren Mood Mingguan")
		fmt.Println("13. Statistik Penyelesaian Tugas")
		fmt.Println("0.  Keluar Aplikasi")
		fmt.Println("====================================================")

		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahMood(&nMood)
		case 2:
			tampilMood(nMood)
		case 3:
			tampilMood(nMood)
			ubahMood(nMood)
		case 4:
			tampilMood(nMood)
			hapusMood(&nMood)
		case 5:
			tambahTugas(&nTugas)
		case 6:
			tampilTugas(nTugas)
		case 7:
			tampilTugas(nTugas)
			ubahStatusTugas(nTugas)
		case 8:
			tampilTugas(nTugas)
			hapusTugas(&nTugas)
		case 9:
			var target int
			fmt.Print("Masukkan prioritas yang dicari (1/2/3): ")
			fmt.Scan(&target)

			selectionSortPrioritas(nTugas, true)
			idx := binSearchTugasPrioritas(nTugas, target)

			if idx != -1 {
				fmt.Printf("\n[Ditemukan] Tugas: %s | Prioritas: %d | Durasi: %d menit\n", daftarTugas[idx].Nama, daftarTugas[idx].Prioritas, daftarTugas[idx].Durasi)
			} else {
				fmt.Println("\nTugas tidak ditemukan.")
			}
		case 10:
			fmt.Print("Pilih urutan (1: Ascending / 2: Descending): ")
			fmt.Scan(&urutPilihan)
			selectionSortPrioritas(nTugas, urutPilihan == 1)
			tampilTugas(nTugas)
		case 11:
			fmt.Print("Pilih urutan (1: Ascending / 2: Descending): ")
			fmt.Scan(&urutPilihan)
			insertionSortDurasi(nTugas, urutPilihan == 1)
			tampilTugas(nTugas)
		case 12:
			statistikMoodMingguan(nMood)
		case 13:
			statistikPenyelesaianTugas(nTugas)
		case 0:
			fmt.Println("Terima kasih telah menggunakan MindFlow!")
			berjalan = false
		default:
			fmt.Println("Pilihan menu tidak valid!")
		}
	}
}