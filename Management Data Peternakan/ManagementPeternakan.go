package main
import (
	"fmt"
	"strings"
	"github.com/xuri/excelize/v2"
)

type HewanTernak struct {
	ID    string
	Hewan string
	Jenis string
	Berat int
}

type Peternakan struct {
	DaftarHewan [100]HewanTernak // Array statis
	JumlahHewan int
}

type StatistikJenis struct {
	BeratMin   int
	BeratMax   int
	TotalBerat int
}

func getJenisMap() map[string][]string { // Fungsi untuk mendapatkan mapping jenis hewan (tidak global)
	return map[string][]string{
		"Sapi":    {"Wagyu", "Angus", "Perah"},
		"Ayam":    {"Kampung", "Broiler", "Petelur"},
		"Kambing": {"Kacang", "Etawa", "Bligon"},
		"Kuda":    {"Impor", "Timor"},
	}
}

func InputHewan(index int, jenisMap map[string][]string) HewanTernak { // Fungsi untuk input data hewan
	var id, hewan, jenis string 
	var berat int

	fmt.Printf("\nHewan ke-%d\n", index)
	fmt.Print("ID (maksimal 4 digit): ")
	fmt.Scan(&id)
	for len(id) > 4 || len(id) == 0 {
		fmt.Println("ID tidak valid. Harus maksimal 4 digit dan tidak boleh kosong.")
		fmt.Print("ID (maksimal 4 digit): ")
		fmt.Scan(&id)
	}

	fmt.Print("Hewan (Sapi/Ayam/Kambing/Kuda): ") // Validasi Hewan tanpa break
	fmt.Scan(&hewan)
	hewan = strings.Title(strings.ToLower(hewan))
	for _, ok := jenisMap[hewan]; !ok; _, ok = jenisMap[hewan] {
		fmt.Println("Hewan tidak valid. Masukkan salah satu: Sapi, Ayam, Kambing, Kuda.")
		fmt.Print("Hewan (Sapi/Ayam/Kambing/Kuda): ")
		fmt.Scan(&hewan)
		hewan = strings.Title(strings.ToLower(hewan))
	}
	fmt.Printf("Jenis yang tersedia untuk %s: %s\n", hewan, strings.Join(jenisMap[hewan], ", ")) // Validasi Jenis tanpa break
	fmt.Print("Pilih Jenis: ")
	fmt.Scan(&jenis)
	jenis = strings.Title(strings.ToLower(jenis))
	valid := false
	for _, j := range jenisMap[hewan] {
		if j == jenis {
			valid = true
			break
		}
	}
	for !valid {
		fmt.Println("Jenis tidak valid. Silakan pilih dari daftar.")
		fmt.Print("Pilih Jenis: ")
		fmt.Scan(&jenis)
		jenis = strings.Title(strings.ToLower(jenis))
		valid = false
		for _, j := range jenisMap[hewan] {
			if j == jenis {
				valid = true
				break
			}
		}
	}

	fmt.Print("Berat (kg): ")
	fmt.Scan(&berat)

	return HewanTernak{ID: id, Hewan: hewan, Jenis: jenis, Berat: berat}
}

func InputData(p *Peternakan) { // Fungsi untuk input data peternakan
	var n int
	fmt.Print("Masukkan jumlah hewan ternak: ")
	fmt.Scan(&n)
	jenisMap := getJenisMap()
	for i := 0; i < n && p.JumlahHewan < 100; i++ {
		p.DaftarHewan[p.JumlahHewan] = InputHewan(p.JumlahHewan+1, jenisMap)
		p.JumlahHewan++
	}
}

func CetakData(p Peternakan) { // Fungsi untuk mencetak data
	fmt.Println("\nData Semua Hewan Ternak:")
	for i := 0; i < p.JumlahHewan; i++ {
		h := p.DaftarHewan[i]
		fmt.Printf("ID: %s | Hewan: %s | Jenis: %s | Berat: %dkg\n", h.ID, h.Hewan, h.Jenis, h.Berat)
	}
}

