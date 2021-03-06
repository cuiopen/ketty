package grpc_proto

import (
	"fmt"
	A "github.com/yyzybb537/ketty/aop"
	COM "github.com/yyzybb537/ketty/common"
	C "github.com/yyzybb537/ketty/context"
	U "github.com/yyzybb537/ketty/url"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type GrpcClient struct {
	A.AopList

	Impl *grpc.ClientConn
	url  U.Url
}

func newGrpcClient(url U.Url) (*GrpcClient, error) {
	c := new(GrpcClient)
	c.url = url
	err := c.dial(url)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (this *GrpcClient) dial(url U.Url) (err error) {
	this.Impl, err = grpc.Dial(url.SAddr, grpc.WithInsecure(),
		grpc.WithPerRPCCredentials(newGrpcMeta(this)))
	return
}

func (this *GrpcClient) Close() {
	this.Impl.Close()
}

func (this *GrpcClient) Invoke(ctx context.Context, handle COM.ServiceHandle, method string, req, rsp interface{}) error {
	ctx = this.invoke(ctx, handle, method, req, rsp)
	return ctx.Err()
}

func (this *GrpcClient) invoke(inCtx context.Context, handle COM.ServiceHandle, method string, req, rsp interface{}) (ctx context.Context) {
	var err error
	ctx = inCtx
	fullMethodName := fmt.Sprintf("/%s/%s", handle.ServiceName(), method)
	aopList := A.GetAop(ctx)
	if aopList != nil {
		ctx = context.WithValue(ctx, "method", fullMethodName)
		ctx = context.WithValue(ctx, "remote", this.url.SAddr)
		metadata := map[string]string{}
		ctx = context.WithValue(ctx, "metadata", metadata)

		for _, aop := range aopList {
			caller, ok := aop.(A.ClientTransportMetaDataAop)
			if ok {
				ctx = caller.ClientSendMetaData(ctx, metadata)
				if ctx.Err() != nil {
					return
				}
			}
		}

		for _, aop := range aopList {
			caller, ok := aop.(A.BeforeClientInvokeAop)
			if ok {
				ctx = caller.BeforeClientInvoke(ctx, req)
				if ctx.Err() != nil {
					return
				}
			}
		}

		defer func() {
			for _, aop := range aopList {
				caller, ok := aop.(A.ClientInvokeCleanupAop)
				if ok {
					caller.ClientCleanup(ctx)
				}
			}
		}()

		for i, _ := range aopList {
			aop := aopList[len(aopList)-i-1]
			caller, ok := aop.(A.AfterClientInvokeAop)
			if ok {
				defer caller.AfterClientInvoke(&ctx, req, rsp)
			}
		}
	}

	err = grpc.Invoke(ctx, fullMethodName, req, rsp, this.Impl)
	if err != nil {
		ctx = C.WithError(ctx, err)
	}

	return
}
