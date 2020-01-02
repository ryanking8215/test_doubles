package fake

type User struct {
	Email string
}

type Account struct {
	passwordHash string
}

func NewAccount(passwordHash string) *Account {
	return &Account{
		passwordHash: passwordHash,
	}
}

func (a Account) PasswordHash() string {
	return a.passwordHash
}

type AccountRepository interface {
	GetPasswordHash(user *User) string
}

type AccountService struct {
	repo AccountRepository
}

func NewAccountService(repo AccountRepository) *AccountService {
	return &AccountService{
		repo: repo,
	}
}

func (s *AccountService) GetPasswordHash(user *User) string {
	return s.repo.GetPasswordHash(user)
}
