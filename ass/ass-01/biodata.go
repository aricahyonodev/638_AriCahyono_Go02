package main

import (
	"fmt"
)

type Student struct {
	number int
	nama string
	alamat string
	pekerjaan string
	alasan string
}

func main(){
		var students = []Student{
		{number : 1, nama: "Ari Cahyono", alamat: "Jl. Tanjung sari", pekerjaan: "Developer", alasan: "ingin belajar back end" },
		{number : 2, nama: "Devi Rukmana", alamat: "Jl. Tandes", pekerjaan: "UI/UX", alasan: "Mencoba hal baru" },
		{number : 3, nama: "Daniel Aden", alamat: "Jl. Balongsari", pekerjaan: "Organis", alasan: "Switch Karir" },
		{number : 4, nama: "Ester", alamat: "Jl. Kupang", pekerjaan: "Akutan", alasan: "Mencoba hal baru" },
		{number : 5, nama: "Kusno", alamat: "Jl. teratai", pekerjaan: "Apoteker", alasan: "Memperdalam bahasa Golang" },
	}
	fmt.Print("Masukkan Nomor Absen: ")
    var input int
    fmt.Scanln(&input)
    findStudent(input, students)
	
}

func findStudent(absenNumber int ,listStudent []Student){
	for _, v := range listStudent {
		if v.number == absenNumber {
			fmt.Println("Nama :", v.nama)
			fmt.Println("Alamat :", v.alamat)
			fmt.Println("Pekerjaan :", v.pekerjaan)
			fmt.Println("Alasan :", v.alasan)
		}
	}
	if absenNumber > 6 {
			fmt.Println("Biodata siswa tidak di temukan")
	}
}