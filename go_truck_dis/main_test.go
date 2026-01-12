// package main

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestMain(t *testing.T) {
// 	t.Run("processTruck", func(t *testing.T) {

// 		t.Run("should load and unload a truck", func(t *testing.T) {
// 			trucks := []*GasTruck{
// 				{id: "truck-01", cargo: 5}, {id: "truck-02", cargo: 10}, {id: "truck-03", cargo: 15}}

// 			eTrucks := []*ElectricTruck{
// 				{id: "electric-truck-01", cargo: 5},
// 			}

// 			for _, truck := range trucks {

// 				err := processTruck(truck)
// 				if err != nil {
// 					t.Fatalf("Error processing truck: %s", err)
// 				}
// 			}
// 			assert.Equal(t, trucks[0].cargo, 6)
// 			assert.Equal(t, trucks[1].cargo, 11)
// 			assert.Equal(t, trucks[2].cargo, 16)

// 			for _, truck := range eTrucks {

// 				err := processTruck(truck)
// 				if err != nil {
// 					t.Fatalf("Error processing electric truck: %s", err)
// 				}

// 				assert.Equal(t, eTrucks[0].cargo, 6)
// 			}
// 		})
// 	})
// }

// // func TestProcessTruck(t *testing.T) {
// //  processTruck()
// // }
