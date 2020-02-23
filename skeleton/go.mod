module github.com/oceango/skeleton

go 1.13

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/jinzhu/gorm v1.9.12
	github.com/oceango/middleware v1.0.0
	github.com/oceango/router v1.0.0
	github.com/oceango/web v1.0.0
	github.com/stretchr/testify v1.3.0 // indirect
	golang.org/x/crypto v0.0.0-20191227163750-53104e6ec876 // indirect
	golang.org/x/sys v0.0.0-20190626221950-04f50cda93cb // indirect
	golang.org/x/text v0.3.2 // indirect
	gopkg.in/ini.v1 v1.51.1 // indirect
)

replace github.com/oceango/middleware => ../middleware

replace github.com/oceango/router => ../router

replace github.com/oceango/web => ../web

replace github.com/oceango/di => ../di
