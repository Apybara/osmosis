package apybara_indexer

import (
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type BlockData struct {
	gorm.Model
	BlockHeight int64  `json:"block_height" gorm:"index:idx_block_height"`
	Type        string `json:"type" gorm:"index:idx_type"`
	Denom       string `json:"denom" gorm:"index:idx_denom"`
	Amount      string `json:"amount"`
}

type TotalReward struct {
	gorm.Model
	EventType   string `json:"event_type" gorm:"index:idx_event_type"`
	BlockHeight int64  `json:"block_height" gorm:"index:idx_block_height"`
	Amount      string `json:"amount"`
	Denom       string `json:"denom"`
}

type RewardDataDelta struct {
	gorm.Model
	AfterBeginBlockerAmount  string `json:"after_begin_blocker_amount"`
	BeforeBeginBlockerAmount string `json:"before_begin_blocker_amount"`
	Denom                    string `json:"denom"`
	BlockHeight              int64  `json:"block_height"`
	Delta                    string `json:"delta"`
	SumDelta500kb            string `json:"sum_delta_500kb"`
	SumDelta75kb             string `json:"sum_delta_75kb"`
	AnnualizeReward500kb     string `json:"annualize_reward_500kb"`
	Timestamp                int64  `json:"timestamp" gorm:"index:idx_timestamp"`
}

type Asset struct {
	gorm.Model
	EventType   string `json:"event_type" gorm:"index:idx_event_type"`
	BlockHeight int64  `json:"block_height" gorm:"index:idx_block_height"`
	Address     string `json:"address" gorm:"index:idx_address"`
	Amount      string `json:"amount"`
	Denom       string `json:"denom"`
}

func ConnectPg(postgresDsn string) (*gorm.DB, error) {
	// use postgres
	var database *gorm.DB
	var err error

	if postgresDsn[:8] == "postgres" {
		database, err = gorm.Open(postgres.Open(postgresDsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	} else {
		database, err = gorm.Open(sqlite.Open(postgresDsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
	}

	if err != nil {
		return nil, err
	}

	database.AutoMigrate(&BlockData{}, &TotalReward{}, &Asset{}, &RewardDataDelta{})

	return database, nil
}
