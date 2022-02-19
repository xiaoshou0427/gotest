module learn.go

go 1.16

require (
	github.com/armstrongli/go-bmi v0.0.0-00010101000000-000000000000
	//github.com/armstrongli/go-bmi v0.0.1
	github.com/spf13/cobra v1.4.0
	learn.go.tools v0.0.0-00010101000000-000000000000
)

replace (
	github.com/armstrongli/go-bmi => github.com/xiaoshou0427/go-bmi v0.0.0-20210904081709-c4b711282417
	//github.com/armstrongli/go-bmi => ./staging/src/github.com/armstrongli/go-bmi
	github.com/spf13/cobra => github.com/spf13/cobra v1.3.0
	learn.go.tools => ../learn.go.tools
)
