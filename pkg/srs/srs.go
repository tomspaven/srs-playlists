package srs

type EaseFactor float32

type SRSConfig struct {
	Algorithm string
	SRSDBConfig
}

type SRSDBConfig struct {
	DBType string
	DBName string
}

type SRS struct {
	Cfg SRSConfig
}

/*func (s *SRS) GetCardsDueNow() ([]*Card, error) {
	db, err := gorm.Open(s.Cfg.DBType, s.Cfg.DBName)
	if err != nil {
		return nil, fmt.Errorf("Unable to open database")
	}
	defer db.Close()
	return nil, nil
}*/

/*func (s *SRS) PersistCards(card []*Card) error {

}*/
