// Package operatingsystems provides the operating systems functionality for
// the CLI
package operatingsystems

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/vultr/govultr/v3"
	"github.com/vultr/vultr-cli/v3/cmd/utils"
	"github.com/vultr/vultr-cli/v3/pkg/cli"
)

var (
	long    = `OS will retrieve available operating systems that can be deployed`
	example = `
	# Example
	vultr-cli os
	`

	listLong    = `List all operating systems available to be deployed on Vultr`
	listExample = `
	# Full example
	vultr-cli os list
		
	# Full example with paging
	vultr-cli os list --per-page=1 --cursor="bmV4dF9fMTI0" 

	# Shortened with alias commands
	vultr-cli o l
	`
)

// Interface for os
type Interface interface {
	validate(cmd *cobra.Command, args []string)
	List() ([]govultr.OS, *govultr.Meta, error)
}

// Options for os
type Options struct {
	Base *cli.Base
}

// NewOSOptions returns Options struct
func NewOSOptions(base *cli.Base) *Options {
	return &Options{Base: base}
}

// NewCmdOS creates cobra command for OS
func NewCmdOS(base *cli.Base) *cobra.Command {
	o := NewOSOptions(base)

	cmd := &cobra.Command{
		Use:     "os",
		Short:   "list available operating systems",
		Aliases: []string{"o"},
		Long:    long,
		Example: example,
	}

	list := &cobra.Command{
		Use:     "list",
		Short:   "list all available operating systems",
		Aliases: []string{"l"},
		Long:    listLong,
		Example: listExample,
		Run: func(cmd *cobra.Command, args []string) {
			o.validate(cmd, args)
			os, meta, err := o.List()
			data := &OSPrinter{OperatingSystems: os, Meta: meta}

			fmt.Println(o.Base.Printer.Output)
			o.Base.Printer.Display(data, err)
		},
	}

	list.Flags().StringP("cursor", "c", "", "(optional) Cursor for paging.")
	list.Flags().IntP("per-page", "p", 100, "(optional) Number of items requested per page. Default is 100 and Max is 500.")

	cmd.AddCommand(list)
	return cmd
}

func (o *Options) validate(cmd *cobra.Command, args []string) {
	o.Base.Args = args
	o.Base.Options = utils.GetPaging(cmd)
	o.Base.Printer.Output = viper.GetString("output")
}

// List all os
func (o *Options) List() ([]govultr.OS, *govultr.Meta, error) {
	list, meta, _, err := o.Base.Client.OS.List(context.Background(), o.Base.Options)
	if err != nil {
		return nil, nil, err
	}

	return list, meta, nil
}
