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

type RepositoryLinsk struct{}

func (r *RepositoryLinsk) SaveHistory(link models.Links) error {
	if link.ID == 0 {
		DBarr.LastID++
		link.ID = DBarr.LastID
	}

	DBarr.DB = append(DBarr.DB, link)
	return nil
}

func (r *RepositoryLinsk) GetSavedLinks() ([]models.Links, error) {
	return DBarr.DB, nil
}

func (r *RepositoryLinsk) GetStatusLinks(links models.Links) ([]models.Links, error) {
	r.SaveHistory(links)
	return DBarr.DB, nil
}
