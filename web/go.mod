module github.com/oceango/web

go 1.13

require github.com/oceango/router v1.0.0

require (
	github.com/jinzhu/gorm v1.9.12 // indirect
	github.com/oceango/di v1.0.0
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.6.2 // indirect
)

replace github.com/oceango/router => ../router

replace github.com/oceango/di => ../di
