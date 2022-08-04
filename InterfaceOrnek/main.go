package main

/*
 */
func main() {

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
