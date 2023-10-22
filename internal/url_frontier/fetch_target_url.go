package url_frontier

import (
	"fmt"
	"net/url"
)

const origin = "https://www.ur-net.go.jp"

func UrlFrontier() ([]*url.URL, error) {
	u, err := url.Parse(fmt.Sprintf("%s/chintai/kanto/tokyo/result/?line=10500&line_station=10500_2284&line_station=10500_1685&line_station=10500_2997&line_station=10500_2922&line_station=10500_2862&line_station=10500_1989&line_station=10500_1482&line=10500&line_station=10500_2284&line_station=10500_1685&line_station=10500_2997&line_station=10500_2922&line_station=10500_2862&line_station=10500_1989&line_station=10500_1482&rent_low=&rent_high=&commonfee=1&rent_low=&rent_high=&commonfee=1&room=3K&room=3LDK&room=3K&room=3LDK&walk=10&walk=10&floorspace_low=&floorspace_high=&floorspace_low=&floorspace_high=&years=&years=&popular_floor_2=1&popular_south=1&popular_internet_2=1&popular_floor_2=1&popular_south=1&popular_internet_2=1&facility_internet=2&facility_internet=2&floor=2&floor=2&position_south=1&position_south=1&todofuken=tokyo", origin))
	if err != nil {
		return nil, err
	}
	return []*url.URL{
		u,
	}, nil
}
