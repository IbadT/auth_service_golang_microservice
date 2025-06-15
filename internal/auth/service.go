package auth

type Service interface {
	Login() error
	Register() error
	RefreshToken() error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) Login() error {
	return s.Login()
}

func (s *service) Register() error {
	return s.Register()
}

func (s *service) RefreshToken() error {
	return s.RefreshToken()
}
