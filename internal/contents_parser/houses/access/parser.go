package access

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type TimeInMinutes struct {
	minutes        int
	movementMethod string
}
type Access struct {
	trainName   string
	stationName string
	walk        TimeInMinutes
	bus         *TimeInMinutes
}

func (a Access) TrainName() string {
	return a.trainName
}

func (a Access) StationName() string {
	return a.stationName
}

func (a Access) Walk() TimeInMinutes {
	return a.walk
}
func (a Access) Bus() *TimeInMinutes {
	return a.bus
}

func ParseAccess(accessText string) ([]Access, error) {
	accesses := strings.Split(accessText, "\n")
	var parsedAccesses []Access
	for i := range accesses {
		var (
			trainName   string
			stationName string
		)
		// 路線名と駅名、徒歩とバスそれぞれの時間を取得する
		// ex)
		// 西武新宿線「西武柳沢」駅 徒歩11～14分→西武新宿線,西武柳沢,徒歩,11
		// JR中央線「吉祥寺」駅バス11分 徒歩6～8分→JR中央線,吉祥寺,バス,11,徒歩,6
		// JR中央線「武蔵境」駅 徒歩20～22分 またはバス5分 徒歩1～4分→JR中央線,武蔵境,徒歩,20 JR中央線,武蔵境,バス,5,徒歩,1

		accesses := strings.Split(accesses[i], " ")
		if len(accesses) < 2 {
			return nil, fmt.Errorf(fmt.Sprintf("access is invalid. splitted result must be 2, got %d. %s", len(accesses), accesses[i]))
		}
		// 徒歩での所要時間をparse
		if !strings.Contains(accesses[1], "徒歩") {
			return nil, fmt.Errorf("access is invalid. '徒歩' not found. %s", accesses[i])
		}
		walkInMinutes, err := parseWalk(accesses[1])
		if err != nil {
			return nil, err
		}
		walk := TimeInMinutes{
			minutes:        walkInMinutes,
			movementMethod: "徒歩",
		}
		// 路線名と駅名+バスでの所要時間をparse
		trainNameAndOthers := strings.Split(accesses[0], "「")
		if len(trainNameAndOthers) != 2 {
			return nil, fmt.Errorf("access is invalid. splitted result must be 2, got %d. %s", len(trainNameAndOthers), accesses[i])
		}
		trainName = trainNameAndOthers[0]
		stationNameAndOther := strings.Split(trainNameAndOthers[1], "」")
		if len(stationNameAndOther) != 2 {
			return nil, fmt.Errorf("access is invalid. splitted result must be 2, got %d. %s", len(stationNameAndOther), accesses[i])
		}
		stationName = stationNameAndOther[0]
		access := Access{
			trainName:   trainName,
			stationName: stationName,
			walk:        walk,
		}
		// 駅名とバスでの所要時間をparser
		// バスでの所要時間をparse
		if strings.Contains(stationNameAndOther[1], "バス") {
			busInMinutes, err := parseBus(stationNameAndOther[1])
			if err != nil {
				return nil, err
			}
			access.bus = &TimeInMinutes{
				minutes:        busInMinutes,
				movementMethod: "バス",
			}
		}
		parsedAccesses = append(parsedAccesses, access)
		if len(accesses) > 2 {
			var (
				walk TimeInMinutes
				bus  *TimeInMinutes
			)
			for i := range accesses[2:] {
				access := accesses[2+i]
				if strings.Contains(access, "バス") {
					// またはバス5分 徒歩1～4分→JR中央線,武蔵境,バス,5,徒歩,1
					re := regexp.MustCompile("バス|分|または")
					busInMinutes := strings.Split(re.ReplaceAllString(access, ""), "～")
					if busInTime, err := strconv.Atoi(busInMinutes[0]); err != nil {
						return nil, fmt.Errorf("access is invalid.  failed to parse bus in minutes. %s", busInMinutes[0])
					} else {
						bus = &TimeInMinutes{
							minutes:        busInTime,
							movementMethod: "バス",
						}
					}
				}
				if strings.Contains(access, "徒歩") {
					// またはバス5分 徒歩1～4分→JR中央線,武蔵境,バス,5,徒歩,1
					re := regexp.MustCompile("徒歩|分|または")
					walkInMinutes := strings.Split(re.ReplaceAllString(access, ""), "～")
					if walkInTime, err := strconv.Atoi(walkInMinutes[0]); err != nil {
						return nil, fmt.Errorf("access is invalid.  failed to parse walk in minutes. %s", walkInMinutes[0])
					} else {
						walk = TimeInMinutes{
							minutes:        walkInTime,
							movementMethod: "徒歩",
						}
					}
				}
			}
			parsedAccesses = append(parsedAccesses, Access{
				trainName:   trainName,
				stationName: stationName,
				walk:        walk,
				bus:         bus,
			})
		}
	}
	return parsedAccesses, nil
}

// parseWalk 徒歩での所要時間をparse
func parseWalk(accessText string) (int, error) {
	walkInMinutes := strings.Split(strings.ReplaceAll(strings.ReplaceAll(accessText, "徒歩", ""), "分", ""), "～")
	walkInTime, err := strconv.Atoi(walkInMinutes[0])
	if err != nil {
		return 0, fmt.Errorf("access is invalid. failed to parse walk in minutes. %s", walkInMinutes[0])
	}
	return walkInTime, nil
}

// parseBus バスでの所要時間をparse
func parseBus(accessText string) (int, error) {
	busInMinutes := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(accessText, "バス", ""), "分", ""), "または", ""), "駅", "")
	busInTime, err := strconv.Atoi(busInMinutes)
	if err != nil {
		return 0, fmt.Errorf("access is invalid.  failed to parse bus in minutes. %s", busInMinutes)
	}
	return busInTime, err
}
