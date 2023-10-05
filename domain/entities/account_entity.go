package entities

import (
	"golang-cc/domain/types"
	"time"
)

// AccountEntity defines an
// Account's fields and its
// value retrieval methods.
type AccountEntity struct {
	id        types.AccountID
	balance   types.AccountBalance
	name      types.AccountName
	secret    types.AccountSecret
	createdAt types.AccountCreatedAt
}

func (acc AccountEntity) ID() types.AccountID {
	return acc.id
}

func (acc AccountEntity) Balance() types.AccountBalance {
	return acc.balance
}

func (acc AccountEntity) Name() types.AccountName {
	return acc.name
}

func (acc AccountEntity) Secret() types.AccountSecret {
	return acc.secret
}

func (acc AccountEntity) CreatedAt() types.AccountCreatedAt {
	return acc.createdAt
}

// NewAccount returns a newly created account entity setting its creation time.
func NewAccount(id types.AccountID, balance types.AccountBalance,
	name types.AccountName, secret types.AccountSecret) AccountEntity {
	return AccountEntity{
		id:        id,
		balance:   balance,
		name:      name,
		secret:    secret,
		createdAt: types.AccountCreatedAt(time.Now()),
	}
}
