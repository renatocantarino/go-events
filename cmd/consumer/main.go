package main

import (
	"fmt"

	"github.com/rabbitmq/amqp091-go"
	"github.com/renatocantarino/events/pkg/mq"
)

func main() {

	fmt.Println("Inicio do consumer")
	ch, err := mq.OpenChannel()
	if err != nil {
		fmt.Println("erro ao abrir o canal")
		panic(err)
	}

	defer ch.Close()

	msgs := make(chan amqp091.Delivery)
	go mq.Consume(ch, msgs, "myqueue")

	//por boa pratica o ack deve ser manual, o processo pode travar um passo antes
	//e perdemos a mensagem
	for item := range msgs {
		fmt.Println("Lendo msg")
		fmt.Println(string(item.Body))
		item.Ack(false)
	}

}
