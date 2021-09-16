package main

import (
	"fmt"
)

func main() {
	//Variáveis do tipo inteiro que têm limites de precisão já pré-definidos
	fmt.Println("INTEIROS SEM SINAL");
	var u1 byte = 255;
	fmt.Println(u1);
	var u2 uint16 = 65535;
	fmt.Println(u2);
	var u3 uint32 = 4294967295;
	fmt.Println(u3);

	fmt.Println("********************************************************");

	fmt.Println("INTEIROS COM SINAL");
	var i1 rune = -128;
	fmt.Println(i1);

	//Variáveis inteiro cujos limites de precisão de acordo com a arquitetura do sistema operacional e processador

	var t1 uint = 6553646786468461841;
	fmt.Println(t1);
	var t2 int = -123123123121991231;
	fmt.Println(t2);

	fmt.Println("********************************************************");

	//Variáveis do tipo float
	var f1 float32 = 2.2;
	fmt.Println(f1);

	fmt.Println("********************************************************");

	//Variáveis do tipo complex
	var c1 complex64 = 3;
	fmt.Println(c1);
	var c2 complex128 = 3;
	fmt.Println(c2);

	fmt.Println("********************************************************");


	//Variáveis do tipo string

	var s1 string = `Mário
	Júnior`;
	fmt.Println(s1);
	fmt.Println(s1[0]);

	fmt.Println("********************************************************");


	//Variáveis do tipo bool
	fmt.Println("BOOLEAN")
	var b1 bool = true;
	var b2 bool = false;
	fmt.Println(b1);
	fmt.Println(b2);

	fmt.Println("********************************************************");


	//Declaração de múltiplas variáveis numa mesma linha de código
	fmt.Println("DECLARANDO MÚLTIPLAS VARIÁVEIS NUMA MESMA LINHA DE CÓDIGO");
	var (
		nome string = "Mário";
		sobrenome string = "Júnior";
		idade int = 25;
	)
	fmt.Println("Olá "+ nome + " " + sobrenome);
	fmt.Println(idade);
}

