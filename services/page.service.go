package services

import "github.com/deepak/module_page/models"

type PageService interface {
	AddPage(*models.Pages) error
	//GetPages(*string) (*models.Pages, error)
	GetAllPages() ([]*models.Pages, error)
	// UpdateUser(*models.User) error
	// DeleteUser(*string) error
}
