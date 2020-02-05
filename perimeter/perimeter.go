package perimeter

type Perimeter interface{
	Circle_P()
	Triangle_P()
	Rectangle_P()
}

type Circle struct{
	Radius float64
}
type Triangle struct{
	S1,S2,S3 float64
	
}
type Rectangle struct{
	Length float64
	Width float64
}

func (c Circle) Circle_P() float64{
	return 2*3.14*c.Radius
}

func (t Triangle) Triangle_P() float64{
	return t.S1+t.S2+t.S3
}

func (r Rectangle) Rectangle_P() float64{
	return 2*(r.Length+r.Width)
}
