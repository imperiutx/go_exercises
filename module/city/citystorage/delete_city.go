package citystorage

func (s *sqlStore) DeleteCity(id int) error {
	db := s.db

	if err := db.Table("cities").
		Where("id = ?", id).
		Update("status", 0).Error; err != nil {
		return err
	}

	return nil
}