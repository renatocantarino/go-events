package main

import (
	"fmt"

	"github.com/renatocantarino/events/pkg/mq"
)

func main() {

	fmt.Println("Inicio do producer")
	ch, err := mq.OpenChannel()
	if err != nil {
		fmt.Println("erro ao abrir o canal")
		panic(err)
	}

	defer ch.Close()

	for i := 0; i < 5; i++ {
		mq.Publish(ch, fmt.Sprintf("%d - from go producer", i))
	}

}
