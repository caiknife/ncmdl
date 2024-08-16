package main

import (
	"os"

	"github.com/caiknife/mp3lister/lib/logger"
	"github.com/caiknife/mp3lister/lib/types"
	"github.com/urfave/cli/v2"
)

func main() {
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
			cookie = dir + "/ncm.txt"
		}
		cookieFile := NewCookieFile(cookie)
		inputLinks.ForEach(func(s string, i int) {
			link, err := NewLink(s, cookieFile)
			if err != nil {
				logger.ConsoleLogger.Errorln(err)
				return
			}
			logger.ConsoleLogger.Infoln(link)
		})
		return nil
	}
}
