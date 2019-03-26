package simpleshape

/*Shape is an interface.*/
type Shape interface {
	Area() float64
}

/*ShapeArea returns the method defined in the Shape interface.*/
func ShapeArea(s Shape) float64 {
	return s.Area()
}
