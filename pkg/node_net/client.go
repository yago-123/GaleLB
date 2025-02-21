package node_net

import (
	"context"
	"fmt"
	"log"

	"github.com/sirupsen/logrus"
	pb "github.com/yago-123/galelb/pkg/consensus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn   *grpc.ClientConn
	client pb.LBNodeManagerClient

	logger *logrus.Logger
}

func New(logger *logrus.Logger, ip string, port int) *Client {
	remoteServer := fmt.Sprintf("%s:%d", ip, port)

	// todo(): we must have an array of remove servers for multi-node load balancer
	conn, err := grpc.NewClient(remoteServer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to load balancer: %v", err)
	}

	client := pb.NewLBNodeManagerClient(conn)

	return &Client{
		conn:   conn,
		client: client,
		logger: logger,
	}
}

func (s *Client) RegisterNode(ctx context.Context) error {
	resp, err := s.client.RegisterNode(ctx, &pb.NodeInfo{
		NodeId: "192.168.1.1",
		Ip:     "192.168.1.1",
		Port:   1234,
	})

	if err != nil {
		return fmt.Errorf("error registering node: %v", err)
	}

	if resp.Success == false {
		return fmt.Errorf("error registering node: %v", resp.Message)
	}

	s.logger.Debugf("registered node with success", resp.Success, resp.Message)
	return nil
}
