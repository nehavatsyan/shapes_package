package main
import(
	"fmt"
	"perimeter"
	"area"
)
func main(){
	var n,s int
	fmt.Println("----- Enter ----- \n 1 : Finding Area \n 2 : Finding Perimeter")
	fmt.Scan(&n)
	fmt.Println ("----- Enter ----- \n 1: Triangle \n 2: Rectangle \n 3: Circle" )
	fmt.Scan(&s)

	switch (s){
		case 1:  switch (n) {
					case 1: fmt.Println("Enter base and height")
					var a,b float64
					fmt.Scan(&a,&b)
					var t area.Triangle
					t.Base = a
					t.Height=b
					fmt.Printf("value of Area is %v \n", t.Triangle_area())
					case 2: fmt.Println("Enter sides")
					var a,b,c float64
					fmt.Scan(&a,&b,&c)
					var t perimeter.Triangle
					t.S1 = a
					t.S2 = b
					t.S3 = c
					fmt.Printf("value of Perimeter is %v \n", t.Triangle_P())
				}
		case 2:	switch (n) {
					case 1: fmt.Println("Enter length and width")
					var a,b float64
					fmt.Scan(&a,&b)
					var p area.Rectangle
					p.Length = a
					p.Width = b
					fmt.Printf("value of Area is %v \n", p.Rectangle_area())
					case 2: fmt.Println("Enter length and width")
					var a,b float64
					fmt.Scan(&a,&b)
					var p perimeter.Rectangle
					p.Length = a
					p.Width = b
					fmt.Printf("value of Perimeter is %v \n", p.Rectangle_P())

				}
		case 3:	switch (n) {
					case 1: fmt.Println("Enter radius of circle")
					var a float64
					fmt.Scan(&a)
					var r area.Circle
					r.Radius = a
					fmt.Printf("value of Area is %v \n", r.Circle_area())
					case 2: fmt.Println("Enter radius of circle")
					var a float64
					fmt.Scan(&a)
					var r perimeter.Circle
					r.Radius = a
					fmt.Printf("value of Perimeter is %v \n", r.Circle_P())
				}
	}
}
