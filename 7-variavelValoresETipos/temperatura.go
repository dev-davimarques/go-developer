// Declarando meu pacote principal
package main
// Importar função fmt
import(
	"fmt"
)
// constante de ebulição em fahrenheit
const ebulicaoF float64 = 212.0 
// Função principal
func main(){
	var temperaturaEmFahrenheit float64 = ebulicaoF
	var temperaturaEmCelsius float64 = (temperaturaEmFahrenheit - 32)*5/9
	fmt.Println("Temperatura de Ebulição da Água em Fº é: ",temperaturaEmFahrenheit)
	fmt.Println("Temperatura de Ebulição da Água em Cº é: ",temperaturaEmCelsius)
}