package accounts

import (
	"errors"
	"golang-cc/domain/entities"
	"golang-cc/domain/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

// accListRepoMock mocks a account
// list interface implementation.
type accListRepoMock struct {
	expectedError    error
	expectedAccounts []entities.AccountEntity
}

// List is the method implementation from the repository interface.
func (acc accListRepoMock) List() ([]entities.AccountEntity, error) {
	return acc.expectedAccounts, acc.expectedError
}

func TestListAccounts(t *testing.T) {

	testcases := []struct {
		name             string
		expectsError     bool
		expectedError    error
		expectedAccounts []entities.AccountEntity
	}{
		{
			name:             "Repository Error",
			expectsError:     true,
			expectedError:    errors.New("repo layer error"),
			expectedAccounts: []entities.AccountEntity{},
		},
		{
			name:             "Empty list",
			expectsError:     false,
			expectedError:    nil,
			expectedAccounts: []entities.AccountEntity{},
		},
		{
			name:          "Populated list",
			expectsError:  false,
			expectedError: nil,
			expectedAccounts: []entities.AccountEntity{
				//TODO: how to set an Account?
				entities.NewAccount(types.AccountID(1),
					types.AccountBalance(0.0),
					types.AccountName("teste-1"),
					types.AccountSecret("secret-1")),
				entities.NewAccount(types.AccountID(2),
					types.AccountBalance(0.0),
					types.AccountName("teste-2"),
					types.AccountSecret("secret-2")),
			},
		},
	}

	for i, tc := range testcases {
		accListRepoMock := accListRepoMock{
			expectedError:    tc.expectedError,
			expectedAccounts: tc.expectedAccounts,
		}

		accListUC := NewAccountListUC(accListRepoMock)

		accounts, err := accListUC.ListAccounts()

		if tc.expectsError {
			if err == nil {
				t.Errorf("CASE %d: expected %s, but got nothing", i, tc.expectedError.Error())
			}

			if err.Error() != tc.expectedError.Error() {
				t.Errorf("CASE %d: expected %s, but got %s", i, tc.expectedError.Error(), err.Error())
			}
		}

		if !tc.expectsError && err != nil {
			t.Errorf("CASE %d: expected nothing, but got %s", i, err.Error())
		}

		// TODO: the Account's struct fields are not exported, they can be accessed
		// through each of the methods
		if ok := assert.ElementsMatch(t, accounts, tc.expectedAccounts); !ok {
			t.Errorf("CASE %d: expected %v, but got %v", i, tc.expectedAccounts, accounts)
		}
	}

}
