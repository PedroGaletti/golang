package loops

import (
	"fmt"
	"math/rand"
	"time"
)

func Loops() {
	// for initialization; condition;
	// postStatement {BODY}
	for x := 1; x <= 5; x++ {
		fmt.Println(x)
	}

	for x := 5; x >= 1; x-- {
		fmt.Println(x)
	}

	fx := 0
	for fx < 5 {
		fmt.Println(fx)
		fx++
	}

	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
}
