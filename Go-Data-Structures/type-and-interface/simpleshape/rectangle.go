package simpleshape

/*Rectangle is the type struct for a rectangle.*/
type Rectangle struct {
	width  float64
	height float64
}

/*NewRectangle is to create a new instance of type Rectangle.*/
func NewRectangle(w float64, h float64) *Rectangle {
	return &Rectangle{width: w, height: h}
}

/*Area is for calcualte the area of a rectangle*/
func (r *Rectangle) Area() float64 {
	return r.width * r.height
}
