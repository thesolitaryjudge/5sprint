package personaldata

import "fmt"

// Personal содержит персональные данные пользователя.
type Personal struct {
	Name   string
	Weight float64
	Height float64
}

// Print выводит данные структуры на экран.
func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f\nРост: %.2f\n", p.Name, p.Weight, p.Height)
}