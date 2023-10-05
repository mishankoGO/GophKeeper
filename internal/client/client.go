// Package client contains Client struct to work with all the clients in clients package.
package client

import (
	"fmt"
	"github.com/mishankoGO/GophKeeper/config"
	"github.com/mishankoGO/GophKeeper/internal/client/clients"
	"github.com/mishankoGO/GophKeeper/internal/client/interceptors"
	"github.com/mishankoGO/GophKeeper/internal/client/interfaces"
	"github.com/mishankoGO/GophKeeper/internal/security"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"strings"
	"time"
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
func NewClient(conf *config.Config, repo interfaces.Repository, security *security.Security) (*Client, error) {
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
		cardsClient := clients.NewCardsClient(nil, repo, security)

		// connect texts client
		textsClient := clients.NewTextsClient(nil, repo)

		// connect binary files client
		bfClient := clients.NewBinaryFilesClient(nil, repo)

		// connect log pass client
		lpClient := clients.NewLogPassesClient(nil, repo)

		return &Client{
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
		cardsClient := clients.NewCardsClient(conn2, repo, security)

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
			conns:             conns}

		//err = client.sync()

		return client, nil
	}
}

func (c *Client) Close() {
	for _, conn := range c.conns {
		conn.Close()
	}
	c.UsersClient.Close()
}

//func (c *Client) sync() error {
//	cards, err := c.BinaryFilesClient.List()
//	return nil
//}

func ping(address string) bool {
	_, err := net.DialTimeout("tcp", address, 2*time.Second)
	if err != nil {
		return false
	}
	return true
}
