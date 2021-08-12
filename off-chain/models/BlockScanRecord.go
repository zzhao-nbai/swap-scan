package models

import (
	"payment-bridge/off-chain/common/utils"
	"payment-bridge/off-chain/database"
)

type BlockScanRecord struct {
	ID                     int64  `json:"id"`
	LastCurrentBlockNumber int64  `json:"last_current_block_number"`
	UpdateAt               string `json:"update_at"`
}

func (self *BlockScanRecord) FindLastCurrentBlockNumber() ([]*BlockScanRecord, error) {
	db := database.GetDB()
	var models []*BlockScanRecord
	err := db.Find(&models).Error
	return models, err
}

//condition :&models.BlockScanRecord{"last_current_block_number": 18000}
//updateFields: map[string]interface{}{"update_at": taskT.ProcessingTime, "last_current_block_number": 18000}
func UpdateBlockScanRecord(whereCondition interface{}, updateFields interface{}) (BlockScanRecord, error) {
	db := database.GetDB()
	var record BlockScanRecord
	utils.GetEpochInMillis()
	err := db.Model(&record).Where("").Update(updateFields).Error
	return record, err
}