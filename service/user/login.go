package user

import (
	"context"
	"crypto/md5"
	"errors"

	dao "github.com/twelveeee/amis-admin-go/dao/mdb/user"
	pb "github.com/twelveeee/amis-admin-go/pb_gen/user"
)

// func GetUserById(id int64) (*User, error) {
// 	return NewUserModel().GetByUserID(id)
// }

func Login(ctx context.Context, req *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {
	username := req.GetUsername()
	password := req.GetPassword()

	user, err := dao.NewUserModel(ctx).GetByUsername(username)
	if err != nil {
		return nil, err
	}
	hash := md5.New()
	hash.Write([]byte(password))
	hashPassword := hash.Sum(nil)

	if user.Password == password || user.Password == string(hashPassword) {
		return &pb.UserLoginResponse{
			UserId: string(user.UserID),
		}, nil
	}

	return nil, errors.New("login failed")
}
