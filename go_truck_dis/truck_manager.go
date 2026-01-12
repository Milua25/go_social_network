package main

import (
	"errors"
)

var ErrTruckNotFound = errors.New("truck not found")

type FleetManager interface {
	AddTruck(id string, cargo int) error
	GetTruck(id string) (*Truck, error)
	RemoveTruck(id string) error
	UpdateTruckCargo(id string, cargo int) error
}

type Truck struct {
	ID    string
	Cargo int
}

type truckManager struct {
	trucks map[string]*Truck
}

func NewTruckManager() truckManager {
	return truckManager{
		trucks: make(map[string]*Truck),
	}
}

func (manager *truckManager) AddTruck(id string, cargo int) error {
	manager.trucks[id] = &Truck{ID: id, Cargo: cargo}
	return nil
}

func (manager *truckManager) RemoveTruck(id string) error {
	delete(manager.trucks, id)
	return nil
}

func (manager *truckManager) GetTruck(id string) (*Truck, error) {
	truck, ok := manager.trucks[id]
	if !ok {
		return nil, ErrTruckNotFound
	}
	return truck, nil
}
func (manager *truckManager) UpdateTruckCargo(id string, cargo int) error {
	manager.trucks[id].Cargo = cargo
	return nil
}
