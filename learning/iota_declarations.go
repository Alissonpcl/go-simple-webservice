package learning

import "fmt"

const (
	first    = iota //will be 0
	second   = iota //will be 1
	autoIota        // get the next iota value implicitly

	added      = iota + 6  //3 + 6 = 9
	multiplied = 2 << iota // iota = 4 so result is -> (2 x 2 x 2 x 2 x 2) = 32
)

const (
	third = iota // iota resets by const block so here is - again
)

func declareIotas() {
	fmt.Println(first, second, autoIota, added, multiplied)
	fmt.Println(third)
}
