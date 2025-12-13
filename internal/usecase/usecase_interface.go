package usecase

type ListProductUseCaseInterface interface {
	Execute() ([]ProductDTO, error)
}

type GetProductUseCaseInterface interface {
	Execute(input ProductInputDTO) (*ProductDTO, error)
}
