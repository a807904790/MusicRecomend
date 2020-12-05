module MusicRecomend/User

go 1.13

require (
	MusicRecomend/proto v0.0.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.1
	github.com/jinzhu/gorm v1.9.16
)

replace (
	MusicRecomend/proto => ../proto
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)
