package ur

import (
	"context"
	"errors"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/kanade0404/rental-house-scraper/internal/infrastructures/chromedp/initialize"
	"github.com/kanade0404/rental-house-scraper/internal/infrastructures/chromedp/run"
)

type InputListRentalHouses struct {
	trainStation map[string]string
	area         map[string]string
}
type SearchRentalHousesOutput struct {
	houses []interface{}
}

func SearchRentalHouses(ctx context.Context, input InputListRentalHouses) (*SearchRentalHousesOutput, error) {
	ctx, timeoutCancel, allocatorCancel, contextCancel := initialize.UR(ctx)
	defer timeoutCancel()
	defer allocatorCancel()
	defer contextCancel()
	var houses []*cdp.Node
	if err := run.UR(ctx, chromedp.Tasks{
		chromedp.Navigate("https://suumo.jp/"),
		chromedp.Nodes("body > div.wrap_main > div.wrapper > section > article > div.js-module_searchs_property.module_searchs_property.js-bukken-key", &houses, chromedp.ByQueryAll),
	}); err != nil {
		return nil, err
	}
	// TODO: housesをパースする
	// TODO: roomsをパースする
	return nil, nil
}

func parseHouse(ctx context.Context, house *cdp.Node) {
	var (
		houseName string
		access    string
		address   string
		rooms     []*cdp.Node
	)
	opts := []chromedp.QueryOption{
		chromedp.ByQuery,
		chromedp.FromNode(house),
	}
	if err := run.UR(ctx, chromedp.Tasks{
		chromedp.Text("div.searchs_property_head > div > div.cassettes_property_contents > div.item_upper > div > div > h2 > a > span.rep_bukken-name", &houseName, opts...),
		chromedp.Text("div.searchs_property_head > div > div.cassettes_property_contents > div.item_upper > div > ul > li:nth-child(1) > div > div.item_maininfolist_body > div.item_maininfolist_text.rep_bukken-access", &access, opts...),
		chromedp.Text("div.searchs_property_head > div > div.cassettes_property_contents > div.item_upper > div > ul > li:nth-child(2) > div > div.item_maininfolist_body > div", &address, opts...),
		chromedp.Nodes("section > div.item_body.module_tables_room.module_tables_property.js-no-room-hidden > table > tbody > tr", &rooms, chromedp.ByQueryAll, chromedp.FromNode(house)),
	}); err != nil {
		panic(err)
	}
	var roomLinks []string
	for i := range rooms {
		roomLink, err := extractRoomLink(ctx, rooms[i])
		if err != nil {
			continue
		}
		roomLinks = append(roomLinks, roomLink)
	}

}
func extractRoomLink(ctx context.Context, room *cdp.Node) (string, error) {
	if len(room.Attributes) != 1 {
		return "", errors.New("room node is invalid")
	}
	var roomLink []*cdp.Node
	if err := run.UR(ctx, chromedp.Tasks{
		chromedp.Nodes("td:nth-child(6) > ul > li.mgt10 > a", &roomLink, chromedp.ByQuery, chromedp.FromNode(room)),
	}); err != nil {
		return "", err
	}
	if len(roomLink) != 1 {
		return "", fmt.Errorf("room link node must be 1. got %d", len(roomLink))

	}
	if href, ok := roomLink[0].Attribute("href"); ok {
		return href, nil
	}
	return "", errors.New("room link node must have href attribute")
}
