// package main

// import (
// 	"context"
// 	"errors"
// 	"fmt"
// 	"sync"
// 	"time"
// )

// var (
// 	ErrNotImplemented = errors.New("not implemented")
// 	ErrTruckNotFound  = errors.New("truck not found")
// )

// type GasTruck struct {
// 	id    string
// 	cargo int
// }

// type ElectricTruck struct {
// 	id      string
// 	cargo   int
// 	battery float64
// }

// type Truck interface {
// 	LoadCargo() error
// 	UnLoadCargo() error
// }

// type foo interface {
// }

// func (t *GasTruck) LoadCargo() error {
// 	t.cargo += 2
// 	return nil
// }

// func (t *GasTruck) UnLoadCargo() error {
// 	t.cargo -= 1
// 	return nil
// }

// func (t *ElectricTruck) LoadCargo() error {
// 	t.cargo += 2
// 	return nil
// }

// func (t *ElectricTruck) UnLoadCargo() error {
// 	t.cargo -= 1
// 	return nil
// }

// // processTruck Handles the loading and unloading of a truck
// func processTruck(ctx context.Context, t Truck) error {
// 	//  fmt.Printf("Processing truck: %s\n", truck.id)

// 	fmt.Println("Processing truck!!!")
// 	fmt.Println("Context value:", ctx.Value("UserId"))

// 	ctx, cancel := context.WithTimeout(ctx, time.Second*2)

// 	defer cancel()

// 	//simulate a delay
// 	// delay := time.Second * 30
// 	// select {
// 	// case <-ctx.Done():
// 	// 	fmt.Println("Error")
// 	// 	return ctx.Err()
// 	// case <-time.After(delay):
// 	// 	break
// 	// }

// 	err := t.LoadCargo()
// 	if err != nil {
// 		fmt.Errorf("Error loading truck: %w", err)
// 		return err
// 	}
// 	err = t.UnLoadCargo()
// 	if err != nil {
// 		fmt.Errorf("Error unloading truck: %w", err)
// 		return err
// 	}
// 	return nil
// }

// func main() {

// 	ctx := context.Background()

// 	ctx = context.WithValue(ctx, "UserId", 42)

// 	trucks := []*GasTruck{
// 		{id: "truck-01"}, {id: "truck-02"}, {id: "truck-03"}}

// 	eTrucks := []*ElectricTruck{
// 		{id: "electric-truck-01"},
// 	}

// 	var electricWg, gasWg sync.WaitGroup
// 	errorsChan := make(chan error, len(eTrucks)+len(trucks))

// 	person := map[string]interface{}{}

// 	person["name"] = "Tiago"
// 	person["Age"] = 42

// 	electricWg.Add(len(eTrucks))

// 	gasWg.Add(len(trucks))

// 	defer close(errorsChan)
// 	for _, truck := range trucks {
// 		go func(t *GasTruck) {
// 			fmt.Printf("Truck %s arrived.\n", truck.id)
// 			defer gasWg.Done()
// 			err := processTruck(ctx, t)
// 			fmt.Printf("Cargo Size: %v\n", t.cargo)
// 			fmt.Printf("Truck %s departed.\n", t.id)
// 			errorsChan <- err
// 		}(truck)

// 		// if err != nil {
// 		// 	log.Fatalf("Error processing truck: %s", err)
// 		// }
// 	}

// 	for _, truck := range eTrucks {

// 		go func(t *ElectricTruck) {
// 			fmt.Printf("Electric Truck %s arrived.\n", truck.id)
// 			defer electricWg.Done()
// 			err := processTruck(ctx, t)
// 			fmt.Printf("Cargo Size: %v\n", t.cargo)
// 			fmt.Printf("Truck %s departed.\n", t.id)
// 			errorsChan <- err
// 		}(truck)

// 		// if err != nil {
// 		// 	log.Fatalf("Error processing electric truck: %s", err)
// 		// }

// 	}

// 	num, ok := person["Age"].(int)
// 	if ok {
// 		fmt.Println(ok)
// 		fmt.Println(num)
// 	}
// 	gasWg.Wait()

// 	electricWg.Wait()

// 	var errs []error

// 	for err := range errorsChan {
// 		errs = append(errs, err)
// 	}

// 	fmt.Println(errs)

// 	// select {
// 	// case err := <-errorsChan:
// 	// 	log.Fatalln(err)
// 	// default:
// 	// 	log.Println("No error")
// 	// }
// }
