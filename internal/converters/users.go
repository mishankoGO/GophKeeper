// Package converters contains function to convert proto data to model data.
package converters

import (
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/models/texts"
	"github.com/mishankoGO/GophKeeper/internal/models/users"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// UserToPBUser converts model user to proto user.
func UserToPBUser(u *users.User) *pb.User {
	return &pb.User{UserId: u.UserID, Login: u.Login, CreatedAt: timestamppb.New(u.CreatedAt)}
}

// PBUserToUser converts proto user to model user.
func PBUserToUser(pbu *pb.User) *users.User {
	return &users.User{UserID: pbu.UserId, Login: pbu.Login, CreatedAt: pbu.CreatedAt.AsTime()}
}

// TextsToPBTexts converts model texts to proto texts.
func TextsToPBTexts(ts []*texts.Texts) ([]*pb.Text, error) {
	var protoTs []*pb.Text

	for _, t := range ts {
		protoT, err := TextToPBText(t)
		if err != nil {
			return nil, err
		}
		protoTs = append(protoTs, protoT)
	}
	return protoTs, nil
}

// PBTextsToTexts converts proto texts to model texts.
func PBTextsToTexts(uid string, protoTs []*pb.Text) ([]*texts.Texts, error) {
	var ts []*texts.Texts
	for _, protoT := range protoTs {
		t, err := PBTextToText(uid, protoT)
		if err != nil {
			return nil, err
		}
		ts = append(ts, t)
	}
	return ts, nil
}
