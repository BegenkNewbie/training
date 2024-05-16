package mobil

import (
	"errors"
	"training/pintu"
	"training/roda"
)

type Mobil struct {
	roda       [4]roda.Roda
	PintuKanan pintu.Pintu
	PintuKiri  pintu.Pintu
}

func (m Mobil) GantiRoda(i int, r roda.JenisRoda) (string, error) {
	if i < 0 || i >= 4 {
		return "", errors.New("gak ada dongkrak")
	}
	m.roda[i] = roda.Roda{Roda: r}
	return "Berhasil diganti dengan" + r.Jenis(), nil
}
