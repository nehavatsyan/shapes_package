package area

type area interface {
	Circle_area()
	Triangle_area()
	Rectangle_area()
}
type Circle struct{
	Radius float64
}
type Triangle struct{
	Base float64
	Height float64
}
type Rectangle struct{
	Length float64
	Width float64
}

func (c Circle) Circle_area() float64{
	return 3.14*c.Radius*c.Radius
} 
func (t Triangle) Triangle_area() float64{
	return 0.5*t.Base*t.Height
}
func (r Rectangle) Rectangle_area() float64{
	return r.Length*r.Width
}
