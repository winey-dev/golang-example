package repo

import (
	v1userpb "github.com/yiaw/grpc-example/protos/v1/user"
	v2userpb "github.com/yiaw/grpc-example/protos/v2/user"
)

type UserRepo struct {
	UserId             string
	Password           string
	Labelkey           string
	Labelvalue         string
	PasswordExpiretime string
	Maxuser            int32
	Locallogin         bool
}

var User map[string]*UserRepo

func init() {
	User = make(map[string]*UserRepo)
}

func MapperV1User(u *v1userpb.UserProto) *UserRepo {
	return &UserRepo{
		UserId:             u.UserId,
		Password:           u.Password,
		Labelkey:           u.Labelkey,
		Labelvalue:         u.Labelvalue,
		PasswordExpiretime: u.PasswordExpiretime,
		Maxuser:            u.Maxuser,
		Locallogin:         u.Locallogin,
	}
}

func MapperV2User(u *v2userpb.UserProto) *UserRepo {
	return &UserRepo{
		UserId:             u.UserId,
		Password:           u.Password,
		Labelkey:           u.Labelkey,
		Labelvalue:         u.Labelvalue,
		PasswordExpiretime: u.PasswordExpiretime,
		Maxuser:            u.Maxuser,
		Locallogin:         u.Locallogin,
	}
}

func ConvertV1User(u *UserRepo) *v1userpb.UserProto {
	return &v1userpb.UserProto{
		UserId:             u.UserId,
		Password:           u.Password,
		Labelkey:           u.Labelkey,
		Labelvalue:         u.Labelvalue,
		PasswordExpiretime: u.PasswordExpiretime,
		Maxuser:            u.Maxuser,
		Locallogin:         u.Locallogin,
	}
}

func ConvertV2User(u *UserRepo) *v2userpb.UserProto {
	return &v2userpb.UserProto{
		UserId:             u.UserId,
		Password:           u.Password,
		Labelkey:           u.Labelkey,
		Labelvalue:         u.Labelvalue,
		PasswordExpiretime: u.PasswordExpiretime,
		Maxuser:            u.Maxuser,
		Locallogin:         u.Locallogin,
	}
}
