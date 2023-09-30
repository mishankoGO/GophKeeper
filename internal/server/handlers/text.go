// Package handlers contains servers interfaces.
// The list of servers:
//     Users, Credentials, BinaryFiles, Cards, Texts, LogPasses
package handlers

import (
	"bytes"
	"context"
	"encoding/json"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/mishankoGO/GophKeeper/internal/converters"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/security"
	"github.com/mishankoGO/GophKeeper/internal/server/interfaces"
)

// NewTexts creates new instance of text server.
func NewTexts(repo interfaces.Repository, security *security.Security) *Texts {
	return &Texts{
		Repo:     repo,
		Security: *security,
	}
}

// Texts contains repository and security instances.
type Texts struct {
	pb.UnimplementedTextsServer
	Repo     interfaces.Repository // repository
	Security security.Security     // security
}

// Insert method inserts text to db.
func (t *Texts) Insert(ctx context.Context, req *pb.InsertTextRequest) (*pb.InsertResponse, error) {
	// convert proto text to model text
	text, err := converters.PBTextToText(req.User.UserId, req.Text)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting text: %v", err)
	}

	// get text
	text_ := text.Text

	// create encoder
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	res := &pb.InsertResponse{IsInserted: false}

	// marshall into bytes
	err = encoder.Encode(string(text_))
	if err != nil {
		return res, status.Errorf(codes.Internal, "error encoding text: %v", err)
	}

	// encrypt data
	encData := t.Security.EncryptData(buf)

	// set encrypted text as Text
	text.Text = encData

	// insert new text to db
	err = t.Repo.InsertT(ctx, text)
	if err != nil {
		return res, status.Errorf(codes.Internal, "error inserting text: %v", err)
	}

	// set status
	res.IsInserted = true

	return res, nil
}

// Get method retrieves text from db.
func (t *Texts) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetTextResponse, error) {
	// convert proto user to model user
	user := converters.PBUserToUser(req.GetUser())
	name := req.GetName()

	// get text from database
	text, err := t.Repo.GetT(ctx, user.UserID, name)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error getting text: %v", err)
	}

	// decrypt text
	decData, err := t.Security.DecryptData(text.Text)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error decrypting text: %v", err)
	}

	// set decrypted text to Text
	text.Text = bytes.Trim(decData, "\"\n")

	// convert text to proto text
	pbText, err := converters.TextToPBText(text)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting text: %v", err)
	}

	// create response
	res := &pb.GetTextResponse{Text: pbText}

	return res, nil
}

// Update method updates text in db.
func (t *Texts) Update(ctx context.Context, req *pb.UpdateTextRequest) (*pb.UpdateTextResponse, error) {
	// convert proto user to model user
	user := converters.PBUserToUser(req.GetUser())

	// convert proto text to model text
	text, err := converters.PBTextToText(user.UserID, req.GetText())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting text: %v", err)
	}

	// get text
	text_ := text.Text

	// create encoder
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)

	// marshall into bytes
	err = encoder.Encode(string(text_))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error encoding text: %v", err)
	}

	// encrypt data
	encData := t.Security.EncryptData(buf)

	// set encrypted file as File
	text.Text = encData

	// update record in db
	updatedText, err := t.Repo.UpdateT(ctx, text)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error updating text: %v", err)
	}

	// convert model text to proto text
	pbText, err := converters.TextToPBText(updatedText)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error converting text: %v", err)
	}

	// create response
	res := &pb.UpdateTextResponse{Text: pbText}

	return res, nil
}

// Delete method deletes text from db.
func (t *Texts) Delete(ctx context.Context, req *pb.DeleteTextRequest) (*pb.DeleteResponse, error) {
	// convert proto user to user and get name
	user := converters.PBUserToUser(req.GetUser())
	name := req.GetName()

	// create response
	res := &pb.DeleteResponse{Ok: false}

	// delete record
	err := t.Repo.DeleteT(ctx, user.UserID, name)
	if err != nil {
		return res, status.Errorf(codes.Internal, "error deleting text: %v", err)
	}

	// set result
	res.Ok = true

	return res, nil
}
