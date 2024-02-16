package service

import "mongo_peg_1/repository"

type LinkExtractService struct {
	PageRepository repository.PageRepository
}

func (l LinkExtractService) Listen() {

}
