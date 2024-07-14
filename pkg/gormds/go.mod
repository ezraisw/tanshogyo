module github.com/ezraisw/tanshogyo/pkg/gormds

go 1.20

require (
	github.com/ezraisw/tanshogyo/pkg/common v0.0.0-00010101000000-000000000000
	github.com/rs/zerolog v1.29.1
	gorm.io/driver/mysql v1.5.1
	gorm.io/gorm v1.25.2
)

require (
	github.com/go-sql-driver/mysql v1.7.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	golang.org/x/sys v0.10.0 // indirect
)

replace github.com/ezraisw/tanshogyo/pkg/common => ../common
