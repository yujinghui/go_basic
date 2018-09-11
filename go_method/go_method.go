package main

type Vertex struct {
	X, Y float64
}

func (this Vertex) Sum() float64 {
	return this.X + this.Y
}

// func main() {
// v := Vertex{3, 4}
// fmt.Print(v.Sum())
// }
