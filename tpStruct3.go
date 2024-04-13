package main

import "fmt"

// Tipe bentukan pegawai dengan komponen (field) nama, gaji_pokok, masa_kerja, dan besar_bonus
type pegawai struct {
	nama                   string
	gaji_pokok, masa_kerja int
	besar_bonus            float64
}

func main() {
	// deklarasi variabel bertipe pegawai
	var p1 pegawai

	// baca data pengawai
	fmt.Scan(&p1.nama, &p1.gaji_pokok, &p1.masa_kerja)

	// hitung bonus dengan memanggil prosedur hitung_bonus
	hitung_bonus(&p1)

	// Cetak besar bonus
	fmt.Printf("Pegawai dengan nama %s mendapatkan bonus sebesar Rp %.0f", p1.nama, p1.besar_bonus)
}

func hitung_bonus(p *pegawai) {
	/* IS: p.nama, p.gaji_pokok, p.masa_kerja terdefinisi
	   Proses: Besar bonus dihitung dengan mengalikan masa kerja dengan angka pengali
	           Jika masa kerja minimal 10 tahun, angka pengalinya 1.5
	           jika masa kerja di bawah 10 tahun hingga 5 tahun, angka pengalinya 0.75
			   di bawah 5 tahun, angka pengalinya 0.5
	   FS: p.besar_bonus berisi nilai
	*/
	if p.masa_kerja >= 10 {
		p.besar_bonus = 1.5 * float64(p.gaji_pokok)
	} else if p.masa_kerja >= 5 && p.masa_kerja < 10 {
		p.besar_bonus = 0.75 * float64(p.gaji_pokok)
	} else {
		p.besar_bonus = 0.5 * float64(p.gaji_pokok)
	}
}
