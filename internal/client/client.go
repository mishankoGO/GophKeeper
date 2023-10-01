// Package client contains Client struct to work with all the clients in clients package.
package client

import (
	"fmt"
	"github.com/mishankoGO/GophKeeper/config"
	"github.com/mishankoGO/GophKeeper/internal/client/clients"
	"github.com/mishankoGO/GophKeeper/internal/client/interceptors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strings"
)

// Client contains configuration file, and client for different kinds of data.
type Client struct {
	conf              *config.Config             // config file
	conns             []*grpc.ClientConn         // array of connections
	UsersClient       *clients.UsersClient       // users client
	CardsClient       *clients.CardsClient       // cards client
	TextsClient       *clients.TextsClient       // texts client
	BinaryFilesClient *clients.BinaryFilesClient // binary files client
	LogPassesClient   *clients.LogPassesClient   // log passes client
}

// NewClient function create new Client instance.
func NewClient(conf *config.Config) (*Client, error) {
	// parse port
	port := ":" + strings.Split(conf.Address, ":")[1]

	// connect to server
	conn1, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("error making connection: %w", err)
	}

	// connect users client
	usersClient := clients.NewUsersClient(conn1)

	// create auth interceptor
	interceptor, err := interceptors.NewAuthInterceptor(usersClient)
	if err != nil {
		return nil, err
	}

	// connect to server
	conn2, err := grpc.Dial(
		port,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(interceptor.Unary()),
	)
	if err != nil {
		return nil, err
	}

	// connect cards client
	cardsClient := clients.NewCardsClient(conn2)

	// connect texts client
	textsClient := clients.NewTextsClient(conn2)

	// connect binary files client
	bfClient := clients.NewBinaryFilesClient(conn2)

	// connect log pass client
	lpClient := clients.NewLogPassesClient(conn2)

	// create connections array
	conns := []*grpc.ClientConn{conn1, conn2}

	return &Client{
		UsersClient:       usersClient,
		CardsClient:       cardsClient,
		TextsClient:       textsClient,
		BinaryFilesClient: bfClient,
		LogPassesClient:   lpClient,
		conns:             conns}, nil
}

func (c *Client) Close() {
	for _, conn := range c.conns {
		conn.Close()
	}
}
