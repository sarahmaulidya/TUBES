package main

import "fmt"

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

var daftarMood = []CatatanMood{}
var daftarTugas = []Tugas{}

func tambahMood() {
	var tanggal string
	var skor int
	var deskripsi string

	fmt.Print("Masukkan tanggal (DD/MM/YYYY): ")
	fmt.Scanln(&tanggal)

	fmt.Print("Skor emosi (1-10): ")
	fmt.Scanln(&skor)

	var dummy string
	fmt.Scanln(&dummy)

	fmt.Print("Deskripsi perasaan: ")
	fmt.Scanln(&deskripsi)

	daftarMood = append(daftarMood, CatatanMood{Tanggal: tanggal, SkorEmosi: skor, Deskripsi: deskripsi})
	fmt.Println("Mood berhasil ditambahkan!")
}

func tampilMood() {
	fmt.Println("\n=== DAFTAR CATATAN SUASANA HATI ===")
	if len(daftarMood) == 0 {
		fmt.Println("Belum ada data.")
		return
	}
	for i, m := range daftarMood {
		fmt.Printf("%d. [%s] Skor: %d | %s\n", i+1, m.Tanggal, m.SkorEmosi, m.Deskripsi)
	}
}

func hapusMood() {
	tampilMood()
	var idx int
	fmt.Print("Nomor mood yang dihapus: ")
	fmt.Scan(&idx)
	idx--

	if idx >= 0 && idx < len(daftarMood) {
		daftarMood = append(daftarMood[:idx], daftarMood[idx+1:]...)
		fmt.Println("Mood dihapus!")
	} else {
		fmt.Println("Nomor tidak valid!")
	}
}

func tambahTugas() {
	var nama string
	var prioritas, durasi int

	fmt.Print("Nama tugas: ")
	fmt.Scanln(&nama)

	fmt.Println("Prioritas: 1=Tinggi, 2=Sedang, 3=Rendah")
	fmt.Print("Pilih prioritas: ")
	fmt.Scan(&prioritas)

	fmt.Print("Durasi (menit): ")
	fmt.Scan(&durasi)

	var dummy string
	fmt.Scanln(&dummy)

	daftarTugas = append(daftarTugas, Tugas{Nama: nama, Prioritas: prioritas, Durasi: durasi, Status: "Belum"})
	fmt.Println("Tugas berhasil ditambahkan!")
}

func tampilTugas() {
	fmt.Println("\n=== DAFTAR TUGAS HARIAN ===")
	if len(daftarTugas) == 0 {
		fmt.Println("Belum ada data.")
		return
	}
	for i, t := range daftarTugas {
		simbol := "○"
		if t.Status == "Selesai" {
			simbol = "✓"
		}
		fmt.Printf("%d. %s %s | Prioritas: %d | %d menit\n", i+1, simbol, t.Nama, t.Prioritas, t.Durasi)
	}
}

func ubahStatusTugas() {
	tampilTugas()
	var idx int
	fmt.Print("Nomor tugas yang diubah: ")
	fmt.Scan(&idx)
	idx--

	if idx >= 0 && idx < len(daftarTugas) {
		daftarTugas[idx].Status = "Selesai"
		fmt.Println("Status diubah menjadi Selesai!")
	} else {
		fmt.Println("Nomor tidak valid!")
	}
}

func urutTugasPrioritas() {
	n := len(daftarTugas)
	for i := 0; i < n-1; i++ {
        minIdx := i
    	for j := i + 1; j < n; j++ {
       		if daftarTugas[j].Prioritas < daftarTugas[minIdx].Prioritas {
            	minIdx = j
        	}
   		}
	daftarTugas[i], daftarTugas[minIdx] = daftarTugas[minIdx], daftarTugas[i]
    }
    fmt.Println("\n[Selesai] Tugas diurutkan berdasarkan Prioritas.")
}

func urutTugasDurasi() {
    n := len(daftarTugas)
    for i := 1; i < n; i++ {
        key := daftarTugas[i]
        j := i - 1
    	for j >= 0 && daftarTugas[j].Durasi > key.Durasi {
        	daftarTugas[j+1] = daftarTugas[j]
            j = j - 1
        }
        daftarTugas[j+1] = key
    }
    fmt.Println("\n[Selesai] Tugas diurutkan berdasarkan Durasi.")
}

