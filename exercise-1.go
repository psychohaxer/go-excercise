package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	// "sort"
	"strconv"
)

// Buku adalah tipe data yang menyimpan informasi buku
type Buku struct {
	Kategori string
	Judul    string
	Ukuran   int
}

// indeksKategori adalah kategori buku dengan urutan sesuai dengan indeks
var indeksKategori = []string{"6", "7", "0", "9", "4", "8", "1", "2", "5", "3"}

// parseInput mengonversi notasi buku menjadi objek buku
func parseInput(input string) Buku {
	tokens := strings.Split(input, "")
	kategori := tokens[0]
	judul := tokens[1]
	ukuranArray := tokens[2:4]
	ukuranStr := strings.Join(ukuranArray, "")

	// konversi ukuran ke int
	ukuran, _ := strconv.Atoi(ukuranStr)
	return Buku{kategori, judul, ukuran}
}

func main(){
	// menyimpan input
	input, inputan := "", ""
	var inputanArr []string
	
	// menyimpan objek buku
	var bukuUnsorted []Buku
	var bukuSorted []Buku

	// membuat scanner
	scanner := bufio.NewScanner(os.Stdin)
	
	// menerima input
	fmt.Println("Masukkan input:")
	for scanner.Scan() {
		input = scanner.Text()

		// berhenti jika input kosong
		if input == "" {
			break
		}

		inputan += input
	}

	inputanArr = strings.Split(inputan, " ")

	// fmt.Println("inputan:", inputan)
	// fmt.Println("inputanArr:", inputanArr)

	// perulangan untuk setiap inputanArr
	for _, notasi := range inputanArr {
		buku := parseInput(notasi)
		bukuUnsorted = append(bukuUnsorted, buku)
	}

	fmt.Println("uns:", bukuUnsorted)
	
	// Urutkan bukuUnsorted berdasarkan Ukuran
	for i := 0; i < len(bukuUnsorted); i++ {
		for j := i + 1; j < len(bukuUnsorted); j++ {
			if bukuUnsorted[i].Ukuran > bukuUnsorted[j].Ukuran {
				bukuUnsorted[i], bukuUnsorted[j] = bukuUnsorted[j], bukuUnsorted[i]
			}
		}
	}

	// Balik urutan bukuUnsorted
	for i := 0; i < len(bukuUnsorted)/2; i++ {
		j := len(bukuUnsorted) - i - 1
		bukuUnsorted[i], bukuUnsorted[j] = bukuUnsorted[j], bukuUnsorted[i]
	}


	// Balik urutan indeksKategori
	// for i := 0; i < len(indeksKategori)/2; i++ {
	// 	j := len(indeksKategori) - i - 1
	// 	indeksKategori[i], indeksKategori[j] = indeksKategori[j], indeksKategori[i]
	// }

	// 3A13 5X19 9Y20 2C18 1N20 3N20 1M21 1F14 9A21 3N21 0E13 5G14 8A23 9E22 3N14

	// perulangan untuk setiap indeksKategori
	for _, kategori := range indeksKategori {

		// perulangan untuk setiap bukuUnsorted
		for _, buku := range bukuUnsorted {
		
			// jika kategori sama dengan kategori buku
			if kategori == buku.Kategori {
				// masukkan Buku ke bukuSorted
				bukuSorted = append(bukuSorted, buku)
			}
		}
	}

	fmt.Println("srt:", bukuSorted)
	
}