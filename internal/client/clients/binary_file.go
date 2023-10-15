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
	"bytes"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/mishankoGO/GophKeeper/internal/client/interfaces"
	"github.com/mishankoGO/GophKeeper/internal/converters"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/models/binary_files"
	"github.com/mishankoGO/GophKeeper/internal/security"
	"github.com/mishankoGO/GophKeeper/pkg/util"
)

// BinaryFilesClient contains binary files client service.
type BinaryFilesClient struct {
	service  pb.BinaryFilesClient
	Security *security.Security
	repo     interfaces.Repository
	offline  bool
}

// NewBinaryFilesClient creates new BinaryFiles client.
func NewBinaryFilesClient(cc *grpc.ClientConn, repo interfaces.Repository) *BinaryFilesClient {
	if cc != nil {
		service := pb.NewBinaryFilesClient(cc)
		return &BinaryFilesClient{service: service, repo: repo, offline: false}

	}
	return &BinaryFilesClient{repo: repo, offline: true}
}

// Insert method inserts new BinaryFiles.
func (c *BinaryFilesClient) Insert(ctx context.Context, req *pb.InsertBinaryFileRequest) (*pb.InsertResponse, error) {
	bf, err := converters.PBBinaryFileToBinaryFile(req.GetUser().GetUserId(), req.GetFile())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting proto binary file to binary file: %w", err)
	}

	// encrypt data
	var buf bytes.Buffer
	buf.Write(bf.File)

	encData := c.Security.EncryptData(buf)

	buf.Reset()

	buf.Write(bf.Extension)
	encExt := c.Security.EncryptData(buf)

	// set encrypted binary file to binary file
	bf.File = encData
	bf.Extension = encExt

	err = c.repo.InsertBF(bf)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error inserting binary file: %v", err)
	}

	if !c.offline {
		req.File.File = bytes.Trim(encData, "\"\n")
		req.File.Extension = bytes.Trim(encExt, "\"\n")
		resp, err := c.service.Insert(ctx, req)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error inserting binary file: %v", err)
		}
		return resp, nil
	}

	return &pb.InsertResponse{IsInserted: true}, nil
}

// Get method retrieves binary file information.
func (c *BinaryFilesClient) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetBinaryFileResponse, error) {
	if c.offline {
		bf, err := c.repo.GetBF(req.GetName())
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error retrieving binary file: %v", err)
		}

		// decrypt data
		decData, err := c.Security.DecryptData(bf.File)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error decrypting data: %v", err)
		}

		decExt, err := c.Security.DecryptData(bf.Extension)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error decrypting data: %v", err)
		}
		// set decrypted binary file to binary file
		bf.File = bytes.Trim(decData, "\"\n")
		bf.Extension = bytes.Trim(decExt, "\"\n")

		protoBinaryFile, err := converters.BinaryFileToPBBinaryFile(bf)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error converting model binary file to proto binary file: %v", err)
		}
		return &pb.GetBinaryFileResponse{File: protoBinaryFile}, nil
	}
	resp, err := c.service.Get(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting response: %v", err)
	}

	// decrypt data
	f := resp.File.File
	decData, err := c.Security.DecryptData(f)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error decrypting data: %v", err)
	}

	decExt, err := c.Security.DecryptData(resp.File.Extension)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error decryption extension: %v", err)
	}

	// set decrypted binary file to binary file
	resp.File.File = bytes.Trim(decData, "\"\n")
	resp.File.Extension = bytes.Trim(decExt, "\"\n")

	return resp, nil
}

// Update method updates binary file information.
func (c *BinaryFilesClient) Update(ctx context.Context, req *pb.UpdateBinaryFileRequest) (*pb.UpdateBinaryFileResponse, error) {
	mFile, err := converters.PBBinaryFileToBinaryFile(req.GetUser().GetUserId(), req.GetFile())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting proto binary file to model binary file: %v", err)
	}

	// encrypt data
	var buf bytes.Buffer
	buf.Write(mFile.File)

	encData := c.Security.EncryptData(buf)

	buf.Reset()
	buf.Write(mFile.Extension)

	encExt := c.Security.EncryptData(buf)

	// set encrypted binary file to Binary file
	mFile.File = encData
	mFile.Extension = encExt

	_, err = c.repo.UpdateBF(mFile)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error updating binary file: %v", err)
	}

	if !c.offline {
		req.File.File = encData
		req.File.Extension = encExt
		resp, err := c.service.Update(ctx, req)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error updating binary file information: %v", err)
		}
		return resp, nil
	}

	return &pb.UpdateBinaryFileResponse{File: req.GetFile()}, nil
}

