package zealy

import (
	"github.com/pagu-project/Pagu/internal/engine/command"
	"github.com/pagu-project/Pagu/internal/entity"
	"github.com/pagu-project/Pagu/internal/repository"
	"github.com/pagu-project/Pagu/pkg/wallet"
)

const (
	CommandName       = "zealy"
	ClaimCommandName  = "claim"
	StatusCommandName = "status"
	HelpCommandName   = "help"
)

type Zealy struct {
	db     *repository.DB
	wallet *wallet.Wallet
}

func NewZealy(db *repository.DB, wallet *wallet.Wallet) Zealy {
	return Zealy{
		db:     db,
		wallet: wallet,
	}
}

func (z *Zealy) GetCommand() command.Command {
	subCmdClaim := command.Command{
		Name: ClaimCommandName,
		Desc: "Claim your Zealy Reward",
		Help: "",
		Args: []command.Args{
			{
				Name:     "address",
				Desc:     "Your Pactus address",
				Optional: false,
			},
		},
		SubCommands: nil,
		AppIDs:      entity.AllAppIDs(),
		Handler:     z.claimHandler,
	}

	subCmdStatus := command.Command{
		Name:        StatusCommandName,
		Desc:        "Status of Zealy reward claims",
		Help:        "",
		Args:        nil,
		SubCommands: nil,
		AppIDs:      entity.AllAppIDs(),
		Handler:     z.statusHandler,
	}

	cmdZealy := command.Command{
		Name:        CommandName,
		Desc:        "Zealy Commands",
		Help:        "",
		Args:        nil,
		AppIDs:      entity.AllAppIDs(),
		SubCommands: make([]command.Command, 0),
		Handler:     nil,
		TargetFlag:  command.TargetMaskMain,
	}

	cmdZealy.AddSubCommand(subCmdClaim)
	cmdZealy.AddSubCommand(subCmdStatus)
	return cmdZealy
}
