package fake

// FakeAccountRepository implements AccountRepository
var _ AccountRepository = (*FakeAccountRepository)(nil)

type FakeAccountRepository struct {
	accounts map[string]*Account // key is user's unique email
}

func NewFakeAccountRepository() *FakeAccountRepository {
	return &FakeAccountRepository{
		accounts: map[string]*Account{
			"john@bmail.com": NewAccount("111"),
			"boby@bmail.com": NewAccount("222"),
		},
	}
}

func (r *FakeAccountRepository) GetPasswordHash(user *User) string {
	a := r.accounts[user.Email]
	return a.PasswordHash()
}
