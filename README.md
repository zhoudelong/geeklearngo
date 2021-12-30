# geeklearngo 作为极客时间提交代码的目录, 目录结构如下:

仓库结构目录说明:  周/代码, 依此类推

Week01:
		需求: go实现计算圆面积
		$ go run circle_area.go
  

Week02:
		需求:
			1：计算多个人的平均体脂
				实现完整的体脂计算器
				连续输入三人的姓名、性别、身高、体重、年龄信息
				计算每个人的 BMI、体脂率
				输出：
				每个人姓名、BMI、体脂率、建议
				总人数、平均体脂率
			
				$ go run bmi.go
		
			2: 判断两条线是否平行
				$ go run line_pra.go
Week03:
		需求:
			1: 	使用 github 上的 lib： github.com/armstrongli/go-bmi 完成体脂计算器。
				本地添加 module 的 replace，并在本地项目扩展  github.com/armstrongli/go-bmi 以支持 BMI、FatRate 的计算。
				使用 vendor 保证代码的完整性与可运行

			2: 
				2.1: 为体脂计算器编写单元测试并完善体脂计算器的验证逻辑。
				2.2: BMI 计算：
					录入正常身高、体重，确保计算结果符合预期；
					录入 0 或负数身高，返回错误；
					录入 0 或负数体重，返回错误。
				2.3: 体脂率计算：
					录入正常 BMI、年龄、性别，确保计算结果符合预期；
					录入非法 BMI、年龄、性别（0、负数、超过 150 的年龄、非男女的性别输入），返回错误；
					录入完整的性别、年龄、身高、体重，确保最终获得的健康建议符合预期。

		code:
		1.第一个作业题代码目录
    		- homework/Week03/fatRate

			作业地址: https://github.com/zhoudelong/geeklearngo/tree/main/homework/Week03/fatRate

			执行:
			go run .\main.go --Name zhoudelong --Sex "男" --Height 1.75 --Weight 73 --Age 27 

			Name: zhoudelong
			Sex: 男                                       
			Height: 1.75                                  
			Weight: 73                                    
			Age: 27                                       
			bmi:23.836734693877553,bfr:0.17804081632653063


		2.第二个作业题代码目录
			- homework/Week03/fatRateTest

			作业地址: 
					bmi: https://github.com/zhoudelong/geeklearngo/tree/main/homework/Week03/fatRateTest/bmi
					fatrate:  https://github.com/zhoudelong/geeklearngo/tree/main/homework/Week03/fatRateTest/bfa
			
			go run bmi_test.go

			go run fatrate_test.go