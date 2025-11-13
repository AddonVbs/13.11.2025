package hendler

import (
	"modul/internal/ServiceLinks"
	"net/http"
)

type ServiceLinskHendler struct {
	service ServiceLinks.LinksServiceInterface
}

func NewLinksServiHendler(s ServiceLinks.LinksServiceInterface) *ServiceLinskHendler {
	return &ServiceLinskHendler{service: s}
}

func (h *ServiceLinskHendler) GetLinkByID(w http.ResponseWriter, r *http.Request) {
	panic("ff")
}

func (h *ServiceLinskHendler) GetLink(w http.ResponseWriter, r *http.Request) {
	panic("ff")
}
