package snippet

import (
	"github.com/rs/zerolog/log"
)

func DealSliceByStep() {
	tarList := make([]int, 1000)

	step := 3
	total := len(tarList)
	for indexNow := 0; indexNow < total; {
		next := indexNow + step
		if next > total {
			next = total
		}
		// Processing, the
		log.Info().Msgf("current progress is [%d-%d] / %d", indexNow, next, total)
		DoList(tarList[indexNow:next])
		indexNow = next
	}
}

func DoList(ins []int) {
	for k := range ins {
		ins[k] = 5
	}
}
