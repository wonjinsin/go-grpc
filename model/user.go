package model

// ValidateUpdateUser ...
func (u *User) ValidateUpdateUser() bool {
	if u.Uid != "" {
		return false
	}
	return true
}

// UpdateUser ...
func (u *User) UpdateUser(user *User) *User {
	u.Email = user.Email
	u.Nick = user.Nick
	return u
}
