package main

func main() {
	canvas := NewCanvas(100, 100)
	canvas.WritePixel(10, 10, NewColor(1, 0, 0))
	canvas.WritePixel(10, 11, NewColor(1, 0, 0))
	canvas.WritePixel(10, 12, NewColor(1, 0, 0))
	canvas.WritePixel(10, 13, NewColor(1, 0, 0))
	canvas.WritePixel(11, 10, NewColor(1, 0, 0))
	canvas.WritePixel(11, 11, NewColor(1, 0, 0))
	canvas.WritePixel(11, 12, NewColor(1, 0, 0))
	canvas.WritePixel(11, 13, NewColor(1, 0, 0))

	canvas.Save()
}
