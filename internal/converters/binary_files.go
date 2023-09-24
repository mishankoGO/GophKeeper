package converters

import (
	"encoding/json"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/models/binary_files"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func PBBinaryFileToBinaryFile(uid string, pbf *pb.BinaryFile) (*binary_files.Files, error) {
	if pbf.Meta != nil {
		var meta = make(map[string]string)
		err := json.Unmarshal(pbf.Meta, &meta)
		if err != nil {
			return nil, status.Error(codes.Internal, "error unmarshalling binary file meta")
		}
		return &binary_files.Files{UserID: uid, Name: pbf.Name, HashFile: pbf.HashFile, UpdatedAt: pbf.UpdatedAt.AsTime(), Meta: meta}, nil
	}
	return &binary_files.Files{UserID: uid, Name: pbf.Name, HashFile: pbf.HashFile, UpdatedAt: pbf.UpdatedAt.AsTime()}, nil
}

func BinaryFileToPBBinaryFile(bf *binary_files.Files) (*pb.BinaryFile, error) {
	if bf.Meta != nil {
		meta, err := json.Marshal(bf.Meta)
		if err != nil {
			return nil, status.Error(codes.Internal, "error marshalling binary file meta")
		}
		return &pb.BinaryFile{Name: bf.Name, HashFile: bf.HashFile, UpdatedAt: timestamppb.New(bf.UpdatedAt), Meta: meta}, nil
	}
	return &pb.BinaryFile{Name: bf.Name, HashFile: bf.HashFile, UpdatedAt: timestamppb.New(bf.UpdatedAt)}, nil
}
