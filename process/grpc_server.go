package process

import (
	"context"
	fmt "fmt"
	logger "highway/customizelog"
	"highway/p2p"
	"highway/process/topic"
)

func ProcessConnection(h *p2p.Host) {
	s := &GRPCService_Server{}
	RegisterProxyRegisterServiceServer(h.GRPC.GetGRPCServer(), s)
}

func (s *GRPCService_Server) ProxyRegister(ctx context.Context, req *ProxyRegisterMsg) (*ProxyRegisterResponse, error) {
	logger.Infof("Receive new request from %v via gRPC", req.GetCommitteePublicKey())
	pairs := []*MessageTopicPair{}
	topicGenerator := new(topic.InsideTopic)
	for _, m := range req.WantedMessages {
		err := topicGenerator.FromMessageType(req.CommitteePublicKey, m)
		responseTopic := ""
		if err == nil {
			responseTopic = topicGenerator.ToString()
			logger.Infof("%v wanted message %v, response %v", req.GetCommitteePublicKey(), responseTopic)
			GlobalPubsub.NewMessage <- responseTopic
		}
		pair := &MessageTopicPair{
			Message: m,
			Topic:   responseTopic,
		}
		pairs = append(pairs, pair)
	}
	fmt.Println(pairs)
	return &ProxyRegisterResponse{Pair: pairs}, nil
}

type GRPCService_Server struct{}
