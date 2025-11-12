package ServiceLinks

import (
	"fmt"
	"modul/internal/models"
	"net/http"
	"strings"
	"time"
)

type LinksService interface {
	GetSaveLinskService() ([]models.Links, error)
	GetStatusLinksService(models.Links) ([]models.Links, error)
	GetStatusLinksServiceById(id int) (models.Links, error)
}

type LinksServer struct {
	repo RepositoryLinsk
}

func NewLinksService(repo RepositoryLinsk) *LinksServer {
	return &LinksServer{repo: repo}
}

func (s *LinksServer) GetSaveLinskService() ([]models.Links, error) {
	return s.repo.GetSavedLinks()
}

func isAliveLinsk(URL string) bool {
	if !strings.HasPrefix(URL, "http://") && !strings.HasPrefix(URL, "https://") {
		URL = "http://" + URL
	}

	fmt.Println(URL)
	client := http.Client{Timeout: 5 * time.Second}

	resp, err := client.Get(URL)

	if err != nil {
		fmt.Println(err)
		return false
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println(URL)
		return true
	}

	fmt.Println(resp.StatusCode, URL)

	return false
}

func (s *LinksServer) GetStatusLinksService(lik models.Links) ([]models.Links, error) {
	if isAliveLinsk(lik.URL) {

		lik.Status = "avalibale"
	} else {
		lik.Status = "not avalibale"
	}

	if err := s.repo.SaveHistory(lik); err != nil {
		return nil, err
	}

	return s.repo.GetSavedLinks()
}

func (s *LinksServer) GetStatusLinksServiceById(id int) (models.Links, error) {
	links, err := s.repo.GetSavedLinks()
	if err != nil {
		return models.Links{}, err
	}

	for i, l := range links {
		if l.ID == id {
			if isAliveLinsk(l.URL) {
				links[i].Status = "avalibale"
			} else {
				links[i].Status = "not avalibale"
			}

			if err := s.repo.SaveHistory(links[i]); err != nil {
				return models.Links{}, err
			}
			return links[i], nil
		}
	}
	return models.Links{}, fmt.Errorf("link with ID %s not found", id)

}
