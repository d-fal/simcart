package server

import "context"

type Server struct {
	ctx context.Context
}

func (s *Server) GetContext() context.Context {
	return s.ctx
}
