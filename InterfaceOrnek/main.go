package main

import "strconv"

/*
 */
func main() {
	//Bu şekilde de yapılabilir, Constructor oluşturularakta yapılabilir
	//ferr := new(Ferari)
	//ferr.Brand = "Ferrari"
}
func NewFerrari() {

}

//interface
type Carface interface {
	Run() bool
	Stop() bool
	Information() string
}

// Base structs
type Car struct {
	Brand string
	Model string
	Color string
	Speed int
	Price float64
}

// Özel üretim olduğu gösterildi
type SpecialProduction struct {
	Special bool
}

// ferrari
type Ferari struct {
	//Ortak olan özelliklerden oluşturuldu ve ferrari struct ında çalıştırıldı
	Car //Composition özelliği ile getirildi
	SpecialProduction
}

func (_ Ferari) Run() bool {
	return true
}

func (_ Ferari) Stop() bool {
	return true
}

func (x *Ferari) Information() string {
	ret := "\t" + x.Brand + " " + x.Model + "\n" + "\t" + "Color:" + x.Color + "\n" + "\t" + "Speed:" + strconv.Itoa(x.Speed) + "\n" + "\t" + "Price :" + strconv.FormatFloat(x.Price, 'g', -1, 64) + "Million"
	add := "Yes"
	if x.Special {
		ret += "\n" + "\t" + "Special:" + add
	}
	return ret
}
