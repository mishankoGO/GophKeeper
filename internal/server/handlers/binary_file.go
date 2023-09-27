package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/mishankoGO/GophKeeper/internal/converters"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/security"
	"github.com/mishankoGO/GophKeeper/internal/server/interfaces"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	// get file
	file := binaryFile.File

	// create encoder
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	res := &pb.InsertResponse{IsInserted: false}

	// marshall into bytes
	err = encoder.Encode(string(file))
	if err != nil {
		return res, status.Errorf(codes.Internal, "error encoding binary file: %v", err)
	}

	// encrypt data
	encData := bf.Security.EncryptData(buf)

	// set encrypted file as File
	binaryFile.File = encData

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

	// set decrypted file to File
	binaryFile.File = bytes.Trim(decData, "\"\n")

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
	// convert proto user to model use
	user := converters.PBUserToUser(req.GetUser())

	// convert proto binary file to model binary file
	binaryFile, err := converters.PBBinaryFileToBinaryFile(user.UserID, req.GetFile())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting binary file: %v", err)
	}

	// get file
	file := binaryFile.File

	// create encoder
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)

	// marshall into bytes
	err = encoder.Encode(string(file))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error encoding binary file: %v", err)
	}

	// encrypt data
	encData := bf.Security.EncryptData(buf)

	// set encrypted file as File
	binaryFile.File = encData

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
