package factory

import (
	"github.com/Alfian57/belajar-golang/internal/model"
	"github.com/bluele/factory-go/factory"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
)

func NewUrlVisitorFactory(urlIDs uuid.UUIDs) *factory.Factory {
	if len(urlIDs) == 0 {
		panic("NewUrlFactory: urlIDs cannot be empty")
	}

	return factory.NewFactory(
		&model.URLVisitor{},
	).SeqInt("ID", func(n int) (any, error) {
		return uuid.New(), nil
	}).Attr("UrlID", func(args factory.Args) (any, error) {
		return urlIDs[gofakeit.Number(0, len(urlIDs)-1)], nil
	}).Attr("IpAddress", func(args factory.Args) (any, error) {
		return gofakeit.IPv4Address(), nil
	}).Attr("UserAgent", func(args factory.Args) (any, error) {
		return gofakeit.UserAgent(), nil
	})
}
