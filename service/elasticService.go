package service

import(
	"context"
	"errors"
	
	"gopkg.in/olivere/elastic.v6"
	"github.com/grupo8hackglobo2019/back-end/model"
)

const(
	indexName    = "chat"
	docType      = "log"
	appName      = "back-end"
)

type ElasticService struct {
	ElasticCLI *elastic.Client
}

func (e *ElasticService) SaveToElastic(ctx context.Context, payload model.Message) error {

	exists, err := e.ElasticCLI.IndexExists(indexName).Do(ctx)
	if err != nil {
		return err
	}

	if !exists {
		res, error := e.ElasticCLI.CreateIndex(indexName).Do(ctx)
		if error != nil {
			return error
		}
		if !res.Acknowledged {
			return errors.New("CreateIndex was not acknowledged. Check that timeout value is correct.")
		}
	}

	_, error := e.ElasticCLI.Index().
		Index(indexName).
		Type(docType).
		BodyJson(payload).
		Do(ctx)

	if error != nil {
		return error
	}

	return nil
}