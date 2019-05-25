package main

import (
	"fmt"
	"math/big"
)

const mod = 1000000007

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for ti := 1; ti <= t; ti++ {
		var N, K, x1, y1, C, D, E1, E2, F int64
		fmt.Scanf("%d %d %d %d %d %d %d %d %d", &N, &K, &x1, &y1, &C, &D, &E1, &E2, &F)
		var totalMod, powerFactorSum int64 // (1^1 + .. + 1^K) + .. + (N^1 + .. + N^K)
		// numFactorSum := 0   // (A_1*N + A_2*(N-2) + .. + A_N*1)
		var prevX, prevY int64
		for i := int64(1); i <= N; i++ {
			var powerFactor int64 // i^1+ i^2 + ... + i^K
			if i == 1 {
				powerFactor = K
			} else {
				// powerFactor = i * int64(math.Pow(float64(i), float64(K))-1) / (i - 1) // Sometimes wrong because of overflow.
				// powerFactor = i * (quickpowMod(i, K, mod) - 1) / (i - 1) // Wrong: can't use modulo because of division and subtraction.
				// power := int64(1)
				// for k := int64(1); k <= K; k++ {
				// 	power = power * i % mod
				// 	powerFactor = (powerFactor + power) % mod
				// // }
				// superMod := (i - 1) * mod
				// powerFactor = i * ((quickpowMod(i, K, superMod) + superMod - 1) / (i - 1) % mod) // overflow
				// powerFactor = i * ((quickpowMod(i, K, superMod) + superMod - 1) / (i - 1) % mod) % mod
				// if (quickpowBigMod(i, K, superMod)+superMod-1)%(i-1) != 0 {
				// 	//fmt.Println("here1", (quickpowMod(i, K, (i-1))+i-2)%(i-1))
				// 	fmt.Println(i, K, (quickpowBigMod(i, K, superMod)+superMod-1)%(i-1))
				// 	fmt.Println("here2", superMod, (quickpowBigMod(i, K, superMod) + superMod - 1), i-1)
				// }
				powerFactor = i * ((quickpowMod(i, K, mod) - 1) * quickpowMod(i-1, mod-2, mod) % mod) % mod
			}
			powerFactorSum = (powerFactor + powerFactorSum) % mod
			currX, currY := prevX, prevY
			if i != 1 {
				currX = (C*prevX + D*prevY + E1) % F
				currY = (D*prevX + C*prevY + E2) % F
			} else {
				currX, currY = x1%F, y1%F
			}
			currA := (currX + currY) % F
			prevX, prevY = currX, currY
			numFactor := currA * (N - i + 1) % mod // A_i*(N-i+1) : A_1 * N, A_2 * (N-1), .., A_N * 1
			totalMod = (totalMod + numFactor*powerFactorSum) % mod
		}
		fmt.Printf("Case #%d: %d\n", ti, totalMod)
	}
}

func quickpowMod(x, n, mod int64) int64 {
	if n == 0 {
		return 1
	} else if n == 1 {
		return x
	}
	part := quickpowMod(x, n/2, mod)
	if n%2 == 0 {
		return part * part % mod
	}
	return part * part % mod * x % mod
}

func quickpow(x float64, n int64) float64 {
	if n == 0 {
		return 1
	} else if n == 1 {
		return x
	}
	part := quickpow(x, n/2)
	if n%2 == 0 {
		return part * part
	}
	return part * part * x
}

func quickpowBigMod(x int64, n, mod int64) int64 {
	if n == 0 {
		return 1
	} else if n == 1 {
		return x
	}
	part := quickpowBigMod(x, n/2, mod)
	bigInt := new(big.Int)
	bigInt.Mul(big.NewInt(part), big.NewInt(part))
	bigInt.Mod(bigInt, big.NewInt(mod))
	// fmt.Println("here", n, bigInt.Int64())
	if n%2 != 0 {
		bigInt.Mul(bigInt, big.NewInt(x))
		bigInt.Mod(bigInt, big.NewInt(mod))
		// fmt.Println("here", n, bigInt.Int64())
	}
	return bigInt.Int64()
}
