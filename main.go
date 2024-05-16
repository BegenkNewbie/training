package main

import (
	"fmt"
	"training/mobil"
	"training/pintu"
	"training/roda"
)

func main() {

	m := mobil.Mobil{
		PintuKanan: pintu.Pintu{Posisi: "kanan"},
		PintuKiri:  pintu.Pintu{Posisi: "kiri"},
	}

	cek, _ := m.GantiRoda(0, roda.Karet{})
	fmt.Println(cek)
	cek, _ = m.GantiRoda(1, roda.Besi{})
	fmt.Println(cek)
	cek, _ = m.GantiRoda(2, roda.Kayu{})
	fmt.Println(cek)
	cek, _ = m.GantiRoda(3, roda.Karet{})
	fmt.Println(cek)
	cek, err := m.GantiRoda(4, roda.Karet{})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(m.PintuKiri.Buka())
	fmt.Println(m.PintuKanan.Buka())

	fmt.Println(m.PintuKanan.Ketuk())
	fmt.Println(m.PintuKiri.Ketuk())
}
