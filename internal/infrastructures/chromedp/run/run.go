package run

import (
	"context"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

// UR
/*
UR賃貸のスクレイピングを行う
*/
func UR(ctx context.Context, actions []chromedp.Action) error {
	return run(ctx, append(chromedp.Tasks{chromedp.Tasks{network.SetBlockedURLS([]string{
		// ロードに時間がかかるので広告系のJSのリクエストをブロックする
		"https://www.googletagmanager.com/gtm.js?id=*",
	})}}, actions...))
}

// run
/*
chromedpでスクレイピングを行う
*/
func run(ctx context.Context, actions []chromedp.Action) error {
	return chromedp.Run(ctx, actions...)
}
