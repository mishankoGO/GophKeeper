package converters

import (
	"encoding/json"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/models/texts"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func PBTextToText(uid string, pt *pb.Text) (*texts.Texts, error) {
	if pt.Meta != nil {
		var meta = make(map[string]string)
		err := json.Unmarshal(pt.Meta, &meta)
		if err != nil {
			return nil, status.Error(codes.Internal, "error unmarshalling text meta")
		}
		return &texts.Texts{UserID: uid, Name: pt.Name, Text: pt.Text, UpdatedAt: pt.UpdatedAt.AsTime(), Meta: meta}, nil
	}
	return &texts.Texts{UserID: uid, Name: pt.Name, Text: pt.Text, UpdatedAt: pt.UpdatedAt.AsTime()}, nil
}

func TextToPBText(t *texts.Texts) (*pb.Text, error) {
	if t.Meta != nil {
		meta, err := json.Marshal(t.Meta)
		if err != nil {
			return nil, status.Error(codes.Internal, "error marshalling text meta")
		}
		return &pb.Text{Name: t.Name, Text: t.Text, UpdatedAt: timestamppb.New(t.UpdatedAt), Meta: meta}, nil
	}
	return &pb.Text{Name: t.Name, Text: t.Text, UpdatedAt: timestamppb.New(t.UpdatedAt)}, nil
}
