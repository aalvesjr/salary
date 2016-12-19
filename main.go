package main

import (
	"fmt"
	"github.com/aalvesjr/calcula_ir/models"
	"log"
	"os"
	"strconv"
)

func main() {
	// Recebe o salário por argumento
	// ./calcula_ir 6120.32
	arg, err := strconv.ParseFloat(os.Args[1], 32)
	if err != nil || arg < 0 {
		log.Fatal("Formato de salário inválido! Tente: 12345.67")
	}
	salario := models.Salario(arg)

	// Calculo do INSS
	aliquotaInss, baseInss := salario.AliquotaEBaseINSS()
	inss := salario.INSS()

	// Calculo do IR
	baseIR := salario.BaseIR()
	aliquotaIR, descontoIR := salario.AliquotaEDescontoIR()
	irSemDesconto := salario.IRSemParcelaDesconto()
	ir := salario.IR()

	fmt.Printf("Salário         => R$ %.2f\n", salario)
	fmt.Printf("Base INSS       => R$ %.2f\n", baseInss)
	fmt.Printf("Aliquota INSS   => %.2f\n", aliquotaInss)
	fmt.Printf("INSS            => R$ %.2f\n", inss)
	fmt.Printf("Base IR         => R$ %.2f\n", baseIR)
	fmt.Printf("Aliquota IR     => %.2f\n", aliquotaIR)
	fmt.Printf("IR sem desconto => R$ %.2f\n", irSemDesconto)
	fmt.Printf("Desconto do IR  => R$ %.2f\n", descontoIR)
	fmt.Printf("valor IR        => R$ %.2f\n", ir)
}
