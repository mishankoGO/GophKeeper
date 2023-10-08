// Package clients contains clients for all the operations.
// It has:
// CredentialsClient for registration.
// UsersClient for login in.
// CardsClient to operate with bank cards.
// LogPassesClient to operate with log passes.
// TextsClient to operate with texts.
// LogPassesClient to operate with log passes.
package clients

import (
	"bytes"
	"context"
	"github.com/mishankoGO/GophKeeper/internal/client/interfaces"
	"github.com/mishankoGO/GophKeeper/internal/converters"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/models/log_passes"
	"github.com/mishankoGO/GophKeeper/internal/security"
	"github.com/mishankoGO/GophKeeper/pkg/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// LogPassesClient contains log pass client service.
type LogPassesClient struct {
	Security *security.Security
	service  pb.LogPassesClient
	repo     interfaces.Repository
	offline  bool
}

// NewLogPassesClient creates new LogPasses client.
func NewLogPassesClient(cc *grpc.ClientConn, repo interfaces.Repository, security *security.Security) *LogPassesClient {
	if cc != nil {
		service := pb.NewLogPassesClient(cc)
		return &LogPassesClient{service: service, repo: repo, Security: security, offline: false}
	}
	return &LogPassesClient{repo: repo, Security: security, offline: true}
}

// Insert method inserts new LogPasses.
func (c *LogPassesClient) Insert(ctx context.Context, req *pb.InsertLogPassRequest) (*pb.InsertResponse, error) {
	lp, err := converters.PBLogPassToLogPass(req.GetUser().GetUserId(), req.GetLogPass())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting proto log pass to model log pass: %v", err)
	}

	// encrypt data
	var buf bytes.Buffer
	buf.Write(lp.Login)

	encLogin := c.Security.EncryptData(buf)

	// set encrypted login as Login
	lp.Login = encLogin

	// encrypt pass
	buf.Reset()

	buf.Write(lp.Password)

	encPass := c.Security.EncryptData(buf)

	// set encrypted pass as Password
	lp.Password = encPass

	err = c.repo.InsertLP(lp)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error inserting log pass: %v", err)
	}

	if !c.offline {
		req.LogPass.Login = encLogin
		req.LogPass.Pass = encPass
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
	if c.offline {
		lp, err := c.repo.GetLP(req.GetName())
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error retrieving log pass: %v", err)
		}

		// decrypt data
		decLogin, err := c.Security.DecryptData(lp.Login)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error decrypting login: %v", err)
		}

		decPass, err := c.Security.DecryptData(lp.Password)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error decrypting pass: %v", err)
		}
		lp.Login = bytes.Trim(decLogin, "\"\n")
		lp.Password = bytes.Trim(decPass, "\"\n")

		protoLP, err := converters.LogPassToPBLogPass(lp)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error converting log pass to proto log pass: %v", err)
		}
		return &pb.GetLogPassResponse{LogPass: protoLP}, nil
	}

	resp, err := c.service.Get(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting log pass information: %v", err)
	}

	// decrypt data
	ll := resp.LogPass.Login
	pp := resp.LogPass.Pass

	decLogin, err := c.Security.DecryptData(ll)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error decrypting login: %v", err)
	}
	decPass, err := c.Security.DecryptData(pp)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error decrypting pass: %v", err)
	}

	resp.LogPass.Login = bytes.Trim(decLogin, "\"\n")
	resp.LogPass.Pass = bytes.Trim(decPass, "\"\n")
	return resp, nil
}

