// Package clients contains clients for all the operations.
// It has:
// CredentialsClient for registration.
// UsersClient for login in.
// CardsClient to operate with bank cards.
// LogPassesClient to operate with log passes.
// TextsClient to operate with texts.
// BinaryFilesClient to operate with binary files.
package clients

import (
	"context"
	"github.com/mishankoGO/GophKeeper/internal/client/interfaces"
	"github.com/mishankoGO/GophKeeper/internal/converters"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// LogPassesClient contains log pass client service.
type LogPassesClient struct {
	service pb.LogPassesClient
	repo    interfaces.Repository
	offline bool
}

// NewLogPassesClient creates new LogPasses client.
func NewLogPassesClient(cc *grpc.ClientConn, repo interfaces.Repository) *LogPassesClient {

	if cc != nil {
		service := pb.NewLogPassesClient(cc)
		return &LogPassesClient{service: service, repo: repo, offline: false}
	}
	return &LogPassesClient{repo: repo, offline: true}
}

// Insert method inserts new LogPasses.
func (c *LogPassesClient) Insert(ctx context.Context, req *pb.InsertLogPassRequest) (*pb.InsertResponse, error) {
	lp, err := converters.PBLogPassToLogPass(req.GetUser().GetUserId(), req.GetLogPass())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting proto log pass to model log pass: %v", err)
	}
	err = c.repo.InsertLP(lp)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error inserting log pass: %v", err)
	}

	if !c.offline {
		resp, err := c.service.Insert(ctx, req)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error inserting log pass: %v", err)
		}

		return resp, nil
	}

	return &pb.InsertResponse{IsInserted: true}, nil
}

// Get method retrieves log pass information.
func (c *LogPassesClient) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetLogPassResponse, error) {
	lp, err := c.repo.GetLP(req.GetName())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error retrieving log pass: %v", err)
	}
	protoLP, err := converters.LogPassToPBLogPass(lp)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting log pass to proto log pass: %v", err)
	}

	//resp, err := c.service.Get(ctx, req)
	//if err != nil {
	//	return nil, status.Errorf(codes.Internal, "error getting log pass information: %v", err)
	//}
	//
	//return resp, nil
	return &pb.GetLogPassResponse{LogPass: protoLP}, nil
}

// Update method updates log pass information.
func (c *LogPassesClient) Update(ctx context.Context, req *pb.UpdateLogPassRequest) (*pb.UpdateLogPassResponse, error) {
	mLogPass, err := converters.PBLogPassToLogPass(req.GetUser().GetUserId(), req.GetLogPass())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting proto log pass to model log pass: %v", err)
	}
	_, err = c.repo.UpdateLP(mLogPass)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error updating log pass: %v", err)
	}
	if !c.offline {
		resp, err := c.service.Update(ctx, req)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error updating log pass information: %v", err)
		}

		return resp, nil
	}
	return &pb.UpdateLogPassResponse{LogPass: req.GetLogPass()}, nil
}

// Delete method deletes log pass.
func (c *LogPassesClient) Delete(ctx context.Context, req *pb.DeleteLogPassRequest) (*pb.DeleteResponse, error) {
	err := c.repo.DeleteLP(req.GetName())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error deleting log pass: %v", err)
	}

	if !c.offline {
		resp, err := c.service.Delete(ctx, req)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error deleting log pass information: %v", err)
		}

		return resp, nil
	}
	return &pb.DeleteResponse{Ok: true}, nil
}
