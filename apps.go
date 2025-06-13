package main

import (
	"bufio"
	"c2-finger/models"
	"fmt"
	"os"
	"strings"

	"github.com/projectdiscovery/gologger"
)

var index = models.IsDB

func findIndex(ishash string) (bool, *string) {
	for _, i := range index {
		if ishash == i.Hash {
			return true, &i.Type
		}
	}
	return false, nil
}

func ishash(stdin string) string {
	data := strings.Fields(stdin)
	return data[len(data)-1]
}
func ishost(stdin string) string {
	data := strings.Fields(stdin)
	return data[len(data)-2]
}

func main() {
	// JARM	192.168.2.1:443	29d14d00029d29d00029d14d29d29df26cd45df722aff084b72395f8d18d1b
	// format JARM IP HASH
	var total_found = 0
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		gologger.Info().Msg("Start Finding The RAT....")

		ishash := ishash(sc.Text())
		ishost := ishost(sc.Text())
		msg := fmt.Sprintf("Host %s", ishost)
		isfind, itype := findIndex(ishash)
		if isfind {
			total_found += 1
			gologger.Info().Str("Fingerprint", ishash).Str("Found RAT", *itype).Msg(msg)
		}
	}

	if total_found == 0 {
		gologger.Info().Msg("[Done] No RAT Found")
	} else {
		gologger.Info().Msg(fmt.Sprintf("[Done] Total Found RAT %v", total_found))

	}

}
