package voucher

import (
	"github.com/pagu-project/Pagu/internal/engine/command"
	"github.com/pagu-project/Pagu/internal/entity"
	"github.com/pagu-project/Pagu/internal/repository"
	"github.com/pagu-project/Pagu/pkg/client"
	"github.com/pagu-project/Pagu/pkg/wallet"
)

const (
	CommandName      = "voucher"
	ClaimCommandName = "claim"
	HelpCommandName  = "help"
)

type Voucher struct {
	db            *repository.DB
	wallet        *wallet.Wallet
	clientManager *client.Mgr
	targetMask    int
}

func NewVoucher(db *repository.DB, wallet *wallet.Wallet, cli *client.Mgr, target int) Voucher {
	return Voucher{
		db:            db,
		wallet:        wallet,
		clientManager: cli,
		targetMask:    target,
	}
}

func (v *Voucher) GetCommand() command.Command {
	subCmdClaim := command.Command{
		Name: ClaimCommandName,
		Desc: "Claim your voucher coins and bond to validator",
		Help: "",
		Args: []command.Args{
			{
				Name:     "code",
				Desc:     "voucher code",
				Optional: false,
			},
			{
				Name:     "address",
				Desc:     "your pactus validator address",
				Optional: false,
			},
		},
		SubCommands: nil,
		AppIDs:      entity.AllAppIDs(),
		Handler:     v.claimHandler,
	}

	cmdVoucher := command.Command{
		Name:        CommandName,
		Desc:        "Voucher Commands",
		Help:        "",
		Args:        nil,
		AppIDs:      entity.AllAppIDs(),
		SubCommands: make([]command.Command, 0),
		Handler:     nil,
		TargetMask:  v.targetMask,
	}

	cmdVoucher.AddSubCommand(subCmdClaim)
	return cmdVoucher
}
