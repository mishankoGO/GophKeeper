// Package client contains Client struct to work with all the clients in clients package.
package client

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/mishankoGO/GophKeeper/config"
	"github.com/mishankoGO/GophKeeper/internal/client/clients"
	"github.com/mishankoGO/GophKeeper/internal/client/interceptors"
	"github.com/mishankoGO/GophKeeper/internal/client/interfaces"
	"github.com/mishankoGO/GophKeeper/internal/converters"
	pb "github.com/mishankoGO/GophKeeper/internal/grpc"
	"github.com/mishankoGO/GophKeeper/internal/models/users"
	"github.com/mishankoGO/GophKeeper/internal/security"
)

// Client contains configuration file, and client for different kinds of data.
type Client struct {
	Security          *security.Security         // security client
	connected         bool                       // connected to server
	conf              *config.Config             // config file
	conns             []*grpc.ClientConn         // array of connections
	UsersClient       *clients.UsersClient       // users client
	CardsClient       *clients.CardsClient       // cards client
	TextsClient       *clients.TextsClient       // texts client
	BinaryFilesClient *clients.BinaryFilesClient // binary files client
	LogPassesClient   *clients.LogPassesClient   // log passes client
}

// NewClient function create new Client instance.
func NewClient(conf *config.Config, repo interfaces.Repository) (*Client, error) {
	// parse port
	port := ":" + strings.Split(conf.Address, ":")[1]

	// ping connection
	// if no connection turn offline regime
	connected := ping(port)
	if !connected {
		log.Println("you are offline")

		// connect users client
		usersClient := clients.NewUsersClient(nil, repo)

		// connect cards client
		cardsClient := clients.NewCardsClient(nil, repo)

		// connect texts client
		textsClient := clients.NewTextsClient(nil, repo)

		// connect binary files client
		bfClient := clients.NewBinaryFilesClient(nil, repo)

		// connect log pass client
		lpClient := clients.NewLogPassesClient(nil, repo)

		return &Client{
			connected:         connected,
			UsersClient:       usersClient,
			CardsClient:       cardsClient,
			TextsClient:       textsClient,
			BinaryFilesClient: bfClient,
			LogPassesClient:   lpClient}, nil
	} else {
		// connect to server
		conn1, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return nil, fmt.Errorf("error making connection: %w", err)
		}

		// connect users client
		usersClient := clients.NewUsersClient(conn1, repo)

		// create auth interceptor
		interceptor, err := interceptors.NewAuthInterceptor(usersClient)
		if err != nil {
			return nil, fmt.Errorf("error creating interceptor: %w", err)
		}

		// connect to server
		conn2, err := grpc.Dial(
			port,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithUnaryInterceptor(interceptor.Unary()),
		)
		if err != nil {
			return nil, fmt.Errorf("error making connection: %w", err)
		}

		// connect cards client
		cardsClient := clients.NewCardsClient(conn2, repo)

		// connect texts client
		textsClient := clients.NewTextsClient(conn2, repo)

		// connect binary files client
		bfClient := clients.NewBinaryFilesClient(conn2, repo)

		// connect log pass client
		lpClient := clients.NewLogPassesClient(conn2, repo)

		// create connections array
		conns := []*grpc.ClientConn{conn1, conn2}

		// create client
		var client = &Client{
			UsersClient:       usersClient,
			CardsClient:       cardsClient,
			TextsClient:       textsClient,
			BinaryFilesClient: bfClient,
			LogPassesClient:   lpClient,
			conns:             conns,
			connected:         connected,
		}

		return client, nil
	}
}

// Close method closes connections.
func (c *Client) Close() {
	for _, conn := range c.conns {
		conn.Close()
	}
	c.UsersClient.Close()
}

// Sync method syncs data between server and client.
func (c *Client) Sync(user *users.User) error {

	if c.connected {
		log.Println("syncing data...")
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		if user != nil {
			protoUser := converters.UserToPBUser(user)

			reqBF := &pb.ListBinaryFileRequest{User: protoUser}
			err := c.BinaryFilesClient.Sync(ctx, reqBF)
			if err != nil {
				return fmt.Errorf("error syncing binary files: %w", err)
			}

			reqC := &pb.ListCardRequest{User: protoUser}
			err = c.CardsClient.Sync(ctx, reqC)
			if err != nil {
				return fmt.Errorf("error syncing cards: %w", err)
			}

			reqLP := &pb.ListLogPassRequest{User: protoUser}
			err = c.LogPassesClient.Sync(ctx, reqLP)
			if err != nil {
				return fmt.Errorf("error syncing log passes: %w", err)
			}

			reqT := &pb.ListTextRequest{User: protoUser}
			err = c.TextsClient.Sync(ctx, reqT)
			if err != nil {
				return fmt.Errorf("error syncing texts: %w", err)
			}
		}

	}
	log.Println("data synced successfully!")

	return nil
}

// ping function checks if client is connected to server.
func ping(address string) bool {
	_, err := net.DialTimeout("tcp", address, 2*time.Second)
	if err != nil {
		return false
	}
	return true
}

// SetSecurity method sets security attribute after login.
func (c *Client) SetSecurity(security *security.Security) {
	c.CardsClient.SetSecurity(security)
	c.TextsClient.SetSecurity(security)
	c.BinaryFilesClient.SetSecurity(security)
	c.LogPassesClient.SetSecurity(security)
}
