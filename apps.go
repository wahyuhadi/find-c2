package main

import (
	"bufio"
	"c2-finger/models"
	"fmt"
	"os"
	"strings"
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
	fmt.Println("[INFO] It might take a while — please enjoy your coffee in the meantime")
	fmt.Println("[>] Start Finding The RAT....")

	for sc.Scan() {

		ishash := ishash(sc.Text())
		ishost := ishost(sc.Text())
		msg := fmt.Sprintf(" -- Host %s", ishost)
		isfind, itype := findIndex(ishash)
		if isfind {
			total_found += 1
			fmt.Println(fmt.Sprintf("%s Found RAT = %s with Fingerprint %s", msg, *itype, ishash))
		}
	}

	if total_found == 0 {
		fmt.Println("[Done] No RAT Found")
	} else {
		fmt.Println(fmt.Sprintf("[Done] Total Found RAT %v", total_found))

	}

	fmt.Println("stop")
}
