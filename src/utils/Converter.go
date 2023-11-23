package utils

import (
	"fmt"
	"math/big"
	"strconv"
)

func FloatToBigInt(val float64) *big.Int {
	bigval := new(big.Float)
	bigval.SetFloat64(val)

	coin := new(big.Float)
	coin.SetInt(big.NewInt(1000000000000000000))

	bigval.Mul(bigval, coin)

	result := new(big.Int)
	bigval.Int(result)

	return result
}

func StringToBigInt(val string) *big.Int {
	amount, err := strconv.ParseFloat(val, 64)
	fmt.Println(amount)
	if err != nil {
		fmt.Println("SetString: error")
		return nil
	}
	amountBigInt := FloatToBigInt(amount)
	fmt.Println(amountBigInt)
	if amountBigInt == nil {
		fmt.Println("SetString: error")
		return nil
	}
	return amountBigInt
}
