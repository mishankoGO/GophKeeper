package converters

import (
	"encoding/json"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/models/log_passes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func PBLogPassToLogPass(uid string, pblp *pb.LogPass) (*log_passes.LogPasses, error) {
	if pblp.Meta != nil {
		var meta = make(map[string]string)
		err := json.Unmarshal(pblp.GetMeta(), &meta)
		if err != nil {
			return nil, status.Error(codes.Internal, "error unmarshalling log pass meta")
		}
		return &log_passes.LogPasses{
			UserID:       uid,
			Name:         pblp.GetName(),
			HashLogin:    pblp.GetLogin(),
			HashPassword: pblp.GetPass(),
			UpdatedAt:    pblp.GetUpdatedAt().AsTime(),
			Meta:         meta}, nil
	}
	return &log_passes.LogPasses{
		UserID:       uid,
		Name:         pblp.GetName(),
		HashLogin:    pblp.GetLogin(),
		HashPassword: pblp.GetPass(),
		UpdatedAt:    pblp.GetUpdatedAt().AsTime(),
	}, nil
}

func LogPassToPBLogPass(lp *log_passes.LogPasses) (*pb.LogPass, error) {
	if lp.Meta != nil {
		meta, err := json.Marshal(lp.Meta)
		if err != nil {
			return nil, status.Error(codes.Internal, "error marshalling log pass meta")
		}
		return &pb.LogPass{
			Name:      lp.Name,
			Login:     lp.HashLogin,
			Pass:      lp.HashPassword,
			UpdatedAt: timestamppb.New(lp.UpdatedAt),
			Meta:      meta}, nil
	}
	return &pb.LogPass{
		Name:      lp.Name,
		Login:     lp.HashLogin,
		Pass:      lp.HashPassword,
		UpdatedAt: timestamppb.New(lp.UpdatedAt),
	}, nil
}
