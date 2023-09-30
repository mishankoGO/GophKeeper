package converters

import (
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/models/users"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func UserToPBUser(u *users.User) *pb.User {
	return &pb.User{UserId: u.UserID, Login: u.Login, CreatedAt: timestamppb.New(u.CreatedAt)}
}

func PBUserToUser(pbu *pb.User) *users.User {
	return &users.User{UserID: pbu.UserId, Login: pbu.Login, CreatedAt: pbu.CreatedAt.AsTime()}
}
