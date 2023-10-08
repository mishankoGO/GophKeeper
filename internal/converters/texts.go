// Package converters contains function to convert proto data to model data.
package converters

import (
	"encoding/json"
	"fmt"

	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/models/texts"
)

// PBTextToText converts proto text to model text.
func PBTextToText(uid string, pt *pb.Text) (*texts.Texts, error) {
	// unmarshall meta if present
	if pt.Meta != nil {
		var meta = make(map[string]string)
		err := json.Unmarshal(pt.Meta, &meta)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling text meta: %w", err)
		}
		return &texts.Texts{UserID: uid, Name: pt.Name, Text: pt.Text, UpdatedAt: pt.UpdatedAt.AsTime(), Meta: meta}, nil
	}
	return &texts.Texts{UserID: uid, Name: pt.Name, Text: pt.Text, UpdatedAt: pt.UpdatedAt.AsTime()}, nil
}

// TextToPBText converts model text to proto text.
func TextToPBText(t *texts.Texts) (*pb.Text, error) {
	// marshall meta if present
	if t.Meta != nil {
		meta, err := json.Marshal(t.Meta)
		if err != nil {
			return nil, fmt.Errorf("error marshalling text meta: %w", err)
		}
		return &pb.Text{Name: t.Name, Text: t.Text, UpdatedAt: timestamppb.New(t.UpdatedAt), Meta: meta}, nil
	}
	return &pb.Text{Name: t.Name, Text: t.Text, UpdatedAt: timestamppb.New(t.UpdatedAt)}, nil
}
