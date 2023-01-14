package simple

type FooRepository struct {
}

type FooService struct {
	*FooRepository
}

func NewFooRepository() *FooRepository {
	return &FooRepository{}
}

func NewFooService(fooRepository *FooRepository) *FooService {
	return &FooService{
		FooRepository: fooRepository,
	}
}
