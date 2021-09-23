package formas

func TestArea(t *testing.T){
	//TDD - Test Driven Development
	//o t.Run() representa 2 funções que vão representar os métodos que estão sendo testados
	t.Run("Retangulo", func(t *testing.T){
		ret := Retangulo{10, 12}
		areaEsperada := float64(120)
		areaRecebida := ret.Area()

		if areaEsperada != areaRecebida {
			t.Errorf("A área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}

	})

	t.Run("Ciculo", func(t *testing.T){
		circ := Circulo{10}

		areaEsperada := float64(math.Pi * 100)
		areaRecebida := circ.Area()

		if areaEsperada != areaRecebida {
			t.Fatalf("A área recebida %f é diferente da esperada %f", areaRecebida, areaEsperada)
		}

	})
}