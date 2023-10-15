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

// NewBinaryFiles is a constructor for BinaryFilesServer interface instance.
func NewBinaryFiles(repo interfaces.Repository, security *security.Security) *BinaryFiles {
	return &BinaryFiles{
		Repo:     repo,
		Security: *security,
	}
}

// BinaryFiles struct realizes BinaryFilesServer interface.
type BinaryFiles struct {
	pb.UnimplementedBinaryFilesServer
	Repo     interfaces.Repository // data storage
	Security security.Security     // cipher component
}

// Insert method encrypts and inserts data to db.
func (bf *BinaryFiles) Insert(ctx context.Context, req *pb.InsertBinaryFileRequest) (*pb.InsertResponse, error) {
	// convert proto binary file to model binary file
	binaryFile, err := converters.PBBinaryFileToBinaryFile(req.User.UserId, req.File)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting proto binary file to model binary file: %v", err)
	}

	// get file
	file := binaryFile.File
	extenstion := binaryFile.Extension

	// create encoder
	var buf bytes.Buffer
	res := &pb.InsertResponse{IsInserted: false}

	// marshall into bytes
	buf.Write(file)

	// encrypt data
	encData := bf.Security.EncryptData(buf)

	buf.Reset()
	buf.Write(extenstion)

	encExtension := bf.Security.EncryptData(buf)

	// set encrypted file as File
	binaryFile.File = encData
	binaryFile.Extension = encExtension

	// insert new binary file to db
	err = bf.Repo.InsertBF(ctx, binaryFile)
	if err != nil {
		return res, status.Errorf(codes.Internal, "error inserting binary file: %v", err)
	}

	// set status
	res.IsInserted = true

	return res, nil
}

// Get method gets and decrypts data from db.
func (bf *BinaryFiles) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetBinaryFileResponse, error) {
	// convert proto user to model user
	user := converters.PBUserToUser(req.GetUser())
	name := req.GetName()

	// get binary file from database
	binaryFile, err := bf.Repo.GetBF(ctx, user.UserID, name)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting binary file: %v", err)
	}

	// decrypt file
	decData, err := bf.Security.DecryptData(binaryFile.File)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error decrypting binary file: %v", err)
	}

	// decrypt extension
	decExt, err := bf.Security.DecryptData(binaryFile.Extension)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error decrypting extension: %v", err)
	}

	// set decrypted file to File
	binaryFile.File = bytes.Trim(decData, "\"\n")
	binaryFile.Extension = bytes.Trim(decExt, "\"\n")

	// convert binary file to proto binary file
	pbBinaryFile, err := converters.BinaryFileToPBBinaryFile(binaryFile)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting binary file: %v", err)
	}

	// create response
	res := &pb.GetBinaryFileResponse{File: pbBinaryFile}

	return res, nil
}

// Update method encrypts new binary file and updates record in db.
func (bf *BinaryFiles) Update(ctx context.Context, req *pb.UpdateBinaryFileRequest) (*pb.UpdateBinaryFileResponse, error) {
	// convert proto user to model user
	user := converters.PBUserToUser(req.GetUser())

	// convert proto binary file to model binary file
	binaryFile, err := converters.PBBinaryFileToBinaryFile(user.UserID, req.GetFile())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting binary file: %v", err)
	}

	// get file
	file := binaryFile.File
	extension := binaryFile.Extension

	// create encoder
	var buf bytes.Buffer

	// marshall into bytes
	buf.Write(file)

	// encrypt data
	encData := bf.Security.EncryptData(buf)

	buf.Reset()
	buf.Write(extension)

	encExtension := bf.Security.EncryptData(buf)

	// set encrypted file as File
	binaryFile.File = encData
	binaryFile.Extension = encExtension

	// update record in db
	updatedBinaryFile, err := bf.Repo.UpdateBF(ctx, binaryFile)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error updating binary file: %v", err)
	}

	// convert model binary file to proto binary file
	pbBinaryFile, err := converters.BinaryFileToPBBinaryFile(updatedBinaryFile)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting binary file: %v", err)
	}

	// create response
	res := &pb.UpdateBinaryFileResponse{File: pbBinaryFile}

	return res, nil
}

// Delete method deletes binary file record from db.
func (bf *BinaryFiles) Delete(ctx context.Context, req *pb.DeleteBinaryFileRequest) (*pb.DeleteResponse, error) {
	// convert proto user to user and get name
	user := converters.PBUserToUser(req.GetUser())
	name := req.GetName()

	// create response
	res := &pb.DeleteResponse{Ok: false}

	// delete record
	err := bf.Repo.DeleteBF(ctx, user.UserID, name)
	if err != nil {
		return res, status.Errorf(codes.Internal, "error deleting binary file: %v", err)
	}

	// set result
	res.Ok = true

	return res, nil
}

// List method lists all binary files in db.
func (bf *BinaryFiles) List(ctx context.Context, req *pb.ListBinaryFileRequest) (*pb.ListBinaryFileResponse, error) {
	// convert proto user to user
	user := converters.PBUserToUser(req.GetUser())

	bfs, err := bf.Repo.ListBF(ctx, user.UserID)
	if err != nil {
		return nil, fmt.Errorf("error listing binary files: %w", err)
	}

	// decrypt files
	for i, bfFile := range bfs {
		decData, err := bf.Security.DecryptData(bfFile.File)
		decExt, err := bf.Security.DecryptData(bfFile.Extension)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error decrypting binary file: %v", err)
		}
		bfs[i].File = bytes.Trim(decData, "\"\n")
		bfs[i].Extension = bytes.Trim(decExt, "\"\n")
	}

	// converts model binary files to proto binary files
	protoBFs, err := converters.BinaryFilesToPBBinaryFiles(bfs)
	if err != nil {
		return nil, fmt.Errorf("error converting binary files: %w", err)
	}

	// create response
	res := &pb.ListBinaryFileResponse{BinaryFiles: protoBFs}
	return res, nil
}
