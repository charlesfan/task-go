package app

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/charlesfan/task-go/entity/config"
	"github.com/charlesfan/task-go/repository/store"
)

var server = newServer()

type Server struct {
	router *Router
}

func (s *Server) Run(c config.Config) error {
	store.Init(c)
	s.router = NewRouter(net.JoinHostPort(c.Server.Host, c.Server.Port))
	s.router.Config()
	s.router.Run()
	return nil
}

func newServer() *Server {
	return &Server{}
}

func NewServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "server",
		Long:         `Here is task-go backend server.`,
		SilenceUsage: true,
		Run: func(cmd *cobra.Command, args []string) {
			config.Init()
			c := config.New()
			server.Run(c)
		},
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}

	cmd.Flags().StringP("config", "c", "", "config file (default is $HOME/.cobra.yaml)")

	viper.BindPFlags(cmd.Flags())

	return cmd
}
