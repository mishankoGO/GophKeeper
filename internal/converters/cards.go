package converters

import (
	"encoding/json"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/models/cards"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func PBCardToCard(uid string, pbc *pb.Card) (*cards.Cards, error) {
	if pbc.Meta != nil {
		var meta = make(map[string]string)
		err := json.Unmarshal(pbc.GetMeta(), &meta)
		if err != nil {
			return nil, status.Error(codes.Internal, "error unmarshalling card meta")
		}
		return &cards.Cards{
			UserID:    uid,
			Name:      pbc.GetName(),
			Card:      pbc.GetCard(),
			UpdatedAt: pbc.GetUpdatedAt().AsTime(),
			Meta:      meta}, nil
	}
	return &cards.Cards{
		UserID:    uid,
		Name:      pbc.GetName(),
		Card:      pbc.GetCard(),
		UpdatedAt: pbc.GetUpdatedAt().AsTime(),
	}, nil
}

func CardToPBCard(c *cards.Cards) (*pb.Card, error) {
	if c.Meta != nil {
		meta, err := json.Marshal(c.Meta)
		if err != nil {
			return nil, status.Error(codes.Internal, "error marshalling card meta")
		}
		return &pb.Card{
			Name:      c.Name,
			Card:      c.Card,
			UpdatedAt: timestamppb.New(c.UpdatedAt),
			Meta:      meta}, nil
	}
	return &pb.Card{
		Name:      c.Name,
		Card:      c.Card,
		UpdatedAt: timestamppb.New(c.UpdatedAt),
	}, nil
}
