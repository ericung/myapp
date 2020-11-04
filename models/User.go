package models

type User struct {
	id    int32
	name  string
	email string
}

func (user User) User(id int32, name string, email string) User {
	user.id = id
	user.name = name
	user.email = email
	return user
}

func (user User) GetId() int32 {
	return user.id
}

func (user User) SetId(id int32) User {
	user.id = id
	return user
}

func (user User) GetName() string {
	return user.name
}

func (user User) SetName(name string) User {
	user.name = name
	return user
}

func (user User) GetEmail() string {
	return user.email
}

func (user User) SetEmail(email string) User {
	user.email = email
	return user
}
