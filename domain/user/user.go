package user

type User struct {
	ID   UserID
	Name Username
}

func NewUser(name string, id *string) (*User, error) {
	userID, err := NewUserID(id)
	if err != nil {
		return nil, err
	}

	username, err := NewUsername(name)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:   *userID,
		Name: *username,
	}, nil
}
