module github.com/No3371/gitwatch

go 1.17

require (
	github.com/No3371/gitwatch/gitlog v0.0.0-00010101000000-000000000000
	github.com/diamondburned/arikawa/v3 v3.0.0-rc.4
	go.uber.org/zap v1.19.1
)

require (
	github.com/gorilla/schema v1.2.0 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/tsuyoshiwada/go-gitcmd v0.0.0-20180205145712-5f1f5f9475df // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac // indirect
)

replace github.com/No3371/gitwatch/gitlog => ./go-gitlog
