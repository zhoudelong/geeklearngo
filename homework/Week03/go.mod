module Week03

go 1.17

require github.com/armstrongli/go-bmi v0.0.1

//扩展github上的包到本地, 继续引用, 可以在本地新增功能
replace github.com/armstrongli/go-bmi => ./staging/go-bmi
