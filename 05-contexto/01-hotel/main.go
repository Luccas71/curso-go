package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	//inicia context em branco
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {

	select {
	// Done -> quando o case der match com o que foi estabelecido no context
	case <-ctx.Done():
		fmt.Println("Hotel booking cancelled. Timeout reached")
		return
	// caso passe antes do time do ctx
	case <-time.After(1 * time.Second):
		fmt.Println("Hotel booked.")
	}
}
