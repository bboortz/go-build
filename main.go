package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/bboortz/go-utils"
	"github.com/bboortz/go-utils/logger"
	"github.com/bboortz/go-utils/stringutil"
	"github.com/urfave/cli"
)

var defaultConfigFile string
var log logger.Logger

type testConfig struct {
	IgnorePackages []string
}

type programConfig struct {
	AppName    string
	AppVersion string
	PkgName    string
	LogLevel   string
	Arch       string
	Test       testConfig
}

func init() {
	defaultConfigFile = "build.toml"
	log = logger.NewLogger().Build()
}

func loadConfig() programConfig {
	log.Debug("Loading configfile: ", defaultConfigFile)
	var config programConfig
	if _, err := toml.DecodeFile(defaultConfigFile, &config); err != nil {
		log.Fatal("Cannot load config file:", defaultConfigFile)
	}
	stringutil.CheckEmpty("AppName", config.AppName)
	stringutil.CheckEmpty("AppVersion", config.AppVersion)
	stringutil.CheckEmpty("PkgName", config.PkgName)

	log.Debug("Successully loaded.")
	return config
}

func buildContainer(imageName, dockerfile string) {
	log.Debug("Building Container: " + imageName)
	log.Debug(" -> using dockerfile: " + dockerfile)
	cmd := fmt.Sprintf("docker build -t %s -f %s .", imageName, dockerfile)
	log.Debug(" -> build command: " + cmd)

	_, _, _ = utils.ExecCommand(cmd)
	log.Debug("successully built.")
}

func cmdBuildContainer(appName string, pkgName string) {
	log.Info("CMD: Building Container: " + appName)
	buildContainer(appName+"-build", "dockerfiles/Dockerfile.build")
	log.Info("CMD successully executed.")
}

func cmdBuildBuildContainer(appName string, pkgName string) {
	log.Info("CMD: Building the Build Container ...")
	err := os.Chdir("./dockerimage")
	if err != nil {
		log.Fatal(err)
	}
	buildContainer("go-build-base", "Dockerfile")
	err = os.Chdir("..")
	if err != nil {
		log.Fatal(err)
	}
	log.Info("CMD successully executed.")
}

func cmdBuildApplication(appName string, pkgName string, arch string) {
	log.Info("CMD: Building Application: " + appName)
	currentDir, _ := os.Getwd()
	const cmdStr string = "docker run -u %s:%s -v %s:%s -w %s -e APP=%s -e ARCH=%s go-build-base /build/build.sh"
	cmd := fmt.Sprintf(cmdStr, "1000", "1000", currentDir, "/go/src/"+pkgName, "/go/src/"+pkgName, appName, arch)
	log.Debug(" -> build command: " + cmd)

	_, _, _ = utils.ExecCommand(cmd)
	log.Info("CMD successully executed.")
}

func cmdTestApplication(appName string, pkgName string, arch string, ignorePackages []string) {
	log.Info("CMD: Testing Application: " + appName)
	currentDir, _ := os.Getwd()
	ignorePackagesStr := strings.Join(ignorePackages, "|")
	const cmdStr string = "docker run -u %s:%s -v %s:%s -w %s -e APP=%s -e ARCH=%s -e 'IGNORE_PACKAGES=%s' go-build-base /build/test.sh"
	cmd := fmt.Sprintf(cmdStr, "1000", "1000", currentDir, "/go/src/"+pkgName, "/go/src/"+pkgName, appName, arch, ignorePackagesStr)
	log.Debug(" -> build command: " + cmd)

	_, _, _ = utils.ExecCommand(cmd)
	log.Info("CMD successully executed.")
}

func main() {
	// load config
	config := loadConfig()
	if config.LogLevel != "" {
		log.SetLevelWithStr(config.LogLevel)
	}
	if config.Arch == "" {
		config.Arch = runtime.GOARCH
	}
	log.Debug("Building Architecture: " + config.Arch)

	// command line parameters
	app := cli.NewApp()
	app.Name = "go-build"
	app.Version = "1"
	app.Usage = "build tool for golang applications"
	app.EnableBashCompletion = true

	app.Commands = []cli.Command{
		{
			Name:    "create",
			Aliases: []string{"c"},
			Usage:   "create a new application",
			Action: func(c *cli.Context) error {
				log.Fatal("Not implemented yet")
				return nil
			},
		},
		{
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "build the project",
			Action: func(c *cli.Context) error {
				cmdBuildApplication(config.AppName, config.PkgName, config.Arch)
				return nil
			},
			Subcommands: []cli.Command{
				{
					Name:    "application",
					Aliases: []string{"a"},
					Usage:   "build the application",
					Action: func(c *cli.Context) error {
						cmdBuildApplication(config.AppName, config.PkgName, config.Arch)
						return nil
					},
				},
				{
					Name:    "container",
					Aliases: []string{"c"},
					Usage:   "build a container which contains the application",
					Action: func(c *cli.Context) error {
						log.Fatal("Not implemented yet")
						cmdBuildContainer(config.AppName, config.PkgName)
						return nil
					},
				},
				{
					Name:    "build-container",
					Aliases: []string{"b"},
					Usage:   "build the build-container for go-build",
					Action: func(c *cli.Context) error {
						cmdBuildBuildContainer(config.AppName, config.PkgName)
						return nil
					},
				},
			},
		},
		{
			Name:    "test",
			Aliases: []string{"t"},
			Usage:   "test the project",
			Action: func(c *cli.Context) error {
				cmdTestApplication(config.AppName, config.PkgName, config.Arch, config.Test.IgnorePackages)
				return nil
			},
		},
	}

	_ = app.Run(os.Args)
}
