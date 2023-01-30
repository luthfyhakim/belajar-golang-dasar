package main

import "fmt"

// Defer function adalah function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function selesai di eksekusi
// Defer function akan selalu dieksekusi walaupun terjadi error di function yang dieksekusi

// Panic function adalah function yang bisa kita gunakan untuk menghentikan program
// Panic function biasanya dipanggil ketika terjadi error pada saat program kita berjalan
// Saat panic function dipanggil, program akan terhenti, namun defer function tetap akan dieksekusi

// Recover adalah function yang bisa kita gunakan untuk menangkap data panic
// Dengan recover proses panic akan terhenti, sehingga program akan tetap berjalan

func endApp() {
	fmt.Println("Aplikasi Selesai")

	// recover() hrs diletakkan pd defer function saat terpanggil
	message := recover() // menangkap data dr panic
	if message != nil {
		fmt.Println("Terjadi Error = ", message)
	}
}

func startApp(value bool) {
	defer endApp() // call endApp() when error or not
	if value {
		panic("Aplikasi Error") // stop program
	} else {
		fmt.Println("Aplikasi Berjalan!")
	}
}

func main() {
	startApp(true)
	fmt.Println("Hakim")
}
