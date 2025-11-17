package logic

import (
	"github.com/Dylanps05/Capacity-Scanner/internal/storage"
)

type DefaultController struct{}

func NewDefaultController(s storage.SiteStorage) Controller {
	return &DefaultController{}
}
