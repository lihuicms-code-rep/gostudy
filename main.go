package main

import "fmt"

func main ()  {
    a1 := SelectFirstArea(0.66)
    fmt.Printf("a1=%d\n", a1)
    a2 := SelectSecondArea(a1, 0.02)
    fmt.Printf("a2=%d\n", a2)
}


func SelectFirstArea(prob float64) int32 {
	bankerProb := 0.47
	playerProb := 0.47

	if prob >= 0 && prob < bankerProb {
		return 1
	} else if prob >= bankerProb && prob < bankerProb+playerProb {
		return 2
	} else if prob >= bankerProb+playerProb && prob < 1 {
		return 3
	}

	return -1
}

//根据第一层押注区域和当前随机概率值挑选第二层押注区域
func SelectSecondArea(arena1 int32, prob float64) int32 {
	switch arena1 {
	case 1:
		bankerPair := 0.1
		playerPair := 0.1
		zlb := 0.8
		return switchArenaByProb(bankerPair, playerPair, zlb, prob, arena1)
	case 2:
		bankerPair := 0.1
		playerPair := 0.1
		xlb := 0.8
		return switchArenaByProb(bankerPair, playerPair, xlb, prob, arena1)
	case 3:
		bankerPair := 0.5
		playerPair := 0.5
		return switchArenaByProb(bankerPair, playerPair, float64(0), prob, arena1)
	}

	return -1
}

func switchArenaByProb(prob1, prob2, prob3, prob float64, arnea1 int32) int32 {
	if prob >= 0 && prob <= prob1 {
		return 4
	} else if prob > prob1 && prob <= prob1+prob2 {
		return 5
	} else if prob > prob1+prob2 && prob <= prob1+prob2+prob3 {
		if arnea1 == 1 {
			return 6
		} else if arnea1 == 2 {
			return 7
		}
	}

	return -1
}