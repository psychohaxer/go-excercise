package main

import (
	"bufio"		// Library bufio digunakan untuk membaca input dari pengguna
	"fmt"		// Library fmt digunakan untuk mencetak pesan ke layar 
	"os"		// Library os digunakan untuk membaca input dari pengguna
	"strconv"   // Library strconv digunakan untuk konversi tipe data
	"strings"   // Library strings digunakan untuk memanipulasi string
)

// Memeriksa apakah sebuah bilangan bulat merupakan bilangan palindrom
func cekPalindrom(num int) bool {
	// Konversi bilangan bulat menjadi string
	str := strconv.Itoa(num)
	length := len(str)

	// Perulangan dilakukan dari indeks awal 0 hingga setengah panjang bilangan 
	// (dibulatkan ke bawah jika panjang bilangan ganjil).
	for i := 0; i < length/2; i++ {

		// Seleksi kondisi digunakan untuk membandingkan digit yang berada di posisi i dengan 
		// digit yang berada di posisi yang berlawanan (posisi length-1-i). Jika digit tidak sama, 
		// itu berarti bilangan tidak palindrom.
		if str[i] != str[length-1-i] {

			// Jika ditemukan digit yang tidak sama, fungsi langsung mengembalikan nilai false, 
			// menandakan bahwa bilangan tersebut bukan bilangan palindrom.
			return false
		}
	}
	return true
}

// Menghitung jumlah bilangan palindrom antara dua bilangan bulat
func hitungPalindrom(start, end int) int {
	count := 0

	for num := start; num <= end; num++ {
		if cekPalindrom(num) {
			count++
		}
	}
	return count
}

// Fungsi main
func main() {
	// Membaca input dari pengguna
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Masukkan angka dalam format 'angka1 angka2':")

	// Perulangan dilakukan selama pengguna memasukkan input
	for scanner.Scan() {
		// Memeriksa apakah pengguna memasukkan baris kosong
		line := scanner.Text()
		if line == "" {
			break
		}

		// Memisahkan input menjadi dua angka
		input := strings.Split(line, " ")

		// Periksa apakah input valid
		// Jika input tidak terdiri dari dua angka, maka input tidak valid
		if len(input) != 2 {
			fmt.Println("Format input tidak valid. Mohon masukkan dua angka dipisahkan oleh satu spasi.")
			continue
		}

		// Konversi ke integer
		start, err1 := strconv.Atoi(input[0])
		end, err2 := strconv.Atoi(input[1])

		// Periksa apakah input valid
		// Jika terjadi kesalahan saat konversi, maka err1 atau err2 akan berisi nilai yang tidak nil
		if err1 != nil || err2 != nil {
			fmt.Println("Format input tidak valid. Mohon masukkan dua angka yang valid.")
			continue
		}

		// Periksa apakah input valid
		// Jika angka pertama < 1 atau angka kedua > 10^9 atau angka kedua <= angka pertama, maka input tidak valid
		if start < 1 || end > 1e9 || start >= end {
			fmt.Println("Batas angka tidak valid. Pastikan angka pertama >= 1 dan angka kedua <= 10^9 dan angka kedua > angka pertama.")
			continue
		}

		// Hitung jumlah bilangan palindrom
		count := hitungPalindrom(start, end)
		fmt.Println(count)
	}

	// Periksa apakah terjadi kesalahan saat membaca input
	if err := scanner.Err(); err != nil {
		fmt.Println("Terjadi kesalahan dalam membaca input:", err)
	}

}
