package main

import(
	"fmt"
	"sort"
)

type dados struct{
	nome string
	idade int
}

type paraNome []dados

// Len -> retorna o comprimento do slice
func (ps paraNome) Len()int{
	return len(ps)
}

// Less -> item na posição i é menor que o item na posição j
func (ps paraNome) Less(i, j int)bool{
	return ps[i].nome < ps[j].nome
}

// Swap -> troca os itens
func (ps paraNome) Swap(i,j int){
	ps[i], ps[j] = ps[j], ps[i]
}

func main(){
	kids := []dados{
		{"Davi", 24},
		{"Diego", 30},
		{"Abner", 14},
	}

	sort.Sort(paraNome(kids))
	fmt.Println(kids)
}