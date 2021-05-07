package grpc

import (
	"context"
	pbhighscore "github.com/Namanl2001/Fun-Game/m-highscore/api/v1"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
)

type Grpc struct {
	address string
	srv     *grpc.Server
}

var HighScore = 9999999999.0

func NewServer(address string) *Grpc {
	return &Grpc{
		address: address,
	}
}

func (g *Grpc) SetHighScore(ctx context.Context, input *pbhighscore.SetHighScoreRequest) (*pbhighscore.SetHighScoreResponse, error) {
	log.Info().Msg("SetHighScore in m-highscore is called")
	HighScore = input.HighScore
	return &pbhighscore.SetHighScoreResponse{
		Set: true,
	}, nil
}

func (g *Grpc) GetHighScore(ctx context.Context, input *pbhighscore.GetHighScoreRequest) (*pbhighscore.GetHighScoreResponse, error) {
	log.Info().Msg("GetHighScore in m-highscore is called")
	return &pbhighscore.GetHighScoreResponse{
		HighScore: HighScore,
	}, nil
}

func (g *Grpc) ListenAndServe() error {
	lis, err := net.Listen("tcp", g.address)
	if err != nil {
		return errors.Wrap(err, "failed to open tcp port")
	}

	serverOpts := []grpc.ServerOption{}

	g.srv = grpc.NewServer(serverOpts...)

	pbhighscore.RegisterGameServer(g.srv, g)

	log.Info().Str("address", g.address).Msg("starting gRPC server for m-highscore microservice")

	err = g.srv.Serve(lis)
	if err != nil {
		return errors.Wrap(err, "failed to start gRPC server for m-highscore microservice")
	}
	return nil
}
