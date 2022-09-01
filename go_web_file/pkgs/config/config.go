package config

import (
	"errors"
	"filrserver/pkgs/model"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// viperConfig 全局配置变量
var (
	ViperConfig *viper.Viper
	Users       model.Usersc
)

// 初始化配置文件相关设置，在 main 包中调用进行初始化加载
func Init() error {
	log.SetFlags(log.Lshortfile | log.LstdFlags) //设置日志行号
	var configfile = pflag.StringP("configfile", "c", "./config.yaml", "user -c set Your congfile")
	pflag.StringP("logfile", "f", "/var/log/fileserver.log", "user -c set Your congfile")
	pflag.StringP("loglevel", "l", "INFO", "user -l set loglevel")
	pflag.ErrHelp = errors.New("")
	pflag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		pflag.PrintDefaults()
		fmt.Fprint(os.Stderr, "优先级: 命令行 > 配置文件 > 命令行默认值")
	}

	pflag.Parse()
	if pflag.NArg() != 0 {
		fmt.Printf("参数 %v 无法解析,请参考以下语法:\n", pflag.Args())
		pflag.Usage()
		fmt.Print("\n")
		os.Exit(1)
	} //
	if !Exists(*configfile) {
		log.Fatalf("no such config directory: %s", *configfile)
	}
	viper.SetConfigFile(*configfile) // path to look for the config file in
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Fatalf("Fatal error config file: %s \n", err)
	}
	viper.BindPFlags(pflag.CommandLine)
	ViperConfig = viper.GetViper()

	if err = validateConfig(ViperConfig); err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
	}
	err = ViperConfig.Unmarshal(&Users)
	if err != nil { // Handle errors reading the config file
		log.Fatalf("Fatal error config file: %s \n", err)
	}
	return nil
}

func Exists(name string) bool {
	_, err := os.Stat(name)
	if err == nil {
		return true
	}
	//if errors.Is(err, os.ErrNotExist) {
	//	return false
	//}
	return false
}

func validateConfig(v *viper.Viper) error {
	var (
		logdir = filepath.Dir(v.GetString("logfile"))
	)

	if !Exists(logdir) {
		return fmt.Errorf("no such directory: logdir: %s, please check configuration", logdir)
	}

	return nil
}
