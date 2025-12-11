package usecase

import "fmt"

type ProductInputDTO struct {
	ID string `json:"id"`
}

type GetProductUseCase struct {
	ProductRepository ProductRepositoryInterface
}

func NewGetProductUseCase(productRepo ProductRepositoryInterface) *GetProductUseCase {
	return &GetProductUseCase{
		ProductRepository: productRepo,
	}
}

func (p *GetProductUseCase) Execute(input ProductInputDTO) (*ProductDTO, error) {
	product, err := p.ProductRepository.GetProduct(input.ID)
	if err != nil {
		return nil, fmt.Errorf("error when get product %w", err)
	}

	return &ProductDTO{
		ID: product.ID,
	}, nil
}
