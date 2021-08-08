package memoryDataProvider

import (
  . "github.com/brickshot/roadtrip/internal/server/types"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("Data", func() {
  BeforeEach(func() {
    resetCars()
  })
  Describe("Creating a car", func() {
    Context("when input is valid", func() {
      It("should store the car", func() {
        car, err := NewCar(Car{Name: "Saab"})
        Expect(err).To(BeNil())
        Expect(cars[car.UUID]).To(Equal(car))
        c, err := GetCar(car.UUID)
        Expect(err).To(BeNil())
        Expect(cars[car.UUID].UUID).To(Equal(c.UUID))
        Expect(cars[car.UUID].Name).To(Equal(c.Name))
      })
    })
  })
  Describe("Getting a car", func() {
    Context("which exists", func() {
      It("should return car", func() {
        newCar, err := NewCar(Car{Name: "Maserati"})
        Expect(err).To(BeNil())
        getCar, err := GetCar(newCar.UUID)
        Expect(err).To(BeNil())
        Expect(getCar).To(Equal(newCar))
      })
    })
  })
})
