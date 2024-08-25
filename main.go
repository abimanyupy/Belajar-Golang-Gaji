package main

import (
	"fmt"
)

const (
	gajiPokok int = 3000000
	istri int = 200000	
)

func main() {
	var anak int 
	var status string

	fmt.Print("Apakah Anda sudah menikah ? (y/t) : ")
	fmt.Scan(&status)

	if status != "y" && status != "t" {
		fmt.Println("input harus y atau t")
		return 
	}

	fmt.Print("Berapa anak yang Anda miliki ? :")
	fmt.Scan(&anak)

	gajiDibayar, err := hitungGaji(status, anak)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Gaji yang harus dibayar :", gajiDibayar)
	

}

func hitungGaji (status string, anak int) (int, error) {
	total := gajiPokok 

	if status == "y" {
		total += istri
	} 

	total += anak * 100000
	potongan := total * 2/100
	gajiDibayar := total - potongan
	return gajiDibayar, nil
}