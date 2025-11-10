package ServiceLinks

import (
	"modul/internal/DBarr"
	"modul/internal/models"
)

type LinksRepository interface {
	SaveHestory(link models.Links) error
	GetSaveLinsk() ([]models.Links, error)
	GetStatusLinks(models.Links) ([]models.Links, error)
}

type RepositoryService struct{}

func (r *RepositoryService) SaveHistory(link models.Links) error {
	DBarr.DB = append(DBarr.DB, link)
	return nil
}

func (r *RepositoryService) GetSavedLinks() ([]models.Links, error) {
	return DBarr.DB, nil
}

func (r *RepositoryService) GetStatusLinks(links models.Links) ([]models.Links, error) {
	r.SaveHistory(links)
	return DBarr.DB, nil
}
