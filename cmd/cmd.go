package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
	"sync"
	"j7go/components"
	"j7go/modules"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "j7go",
	Short: "j7go is a test go framework",
	Long:  `j7go is a test go framework`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		components.E, err = components.NewEngine(cfgFile)
		if err != nil {
			panic(fmt.Errorf("new engine err %s", err))
		}

		components.L.Info("Using config file:", zap.String(`path`, components.E.Opts.Config.ConfigFileUsed()))

		//注册 模块 路由
		modules.RegisterModules(components.E) //register server

		waiter := sync.WaitGroup{}
		for range components.E.Server {
			waiter.Add(1)
		}
		go components.E.Run(&waiter)	//启动服务//注册etcd
		waiter.Wait()
		//fmt.Println(os.Getpid())

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(fmt.Errorf("new engine err %s", err))
	}
}

func init() {
	cobra.OnInitialize(initCfgFile)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./conf/{environment}/config.yml)")

}

// initConfig reads in config file and ENV variables if set.
func initCfgFile() {
	if cfgFile == "" {
		env := os.Getenv("RUNTIME_ENV")
		if env == "" {
			env = "default"
		}
		cfgFile = fmt.Sprintf("./conf/%s/config.yml", env)
	}
}
