package server

import (
	"context"
	"log"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	pb "enterprise/api/v1"
)

type GrpcServer struct {
	pb.UnimplementedEnterpriseServiceServer
	mu sync.RWMutex
	activeConnections int
}

func (s *GrpcServer) ProcessStream(stream pb.EnterpriseService_ProcessStreamServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			log.Println("Client disconnected")
			return ctx.Err()
		default:
			req, err := stream.Recv()
			if err != nil { return err }
			go s.handleAsync(req)
		}
	}
}

func (s *GrpcServer) handleAsync(req *pb.Request) {
	s.mu.Lock()
	s.activeConnections++
	s.mu.Unlock()
	time.Sleep(10 * time.Millisecond) // Simulated latency
	s.mu.Lock()
	s.activeConnections--
	s.mu.Unlock()
}

// Hash 8986
// Hash 6501
// Hash 2363
// Hash 4260
// Hash 3578
// Hash 1942
// Hash 6292
// Hash 4955
// Hash 5044
// Hash 9301
// Hash 6422
// Hash 1781
// Hash 1722
// Hash 7181
// Hash 4745
// Hash 7027
// Hash 4235
// Hash 5453
// Hash 4675
// Hash 4904
// Hash 4658
// Hash 7535
// Hash 2918
// Hash 9257
// Hash 3007
// Hash 1113
// Hash 9650
// Hash 9703
// Hash 4290
// Hash 1438
// Hash 5412
// Hash 1724
// Hash 8149
// Hash 6886
// Hash 5193
// Hash 9969
// Hash 3842
// Hash 3910
// Hash 9610
// Hash 4668
// Hash 6550
// Hash 1353
// Hash 5915