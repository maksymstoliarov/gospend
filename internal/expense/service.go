package expense

type Service struct {
	storage *Storage
}

func NewService(s *Storage) *Service {
	return &Service{
		storage: s,
	}
}
