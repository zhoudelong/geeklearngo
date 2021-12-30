package bmi

import "testing"

// TestCalcBMI_Case1: 录入正常身高、体重，确保计算结果符合预期；
func TestCalcBMI_Case1(t *testing.T) {
	// 实际计算得出 体重70kg,身高1.75m bmi计算结果为 23.836734693877553
	weight, height := 73.0, 1.75
	bmi, err := CalcBMI(weight, height)

	if err != nil {
		//  t.Error/t.Errorf 输出可能的错误, 遇到错误不停, 直到执行全部完
		t.Errorf("您输入的身高%v,或者体重%v 异常,请重新输入", height, weight)
	}

	if bmi != 23.836734693877553 {
		t.Errorf("预期结果是22.9,但计算结果是%v,不符合预期。\n", bmi)
	}
}

// TestCalcBMI_Case2 录入 0 或负数身高，返回错误；
func TestCalcBMI_Case2(t *testing.T) {
	// 录入0身高,返回错误.
	{
		_, err := CalcBMI(70, 0)
		if err != nil {
			t.Errorf("身高数值输入异常: %v\n", err)
		}
	}

	//录入负数身高，返回错误。
	{
		_, err := CalcBMI(70, -2)
		if err != nil {
			t.Errorf("身高数值输入异常: %v\n", err)
		}
	}

}

// TestCalcBMI_Case3 录入 0 或负数体重，返回错误。
func TestCalcBMI_Case3(t *testing.T) {
	{
		_, err := CalcBMI(0, 1.75)
		if err != nil {
			t.Errorf("体重数值输入异常: %v\n", err)
		}
	}

	//录入负数体重，返回错误。
	{
		_, err := CalcBMI(-2, 1.75)
		if err != nil {
			t.Errorf("体重数值输入异常: %v\n", err)
		}
	}
}
