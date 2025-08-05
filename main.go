// Ponto de partida de um programa Go, um pacote obrigatório que define a função main() como ponto inicial para compilar e executar o código 
package main

// Importando alguns pacotes para usar durante o programa que são responsáveis pela entrada e saida de dados e para deletar ítens de um array
import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

// Função principal do nosso programa, tudo roda aqui, como em C
func main(){
	i := 0
	var numberTaskView int

	// Declarando um array em Go; caso fosse colocar valores nele, colocaria dentro das {}
	taskList := []string {}
	
	for i != 1{
		// Printando o menu com 5 opções usando o pacote 'fmt'
		fmt.Println("#-------------------#")
		fmt.Println("| - MENU ToDoList - |")
		fmt.Println("| -- 1. Add task -- |")
		fmt.Println("| -- 2. List     -- |")
		fmt.Println("| -- 3. Upd task -- |")
		fmt.Println("| -- 4. Del task -- |")
		fmt.Println("| -- 5. Close    -- |")
		fmt.Println("#-------------------#")
		
		fmt.Println("Escolha uma opção (Nº):")

		// Realizando a captura de um valor do teclado (limitado a uma única entrada antes de espaços)
		fmt.Scan(&numberTaskView)
		
		switch numberTaskView{ // Switch case do Go
		case 1:
			var taskTemp string

			// Cria um novo leitor
			reader := bufio.NewReader(os.Stdin)

			// Essa leitura descarta qualquer entrada pendente (basicamente, limpa o buffer do teclado)
			_, _ = reader.ReadString('\n')

			fmt.Println("Digite sua tarefa: ")

			// Lê a entrada do teclado (incluindo o '\n') e armazena na variável 'taskTemp'
			// O '_' ignora o valor de erro que a função retorna, pois não estamos tratando erros aqui
			taskTemp, _ = reader.ReadString('\n')

			// Insere a nova tarefa no array 
			taskList = append(taskList, taskTemp)
			fmt.Println("Tarefa registrada com sucesso!")
		case 2:
			i := 0
			fmt.Println("-- LISTAS DE TAREFAS --")
			for i < len(taskList){
			
				// Assim como no C, substitui o valor do %d para um valor inteiro. Por isso usei o Printf
				fmt.Printf("Tarefa %d: ", i+1)
				fmt.Print(taskList[i])
				i++
			}
			fmt.Println("")
			fmt.Println("")
		case 3:
			var indexTaskUpdate int
			var newTask string
			fmt.Println("Qual o Nº da tarefa que você quer editar?")
			fmt.Scan(&indexTaskUpdate)

			// Verifica se o index requerido não é menor ou maior que o tamanho do próprio Array de tarefas
			if indexTaskUpdate < 1 || indexTaskUpdate > len(taskList){
				fmt.Println("Indicie Inválido!")
			} else {
				reader := bufio.NewReader(os.Stdin)
				_, _ = reader.ReadString('\n')

				fmt.Println("Digite sua nova tarefa:")
				newTask, _ = reader.ReadString('\n')

				taskList[indexTaskUpdate-1] = newTask

				fmt.Println("Tarefa atualizada com sucesso!")
			}
		case 4:
			var indexTaskDelete int
		
			fmt.Println("Qual o Nº da tarefa que você quer deletar?")
			fmt.Scan(&indexTaskDelete)

			if indexTaskDelete < 1 || indexTaskDelete > len(taskList){
				fmt.Println("Indicie Inválido!")
			} else {
				taskList = slices.Delete(taskList, indexTaskDelete-1, indexTaskDelete)
				fmt.Println("Tarefa deletada com sucesso!")
			}
		case 5:
			// Assim que o usuário quiser encerrar, ele digita 5 e i agora passa a valer 1 e o loop para
			fmt.Println("Até a próxima! Tchau!")
			i = 1
		default:
			// Caso o usuário digite uma opção não descrita
			fmt.Println("Opção inválida!")
		}
	}
}