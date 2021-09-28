package main

import (
	"fmt"
	"runtime"
)

func main() {
	//Go é uma linguagem concorrente - intercala processos ao mesmo tempo, oq pode ocorrer em single core
	//Paralelismo - Executa o código simultaneamento em processadores físicos diferetes (multicore)

	fmt.Println(runtime.NumCPU())

}