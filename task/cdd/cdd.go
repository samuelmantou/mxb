package cdd

import (
	"github.com/chromedp/chromedp"
)

func Task() chromedp.Tasks {
	return []chromedp.Action{
		chromedp.Click(`#mms-main > div.top-data-panel > div:nth-child(1) > p.top-data-panel__card__content`),
	}
}
