module github.com/WeCanRun/gin-blog

go 1.15

replace (
	github.com/WeCanRun/gin-blog/conf => ./conf
	github.com/WeCanRun/gin-blog/middleware => ./middleware
	github.com/WeCanRun/gin-blog/models => ./models
	github.com/WeCanRun/gin-blog/pkg/e => ./pkg/e
	github.com/WeCanRun/gin-blog/pkg/logging => ./pkg/logging
	github.com/WeCanRun/gin-blog/pkg/setting => ./pkg/setting
	github.com/WeCanRun/gin-blog/pkg/util => ./pkg/util
	github.com/WeCanRun/gin-blog/pkg/file => ./pkg/file
	github.com/WeCanRun/gin-blog/routers => ./routers
)

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.62.0
	github.com/go-playground/assert/v2 v2.0.1
	github.com/jinzhu/gorm v1.9.16
	github.com/robfig/cron v1.2.0
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/unknwon/com v1.0.1
	gopkg.in/ini.v1 v1.62.0 // indirect
)

replace (
	golang.org/x/net => github.com/golang/net v0.0.0-20200324143707-d3edc9973b7e
	golang.org/x/sys => github.com/golang/sys v0.0.0-20200323222414-85ca7c5b95cd
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190611222205-d73e1c7e250b
)
