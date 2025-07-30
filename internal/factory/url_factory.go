package factory

import (
	"time"

	"github.com/Alfian57/belajar-golang/internal/model"
	"github.com/bluele/factory-go/factory"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
)

func NewUrlFactory(userIDs uuid.UUIDs) *factory.Factory {
	if len(userIDs) == 0 {
		panic("NewUrlFactory: userIDs cannot be empty")
	}

	return factory.NewFactory(
		&model.Url{},
	).SeqInt("ID", func(n int) (any, error) {
		return uuid.New(), nil
	}).Attr("ShortUrl", func(args factory.Args) (any, error) {
		return gofakeit.Word(), nil
	}).Attr("LongUrl", func(args factory.Args) (any, error) {
		return gofakeit.URL(), nil
	}).Attr("UserID", func(args factory.Args) (any, error) {
		return userIDs[gofakeit.Number(0, len(userIDs)-1)], nil
	}).Attr("ExpiredAt", func(args factory.Args) (any, error) {
		return time.Now().Add(24 * 2 * time.Hour), nil
	})
}