// Delete method deletes binary file.
func (c *BinaryFilesClient) Delete(ctx context.Context, req *pb.DeleteBinaryFileRequest) (*pb.DeleteResponse, error) {
	err := c.repo.DeleteBF(req.GetName())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error deleting binary file: %v", err)
	}

	if !c.offline {
		resp, err := c.service.Delete(ctx, req)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error deleting binary file information: %v", err)
		}

		return resp, nil
	}

	return &pb.DeleteResponse{Ok: true}, nil
}

// List method to list all binary files.
func (c *BinaryFilesClient) List(ctx context.Context, req *pb.ListBinaryFileRequest) (*pb.ListBinaryFileResponse, []*binary_files.Files, error) {

	// client repo
	bfs, err := c.repo.ListBF()
	if err != nil {
		return nil, nil, status.Errorf(codes.Internal, "error listing binary files: %v", err)
	}

	resp, err := c.service.List(ctx, req)
	if err != nil {
		return nil, nil, fmt.Errorf("error listing data from server: %v", err)
	}
	return resp, bfs, nil
}

// Sync method to sync binary files between dbs.
func (c *BinaryFilesClient) Sync(ctx context.Context, req *pb.ListBinaryFileRequest) error {
	serverBFs, clientBFs, err := c.List(ctx, req)
	if err != nil {
		return err
	}

	// arrays of names for future syncing
	clientNames := make([]string, len(clientBFs))
	serverNames := make([]string, len(serverBFs.GetBinaryFiles()))

	// flag which shows which db has the latest data.
	// if flag set to "server", it means server has fresher data.
	var dataPrimary string
	if c.offline {
		dataPrimary = "client"
	} else {
		dataPrimary = "server"
	}

	// update cycle
	for _, cbf := range clientBFs {
		cname := cbf.Name
		clientNames = append(clientNames, cname)
		for _, sbf := range serverBFs.GetBinaryFiles() {
			sname := sbf.GetName()
			serverNames = append(serverNames, sname)

			// update common files
			if sname == cname {
				if cbf.UpdatedAt.After(sbf.UpdatedAt.AsTime()) {
					// convert binary file to proto binary file
					protoBF, err := converters.BinaryFileToPBBinaryFile(cbf)
					if err != nil {
						return err
					}

					// update server binary file
					reqS := &pb.UpdateBinaryFileRequest{User: req.GetUser(), File: protoBF}
					_, err = c.service.Update(ctx, reqS)
					if err != nil {
						return err
					}
				} else {
					// convert proto binary file to binary file
					bf, err := converters.PBBinaryFileToBinaryFile(req.GetUser().GetUserId(), sbf)
					if err != nil {
						return err
					}

					// update client binary file
					_, err = c.repo.UpdateBF(bf)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	if dataPrimary == "server" {
		// insert missing server binary files to client
		for _, sbf := range serverBFs.GetBinaryFiles() {
			// convert proto binary file to model binary file
			bf, err := converters.PBBinaryFileToBinaryFile(req.GetUser().GetUserId(), sbf)
			if err != nil {
				return err
			}

			if !util.StringInSlice(sbf.Name, clientNames) {
				// insert missing binary file to client db
				err = c.repo.InsertBF(bf)
				if err != nil {
					return err
				}
			}
		}

		// delete files
		for _, cbf := range clientBFs {
			if !util.StringInSlice(cbf.Name, serverNames) {
				// delete binary files absent in server
				err = c.repo.DeleteBF(cbf.Name)
				if err != nil {
					return err
				}
			}
		}
	} else if dataPrimary == "client" {
		// insert missing client binary files to binary file
		for _, cbf := range clientBFs {
			// convert model binary file to proto binary file
			protoBF, err := converters.BinaryFileToPBBinaryFile(cbf)
			if err != nil {
				return err
			}

			if !util.StringInSlice(cbf.Name, serverNames) {
				// insert missing binary file to server db
				reqS := &pb.InsertBinaryFileRequest{User: req.GetUser(), File: protoBF}
				_, err = c.service.Insert(ctx, reqS)
				if err != nil {
					return err
				}
			}
		}

		// delete files
		for _, sbf := range serverBFs.GetBinaryFiles() {
			if !util.StringInSlice(sbf.GetName(), clientNames) {
				// delete binary files absent in client
				resD := &pb.DeleteBinaryFileRequest{User: req.GetUser(), Name: sbf.GetName()}
				_, err = c.service.Delete(ctx, resD)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// SetSecurity method to set security attribute.
func (c *BinaryFilesClient) SetSecurity(security *security.Security) {
	c.Security = security
}
