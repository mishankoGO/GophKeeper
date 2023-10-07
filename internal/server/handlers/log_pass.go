// Package handlers contains servers interfaces.
// The list of servers:
//     Users, Credentials, BinaryFiles, Cards, Texts, LogPasses
package handlers

import (
	"bytes"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/mishankoGO/GophKeeper/internal/converters"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/security"
	"github.com/mishankoGO/GophKeeper/internal/server/interfaces"
)

func NewLogPasses(repo interfaces.Repository, security *security.Security) *LogPasses {
	return &LogPasses{
		Repo:     repo,
		Security: *security,
	}
}

// LogPasses struct realizes LogPassesServer interface.
type LogPasses struct {
	pb.UnimplementedLogPassesServer
	Repo     interfaces.Repository // data storage
	Security security.Security     // cipher component
}

// Insert method encrypts and inserts data to db.
func (lp *LogPasses) Insert(ctx context.Context, req *pb.InsertLogPassRequest) (*pb.InsertResponse, error) {
	// convert proto log pass to model log pass
	logPass, err := converters.PBLogPassToLogPass(req.User.UserId, req.GetLogPass())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting log pass: %v", err)
	}

	// get login and password
	login := logPass.Login
	password := logPass.Password

	// create encoder
	var buf bytes.Buffer
	res := &pb.InsertResponse{IsInserted: false}

	// marshal into bytes
	buf.Write(login)

	// encrypt login
	encLogin := lp.Security.EncryptData(buf)

	// marshal into bytes
	buf.Write(password)

	// encrypt password
	encPassword := lp.Security.EncryptData(buf)

	// set encrypted log pass as Login and Password
	logPass.Login = encLogin
	logPass.Password = encPassword

	// insert new log pass to db
	err = lp.Repo.InsertLP(ctx, logPass)
	if err != nil {
		return res, status.Errorf(codes.Internal, "error inserting log pass: %v", err)
	}

	// set status
	res.IsInserted = true
	return res, nil
}

// Get method gets and decrypts data from db.
func (lp *LogPasses) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetLogPassResponse, error) {
	// convert proto user to model user
	user := converters.PBUserToUser(req.GetUser())
	name := req.GetName()

	// get log pass from database
	logPass, err := lp.Repo.GetLP(ctx, user.UserID, name)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting log pass: %v", err)
	}

	// decrypt login and password
	decLogin, err := lp.Security.DecryptData(logPass.Login)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error decrypting login: %v", err)
	}
	decPassword, err := lp.Security.DecryptData(logPass.Password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error decrypting password: %v", err)
	}

	// set decrypted login and password as Login and Password
	logPass.Login = bytes.Trim(decLogin, "\"\n")
	logPass.Password = bytes.Trim(decPassword, "\"\n")

	// convert log pass to proto log pass
	pbLogPass, err := converters.LogPassToPBLogPass(logPass)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting log pass: %v", err)
	}

	// create response
	res := &pb.GetLogPassResponse{LogPass: pbLogPass}

	return res, nil
}

// Update method encrypts new log pass and updates record in db.
func (lp *LogPasses) Update(ctx context.Context, req *pb.UpdateLogPassRequest) (*pb.UpdateLogPassResponse, error) {
	// convert proto user to model user
	user := converters.PBUserToUser(req.GetUser())

	// convert proto log pass to model log pass
	logPass, err := converters.PBLogPassToLogPass(user.UserID, req.GetLogPass())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting log pass: %v", err)
	}

	// get login and password
	login := logPass.Login
	password := logPass.Password

	// create encoder
	var buf bytes.Buffer

	// marshall into bytes
	buf.Write(login)

	// encrypt login
	encLogin := lp.Security.EncryptData(buf)

	buf.Reset()

	// marshall into bytes
	buf.Write(password)

	// encrypt password
	encPassword := lp.Security.EncryptData(buf)

	// set encrypted login and password as Login and Password
	logPass.Login = encLogin
	logPass.Password = encPassword

	// update record in db
	updatedLogPass, err := lp.Repo.UpdateLP(ctx, logPass)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error updating log pass: %v", err)
	}

	// convert model log pass to proto log pass
	pbLogPass, err := converters.LogPassToPBLogPass(updatedLogPass)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting log pass: %v", err)
	}

	// create response
	res := &pb.UpdateLogPassResponse{LogPass: pbLogPass}

	return res, nil
}

// Delete method deletes binary file record from db.
func (lp *LogPasses) Delete(ctx context.Context, req *pb.DeleteLogPassRequest) (*pb.DeleteResponse, error) {
	// convert proto user to user and get name
	user := converters.PBUserToUser(req.GetUser())
	name := req.GetName()

	// create response
	res := &pb.DeleteResponse{Ok: false}

	// delete record
	err := lp.Repo.DeleteLP(ctx, user.UserID, name)
	if err != nil {
		return res, status.Errorf(codes.Internal, "error deleting log pass: %v", err)
	}

	// set result
	res.Ok = true

	return res, nil
}

// List method lists all log passes in db.
func (lp *LogPasses) List(ctx context.Context, req *pb.ListLogPassRequest) (*pb.ListLogPassResponse, error) {
	// convert proto user to user
	user := converters.PBUserToUser(req.GetUser())

	lps, err := lp.Repo.ListLP(ctx, user.UserID)
	if err != nil {
		return nil, fmt.Errorf("error listing log passes: %w", err)
	}

	// decrypt log passes
	for i, logPass := range lps {
		decLogin, err := lp.Security.DecryptData(logPass.Login)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error decrypting login: %v", err)
		}

		decPass, err := lp.Security.DecryptData(logPass.Password)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error decrypting password: %v", err)
		}

		lps[i].Login = bytes.Trim(decLogin, "\"\n")
		lps[i].Password = bytes.Trim(decPass, "\"\n")
	}

	// converts model log passes to proto log passes
	protoLPs, err := converters.LogPassesToPBLogPasses(lps)
	if err != nil {
		return nil, fmt.Errorf("error converting log passes: %w", err)
	}

	// create response
	res := &pb.ListLogPassResponse{LogPasses: protoLPs}
	return res, nil
}
