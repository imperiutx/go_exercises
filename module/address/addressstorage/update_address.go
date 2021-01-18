package addressstorage

import (
	"context"

	"github.com/imperiustx/go_excercises/common"
	"github.com/imperiustx/go_excercises/module/address/addressmodel"
)

func (s *sqlStore) UpdateAddress(ctx context.Context, id int, data *addressmodel.AddressUpdate) error {
	db := s.db

	if err := db.Table(data.TableName()).
		Where("id = ?", id).
		Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
