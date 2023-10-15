// Package clients contains clients for all the operations.
// It has:
// CredentialsClient for registration.
// UsersClient for login in.
// CardsClient to operate with bank cards.
// LogPassesClient to operate with log passes.
// TextsClient to operate with texts.
// TextsClient to operate with texts.
package clients

import (
	"bytes"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/mishankoGO/GophKeeper/internal/client/interfaces"
	"github.com/mishankoGO/GophKeeper/internal/converters"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/models/texts"
	"github.com/mishankoGO/GophKeeper/internal/security"
	"github.com/mishankoGO/GophKeeper/pkg/util"
)

// TextsClient contains text client service.
type TextsClient struct {
	service  pb.TextsClient
	Security *security.Security
	repo     interfaces.Repository
	offline  bool
}

// NewTextsClient creates new Texts client.
func NewTextsClient(cc *grpc.ClientConn, repo interfaces.Repository) *TextsClient {
	if cc != nil {
		service := pb.NewTextsClient(cc)
		return &TextsClient{service: service, repo: repo, offline: false}
	}
	return &TextsClient{repo: repo, offline: true}
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
func (c *TextsClient) List(ctx context.Context, req *pb.ListTextRequest) (*pb.ListTextResponse, []*texts.Texts, error) {
	ts, err := c.repo.ListT()
	if err != nil {
		return nil, nil, status.Errorf(codes.Internal, "error listing texts: %v", err)
	}

	resp, err := c.service.List(ctx, req)
	if err != nil {
		return nil, nil, status.Errorf(codes.Internal, "error listing data from server: %v", err)
	}
	return resp, ts, nil
}

// Sync method to sync texts between dbs.
func (c *TextsClient) Sync(ctx context.Context, req *pb.ListTextRequest) error {
	serverTs, clientTs, err := c.List(ctx, req)
	if err != nil {
		return err
	}

	// arrays of names for future syncing
	clientNames := make([]string, len(clientTs))
	serverNames := make([]string, len(serverTs.GetTexts()))

	// flag which shows which db has the latest data.
	// if flag set to "server", it means server has fresher data.
	var dataPrimary string
	if c.offline {
		dataPrimary = "client"
	} else {
		dataPrimary = "server"
	}

	// update cycle
	for _, ct := range clientTs {
		cname := ct.Name
		clientNames = append(clientNames, cname)
		for _, st := range serverTs.GetTexts() {
			sname := st.GetName()
			serverNames = append(serverNames, sname)

			// update common files
			if sname == cname {
				if ct.UpdatedAt.After(st.UpdatedAt.AsTime()) {
					// convert text to proto text
					protoT, err := converters.TextToPBText(ct)
					if err != nil {
						return err
					}

					// update server text
					reqS := &pb.UpdateTextRequest{User: req.GetUser(), Text: protoT}
					_, err = c.service.Update(ctx, reqS)
					if err != nil {
						return err
					}
					dataPrimary = "client"
				} else {
					// convert proto text to text
					t, err := converters.PBTextToText(req.GetUser().GetUserId(), st)
					if err != nil {
						return err
					}

					// update client text
					_, err = c.repo.UpdateT(t)
					if err != nil {
						return err
					}
				}
			}
		}
	}

	if dataPrimary == "server" {
		// insert missing server texts to client
		for _, st := range serverTs.GetTexts() {
			// convert proto text to model text
			t, err := converters.PBTextToText(req.GetUser().GetUserId(), st)
			if err != nil {
				return err
			}

			if !util.StringInSlice(st.Name, clientNames) {
				// insert missing text to client db
				err = c.repo.InsertT(t)
				if err != nil {
					return err
				}
			}
		}

		// delete files
		for _, ct := range clientTs {
			if !util.StringInSlice(ct.Name, serverNames) {
				// delete texts absent in server
				err = c.repo.DeleteT(ct.Name)
				if err != nil {
					return err
				}
			}
		}
	} else if dataPrimary == "client" {
		// insert missing client texts to text
		for _, ct := range clientTs {
			// convert model text to proto text
			protoT, err := converters.TextToPBText(ct)
			if err != nil {
				return err
			}

			if !util.StringInSlice(ct.Name, serverNames) {
				// insert missing text to server db
				reqS := &pb.InsertTextRequest{User: req.GetUser(), Text: protoT}
				_, err = c.service.Insert(ctx, reqS)
				if err != nil {
					return err
				}
			}
		}

		// delete files
		for _, st := range serverTs.GetTexts() {
			if !util.StringInSlice(st.GetName(), clientNames) {
				// delete texts absent in client
				resD := &pb.DeleteTextRequest{User: req.GetUser(), Name: st.GetName()}
				_, err = c.service.Delete(ctx, resD)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// SetSecurity method to set security attribute.
func (c *TextsClient) SetSecurity(security *security.Security) {
	c.Security = security
}
