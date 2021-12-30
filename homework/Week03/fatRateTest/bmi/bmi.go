package bmi

import (
	"fmt"
)

func CalcBMI(weightKG, heightM float64) (bmi float64, err error) {
	if weightKG < 0 {
		err = fmt.Errorf("体重不能为负数.")
		// errors.New("Name is null...")
		return
	}

	if weightKG == 0 {
		err = fmt.Errorf("体重不能为0.")
		return
	}

	if heightM < 0 {
		err = fmt.Errorf("身高不能为负数.")
		return
	}
	if heightM == 0 {
		err = fmt.Errorf("身高不能为0.")
		return
	}
	if heightM > 2.5 {
		err = fmt.Errorf("身高不能大于2.5m.")
		return
	}

	bmi = weightKG / (heightM * heightM)

	if bmi == 0 {
		err = fmt.Errorf("bmi计算结果不能为0")
		return
	} else if bmi < 0 {
		err = fmt.Errorf("bmi计算结果不能为负数。")
		return
	}

	return
}
