package grpc

import (
	"context"
	pbgameengine "github.com/Namanl2001/Fun-Game/m-game-engine/api/v1"
	"github.com/Namanl2001/Fun-Game/m-game-engine/internal/server/logic"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
)

type Grpc struct {
	address string
	srv     *grpc.Server
}

func NewServer(address string) *Grpc {
	return &Grpc{
		address: address,
	}
}

func (g *Grpc) GetSize(ctx context.Context, input *pbgameengine.GetSizeRequest) (*pbgameengine.GetSizeResponse, error) {
	log.Info().Msg("GetSize in m-game-engine called")
	return &pbgameengine.GetSizeResponse{
		Size: logic.GetSize(),
	}, nil
}

func (g *Grpc) SetScore(ctx context.Context, input *pbgameengine.SetScoreRequest) (*pbgameengine.SetScoreResponse, error) {
	log.Info().Msg("SetScore in m-game-engine called")
	set := logic.SetScore(input.Score)
	return &pbgameengine.SetScoreResponse{
		Set: set,
	}, nil
}

func (g *Grpc) ListenAndServe() error {
	lis, err := net.Listen("tcp", g.address)
	if err != nil {
		return errors.Wrap(err, "failed to open tcp port")
	}

	serverOpts := []grpc.ServerOption{}

	g.srv = grpc.NewServer(serverOpts...)

	pbgameengine.RegisterGameEngineServer(g.srv, g)

	log.Info().Str("address", g.address).Msg("starting gRPC server for m-game-engine microservice")

	err = g.srv.Serve(lis)
	if err != nil {
		return errors.Wrap(err, "failed to start gRPC server for m-game-engine microservice")
	}
	return nil
}
