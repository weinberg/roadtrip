package memoryData

import (
  . "github.com/brickshot/roadtrip/internal/server"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("Data", func() {
  var p MemoryProvider = MemoryProvider{}
  BeforeEach(func() {
    resetCars()
  })
  Describe("Creating a car", func() {
    Context("when input is valid", func() {
      It("should store the car", func() {
        car, err := p.NewCar(Car{Name: "Saab"})
        Expect(err).To(BeNil())
        Expect(cars[car.UUID]).To(Equal(car))
        c, err := p.GetCar(car.UUID)
        Expect(err).To(BeNil())
        Expect(cars[car.UUID].UUID).To(Equal(c.UUID))
        Expect(cars[car.UUID].Name).To(Equal(c.Name))
      })
    })
  })
  Describe("Getting a car", func() {
    Context("which exists", func() {
      It("should return car", func() {
        newCar, err := p.NewCar(Car{Name: "Maserati"})
        Expect(err).To(BeNil())
        getCar, err := p.GetCar(newCar.UUID)
        Expect(err).To(BeNil())
        Expect(getCar).To(Equal(newCar))
      })
    })
  })
})
