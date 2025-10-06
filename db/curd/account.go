package curd

import (
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Account struct {
	ID      uint `gorm:"primaryKey"`
	Balance int64
}

type Transaction struct {
	ID            uint `gorm:"primaryKey;autoIncrement"`
	FromAccountID uint
	ToAccountID   uint
	Amount        int64
}

var (
	ErrInsufficientBalance = errors.New("余额不足")
)

// Transfer 从 fromID 向 toID 转 amount 分（0.01 元）
func Transfer(db *gorm.DB, fromID, toID uint, amount int64) error {
	if amount <= 0 {
		return errors.New("转账金额必须为正")
	}
	if fromID == toID {
		return errors.New("不能给自己转账")
	}

	return db.Transaction(func(tx *gorm.DB) error {
		var fromAcc, toAcc Account

		// 1. 同时加锁，按 ID 升序拿锁 → 避免死锁
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("id IN ?", []uint{fromID, toID}).Order("id ASC").
			Find(&[]Account{fromAcc, toAcc}).Error; err != nil {
			return err
		}
		// 重新赋值（Find 到切片里，上面只是占位）
		if err := tx.Where("id = ?", fromID).Take(&fromAcc).Error; err != nil {
			return err
		}
		if fromAcc.Balance < amount {
			return ErrInsufficientBalance
		}

		// 2. 扣款（纯 SQL，跳过 Hook/UpdatedAt）
		if err := tx.Model(&Account{}).Where("id = ?", fromID).
			UpdateColumn("balance", gorm.Expr("balance - ?", amount)).Error; err != nil {
			return err
		}

		// 3. 入账
		if err := tx.Model(&Account{}).Where("id = ?", toID).
			UpdateColumn("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
			return err
		}

		// 4. 写流水
		return tx.Create(&Transaction{
			FromAccountID: fromID,
			ToAccountID:   toID,
			Amount:        amount,
		}).Error
	})
}
