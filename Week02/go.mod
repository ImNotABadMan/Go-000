module Go-000/Week02

go 1.15

require (
	Go-000/Week02/dao v1.0.0
	Go-000/Week02/mysql v1.0.0
)

replace (
	Go-000/Week02/dao => ./dao
	Go-000/Week02/mysql => ./mysql
)
