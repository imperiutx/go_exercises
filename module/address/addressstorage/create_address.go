package addressstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/address/addressmodel"
)

func (s *sqlStore) CreateAddress(ctx context.Context, data *addressmodel.AddressCreate) error {
	db := s.db

	if err := db.Table(data.TableName()).
		Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
