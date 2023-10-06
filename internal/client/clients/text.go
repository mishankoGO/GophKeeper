// Package clients contains clients for all the operations.
// It has:
// CredentialsClient for registration.
// UsersClient for login in.
// CardsClient to operate with bank cards.
// LogPassesClient to operate with log passes.
// TextsClient to operate with texts.
// BinaryFilesClient to operate with binary files.
package clients

import (
	"bytes"
	"context"
	"github.com/mishankoGO/GophKeeper/internal/client/interfaces"
	"github.com/mishankoGO/GophKeeper/internal/converters"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/security"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TextsClient contains text client service.
type TextsClient struct {
	service  pb.TextsClient
	Security *security.Security
	repo     interfaces.Repository
	offline  bool
}

// NewTextsClient creates new Texts client.
func NewTextsClient(cc *grpc.ClientConn, repo interfaces.Repository, security *security.Security) *TextsClient {
	if cc != nil {
		service := pb.NewTextsClient(cc)
		return &TextsClient{service: service, Security: security, repo: repo, offline: false}
	}
	return &TextsClient{repo: repo, Security: security, offline: true}
}

// Insert method inserts new Texts.
func (c *TextsClient) Insert(ctx context.Context, req *pb.InsertTextRequest) (*pb.InsertResponse, error) {
	t, err := converters.PBTextToText(req.GetUser().GetUserId(), req.GetText())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting proto text to text: %v", err)
	}

	// encrypt data
	var buf bytes.Buffer
	buf.Write(t.Text)

	encData := c.Security.EncryptData(buf)

	// set encrypted text to Text
	t.Text = encData

	err = c.repo.InsertT(t)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error inserting text: %v", err)
	}

	if !c.offline {
		req.Text.Text = encData
		resp, err := c.service.Insert(ctx, req)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error inserting text: %v", err)
		}
		return resp, nil
	}
	return &pb.InsertResponse{IsInserted: true}, nil
}

// Get method retrieves text information.
func (c *TextsClient) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetTextResponse, error) {
	if c.offline {
		t, err := c.repo.GetT(req.GetName())
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error retrieving text: %v", err)
		}

		// decrypt data
		decData, err := c.Security.DecryptData(t.Text)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error decrypting data: %v", err)
		}

		// set decrypted text to Text
		t.Text = bytes.Trim(decData, "\"\n")

		protoT, err := converters.TextToPBText(t)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error converting model text to proto text: %v", err)
		}

		return &pb.GetTextResponse{Text: protoT}, nil
	}
	resp, err := c.service.Get(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting text information: %v", err)
	}

	// decrypt data
	tt := resp.Text.Text
	decData, err := c.Security.DecryptData(tt)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error decrypting data: %v", err)
	}

	// set decrypted text to Text
	resp.Text.Text = bytes.Trim(decData, "\"\n")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error retrieving text: %v", err)
	}
	return resp, nil
}

// Update method updates text information.
func (c *TextsClient) Update(ctx context.Context, req *pb.UpdateTextRequest) (*pb.UpdateTextResponse, error) {
	mText, err := converters.PBTextToText(req.GetUser().GetUserId(), req.GetText())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting proto text to model text: %v", err)
	}

	// encrypt data
	var buf bytes.Buffer
	buf.Write(mText.Text)

	encData := c.Security.EncryptData(buf)

	// set encrypted text as Text
	mText.Text = encData

	_, err = c.repo.UpdateT(mText)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error updating text: %v", err)
	}

	if !c.offline {
		req.Text.Text = encData
		resp, err := c.service.Update(ctx, req)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error updating text information: %v", err)
		}
		return resp, nil
	}

	return &pb.UpdateTextResponse{Text: req.GetText()}, nil
}

// Delete method deletes text.
func (c *TextsClient) Delete(ctx context.Context, req *pb.DeleteTextRequest) (*pb.DeleteResponse, error) {
	err := c.repo.DeleteT(req.GetName())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error deleting text: %v", err)
	}

	if !c.offline {
		resp, err := c.service.Delete(ctx, req)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error deleting text information: %v", err)
		}

		return resp, nil
	}
	return &pb.DeleteResponse{Ok: true}, nil
}

// List method to list all texts.
func (c *TextsClient) List(ctx context.Context) (*pb.ListTextResponse, error) {
	ts, err := c.repo.ListT()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error listing texts: %v", err)
	}

	pbTs := make([]*pb.Text, len(*ts))
	for _, lp := range *ts {
		pbT, err := converters.TextToPBText(&lp)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error converting text: %v", err)
		}
		// decrypt data
		decData, err := c.Security.DecryptData(pbT.Text)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error decrypting text: %v", err)
		}

		pbT.Text = bytes.Trim(decData, "\"\n")
		pbTs = append(pbTs, pbT)
	}
	resp := &pb.ListTextResponse{Texts: pbTs}
	return resp, err
}
