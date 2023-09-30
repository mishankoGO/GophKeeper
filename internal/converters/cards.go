// Package converters contains function to convert proto data to model data.
package converters

import (
	"bytes"
	"encoding/json"
	"fmt"

	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/models/cards"
)

// PBCardToCard converts proto card to model card.
func PBCardToCard(uid string, pbc *pb.Card) (*cards.Cards, error) {
	// unmarshall meta if present
	if pbc.Meta != nil && !bytes.Equal(pbc.Meta, []byte("")) {
		var meta = make(map[string]string)
		err := json.Unmarshal(pbc.GetMeta(), &meta)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling card meta: %w", err)
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

// CardToPBCard converts model card to proto card.
func CardToPBCard(c *cards.Cards) (*pb.Card, error) {
	// marshall meta if present
	if c.Meta != nil {
		meta, err := json.Marshal(c.Meta)
		if err != nil {
			return nil, fmt.Errorf("error marshalling card meta: %w", err)
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
