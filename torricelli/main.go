package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"log"
	"strconv"
	"strings"
)

//Torricelli Variáveis 
type Torricelli struct {
	Vi, A, S, Si float64
}

//VelocidadeFinal Cálculo da velocidade final utilizando a 
func (t Torricelli) VelocidadeFinal() float64 {
	return math.Sqrt(t.Vi*t.Vi + 2*t.A*(t.S-t.Si))
}


func lerDados()(Vi, A, S, Si float64){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Velocidade inicial(Vi[m]): ")
	text, err := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\r\n")
	checkErr("Erro ao receber Vi", err)
	Vi, err = strconv.ParseFloat(text, 64)
	checkErr("Erro na conversão de string para float", err)
	fmt.Print("Aceleração(A)[m²]: ")
	text, err = reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\r\n")
	checkErr("Erro ao receber A", err)
	A, err = strconv.ParseFloat(text, 64)
	checkErr("Erro na conversão de string para float", err)
	fmt.Print("Espaço final(S)[m]: ")
	text, err = reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\r\n")
	checkErr("Erro ao receber S", err)
	S, err = strconv.ParseFloat(text, 64)
	checkErr("Erro na conversão de string para float", err)
	fmt.Print("Espaço inicial(Si)[m]: ")
	text, err = reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\r\n")
	checkErr("Erro ao receber Si", err)
	Si, err = strconv.ParseFloat(text, 64)
	checkErr("Erro na conversão de string para float", err)
	return Vi, A, S, Si
}

func checkErr(msg string, err error) {
    if err != nil {
        log.Fatal(msg, err.Error())
    }
}

func main() {
	fmt.Println("Cálculo da velocidade final através da Equação de Torricelli V² = Vi² + 2.A.ΔS")
	Vi, A, S, Si := lerDados()
	t := Torricelli{Vi, A, S, Si}
	fmt.Printf("Velocidade final (v): %f m/s", t.VelocidadeFinal())
}
