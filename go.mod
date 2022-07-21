module github.com/WeCanRun/gin-blog

go 1.15

replace (
	github.com/WeCanRun/gin-blog/conf => ./conf
	github.com/WeCanRun/gin-blog/global/constants => ./global/constants
	github.com/WeCanRun/gin-blog/global/errcode => ./global/errcode
	github.com/WeCanRun/gin-blog/internal/middleware => ./internal/middleware
	github.com/WeCanRun/gin-blog/internal/model => ./internal/model
	github.com/WeCanRun/gin-blog/internal/router => ./internal/router
	github.com/WeCanRun/gin-blog/internal/service => ./internal/service
	github.com/WeCanRun/gin-blog/internal/service/cache_service => ./internal/service/cache_service
	github.com/WeCanRun/gin-blog/pkg/file => ./pkg/file
	github.com/WeCanRun/gin-blog/pkg/logging => ./pkg/logging
	github.com/WeCanRun/gin-blog/pkg/setting => ./pkg/setting
	github.com/WeCanRun/gin-blog/pkg/util => ./pkg/util
)

require (
	github.com/360EntSecGroup-Skylar/excelize/v2 v2.3.1
	github.com/boombuler/barcode v1.0.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fsnotify/fsnotify v1.5.4
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/assert/v2 v2.0.1
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0
	github.com/gomodule/redigo v1.8.2
	github.com/jinzhu/gorm v1.9.16
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/spf13/viper v1.12.0
	github.com/tealeg/xlsx v1.0.5
	github.com/unknwon/com v1.0.1
)

replace (
	golang.org/x/net => github.com/golang/net v0.0.0-20200324143707-d3edc9973b7e
	golang.org/x/sys => github.com/golang/sys v0.0.0-20200323222414-85ca7c5b95cd
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190611222205-d73e1c7e250b
)
