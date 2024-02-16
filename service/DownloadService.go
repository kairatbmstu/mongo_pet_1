package service

import "mongo_peg_1/repository"

type DownloadService struct {
	PageRepository repository.PageRepository
}

func (d DownloadService) Listen() {

}
