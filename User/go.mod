module MusicRecomend/User

go 1.13

require MusicRecomend/proto v0.0.0 // indirect

replace (
	MusicRecomend/proto => ../proto
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)
