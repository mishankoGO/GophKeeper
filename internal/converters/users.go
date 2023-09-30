// Package converters contains function to convert proto data to model data.
package converters

import (
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
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
