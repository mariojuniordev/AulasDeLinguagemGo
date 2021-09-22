package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	//Mapas são estruturas equivalentes aos dicionários
	//Syntax:
	//<nome do mapa> := make(map[<tipo do dado que representa a chave do mapa>] <tipo do valor item do mapa>)
	mapaCursos := make(map[string]int);
	
	//<nome do scanner> := bufio.NewScanner(<local de onde serão lidos os dados>)
	scanner := bufio.NewScanner(os.Stdin);

	curso := "";
	
	//Loop for (while) usado para gerar os dados do mapa
	for curso != "q" {
		var cargaHoraria int;
		fmt.Print("Digite um curso ou digite q para sair: ");
		//O método .Scan() irá escanear os dados do input do usuário 
		scanner.Scan(); //até q o usuário digite enter
		
		//Para encapsular os dados do nome dos cursos dentro da variável "curso"
		//utilizaremos o método .Text() do scanner que inicializamos anteriormente
		//O método .Text() retorna os dados escaneados 
		curso = scanner.Text() //retorna os dados

		if curso != "q" {
			fmt.Print("Digite a carga horária do curso: ");
			fmt.Scanf("%d", &cargaHoraria);
			//Encapsulamento de dados semelhante a objetos, por exemplo
			mapaCursos[curso] = cargaHoraria;
		}
		
	}

	fmt.Println("Cursos registrados: ");
	//Estrutura equivalente ao forEach()
	for nomeCurso, cargaHoraria := range mapaCursos {
		fmt.Printf("- %s: %d h \n", nomeCurso, cargaHoraria);
	}


	//Estrutura para exclusão de curso

	curso = "";
	for curso != "q" {
		fmt.Print("Digite o nome do curso a ser excluído ou digite q para cancelar: ");
		scanner.Scan();
		curso = scanner.Text();
		if curso != "q" {
			cargaHoraria, cursoExiste := mapaCursos[curso];
			if cursoExiste {
				//A função delete recebe como parâmetro a estrutura de onde ela vai deletar algo
				//e a chave do que ela vai deletar
				//Syntax: delete(<nome da estrutura - Array, Mapa...>, <nome da chave>)
				delete(mapaCursos, curso);
				fmt.Printf("O curso %s com %d h foi removido \n", curso, cargaHoraria);
			} else {
				fmt.Println("O curso digitado não existe.")
			}
		}
	}

	fmt.Println("Cursos registrados: ");
	//Estrutura equivalente ao forEach()
	for nomeCurso, cargaHoraria := range mapaCursos {
		fmt.Printf("- %s: %d h \n", nomeCurso, cargaHoraria);
	}

}