func BinarySearchBerat(p Peternakan, berat int) int { // Fungsi Binary Search berdasarkan berat (data harus diurutkan terlebih dahulu)
	left := 0
	right := p.JumlahHewan - 1
	for left <= right {
		mid := (left + right) / 2
		if p.DaftarHewan[mid].Berat == berat {
			return mid
		} else if p.DaftarHewan[mid].Berat < berat {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func CariBerat(p Peternakan, berat int) { // Fungsi untuk mencari berdasarkan berat dengan Binary Search
	SelectionSortBerat(&p, true) // Urutkan terlebih dahulu agar Binary Search bisa digunakan
	fmt.Printf("\nHasil pencarian berat %dkg:\n", berat)
	index := BinarySearchBerat(p, berat)
	if index != -1 {
		h := p.DaftarHewan[index]
		fmt.Printf("ID: %s | Hewan: %s | Jenis: %s\n", h.ID, h.Hewan, h.Jenis)
	} else {
		fmt.Println("Data dengan berat tersebut tidak ditemukan.")
	}
}

func SelectionSortBerat(p *Peternakan, ascending bool) { // Fungsi Selection Sort berdasarkan berat
	for i := 0; i < p.JumlahHewan-1; i++ {
		extremeIdx := i
		for j := i + 1; j < p.JumlahHewan; j++ {
			if ascending {
				if p.DaftarHewan[j].Berat < p.DaftarHewan[extremeIdx].Berat {
					extremeIdx = j
				}
			} else {
				if p.DaftarHewan[j].Berat > p.DaftarHewan[extremeIdx].Berat {
					extremeIdx = j
				}
			}
		}
		p.DaftarHewan[i], p.DaftarHewan[extremeIdx] = p.DaftarHewan[extremeIdx], p.DaftarHewan[i]
	}
}

func InsertionSortID(p *Peternakan, ascending bool) { // Fungsi Insertion Sort berdasarkan ID
	for i := 1; i < p.JumlahHewan; i++ {
		key := p.DaftarHewan[i]
		j := i - 1
		for j >= 0 {
			if ascending {
				if p.DaftarHewan[j].ID > key.ID {
					p.DaftarHewan[j+1] = p.DaftarHewan[j]
					j--
				} else {
					break
				}
			} else {
				if p.DaftarHewan[j].ID < key.ID {
					p.DaftarHewan[j+1] = p.DaftarHewan[j]
					j--
				} else {
					break
				}
			}
		}
		p.DaftarHewan[j+1] = key
	}
}

func HitungStatistik(p Peternakan, hewan string) StatistikJenis { // Fungsi untuk menghitung statistik per jenis hewan
	var stat StatistikJenis
	count := 0

	for i := 0; i < p.JumlahHewan; i++ {
		if p.DaftarHewan[i].Hewan == hewan {
			berat := p.DaftarHewan[i].Berat
			stat.TotalBerat += berat
			if count == 0 {
				stat.BeratMin = berat
				stat.BeratMax = berat
			} else {
				if berat < stat.BeratMin {
					stat.BeratMin = berat
				}
				if berat > stat.BeratMax {
					stat.BeratMax = berat
				}
			}
			count++
		}
	}

	if count == 0 {
		stat.BeratMin = 0
		stat.BeratMax = 0
		stat.TotalBerat = 0
	}

	return stat
}

func EditData(p *Peternakan, id string) { // Fungsi untuk edit data
	for i := 0; i < p.JumlahHewan; i++ {
		if p.DaftarHewan[i].ID == id {
			fmt.Println("Data lama ditemukan, silakan input data baru:")
			jenisMap := getJenisMap()
			p.DaftarHewan[i] = InputHewan(i+1, jenisMap)
			fmt.Println("Data berhasil diperbarui.")
			return
		}
	}
	fmt.Println("Data dengan ID tersebut tidak ditemukan.")
}

func HapusData(p *Peternakan, id string) { // Fungsi untuk hapus data
	for i := 0; i < p.JumlahHewan; i++ {
		if p.DaftarHewan[i].ID == id {
			for j := i; j < p.JumlahHewan-1; j++ {
				p.DaftarHewan[j] = p.DaftarHewan[j+1]
			}
			p.JumlahHewan--
			fmt.Println("Data berhasil dihapus.")
			return
		}
	}
	fmt.Println("Data dengan ID tersebut tidak ditemukan.")
}

func TambahHewan(p *Peternakan) { // Fungsi untuk menambah hewan baru
	if p.JumlahHewan >= 100 {
		fmt.Println("Kapasitas maksimum tercapai.")
		return
	}
	jenisMap := getJenisMap()
	p.DaftarHewan[p.JumlahHewan] = InputHewan(p.JumlahHewan+1, jenisMap)
	p.JumlahHewan++
	fmt.Println("Data berhasil ditambahkan.")
}

func ExportToExcel(p Peternakan, filename string) {
	f := excelize.NewFile()
	sheet := "DataHewan"
	f.NewSheet(sheet)

	// Header kolom
	headers := []string{"ID", "Hewan", "Jenis", "Berat"}
	for i, h := range headers {
		cell, _ := excelize.CoordinatesToCellName(i+1, 1)
		f.SetCellValue(sheet, cell, h)
	}

	// Isi data
	for i := 0; i < p.JumlahHewan; i++ {
		h := p.DaftarHewan[i]
		f.SetCellValue(sheet, fmt.Sprintf("A%d", i+2), h.ID)
		f.SetCellValue(sheet, fmt.Sprintf("B%d", i+2), h.Hewan)
		f.SetCellValue(sheet, fmt.Sprintf("C%d", i+2), h.Jenis)
		f.SetCellValue(sheet, fmt.Sprintf("D%d", i+2), h.Berat)
	}

	// Hapus sheet default
	f.DeleteSheet("Sheet1")

	// Simpan ke file
	if err := f.SaveAs(filename); err != nil {
		fmt.Println("Gagal menyimpan file Excel:", err)
	} else {
		fmt.Println("Data berhasil diekspor ke file", filename)
	}
}

func ImportFromExcel(p *Peternakan, filename string) {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println("Gagal membuka file:", err)
		return
	}

	rows, err := f.GetRows("DataHewan")
	if err != nil {
		fmt.Println("Gagal membaca sheet:", err)
		return
	}

	p.JumlahHewan = 0 // Reset jumlah
	for i, row := range rows {
		if i == 0 {
			continue // Lewati header
		}
		if len(row) < 4 {
			continue // Lewati baris yang tidak lengkap
		}

		var h HewanTernak
		h.ID = row[0]
		h.Hewan = row[1]
		h.Jenis = row[2]
		fmt.Sscanf(row[3], "%d", &h.Berat)

		if p.JumlahHewan < len(p.DaftarHewan) {
			p.DaftarHewan[p.JumlahHewan] = h
			p.JumlahHewan++
		} else {
			fmt.Println("Kapasitas maksimum tercapai saat impor.")
			break
		}
	}

	fmt.Println("Data berhasil diimpor dari file", filename)
}


func main() { // Fungsi utama
	var peternakan Peternakan
	InputData(&peternakan)

	for {
		fmt.Println("\n=== MENU ===")
		fmt.Println("1. Tampilkan Data")
		fmt.Println("2. Cari berdasarkan Berat")
		fmt.Println("3. Urutkan berdasarkan Berat")
		fmt.Println("4. Urutkan berdasarkan ID")
		fmt.Println("5. Tampilkan Statistik per Jenis")
		fmt.Println("6. Edit Data")
		fmt.Println("7. Hapus Data")
		fmt.Println("8. Tambah Hewan Baru")
		fmt.Println("9. Ekspor ke Excel")
		fmt.Println("10. Impor dari Excel")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			CetakData(peternakan)
		case 2:
			var berat int
			fmt.Print("Masukkan berat yang dicari: ")
			fmt.Scan(&berat)
			CariBerat(peternakan, berat)
		case 3:
			var order string
			fmt.Print("Urutkan ascending (a) atau descending (d): ")
			fmt.Scan(&order)
			ascending := strings.ToLower(order) == "a"
			SelectionSortBerat(&peternakan, ascending)
			fmt.Println("Data telah diurutkan berdasarkan berat.")
			CetakData(peternakan)
		case 4:
			var order string
			fmt.Print("Urutkan ascending (a) atau descending (d): ")
			fmt.Scan(&order)
			ascending := strings.ToLower(order) == "a"
			InsertionSortID(&peternakan, ascending)
			fmt.Println("Data telah diurutkan berdasarkan ID.")
			CetakData(peternakan)
		case 5:
			var hewan string
			fmt.Println("Pilih jenis hewan untuk statistik (Sapi/Ayam/Kambing/Kuda): ")
			fmt.Scan(&hewan)
			hewan = strings.Title(strings.ToLower(hewan))
			validHewan := []string{"Sapi", "Ayam", "Kambing", "Kuda"}
			valid := false
			for _, h := range validHewan {
				if h == hewan {
					valid = true
					break
				}
			}
			if !valid {
				fmt.Println("Jenis hewan tidak valid.")
				break
			}
			stats := HitungStatistik(peternakan, hewan)
			fmt.Printf("\nStatistik %s: Berat Minimal: %dkg, Berat Maksimal: %dkg, Total Berat: %dkg\n",
				hewan, stats.BeratMin, stats.BeratMax, stats.TotalBerat)
		case 6:
			var id string
			fmt.Print("Masukkan ID hewan yang ingin diedit: ")
			fmt.Scan(&id)
			EditData(&peternakan, id)
		case 7:
			var id string
			fmt.Print("Masukkan ID hewan yang ingin dihapus: ")
			fmt.Scan(&id)
			HapusData(&peternakan, id)
		case 8:
			TambahHewan(&peternakan)
		case 9:
	fmt.Print("Masukkan nama file (contoh: data.xlsx): ")
	var filename string
	fmt.Scan(&filename)
	ExportToExcel(peternakan,filename)	
	case 10:
	fmt.Print("Masukkan nama file (contoh: data.xlsx): ")
	var filename string
	fmt.Scan(&filename)
	ImportFromExcel(&peternakan, filename)
		case 0:
			fmt.Println("Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}		
	}
}	