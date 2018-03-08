package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/laidingqing/KeepJobs/accounts/model"
	"github.com/laidingqing/KeepJobs/accounts/mongo"
	"github.com/laidingqing/KeepJobs/common/util"
	pb "github.com/laidingqing/KeepJobs/pb"
)

// RPCAccountServer is used to implement user_service.UserServiceServer.
type RPCAccountServer struct{}

var (
	accountManager = mongo.NewAccountManager()
)

// CreateAccount implements account_service.UserServiceServer
func (s *RPCAccountServer) CreateAccount(context context.Context, request *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	fmt.Printf("Receive is %s\n", time.Now())
	rev, err := accountManager.Insert(model.Account{
		UserName: request.Username,
		Password: util.CalculatePassHash(request.Password, request.Username),
		Salt:     request.Username,
	})

	if err != nil {
		return nil, err
	}

	return &pb.CreateAccountResponse{Account: &pb.Account{
		Id:       rev.ID.Hex(),
		Username: rev.UserName,
	}}, nil
}

// GetAccount implements account_service.UserServiceServer
func (s *RPCAccountServer) GetAccount(context context.Context, request *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	log.Printf("Get user account for: %v", request.Id)
	return &pb.GetAccountResponse{Account: &pb.Account{
		Username: "laidingqing",
	}}, nil
}

// GetByUsername implements account_service.UserServiceServer
func (s *RPCAccountServer) GetByUsername(context context.Context, request *pb.GetByUsernameRequest) (*pb.GetByUsernameResponse, error) {
	log.Printf("Get user by username for: %v", request.Username)
	res, err := accountManager.FindByUserName(request.Username)
	if err != nil {
		return &pb.GetByUsernameResponse{}, err
	}
	return &pb.GetByUsernameResponse{Account: &pb.Account{
		Id:       res.ID.Hex(),
		Username: res.UserName,
		Password: res.Password,
	}}, nil
}

//UpdateToken 更新当前会话TOKEN
func (s *RPCAccountServer) UpdateToken(ctx context.Context, in *pb.UpdateSessionTokenRequest) (*pb.UpdateSessionTokenResonse, error) {
	err := accountManager.UpdateCurrentToken(in.Accountid, in.Token)
	if err != nil {
		log.Printf("UpdateCurrentToken Err: %s", err.Error())
		return &pb.UpdateSessionTokenResonse{Updated: false}, err
	}
	return &pb.UpdateSessionTokenResonse{Updated: true}, nil
}

//GetByToken get Account by token
func (s *RPCAccountServer) GetByToken(ctx context.Context, in *pb.GetByTokenRequest) (*pb.GetByTokenResponse, error) {
	acct, err := accountManager.FindAccountByToken(in.Token)
	if err != nil {
		return &pb.GetByTokenResponse{}, err
	}
	return &pb.GetByTokenResponse{
		Account: dencodeAccountInfo(acct)}, nil
}
