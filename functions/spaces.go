package piscine 

func Spaces(align string, NlineSlice int) string {
	var spaces string
	Nspaces := width() - NlineSlice
	if align == "center" {
		Nspaces/=2
	}
	for i := 1; i < Nspaces; i++ {
		spaces += " "
	}
	return spaces
}