package pintu

type Pintu struct {
	Posisi string
}

func (p Pintu) Ketuk() string {
	switch p.Posisi {
	case "kanan":
		return "Knock"
	case "kiri":
		return "Beep"
	}

	return ""
}

func (p Pintu) Buka() string {
	switch p.Posisi {
	case "kiri":
		return "Knock"
	case "kanan":
		return "Beep"
	}

	return ""
}