func cariPrioritas(target int) {
    low := 0
    high := len(daftarTugas) - 1
    ketemu := -1

    for low <= high {
        mid := (low + high) / 2
        if daftarTugas[mid].Prioritas == target {
            ketemu = mid
            break
        } else if daftarTugas[mid].Prioritas < target {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }

    if ketemu != -1 {
        t := daftarTugas[ketemu]
        fmt.Printf("Ditemukan! Tugas: %s | Durasi: %d menit\n", t.Nama, t.Durasi)
    } else {
        fmt.Println("Tugas dengan prioritas tersebut tidak ditemukan.")
    }
}

func statistikMoodMingguan() {
	fmt.Println("\n=== STATISTIK TREN MOOD MINGGUAN ===")

	if len(daftarMood) == 0 {
		fmt.Println("Belum ada data mood.")
		return
	}

	total := 0

	for i, m := range daftarMood {
		fmt.Printf("%d. %s | Skor Emosi: %d | ", i+1, m.Tanggal, m.SkorEmosi)

		for j := 0; j < m.SkorEmosi; j++ {
			fmt.Print("*")
		}

		fmt.Println()
		total += m.SkorEmosi
	}

	rataRata := float64(total) / float64(len(daftarMood))

	fmt.Printf("\nRata-rata skor emosi minggu ini: %.2f\n", rataRata)

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

func statistikPenyelesaianTugas() {
	fmt.Println("\n=== STATISTIK PENYELESAIAN TUGAS ===")

	if len(daftarTugas) == 0 {
		fmt.Println("Belum ada data tugas.")
		return
	}

	selesai := 0
	belum := 0

	for _, t := range daftarTugas {
		if t.Status == "Selesai" {
			selesai++
		} else {
			belum++
		}
	}

	total := len(daftarTugas)

	persenSelesai := float64(selesai) / float64(total) * 100
	persenBelum := float64(belum) / float64(total) * 100

	fmt.Printf("Total tugas           : %d\n", total)
	fmt.Printf("Tugas selesai         : %d (%.2f%%)\n", selesai, persenSelesai)
	fmt.Printf("Tugas belum selesai   : %d (%.2f%%)\n", belum, persenBelum)

	fmt.Print("Grafik selesai        : ")
	for i := 0; i < selesai; i++ {
		fmt.Print("*")
	}

	fmt.Print("\nGrafik belum selesai  : ")
	for i := 0; i < belum; i++ {
		fmt.Print("*")
	}

	fmt.Println()
}

func main() {
	var pilihan int

	for {
		fmt.Println("\n===================================")
		fmt.Println("         APLIKASI MINDFLOW")
		fmt.Println("===================================")
		fmt.Println("1. Tambah Mood")
		fmt.Println("2. Tampilkan Mood")
		fmt.Println("3. Hapus Mood")
		fmt.Println("4. Tambah Tugas")
		fmt.Println("5. Tampilkan Tugas")
		fmt.Println("6. Ubah Status Tugas")
		fmt.Println("7. Cari Tugas Berdasarkan Prioritas")
		fmt.Println("8. Urutkan Tugas Berdasarkan Prioritas")
		fmt.Println("9. Urutkan Tugas Berdasarkan Durasi")
		fmt.Println("10. Statistik Mood Mingguan")
		fmt.Println("11. Statistik Penyelesaian Tugas")
		fmt.Println("0. Keluar")

		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihan)

		switch pilihan {

		case 1:
			tambahMood()

		case 2:
			tampilMood()

		case 3:
			hapusMood()

		case 4:
			tambahTugas()

		case 5:
			tampilTugas()

		case 6:
			ubahStatusTugas()

		case 7:
			var target int
			fmt.Print("Masukkan prioritas yang dicari: ")
			fmt.Scanln(&target)
			cariPrioritas(target)

		case 8:
			urutTugasPrioritas()
			tampilTugas()

		case 9:
			urutTugasDurasi()
			tampilTugas()

		case 10:
			statistikMoodMingguan()

		case 11:
			statistikPenyelesaianTugas()

		case 0:
			fmt.Println("Terima kasih telah menggunakan MindFlow.")
			return

		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}