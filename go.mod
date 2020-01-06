module github.com/m-brady/deals

go 1.13

replace github.com/m-brady/deals/pkg/flyers/firestore => ./pkg/flyers/firestore

replace github.com/m-brady/deals/pkg/flyers/flipp => ./pkg/flyers/flipp

require (
	cloud.google.com/go v0.50.0 // indirect
	cloud.google.com/go/firestore v1.1.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/groupcache v0.0.0-20191227052852-215e87163ea7 // indirect
	github.com/google/go-cmp v0.3.1 // indirect
	github.com/hashicorp/golang-lru v0.5.3 // indirect
	github.com/jstemmer/go-junit-report v0.9.1 // indirect
	go.opencensus.io v0.22.2 // indirect
	golang.org/x/exp v0.0.0-20191227195350-da58074b4299 // indirect
	golang.org/x/net v0.0.0-20191209160850-c0dbc17a3553 // indirect
	golang.org/x/oauth2 v0.0.0-20191202225959-858c2ad4c8b6 // indirect
	golang.org/x/sys v0.0.0-20200103143344-a1369afcdac7 // indirect
	golang.org/x/tools v0.0.0-20200103221440-774c71fcf114 // indirect
	google.golang.org/api v0.15.0
	google.golang.org/appengine v1.6.5 // indirect
	google.golang.org/genproto v0.0.0-20191230161307-f3c370f40bfb // indirect
	google.golang.org/grpc v1.26.0 // indirect
)
