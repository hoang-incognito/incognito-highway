package route

import (
	"highway/route/mocks"
	"testing"
	"time"

	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestEnlistLoop(t *testing.T) {
	h, _ := setupHost()
	publisher := &mocks.Publisher{}
	publisher.On("Publish", mock.Anything, mock.Anything).Return(nil)
	connector := &Connector{
		host:      h,
		publisher: publisher,
		stop:      make(chan int),
	}
	go connector.Start()
	time.Sleep(5 * time.Second)
	connector.stop <- 1
	publisher.AssertNumberOfCalls(t, "Publish", 2)
}

func TestCloseBothStreams(t *testing.T) {
	h, net := setupHost()
	net.On("ClosePeer", mock.Anything).Return(nil)
	keeper := &mocks.ConnKeeper{}
	keeper.On("CloseConnection", mock.Anything).Return(nil)
	connector := &Connector{
		host: h,
		hwc:  keeper,
	}
	connector.closePeer(peer.ID(""))

	net.AssertNumberOfCalls(t, "ClosePeer", 1)
	keeper.AssertNumberOfCalls(t, "CloseConnection", 1)
}

func TestBroadcastEnlistMsg(t *testing.T) {
	publisher := &mocks.Publisher{}
	publisher.On("Publish", mock.Anything, mock.Anything).Return(nil)
	connector := &Connector{
		addrInfo:      peer.AddrInfo{},
		supportShards: []byte{},
		rpcUrl:        "",
		publisher:     publisher,
	}
	err := connector.enlist()
	assert.Nil(t, err)
	publisher.AssertNumberOfCalls(t, "Publish", 1)
}
