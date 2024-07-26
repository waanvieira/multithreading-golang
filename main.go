package main

import (
	"fmt"
	"time"
)

// Receber um chanel como recebedor de daddos é <-
// Nesse caso o nosso worker é o processador de dados, então a seta <- no começo do chan indica que aqui processa os dados, no caso esvazia o canal
// -> a seta indicando par a direita significa que o canal envia os dados para o canal
func worker(workerId int, data <-chan int) {
	for x := range data {
		fmt.Println("Worker %d received %d/n", workerId, x)
		time.Sleep(time.Second)
	}
}

// O main é a 1° thread que o go já cria
func main() {
	// Declaramos o nosso canal que vai ser executado
	data := make(chan int)
	// Declaramos a quantidade de threads que vamos ter
	QtdWorkers := 10

	// Aqui é algo comum de se ver, criamos um for para criar a quantidade de workers que definimos na quantidade
	// Inicialia os workers, aqui iniciamos no caso 10 workers para serem usados para o processamento, para receber a carga
	for i := 0; i < QtdWorkers; i++ {
		// Aqui obrigatoriamente temos que indicar a palavra reservada 'go' para criar uma nova thread, ou seja, cada vez que colocamos a palavra go criamos
		// Uma nova thread para fazer o processamento de dados, se não indicarmos o go vamos receber um deadlock
		go worker(i, data)
	}

	// Iniciando o worker
	// go worker(1, data)
	// Criando mais uma thred para o nosso processamento
	// go worker(2, data)

	// Aqui seria um exemplo de receber requisições web
	// Aqui começamos a receber as requisições que vamos processar
	// Aqui no caso temos 10 workers para receber essa carga e o próprio GO vai se encarregar de fazer o gerenciamento desses workers e balancear as requisições
	for i := 0; i < 100000; i++ {
		// Aqui estamos enchendo o nosso canal
		data <- i
	}
}

// ###################### Buffer
// É comum fazer da seguinte forma, uma thread especifica manda uma mensagem para um canal, depois outra rotina, outra thread vai ler esse canal
// só depois que essa thread ler esse canal que pode ser enviado mais dados, no caso mais 1 dado para esse canal
// Aqui a forma de enviar mais de 1 dado para o mesmo canal sem que a outra thread já tenha lido
// MAS É IMPORTANTE EVITAR ESSA PRATICA DE buffer, só usar fazendo um benchmarq para ver se funciona, é obritório o computador ter recurso suficiente para
// Conseguir ler as mensangens, se não tiver só vai consumir recursos da máquina
func adicionandMaisDeUmaMensagemNoMesmoCanal() {
	// Fazendo desse modo vai dar deadloack, porque quando for enviar o "World" o canal ainda está cheio, então não vai funcionar
	chErro := make(chan string)
	chErro <- "hello"
	chErro <- "World"

	println(<-chErro)
	println(<-chErro)
	// Para funcionar basta passar um parametro depois do tipo do channel que funciona
	chQueFunciona := make(chan string, 2)
	chQueFunciona <- "hello"
	chQueFunciona <- "World"

	println(<-chQueFunciona)
	println(<-chQueFunciona)
}
