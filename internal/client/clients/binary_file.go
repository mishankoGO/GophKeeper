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
	"github.com/mishankoGO/GophKeeper/internal/client/interfaces"
	"github.com/mishankoGO/GophKeeper/internal/converters"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/security"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// BinaryFilesClient contains binary files client service.
type BinaryFilesClient struct {
	service  pb.BinaryFilesClient
	Security *security.Security
	repo     interfaces.Repository
	offline  bool
}

// NewBinaryFilesClient creates new BinaryFiles client.
func NewBinaryFilesClient(cc *grpc.ClientConn, repo interfaces.Repository, security *security.Security) *BinaryFilesClient {
	if cc != nil {
		service := pb.NewBinaryFilesClient(cc)
		return &BinaryFilesClient{service: service, Security: security, repo: repo, offline: false}

	}
	return &BinaryFilesClient{repo: repo, Security: security, offline: true}
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

	// set encrypted binary file to binary file
	bf.File = encData

	err = c.repo.InsertBF(bf)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error inserting binary file: %v", err)
	}

	if !c.offline {
		req.File.File = bytes.Trim(encData, "\"\n")
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

		// set decrypted binary file to binary file
		bf.File = bytes.Trim(decData, "\"\n")

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

	// set decrypted binary file to binary file
	resp.File.File = bytes.Trim(decData, "\"\n")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error retrieving binary file: %v", err)
	}
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

	// set encrypted binary file to Binary file
	mFile.File = encData

	_, err = c.repo.UpdateBF(mFile)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error updating binary file: %v", err)
	}

	if !c.offline {
		req.File.File = encData
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
func (c *BinaryFilesClient) List(ctx context.Context) (*pb.ListBinaryFileResponse, error) {
	bfs, err := c.repo.ListBF()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error listing binary files: %v", err)
	}

	pbBFs := make([]*pb.BinaryFile, len(*bfs))
	for _, bf := range *bfs {
		pbBF, err := converters.BinaryFileToPBBinaryFile(&bf)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error converting binary file: %v", err)
		}
		// decrypt data
		decData, err := c.Security.DecryptData(pbBF.File)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error decrypting data: %v", err)
		}

		// set decrypted binary file to File
		pbBF.File = bytes.Trim(decData, "\"\n")
		pbBFs = append(pbBFs, pbBF)
	}
	resp := &pb.ListBinaryFileResponse{BinaryFiles: pbBFs}
	return resp, err
}
