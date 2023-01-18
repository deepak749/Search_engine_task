package services

import (
	"context"
	"errors"

	"github.com/deepak/module_page/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PageServiceImpl struct {
	pagescollection *mongo.Collection
	ctx             context.Context
}

func NewPageService(pagescollection *mongo.Collection, ctx context.Context) PageService {
	return &PageServiceImpl{
		pagescollection: pagescollection,
		ctx:             ctx,
	}
}

func (u *PageServiceImpl) AddPage(page *models.Pages) error {
	_, err := u.pagescollection.InsertOne(u.ctx, page)
	return err
}

//func (u *PageServiceImpl) GetPages(name *string) {
//var page *models.Pages
//query := bson.D{bson.E{Key: "name", Value: name}}
//err := u.pagescollection.FindOne(u.ctx, nil).Decode(&page)
//return nil
//}

func (u *PageServiceImpl) GetAllPages() ([]*models.Pages, error) {
	var pages []*models.Pages
	cursor, err := u.pagescollection.Find(u.ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(u.ctx) {
		var page models.Pages
		err := cursor.Decode(&page)
		if err != nil {
			return nil, err
		}
		pages = append(pages, &page)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(u.ctx)

	if len(pages) == 0 {
		return nil, errors.New("documents not found")
	}
	return pages, nil
}
