package service




func (s *Service) GetId() {
	s.alloc.Mu.Lock()
	defer s.alloc.Mu.Unlock()
}
