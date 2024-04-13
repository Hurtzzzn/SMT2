package main 
import "fmt"

func hitungMenang(g,k int, jm *int){
	if g>k{
     *jm+=1
	
	}
}

func hitungDraw(g,k int, jd*int){
	if g==k{
		*jd+=1
	}
}

func hitungKalah(g,k int, jk *int){
	if g<k{
		*jk+=1
	}

}

func hitungGolKegolanSelisih(g,k int, jg,jkm,jsg *int ){
	*jg+=g 
	*jkm+=k
	*jsg= *jg-*jkm
}

func hitungJumPoint(jp*int){
	*jp= *jp

}

func main(){
	var i,n, g, k, jm, jd, jk,jg, jsg, jp, jkm int 
	fmt.Scan(&n)
	for i=1; i<=n; i++{
		fmt.Scan(&g, &k)
		hitungMenang(g,k, &jm)
		hitungDraw(g,k, &jd)
		hitungKalah(g,k, &jk)
		hitungGolKegolanSelisih(g,k, &jg, &jkm, &jsg )
		hitungJumPoint(&jp)
		
	}
	fmt.Println(n,jm, jd, jk, jg, jkm,jsg,jp)
}