package piscine

func Spaces(align string, NlineSlice int) (string, bool) {
	var spaces string
	flag:=false
	Nspaces := width() - NlineSlice
	if align == "center" {
		if Nspaces%2 == 1 {
			flag = true
		} 
		Nspaces /= 2
	}
	for i := 2; i < Nspaces; i++ {
		spaces += " "
	}
	return spaces, flag
}
