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

// AccoutListUC holds the repository access
// interface and the actual accounts list method.
type AccountListUC struct {
	accountListRepository AccountListRepository
}

// NewAccountListUC returns a concrete account list use case.
func NewAccountListUC(accountListRepository AccountListRepository) AccountListUC {
	return AccountListUC{
		accountListRepository: accountListRepository,
	}
}

// ListAccounts implements the list accounts use case.
func (uc AccountListUC) ListAccounts() ([]entities.AccountEntity, error) {
	accs, err := uc.accountListRepository.List()
	if err != nil {
		return nil, err
	}
	return accs, nil
}
