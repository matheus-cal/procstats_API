package parserstats

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type ProcStat struct {
	Id                 string
	User, System, Idle int
}

func Scan(reader io.Reader) ([]ProcStat, error) {
	scan := bufio.NewScanner(reader)
	scan.Scan()

	var stats []ProcStat

	for scan.Scan() {
		const NUM_FIELDS = 11
		fields := strings.Fields(string(scan.Bytes()))

		if len(fields) != NUM_FIELDS {
			continue
		}

		if !strings.HasPrefix(fields[0], "cpu") {
			continue
		}

		var times [3]int

		for i, index := range []int{1, 3, 4} {
			value, err := strconv.Atoi(fields[index])
			if err != nil {
				return nil, err
			}

			times[i] = value
		}

		stats = append(stats, ProcStat{
			Id:     fields[0],
			User:   times[0],
			System: times[1],
			Idle:   times[2],
		})
	}

	if err := scan.Err(); err != nil {
		return nil, err
	}

	return stats, nil
}
