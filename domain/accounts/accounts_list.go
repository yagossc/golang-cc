package accounts

import (
	"golang-cc/domain/entities"
)

// AccountRepoLister provides the
// interface/access between the use
// case layer and the repository layer.
type AccountListRepository interface {
	List() ([]entities.AccountEntity, error)
}

// AccoutListUC defines the list
// method access for this use case.
type AccountListUC struct {
	accountListRepository AccountListRepository
}

// NewAccountListUC returns a concrete Account List use case.
func NewAccountListUC(accountListRepository AccountListRepository) AccountListUC {
	return AccountListUC{
		accountListRepository: accountListRepository,
	}
}

// ListAccounts implements the use case
// interface method 'ListAccounts()'.
func (uc AccountListUC) ListAccounts() ([]entities.AccountEntity, error) {
	accs, err := uc.accountListRepository.List()
	if err != nil {
		return nil, err
	}
	return accs, nil
}
