package memoryDataProvider

import (
  . "github.com/brickshot/roadtrip/internal/server/types"
  "github.com/google/uuid"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("Data", func() {
  Describe("Storing a character", func() {
    var (
      character Character
    )
    BeforeEach(func() {
      resetCharacters()
      character = Character{
        Name: "Josh Weinberg",
      }
    })
    Context("When the UUID is valid", func() {
      BeforeEach(func() {
        character.UUID = uuid.NewString()
      })
      It("should store the character", func() {
        StoreCharacter(character)
        c, _ := GetCharacter(character.UUID)
        Expect(characters[character.UUID]).To(Equal(c))
      })
    })
    Context("When the UUID is invalid", func() {
      BeforeEach(func() {
        character.UUID = "XYZZY"
      })
      It("should error", func() {
        var err error
        err = StoreCharacter(character)
        Expect(err).NotTo(BeNil())
        Expect(err.Error()).To(Equal("Invalid UUID"))

        _, err = GetCharacter(character.UUID)
        Expect(err).NotTo(BeNil())
        Expect(err.Error()).To(Equal("Not found"))
      })
    })
  })

  Describe("Getting a character", func() {
    var (
      character Character
    )
    BeforeEach(func() {
      resetCharacters()
      character = Character{
        Name: "Josh Weinberg",
        UUID: uuid.NewString(),
      }
    })
    Context("When character exists", func() {
      It("should return the character", func() {
        StoreCharacter(character)
        Expect(characters[character.UUID]).To(Equal(character))
      })
    })
    Context("When character does not exists", func() {
      It("should error", func() {
        StoreCharacter(character)
        _, err := GetCharacter(uuid.NewString())
        Expect(err.Error()).To(Equal("Not found"))
      })
    })
  })

  Describe("Setting a character's car", func() {
    var (
      character Character
    )
    BeforeEach(func() {
      resetCharacters()
      character = Character{
        Name: "Josh Weinberg",
        UUID: uuid.NewString(),
      }
      StoreCharacter(character)
    })
    Context("When input is valid", func() {
      It("should return the car on the character", func() {
        car, _ := NewCar(Car{Name: "Ford"})
        char, err := SetCar(character.UUID, car.UUID)
        Expect(err).To(BeNil())
        Expect(char.Car.UUID).To(Equal(car.UUID))
      })
      It("should allow unsetting the car", func() {
        car, _ := NewCar(Car{Name: "Ford"})
        c,_ := SetCar(character.UUID, car.UUID)
        Expect(c.Car.UUID).To(Equal(car.UUID))
        char, err := SetCar(character.UUID, "")
        Expect(err).To(BeNil())
        Expect(char.Car).To(BeNil())
      })
    })
  })
})
