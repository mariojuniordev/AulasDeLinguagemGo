// TESTE DE UNIDADE - É UM TESTE QUE TESTA A MENOR PARTE DO SEU CÓDIGO, NESTE CASO, UMA FUNÇÃO

package enderecos_test

import (
	. "introducao-testes/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado string
}

//Por padrão, as funções de teste devem começar com TestxxxxXxXxXXxx
func TestTipoDeEndereco(t *testing.T){
	//Método que faz os testes rodarem em paralelo
	t.Parallel()
	
	//Slice com Endereços e Retornos esperados
	cenariosDeTeste := []cenarioDeTeste {
		{ "Rua ABC", "Rua" },
		{ "Avenida Paulista", "Avenida" },
		{ "Rodovia dos Imigrantes", "Rodovia" },
		{ "Praça das Rosas", "Tipo Inválido" },
		{ "Estrada Qualquer", "Estrada" },
		{ "RUA DOS BOBOS", "Rua" },
		{ "AVENIDA REBOUÇAS", "Avenida" },
		{ "", "Tipo Inválido" },
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
		tipoDeEnderecoRecebido, cenario.retornoEsperado)
		}
	}

}

func TestQualquer(t *testing.T){
	//É necessário colocar o x.Paralell() em todas as funções q vão ser rodadas em paralelo
	t.Parallel()

	if 1 > 2 {
		t.Errorf("Teste quebrou!")
	}
}