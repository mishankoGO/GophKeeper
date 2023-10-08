package converters

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mishankoGO/GophKeeper/internal/models/log_passes"

	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
)

// PBLogPassToLogPass converts proto log pass to model log pass.
func PBLogPassToLogPass(uid string, pblp *pb.LogPass) (*log_passes.LogPasses, error) {
	// unmarshall meta if present
	if pblp.Meta != nil && bytes.Equal(pblp.Meta, []byte("")) {
		var meta = make(map[string]string)
		err := json.Unmarshal(pblp.GetMeta(), &meta)
		if err != nil {
			return nil, fmt.Errorf("error unmarshalling log pass meta: %w", err)
		}
		return &log_passes.LogPasses{
			UserID:    uid,
			Name:      pblp.GetName(),
			Login:     pblp.GetLogin(),
			Password:  pblp.GetPass(),
			UpdatedAt: pblp.GetUpdatedAt().AsTime(),
			Meta:      meta}, nil
	}
	return &log_passes.LogPasses{
		UserID:    uid,
		Name:      pblp.GetName(),
		Login:     pblp.GetLogin(),
		Password:  pblp.GetPass(),
		UpdatedAt: pblp.GetUpdatedAt().AsTime(),
	}, nil
}

// LogPassToPBLogPass converts model log pass to proto log pass.
func LogPassToPBLogPass(lp *log_passes.LogPasses) (*pb.LogPass, error) {
	// marshall meta if present
	if lp.Meta != nil {
		meta, err := json.Marshal(lp.Meta)
		if err != nil {
			return nil, fmt.Errorf("error marshalling log pass meta: %w", err)
		}
		return &pb.LogPass{
			Name:      lp.Name,
			Login:     lp.Login,
			Pass:      lp.Password,
			UpdatedAt: timestamppb.New(lp.UpdatedAt),
			Meta:      meta}, nil
	}
	return &pb.LogPass{
		Name:      lp.Name,
		Login:     lp.Login,
		Pass:      lp.Password,
		UpdatedAt: timestamppb.New(lp.UpdatedAt),
	}, nil
}

// LogPassesToPBLogPasses converts model logpasses to proto logpasses.
func LogPassesToPBLogPasses(lps []*log_passes.LogPasses) ([]*pb.LogPass, error) {
	var protoLPs []*pb.LogPass

	for _, lp := range lps {
		protoLP, err := LogPassToPBLogPass(lp)
		if err != nil {
			return nil, err
		}
		protoLPs = append(protoLPs, protoLP)
	}
	return protoLPs, nil
}

// PBLogPassesToLogPasses converts proto logpasses to model logpasses.
func PBLogPassesToLogPasses(uid string, protoLPs []*pb.LogPass) ([]*log_passes.LogPasses, error) {
	var lps []*log_passes.LogPasses
	for _, protoLP := range protoLPs {
		lp, err := PBLogPassToLogPass(uid, protoLP)
		if err != nil {
			return nil, err
		}
		lps = append(lps, lp)
	}
	return lps, nil
}
