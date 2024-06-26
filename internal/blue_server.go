package internal

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"blue/bsp"
	"blue/cluster"
	"blue/common/timewheel"
	"blue/config"
)

const version_ = "blue v0.1"

var clusterConf = config.BC.ClusterConfig

type Exec interface {
	ExecChain(*Context)
}

type ServerInter interface {
	Exec
	Handle(context.Context, net.Conn)
	Close()
}

// BlueServer implements tcp.Handler and serves as a redis server
type BlueServer struct {
	activeConn sync.Map
	db         []*DB
	closed     atomic.Int32
	cc         *cluster.Cluster
}

func NewBlueServer(dbs ...*DB) *BlueServer {
	b := &BlueServer{
		db:         make([]*DB, len(dbs)),
		activeConn: sync.Map{},
	}

	for i := 0; i < len(dbs); i++ {
		b.db[i] = dbs[i]
	}

	if clusterConf.OpenCluster() {
		b.initClu()
	}

	return b
}

func (svr *BlueServer) initClu() {
	svr.cc = cluster.NewCluster(
		clusterConf.TryTimes,
		clusterConf.Port,
		"",
		time.Duration(clusterConf.DialTimeout)*time.Second)

	// 发送本地地址到集群
	svr.cc.Notify(svr.cc.LocalAddr())

	// 获取集群地址
	addrs := svr.cc.GetClusterAddrs(clusterConf.ClusterAddr)

	// 注册集群地址
	svr.cc.Register(addrs...)
}

// Handle receives and executes redis commands
func (svr *BlueServer) Handle(ctx context.Context, conn net.Conn) {
	if svr.isClose() {
		_ = conn.Close()
		return
	}

	client := NewContext(ctx, conn)
	svr.activeConn.Store(client, struct{}{})
	defer func() {
		svr.closeClient(client)
	}()

	canCtx, cancelFunc := context.WithCancel(*client)
	bch, errch := bsp.BspProtos(canCtx, conn)
	defer func() {
		cancelFunc()
		close(bch)
		close(errch)
	}()

	if svr.isCluster() {
		svr.clusterHandle(client, bch, errch)
	} else {
		svr.localHandle(client, bch, errch)
	}

}

func (svr *BlueServer) localHandle(ctx *Context, bch chan *bsp.BspProto, errch chan *bsp.ErrResp) {
	for !ctx.isClose() {
		timewheel.Delay(ctx.maxActive, ctx.cliToken, func() {
			svr.closeClient(ctx)
		})

		select {
		case <-ctx.Done():
			return
		case req := <-bch:
			fmt.Printf("local req:[%+v]\n", req)
			ctx.request = req
			ctx.response = bsp.Reply(nil)

			svr.ExecChain(ctx)
			_, _ = ctx.Reply()
			bsp.PutBspProto(req)

		case err := <-errch:
			if !errors.Is(err, bsp.RequestEnd) {
				ctx.response = err
				_, _ = ctx.Reply()
			}
		}
	}
}

func (svr *BlueServer) clusterHandle(ctx *Context, bch chan *bsp.BspProto, errch chan *bsp.ErrResp) {
	for !ctx.isClose() {
		select {
		case <-ctx.Done():
			return
		case req := <-bch:
			fmt.Printf("cluster req:[%+v]\n", req)
			ctx.request = req
			ctx.response = bsp.Reply(nil)

			res, ok := svr.cc.Dial(ctx.request)
			if !ok {
				svr.ExecChain(ctx)
			} else {
				ctx.response = bsp.NewClusterReply(res)
			}

			_, _ = ctx.Reply()
			bsp.PutBspProto(req)

		case err := <-errch:
			if !errors.Is(err, bsp.RequestEnd) {
				ctx.response = err
				_, _ = ctx.Reply()
			}
		}
	}
}

func (svr *BlueServer) isCluster() bool {
	return svr.cc != nil
}

func (svr *BlueServer) isClose() bool {
	return svr.closed.Load() == 1
}

func (svr *BlueServer) closeClient(client *Context) {
	if client == nil || client.isClose() {
		return
	}
	client.Close()
	svr.activeConn.Delete(client)
}

// Close stops handler
func (svr *BlueServer) Close() {
	svr.closed.Add(1)

	svr.activeConn.Range(func(key interface{}, _ interface{}) bool {
		client := key.(*Context)
		client.Close()
		return true
	})
}
