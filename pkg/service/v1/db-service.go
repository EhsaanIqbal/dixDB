package v1

import (
	"context"

	v1 "github.com/ehsaaniqbal/dixdb/pkg/api/v1"

	m "github.com/ehsaaniqbal/dixdb/pkg/db/memory"
)

type Server struct {
	db m.MemoryStore
}

func NewServer(db m.MemoryStore) *Server {
	return &Server{db: db}
}

func (s *Server) Set(ctx context.Context, req *v1.SetRequest) (*v1.StatusResponse, error) {
	s.db.Set(req.Key, req.Value)
	return &v1.StatusResponse{
		Status: 0,
	}, nil
}

func (s *Server) Get(ctx context.Context, req *v1.GetRequest) (*v1.GetResponse, error) {
	value := s.db.Get(req.Key)
	return &v1.GetResponse{
		Value: value,
	}, nil
}

func (s *Server) Delete(ctx context.Context, req *v1.DeleteRequest) (*v1.StatusResponse, error) {
	s.db.Delete(req.Key)
	return &v1.StatusResponse{
		Status: 0,
	}, nil
}
