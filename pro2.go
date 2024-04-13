package main
import "fmt"
func main() {
	var pilih int 
	var stop bool 
	for !stop{
		fmt.Println()
		fmt.Println("menu()")
		fmt.Println ("Pilih(1/2/3/4)?")
		fmt.Scan(&pilih)
		switch(pilih){
			case 1: hitungJumlah()
			case 2: hitungKali()
			case 3: hitungBagi()
		}
		stop= pilih==4

	}
}

func hitungJumlah(){
	var bil1, bil2, jumlah int 
	fmt.Scan(&bil1, bil2)
	jumlah= bil1+bil2
	fmt.Println( "Hasil Penjumlahan:",jumlah)

	
}

func hitungKali(){
	var bil1, bil2, kali int 
	fmt.Scan(&bil1, &bil2)
	kali= bil1 *bil2
	fmt.Println("Hasil Perkalian:",kali)

}

func hitungBagi(){
	var bil1, bil2, bagi float64 
	fmt.Scan(&bil1, &bil2)
	bagi= bil1/bil2
	fmt.Println("Hasil Pembagian:",bagi)

}
	
