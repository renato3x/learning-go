package main

import "fmt"

func buy(work1, work2 bool) (bool, bool, bool) {
  buyTv50 := work1 && work2
  buyTv32 := work1 != work2
  buyIceCream := work1 || work2

  return buyTv50, buyTv32, buyIceCream
}

func main() {
  tv50, tv32, iceCream := buy(true, true)

  fmt.Printf("Tv 50: %t, TV 32: %t, Ice cream: %t, Healthy: %t\n", tv50, tv32, iceCream, !iceCream)
}
