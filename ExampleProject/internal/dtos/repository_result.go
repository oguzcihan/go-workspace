package dtos

type RepositoryResult struct {
	//interface olduğu onun type i bilinmeze düşer
	Result interface{}
	Error  error
}
