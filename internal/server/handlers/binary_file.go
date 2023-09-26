package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/mishankoGO/GophKeeper/internal/converters"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/pki/sender"
	"github.com/mishankoGO/GophKeeper/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//func NewBinaryFiles(repository.Repository) (*BinaryFiles, error) {
//
//}

type BinaryFiles struct {
	pb.UnimplementedBinaryFilesServer
	Repo repository.Repository
}

func (bf *BinaryFiles) Insert(ctx context.Context, req *pb.InsertBinaryFileRequest) (*pb.InsertResponse, error) {
	binaryFile, err := converters.PBBinaryFileToBinaryFile(req.User.UserId, req.File)
	res := &pb.InsertResponse{IsInserted: false}
	if err != nil {
		return res, status.Errorf(codes.Internal, "error converting binary file: %v", err)
	}

	err = bf.Repo.InsertBF(ctx, binaryFile)
	if err != nil {
		return res, status.Errorf(codes.Internal, "error inserting binary file: %v", err)
	}

	res.IsInserted = true
	return res, nil
}

func (bf *BinaryFiles) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetBinaryFileResponse, error) {
	user := converters.PBUserToUser(req.GetUser())
	name := req.GetName()

	binaryFile, err := bf.Repo.GetBF(ctx, user.UserID, name)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting binary file: %v", err)
	}

	pbBinaryFile, err := converters.BinaryFileToPBBinaryFile(binaryFile)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting binary file: %v", err)
	}

	res := &pb.GetBinaryFileResponse{File: pbBinaryFile}

	return res, nil
}

func (bf *BinaryFiles) Update(ctx context.Context, req *pb.UpdateBinaryFileRequest) (*pb.UpdateBinaryFileResponse, error) {
	user := converters.PBUserToUser(req.GetUser())
	file, err := converters.PBBinaryFileToBinaryFile(user.UserID, req.GetFile())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting binary file: %v", err)
	}

	binaryFile, err := bf.Repo.UpdateBF(ctx, file)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error updating binary file: %v", err)
	}

	pbBinaryFile, err := converters.BinaryFileToPBBinaryFile(binaryFile)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting binary file: %v", err)
	}

	res := &pb.UpdateBinaryFileResponse{File: pbBinaryFile}

	return res, nil
}

func (bf *BinaryFiles) Delete(ctx context.Context, req *pb.DeleteBinaryFileRequest) (*pb.DeleteResponse, error) {
	user := converters.PBUserToUser(req.GetUser())
	name := req.GetName()

	res := &pb.DeleteResponse{Ok: false}

	err := bf.Repo.DeleteBF(ctx, user.UserID, name)
	if err != nil {
		return res, status.Errorf(codes.Internal, "error deleting binary file: %v", err)
	}

	res.Ok = true
	return res, nil
}

// encryptMessage is a helper function to encrypt incoming data with cryptoKey.
func encryptMessage(cryptoKey string, buf bytes.Buffer, encoder json.Encoder) error {
	// if crypto key is present, send encrypted message
	if cryptoKey != "" {
		encMetric, err := sender.Encrypt(cryptoKey, buf.Bytes())
		if err != nil {
			return fmt.Errorf("failed encrypting metric: %w", err)
		}
		buf.Reset()

		err = encoder.Encode(encMetric)
		if err != nil {
			return fmt.Errorf("failed encoding encrypted metric: %w", err)
		}
	}
	return nil
}
