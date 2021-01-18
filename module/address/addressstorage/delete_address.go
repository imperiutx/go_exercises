package addressstorage

func (s *sqlStore) DeleteAddress(id int) error {
	db := s.db

	if err := db.Table("addresses").
		Where("id = ?", id).
		Update("status", 0).Error; err != nil {
		return err
	}

	return nil
}
