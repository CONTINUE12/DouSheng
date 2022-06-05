package service

import (
	"douyin/repository"
	"errors"
)

type UserFlow struct {
	name     string
	password string

	uid int64
}

func NewUserFlow(name, password string) *UserFlow {
	return &UserFlow{
		name:     name,
		password: password,
	}
}

// ==================================================用户注册功能模块=====================================================

func Register(name, password string) (int64, error) {
	return NewUserFlow(name, password).Register()
}

func (f *UserFlow) Register() (int64, error) {
	if err := f.CreateUser(); err != nil {
		return 0, errors.New("create user failed")
	}
	return f.uid, nil
}

func (f *UserFlow) CreateUser() error {
	user := &repository.User{
		Name:     f.name,
		Password: f.password,
	}
	if err := repository.NewUserDaoInstance().CreateUser(user); err != nil {
		return err
	}
	f.uid = user.Id
	return nil
}

// ================================================用户登录功能模块=======================================================

func Login(name, string string) (int64, error) {
	return NewUserFlow(name, string).Login()
}

func (f *UserFlow) Login() (int64, error) {
	if err := f.QueryUser(); err != nil {
		return 0, errors.New("query user failed")
	}
	return f.uid, nil
}

func (f *UserFlow) QueryUser() error {
	user := &repository.User{
		Name:     f.name,
		Password: f.password,
	}
	if err := repository.NewUserDaoInstance().QueryUser(user); err != nil {
		return err
	}
	f.uid = user.Id
	return nil
}

// ================================================登录验证模块==========================================================

func Verify(uid int64) error {
	if err := repository.NewUserDaoInstance().QueryUserId(uid); err != nil {
		return err
	}
	return nil
}
