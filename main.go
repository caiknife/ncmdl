package main

import (
	"os"
	"path/filepath"

	"github.com/caiknife/mp3lister/lib/logger"
	"github.com/caiknife/mp3lister/lib/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var defaultPoolSize = PoolSize

var appLogger *logrus.Logger

func init() {
	appLogger = logger.NewConsoleLogger(
		logger.DisableTimestamp(true),
	)
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			appLogger.Fatalln("程序发生了异常", r)
		}
	}()

	if err := newApp().Run(os.Args); err != nil {
		appLogger.Fatalln(err)
		return
	}
}

func newApp() *cli.App {
	app := &cli.App{
		Name:  "网易云歌曲下载工具",
		Usage: "支持下载歌单、专辑、单曲",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "cookie",
				Aliases: []string{"c"},
				Usage:   "被加载的cookie文件名",
				Value:   "",
			},
			&cli.BoolFlag{
				Name:    "info",
				Aliases: []string{"i"},
				Usage:   "打印下载信息，所有歌曲的相关信息",
				Value:   false,
			},
			&cli.BoolFlag{
				Name:    "tmp",
				Aliases: []string{"t"},
				Usage:   "下载到当前目录下的tmp目录，而不是下载到当前目录",
				Value:   false,
			},
			&cli.IntFlag{
				Name:    "pool",
				Aliases: []string{"p"},
				Usage:   "使用多任务下载，默认的任务池大小",
				Value:   PoolSize,
			},
		},
		Action: action(),
	}

	return app
}

const (
	ErrInputLinksAreEmpty types.Error = "请输入歌单、专辑、单曲链接"
)

func action() cli.ActionFunc {
	return func(c *cli.Context) error {
		inputLinks := types.Slice[string](c.Args().Slice())
		if inputLinks.IsEmpty() {
			return ErrInputLinksAreEmpty
		}

		cookie := c.String("cookie")
		if cookie == "" {
			dir, err := os.UserHomeDir()
			if err != nil {
				err = errors.WithMessage(err, "user home dir")
				return err
			}
			cookie = filepath.Join(dir, "ncm.txt")
		}

		cookieFile := NewCookieFile(cookie)
		info := c.Bool("info")
		tmp := c.Bool("tmp")

		inputLinks.ForEach(func(s string, i int) {
			link, err := NewLink(
				s,
				LinkOptionCookieFile(cookieFile),
				LinkOptionOnlyInfo(info),
				LinkOptionTmp(tmp),
			)
			if err != nil {
				err = errors.WithMessage(err, "new link")
				appLogger.Errorln(err)
				return
			}

			err = link.Download()
			if err != nil {
				err = errors.WithMessage(err, "link download")
				appLogger.Errorln(err)
			}
		})

		return nil
	}
}
