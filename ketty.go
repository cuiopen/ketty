package ketty

import (
	U "github.com/yyzybb537/ketty/url"
	P "github.com/yyzybb537/ketty/protocol"
	B "github.com/yyzybb537/ketty/balancer"
	A "github.com/yyzybb537/ketty/aop"
	_ "github.com/yyzybb537/ketty/protocol/grpc"
	_ "github.com/yyzybb537/ketty/protocol/http"
)

type Dummy interface{}
type Client P.Client
type Server P.Server

// @主要组件
//type Protocol interface {}
//type Balancer interface {}
//type Client interface {}
//type Server interface {}

// @sUrl:   protocol://ip[:port][,ip[:port]]/path
// E.g:
//    http://127.0.0.1:8030/path
//    https://127.0.0.1:8030
//    tcp://127.0.0.1:8030
//    grpc://127.0.0.1:8030,127.0.0.1:8031
//
// @sBalanceUrl:  protocol://ip[:port][,ip[:port]]/path
//    etcd://127.0.0.1:2379/path
func Listen(sUrl, sDriverUrl string) (server Server, err error) {
	url, err := U.UrlFromString(sUrl)
	if err != nil {
		return
	}

	driverUrl, err := U.UrlFromString(sDriverUrl)
	if err != nil {
		return
	}

	proto, err := P.GetProtocol(url.Protocol)
	if err != nil {
		return
	}

	server, err = proto.CreateServer(url, driverUrl)
	server.AddAop(A.DefaultAop().GetAop()...)
	return
}

func Dial(sUrl, sBalancer string) (client Client, err error) {
	url, err := U.UrlFromString(sUrl)
	if err != nil {
		return
	}

	balancer, err := B.GetBalancer(sBalancer)
	if err != nil {
		return
	}

	clients := newClients(url, balancer, nil)
	err = clients.dial()
	if err != nil {
		return
    }

	client = clients
	client.AddAop(A.DefaultAop().GetAop()...)
	return
}

