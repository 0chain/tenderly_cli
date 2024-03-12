package common

import (
	"fmt"

	"github.com/0chain/tenderly_cli/internal/logging"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	ErrValueRetrieval         = errors.New("err happened during value retrieval")
	ErrInvalidArgumentsAmount = errors.New("given amount of arguments is not enough")
	ErrUnknownArgumentValue   = errors.New("err happened because of unknown argument value")
)

const (
	OptionHost   = "host"
	OptionWallet = "wallet"
	OptionToken  = "token"
	OptionAmount = "amount"
)

type Option struct {
	name                          string
	value                         interface{}
	typename, usage, missingError string
	required                      bool
}

// CreateWalletOption creates wallet command handler option.
func CreateWalletOption() *Option {
	return &Option{
		name:         OptionWallet,
		value:        "",
		typename:     "string",
		usage:        "Ethereum wallet address",
		missingError: "Ethereum wallet address should be set",
		required:     true,
	}
}

// CreateTokenOption creates token command handler option.
func CreateTokenOption() *Option {
	return &Option{
		name:         OptionToken,
		value:        float64(0),
		typename:     "float64",
		usage:        "token address of the ERC20 smart contract",
		missingError: "token address of the ERC20 smart contract where balance should be set",
		required:     true,
	}
}

// CreateAmountOption creates amount command handler option.
func CreateAmountOption() *Option {
	return &Option{
		name:         OptionAmount,
		value:        int64(0),
		typename:     "int64",
		usage:        "amount of tokens to be set",
		missingError: "amount of tokens to be set for the previously given resource",
		required:     true,
	}
}

type Arg struct {
	typeName  string
	fieldName string
	value     interface{}
}

func GetHost(args []*Arg) string {
	return getString(args, OptionHost)
}

func GetWallet(args []*Arg) string {
	return getString(args, OptionWallet)
}

func GetToken(args []*Arg) string {
	return getString(args, OptionToken)
}

func GetAmount(args []*Arg) uint64 {
	return uint64(getInt64(args, OptionAmount))
}

func getString(args []*Arg, name string) string {
	if len(args) == 0 {
		logging.Logger.Fatal(ErrInvalidArgumentsAmount.Error())
	}

	for _, arg := range args {
		if arg.fieldName == name {
			return (arg.value).(string)
		}
	}

	logging.Logger.Fatal(errors.Wrap(ErrValueRetrieval, name).Error())

	return ""
}

func getFloat64(args []*Arg, name string) float64 {
	if len(args) == 0 {
		logging.Logger.Fatal(ErrInvalidArgumentsAmount.Error())
	}

	for _, arg := range args {
		if arg.fieldName == name {
			return (arg.value).(float64)
		}
	}

	logging.Logger.Fatal(errors.Wrap(ErrValueRetrieval, name).Error())

	return 0
}

func getInt64(args []*Arg, name string) int64 {
	if len(args) == 0 {
		logging.Logger.Fatal(ErrInvalidArgumentsAmount.Error())
	}

	for _, arg := range args {
		if arg.fieldName == name {
			return (arg.value).(int64)
		}
	}

	logging.Logger.Fatal(errors.Wrap(ErrValueRetrieval, name).Error())

	return 0
}

// CreateCommand provides factory access for command initialization.
func CreateCommand(use, short, long string, functor func(...*Arg), opts ...*Option) *cobra.Command {
	opts = append(opts, &Option{
		name:         "host",
		value:        "",
		typename:     "string",
		usage:        "JSON-RPC url for the selected provider",
		missingError: "JSON-RPC url for the selected provider not specified",
		required:     true,
	})

	command := createCommandHandler(
		use,
		short,
		long,
		func(parameters ...*Arg) {
			functor(parameters...)
		},
		opts)

	appendOptions(command, opts...)

	return command
}

// createCommandHandler creates raw command handler.
func createCommandHandler(use, short, long string, functor func(...*Arg), opts []*Option) *cobra.Command {
	var cobraCommand = &cobra.Command{
		Use:   use,
		Short: short,
		Long:  long,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			fflags := cmd.Flags()

			var parameters []*Arg

			for _, opt := range opts {
				if !fflags.Changed(opt.name) && opt.required {
					logging.Logger.Fatal(opt.missingError)
				}

				var arg *Arg
				switch opt.typename {
				case "string":
					optValue, err := fflags.GetString(opt.name)
					if err != nil {
						logging.Logger.Fatal(err.Error())
					}
					arg = &Arg{
						typeName:  opt.typename,
						fieldName: opt.name,
						value:     optValue,
					}
				case "int64":
					optValue, err := fflags.GetInt64(opt.name)
					if err != nil {
						logging.Logger.Fatal(err.Error())
					}
					arg = &Arg{
						typeName:  opt.typename,
						fieldName: opt.name,
						value:     optValue,
					}
				case "float64":
					optValue, err := fflags.GetFloat64(opt.name)
					if err != nil {
						logging.Logger.Fatal(err.Error())
					}
					arg = &Arg{
						typeName:  opt.typename,
						fieldName: opt.name,
						value:     optValue,
					}
				case "int":
					optValue, err := fflags.GetInt(opt.name)
					if err != nil {
						logging.Logger.Fatal(err.Error())
					}
					arg = &Arg{
						typeName:  opt.typename,
						fieldName: opt.name,
						value:     optValue,
					}
				default:
					logging.Logger.Fatal(errors.Wrap(ErrUnknownArgumentValue, fmt.Sprintf("%v", opt.value)).Error())
				}

				parameters = append(parameters, arg)
			}

			functor(parameters...)
		},
	}

	return cobraCommand
}

// appendOptions appends options to the given command handler.
func appendOptions(command *cobra.Command, opts ...*Option) {
	for _, opt := range opts {
		switch opt.typename {
		case "string":
			command.PersistentFlags().String(opt.name, opt.value.(string), opt.usage)
		case "int64":
			command.PersistentFlags().Int64(opt.name, opt.value.(int64), opt.usage)
		case "float64":
			command.PersistentFlags().Float64(opt.name, opt.value.(float64), opt.usage)
		case "int":
			command.PersistentFlags().Int(opt.name, opt.value.(int), opt.usage)
		}

		if opt.required {
			_ = command.MarkFlagRequired(opt.name)
		}
	}
}
