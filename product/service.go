package product

type Service interface {
	GetProductById(param *getProductByIdRequest) (*Product, error)
}

type service struct {
	repo Respository
}

func NewService(repo Respository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetProductById(param *getProductByIdRequest) (*Product, error) {
	return s.repo.GetProductById(param.ProductID)
}