// Update method updates log pass information.
func (c *LogPassesClient) Update(ctx context.Context, req *pb.UpdateLogPassRequest) (*pb.UpdateLogPassResponse, error) {
	mLogPass, err := converters.PBLogPassToLogPass(req.GetUser().GetUserId(), req.GetLogPass())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting proto log pass to model log pass: %v", err)
	}

	// encrypt data
	var buf bytes.Buffer
	buf.Write(mLogPass.Login)

	encLogin := c.Security.EncryptData(buf)

	buf.Reset()

	buf.Write(mLogPass.Password)

	encPass := c.Security.EncryptData(buf)

	mLogPass.Login = encLogin
	mLogPass.Password = encPass

	_, err = c.repo.UpdateLP(mLogPass)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error updating log pass: %v", err)
	}

	if !c.offline {
		req.LogPass.Login = encLogin
		req.LogPass.Pass = encPass
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

// List method to list all log passes.
func (c *LogPassesClient) List(ctx context.Context, req *pb.ListLogPassRequest) (*pb.ListLogPassResponse, []*log_passes.LogPasses, error) {
	lps, err := c.repo.ListLP()
	if err != nil {
		return nil, nil, status.Errorf(codes.Internal, "error listing log passes: %v", err)
	}

	//pbLPs := make([]*pb.LogPass, len(lps))
	//for _, lp := range lps {
	//	pbLP, err := converters.LogPassToPBLogPass(lp)
	//	if err != nil {
	//		return nil, status.Errorf(codes.Internal, "error converting log pass: %v", err)
	//	}
	//	// decrypt data
	//	decLogin, err := c.Security.DecryptData(pbLP.Login)
	//	if err != nil {
	//		return nil, status.Errorf(codes.Internal, "error decrypting data: %v", err)
	//	}
	//	decPass, err := c.Security.DecryptData(pbLP.Pass)
	//	if err != nil {
	//		return nil, status.Errorf(codes.Internal, "error decrypting data: %v", err)
	//	}
	//
	//	pbLP.Login = bytes.Trim(decLogin, "\"\n")
	//	pbLP.Pass = bytes.Trim(decPass, "\"\n")
	//	pbLPs = append(pbLPs, pbLP)
	//}

	resp, err := c.service.List(ctx, req)
	if err != nil {
		return nil, nil, status.Errorf(codes.Internal, "error listing log passes: %v", err)
	}
	return resp, lps, err
}

// Sync method to sync log passes between dbs.
func (c *LogPassesClient) Sync(ctx context.Context, req *pb.ListLogPassRequest) error {
	serverLPs, clientLPs, err := c.List(ctx, req)
	if err != nil {
		return err
	}

	// arrays of names for future syncing
	clientNames := make([]string, len(clientLPs))
	serverNames := make([]string, len(serverLPs.GetLogPasses()))

	// flag which shows which db has the latest data.
	// if flag set to "server", it means server has fresher data.
	dataPrimary := "client"

	// update cycle
	for _, clp := range clientLPs {
		cname := clp.Name
		clientNames = append(clientNames, cname)
		for _, slp := range serverLPs.GetLogPasses() {
			sname := slp.GetName()
			serverNames = append(serverNames, sname)

			// update common log passes
			if sname == cname {
				if clp.UpdatedAt.After(slp.UpdatedAt.AsTime()) {
					// convert log pass to proto log pass
					protoLP, err := converters.LogPassToPBLogPass(clp)
					if err != nil {
						return err
					}

					// update server log pass
					reqS := &pb.UpdateLogPassRequest{User: req.GetUser(), LogPass: protoLP}
					_, err = c.service.Update(ctx, reqS)
					if err != nil {
						return err
					}
					dataPrimary = "client"
				} else {
					// convert proto log pass to log pass
					lp, err := converters.PBLogPassToLogPass(req.GetUser().GetUserId(), slp)
					if err != nil {
						return err
					}

					// update client log pass
					_, err = c.repo.UpdateLP(lp)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	if dataPrimary == "server" {
		// insert missing server log passes to client
		for _, slp := range serverLPs.GetLogPasses() {
			// convert proto log pass to model log pass
			lp, err := converters.PBLogPassToLogPass(req.GetUser().GetUserId(), slp)
			if err != nil {
				return err
			}

			if !util.StringInSlice(slp.Name, clientNames) {
				// insert missing log pass to client db
				err = c.repo.InsertLP(lp)
				if err != nil {
					return err
				}
			}
		}

		// delete log passes
		for _, clp := range clientLPs {
			if !util.StringInSlice(clp.Name, serverNames) {
				// delete log passes absent in server
				err = c.repo.DeleteLP(clp.Name)
				if err != nil {
					return err
				}
			}
		}
	} else if dataPrimary == "client" {
		// insert missing client log passes to log pass
		for _, clp := range clientLPs {
			// convert model log pass to proto log pass
			protoLP, err := converters.LogPassToPBLogPass(clp)
			if err != nil {
				return err
			}

			if !util.StringInSlice(clp.Name, serverNames) {
				// insert missing log pass to server db
				reqS := &pb.InsertLogPassRequest{User: req.GetUser(), LogPass: protoLP}
				_, err = c.service.Insert(ctx, reqS)
				if err != nil {
					return err
				}
			}
		}

		// delete log passes
		for _, slp := range serverLPs.GetLogPasses() {
			if !util.StringInSlice(slp.GetName(), clientNames) {
				// delete log passes absent in client
				resD := &pb.DeleteLogPassRequest{User: req.GetUser(), Name: slp.GetName()}
				_, err = c.service.Delete(ctx, resD)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
