package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main(){
	i := 0
	var numberTaskView int
	taskList := []string {}
	
	for i != 1{
		fmt.Println("#-------------------#")
		fmt.Println("| - MENU ToDoList - |")
		fmt.Println("| -- 1. Add task -- |")
		fmt.Println("| -- 2. List     -- |")
		fmt.Println("| -- 3. Upd task -- |")
		fmt.Println("| -- 4. Del task -- |")
		fmt.Println("| -- 5. Close    -- |")
		fmt.Println("#-------------------#")
		
		fmt.Println("Escolha uma opção (Nº):")
		fmt.Scan(&numberTaskView)
		
		switch numberTaskView{
		case 1:
			var taskTemp string

			reader := bufio.NewReader(os.Stdin)
			_, _ = reader.ReadString('\n')

			fmt.Println("Digite sua tarefa: ")
			taskTemp, _ = reader.ReadString('\n')

			taskList = append(taskList, taskTemp)
			fmt.Println("Tarefa registrada com sucesso!")
		case 2:
			i := 0
			fmt.Println("-- LISTAS DE TAREFAS --")
			for i < len(taskList){
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
			i = 1
		default:
			fmt.Println("Opção inválida!")
		}
	}
}