package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	n := flag.Int("n", 0, "Schwangerschaftswoche")
	f := flag.Int("f", 0, "Schwangerschaftswoche")
	flag.Parse()

	nowTime := time.Now()

	futureTime := nowTime.AddDate(0, 0, 7 * (*f - *n))

	fmt.Printf("Du bist in der %d Schwangerschaftswoche und möchtest wissen wann die %d Schwangerschaftswoche ist.\n", *n, *f)

	fmt.Printf("Heute ist der %s.\n", nowTime.Format("02.01.2006"))

	fmt.Printf("Die %d Schwangerschaftswoche ist am %s.\n", *f, futureTime.Format("02.01.2006"))

}
