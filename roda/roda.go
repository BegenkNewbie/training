package roda

type JenisRoda interface {
	Jenis() string
}

type Roda struct {
	Roda JenisRoda
}

type Kayu struct{}

func (k Kayu) Jenis() string {
	return "Ini kayu"
}

type Besi struct{}

func (b Besi) Jenis() string {
	return "ini Besi"
}

type Karet struct{}

func (k Karet) Jenis() string {
	return "ini Karet"
}
