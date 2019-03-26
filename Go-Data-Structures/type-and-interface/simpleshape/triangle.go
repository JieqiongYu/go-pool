package simpleshape

/*Triangle is a struct.*/
type Triangle struct {
	base   float64
	height float64
}

/*NewTriangle is a method to return a Tirangle. */
// *Triangle means it returns a pointer to one instance of Triangle type
func NewTriangle(b float64, h float64) *Triangle {
	// return a newly created instance of the type
	return &Triangle{base: b, height: h}
}

/*Area is a method to calculate the area of a triangle*/
func (t *Triangle) Area() float64 {
	return (t.base * t.height) / 2
}
