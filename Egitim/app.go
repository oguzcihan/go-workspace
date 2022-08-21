package main

import "Egitim/Restfull"

func main() {
	//Conditionals.Demo()
	//Loops.Demo()
	//Loops.Workshop()
	//Loops.Workshop2()
	//Loops.Workshop3()
	//Arrays.Demo()
	//Arrays.Demo2()
	//Arrays.Demo3()
	//Slices.Demo()
	//Slices.Demo2()

	//Maps.Demo()
	//Range.Demo()
	//Range.WorkShop()
	//Range.Demo2()

	//sayi := 20
	//Pointers.Demo(&sayi) //* adres & değeri
	//fmt.Println("Maindeki sayi", sayi)
	//
	////arrayler bellekteki adresle giderler
	//sayilar := []int{1, 2, 4}
	//Pointers.Demo2(sayilar)
	//fmt.Println("Maindeki sayi demo2", sayilar[0])
	//fmt.Println("Maindeki sayi demo2", sayilar[1])
	//fmt.Println("Maindeki sayi demo2", sayilar)

	//Struct eğitimi
	//Structs.Demo()
	//Structs.Demo2()

	//Go routines eğitimi
	//İşlemin benzer zamanlarda yapılmasını istersek başına go yazarız
	//go GoRoutines.CiftSayilar()
	//go GoRoutines.TekSayilar()
	//time.Sleep(2 * time.Second)
	//fmt.Println("Bitti")

	//Channels eğitimi
	//cift/tek sayılar için değişken tanımlandı
	//ciftSayiToplamCn := make(chan int)
	//tekSayiToplamCn := make(chan int)
	//go Channels.CiftSayilar(ciftSayiToplamCn)
	//go Channels.TekSayilar(tekSayiToplamCn)

	//channel içindeki değerler değişkene aktarıldı
	//asenktron ortam olmalıydı
	//channel go routine yaşam döngüsünü kontrol etmeyi sağlıyor
	//ciftSayiToplam, tekSayiToplam := <-ciftSayiToplamCn, <-tekSayiToplamCn
	//carpim := ciftSayiToplam * tekSayiToplam
	//fmt.Println("Çarpım:", carpim)

	//Go interface ve polimorfizm kullanımı
	//Interfaces.Demo()
	//Interfaces.Demo2()

	//defer özellikleri incelendi
	//DeferStatement.B()
	//DeferStatement.Test()
	//DeferStatement.Demo3()

	//ErrorHandling.Demo()
	//Interfaces.Demo3() //type assertion teemelleri
	//ErrorHandling.Demo1()
	//ErrorHandling.Demo3(101)

	//string eğitim
	//StringFunctions.Demo()
	//StringFunctions.Demo2()

	//Restfull eğitimi
	Restfull.Demo()
	Restfull.Demo2()
}
