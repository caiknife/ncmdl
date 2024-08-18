package main

import (
	"os"
	"path/filepath"

	"github.com/caiknife/mp3lister/lib/logger"
	"github.com/caiknife/mp3lister/lib/types"
	"github.com/urfave/cli/v2"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			logger.ConsoleLogger.Fatalln("程序发生了异常", r)
		}
	}()

	if err := newApp().Run(os.Args); err != nil {
		logger.ConsoleLogger.Fatalln(err)
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
				Usage:   "Cookie文件名",
				Value:   "",
			},
			&cli.BoolFlag{
				Name:    "dryrun",
				Aliases: []string{"d"},
				Usage:   "演习模式",
				Value:   false,
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
				return err
			}
			cookie = filepath.Join(dir, "ncm.txt")
		}
		cookieFile := NewCookieFile(cookie)
		dryRun := c.Bool("dryrun")
		inputLinks.ForEach(func(s string, i int) {
			link, err := NewLink(s, cookieFile, dryRun)
			if err != nil {
				logger.ConsoleLogger.Errorln(err)
				return
			}
			SetRequestDataCookie(link.CookieFile.ToHttpCookie())
			err = link.Download()
			if err != nil {
				logger.ConsoleLogger.Errorln(err)
			}
		})
		return nil
	}
}
