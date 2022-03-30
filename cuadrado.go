package figuras

//estructura
type Cuadrado struct {
	Lado float64
}

//metodo
func (cua *Cuadrado) area() float64 {
	return cua.Lado * cua.Lado
}

//metodo
func (cua *Cuadrado) perimetro() float64 {
	return 4 * cua.Lado
}
