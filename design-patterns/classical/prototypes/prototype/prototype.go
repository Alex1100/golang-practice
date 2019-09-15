package prototype


/*
 *
 * Prototype Pattern
 *
 * It is a powerful tool to build caches and default
 * objects
 *
 * Some patterns can overlap a bit, but they
 * have small differences
 */

import (
  "errors"
  "fmt"
)


type ShirtCloner interface {
  GetClone(m int) (ItemInfoGetter, error)
}

const (
  White = 1
  Black = 2
  Blue  = 3
)

type ShirtsCache struct{}
type ItemInfoGetter interface {
  GetInfo() string
}

type ShirtColor byte
type Shirt struct {
  Price float32
  SKU string
  Color ShirtColor
}

func (s *Shirt) GetInfo() string {
  return fmt.Sprintf("Shirt with SKU '%s' and Color id %d that costs %f\n", s.SKU, s.Color, s.Price)
}

func GetShirtsCloner() ShirtCloner {
  var newShirtCloner *ShirtsCache = &ShirtsCache{}
  return newShirtCloner
}

var whitePrototype *Shirt = &Shirt{
  Price: 15.00,
  SKU: "1234",
  Color: White,
}

var blackPrototype *Shirt = &Shirt{
  Price: 25.00,
  SKU: "1235",
  Color: Black,
}

var bluePrototype *Shirt = &Shirt{
  Price: 20.00,
  SKU: "1236",
  Color: Blue,
}

func (i *Shirt) GetPrice() float32 {
  return i.Price
}


func (s *ShirtsCache) GetClone(m int) (ItemInfoGetter, error) {
  switch m {
  case White:
    newItem := *whitePrototype
    return &newItem, nil
  case Black:
    newItem := *blackPrototype
    return &newItem, nil
  case Blue:
    newItem := *bluePrototype
    return &newItem, nil
  default:
    return nil, errors.New("Not implemented yet")
  }
}