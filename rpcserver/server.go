package rpcserver

import (
	"fmt"
	"net"
	"net/rpc"
)

type RpcServer struct {
	// peers    map[string]*peer // list peers which are still pinging to bootnode continuously
	// peersMtx sync.Mutex
	server *rpc.Server
	Config *RpcServerConfig // config for RPC server
	pmap   PeerMap
}

func NewRPCServer(
	conf *RpcServerConfig,
	pmap PeerMap,
) (
	*RpcServer,
	error,
) {
	rpcServer := new(RpcServer)
	rpcServer.server = rpc.NewServer()
	rpcServer.Config = conf
	rpcServer.pmap = pmap
	return rpcServer, nil
}

func (rpcServer *RpcServer) Start() error {
	handler := &Handler{
		rpcServer: rpcServer,
	}
	rpcServer.server.Register(handler)
	listenAddr := fmt.Sprintf(":%d", rpcServer.Config.Port)
	listenner, err := net.Listen("tcp", listenAddr)
	if err != nil {
		logger.Errorf("listen in address %v error: %v\n", listenAddr, err)
		return err
	}
	rpcServer.server.Accept(listenner)
	listenner.Close()
	return nil
}
