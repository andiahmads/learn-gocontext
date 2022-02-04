package contextwithvalue

import (
	"context"
	"fmt"
	"testing"
)

/*
context with value
pada saat awal membuat context, context tidak memiliki value.
kita bisa menambahkan sebua value dengan data pair(key-value) kedalam context.
saat menambah value ke context, secara otomatis akan tercipta child context baru, artinya
original contextnya tidak akan berubah sama sekali.
untuk membuat menambahkan value ke context, kita bisa menggunakan method context.WithValue(parent,key,value) */

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")
	contextG := context.WithValue(contextF, "g", "G")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextG)

	fmt.Println(contextF.Value("f"))
	fmt.Println(contextF.Value("c"))
	fmt.Println(contextE.Value("e"))

}
