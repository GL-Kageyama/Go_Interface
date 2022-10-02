package dog

import (
    "../cat"
)

type Dog struct {}

func NewDog() *Dog {
    return new(Dog)
}

func (d *Dog) GetDog() string {
    cat := &cat.Cat{ CatObj: NewDogRoar() }
    // GetCat -> DogInterface -> GetDogRoar
    return cat.GetCat()
}
