module Go-000/Week04

go 1.15

require (
	Go-000/Week04/server v1.0.0
	Go-000/Week04/config v1.0.0
)


replace (
	Go-000/Week04/server => ./server
	Go-000/Week04/config => ./config
)

