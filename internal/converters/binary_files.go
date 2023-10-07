// Package converters contains function to convert proto data to model data.
package converters

import (
	"encoding/json"
	"fmt"

	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/models/binary_files"
)

// PBBinaryFileToBinaryFile converts proto binary file to model binary file.
func PBBinaryFileToBinaryFile(uid string, pbf *pb.BinaryFile) (*binary_files.Files, error) {
	// unmarshall meta if present
	if pbf.Meta != nil {
		var meta = make(map[string]string)
		err := json.Unmarshal(pbf.Meta, &meta)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling binary file meta: %w", err)
		}
		return &binary_files.Files{UserID: uid, Name: pbf.Name, File: pbf.File, UpdatedAt: pbf.UpdatedAt.AsTime(), Meta: meta}, nil
	}
	return &binary_files.Files{UserID: uid, Name: pbf.Name, File: pbf.File, UpdatedAt: pbf.UpdatedAt.AsTime()}, nil
}

// BinaryFileToPBBinaryFile converts model binary file to proto binary file.
func BinaryFileToPBBinaryFile(bf *binary_files.Files) (*pb.BinaryFile, error) {
	// marshall meta if present
	if bf.Meta != nil {
		meta, err := json.Marshal(bf.Meta)
		if err != nil {
			return nil, fmt.Errorf("error marshalling binary file meta: %w", err)
		}
		return &pb.BinaryFile{Name: bf.Name, File: bf.File, UpdatedAt: timestamppb.New(bf.UpdatedAt), Meta: meta}, nil
	}
	return &pb.BinaryFile{Name: bf.Name, File: bf.File, UpdatedAt: timestamppb.New(bf.UpdatedAt)}, nil
}

// BinaryFilesToPBBinaryFiles converts model binary files to proto binary files.
func BinaryFilesToPBBinaryFiles(bfs []*binary_files.Files) ([]*pb.BinaryFile, error) {
	var protoBFs []*pb.BinaryFile

	for _, bf := range bfs {
		protoBF, err := BinaryFileToPBBinaryFile(bf)
		if err != nil {
			return nil, err
		}
		protoBFs = append(protoBFs, protoBF)
	}
	return protoBFs, nil
}

// PBBinaryFilesToBinaryFiles converts proto binary files to model binary files.
func PBBinaryFilesToBinaryFiles(uid string, protoBFs []*pb.BinaryFile) ([]*binary_files.Files, error) {
	var bfs []*binary_files.Files
	for _, protoBF := range protoBFs {
		bf, err := PBBinaryFileToBinaryFile(uid, protoBF)
		if err != nil {
			return nil, err
		}
		bfs = append(bfs, bf)
	}
	return bfs, nil
}
