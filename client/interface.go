package client

import (
	pactus "github.com/pactus-project/pactus/www/grpc/gen/go"
)

type IClient interface {
	GetBlockchainInfo() (*pactus.GetBlockchainInfoResponse, error)
	GetBlockchainHeight() (uint32, error)
	LastBlockTime() (uint32, uint32, error)
	GetNetworkInfo() (*pactus.GetNetworkInfoResponse, error)
	IsValidator(string) (bool, error)
	GetValidatorInfo(string) (*pactus.GetValidatorResponse, error)
	GetValidatorInfoByNumber(int32) (*pactus.GetValidatorResponse, error)
	Close() error
}