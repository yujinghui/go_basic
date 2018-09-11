package main

type Vertex1 struct {
	X, Y float64
}

func Sum(v Vertex1) float64 {
	return v.X + v.Y
}

// func main() {
// var v = Vertex1{3, 4}
// fmt.Print(Sum(v))
// }
