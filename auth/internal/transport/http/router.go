package http

func (s *Server) InitRouter() {
	s.HTTP.GET("/test", s.handler.Test)
}
