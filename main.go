package main

import (
	"fmt"
	"github.com/aalvesjr/calcula_ir/models"
	"log"
	"os"
	"strconv"
)

func main() {
	basesAliquotaINSS := []float32{1556.94, 2594.92, 5189.82}
	basesAliquotaIR := []float32{1903.98, 2826.65, 3751.05, 4664.68}
	parcelasDescontoIR := []float32{142.8, 354.8, 636.13, 869.36}

	// Recebe o salário por argumento
	// ./calcula_ir 6120.32
	arg, err := strconv.ParseFloat(os.Args[1], 32)
	if err != nil || arg < 0 {
		log.Fatal("Formato de salário inválido! Tente: 12345.67")
	}
	salario := float32(arg)

	// Calculo do INSS
	aliquotaInss := models.AliquotaINSS(salario, basesAliquotaINSS)
	inss := models.INSS(salario, aliquotaInss, basesAliquotaINSS[2])

	// Calculo do IR
	baseIR := models.BaseIR(salario, inss)
	aliquotaIR, descontoIR := models.AliquotaEDescontoIR(baseIR, parcelasDescontoIR, basesAliquotaIR)
	irSemDesconto := models.IRSemParcelaDesconto(baseIR, aliquotaIR)
	ir := models.IR(irSemDesconto, descontoIR)

	fmt.Printf("Salário => %.2f\n", salario)
	fmt.Printf("INSS => %.2f\n", inss)
	fmt.Printf("Base IR => %.2f\n", baseIR)
	fmt.Printf("Aliquota IR => %.2f\n", aliquotaIR)
	fmt.Printf("IR sem desconto => %.2f\n", irSemDesconto)
	fmt.Printf("Desconto do IR => %.2f\n", descontoIR)
	fmt.Printf("valor IR => %.2f\n", ir)
}
