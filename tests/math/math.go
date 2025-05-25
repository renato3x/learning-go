package math

import (
	"fmt"
	"strconv"
)

func Average(values ...float64) float64 {
  sum := 0.0
  for _, value := range values {
    sum += value
  }

  average := sum / float64(len(values))
  roundedAverage, _ := strconv.ParseFloat((fmt.Sprintf("%.2f", average)), 64)

  return roundedAverage
}
