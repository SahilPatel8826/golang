package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	ID     int
	Status string
	mu     sync.Mutex
}

var (
	totalUpdates int
	updateMutex  sync.Mutex
)

func main() {
	// orders := generateOrders(20)
	var wg sync.WaitGroup
	wg.Add(3)
	// go func() {
	// 	defer wg.Done()
	// 	processOrders(orders)
	// }()
	orderChan := make(chan *Order)
	go func() {
		defer wg.Done()
		for _, order := range generateOrders(20) {
			orderChan <- order
		}
		close(orderChan)
		fmt.Println("done generating orders")
	}()

	go processOrders(orderChan, &wg)
	wg.Wait()
	// reportOrderStatus(orders)
	fmt.Println("All operations completed.Existing.")
	// fmt.Println(totalUpdates)
}

func updateOrderStatuses(order *Order) {
	order.mu.Lock()
	time.Sleep(
		time.Duration(rand.Intn(500)) *
			time.Millisecond,
	)
	status := []string{
		"Processing", "Shipped", "Delivered",
	}[rand.Intn(3)]
	order.Status = status
	fmt.Printf(
		"Updated order %d status: %s\n",
		order.ID, status,
	)
	order.mu.Unlock()
	updateMutex.Lock()
	defer updateMutex.Unlock()
	currentUpdates := totalUpdates
	time.Sleep(5 * time.Millisecond)
	totalUpdates = currentUpdates + 1

}

func processOrders(orderChan <-chan *Order,
	wg *sync.WaitGroup) {
	defer wg.Done()
	for order := range orderChan {
		time.Sleep(
			time.Duration(rand.Intn(500)) *
				time.Millisecond,
		)
		fmt.Println("Processing order %d\n", order.ID)
	}
}

func generateOrders(count int) []*Order {
	orders := make([]*Order, count)
	for i := 0; i < count; i++ {
		orders[i] = &Order{ID: i + 1,
			Status: "Pending"}
	}
	return orders
}

func reportOrderStatus(orders []*Order) {

	fmt.Println("\n--- Order Status Report ---")
	for _, order := range orders {
		fmt.Printf(
			"Order %d: %s\n",
			order.ID, order.Status,
		)
	}
	fmt.Println("------------------\n")
}
