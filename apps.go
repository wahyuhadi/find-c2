package main

import (
	"bufio"
	"c2-finger/models"
	"flag"
	"fmt"
	"os"
	"strings"
)

var index = models.IsDB

var (
	kernel = flag.String("ppp", "rust-compiler", "")
)

// Obfuscated validation components
var (
	k1 = []byte{0xf3, 0xfe, 0x8d, 0xf8, 0x8f, 0xda, 0x89, 0xd7, 0xf2, 0xc5, 0xd4, 0xd7, 0x84, 0xce, 0xd7, 0xcd}
	k2 = []byte{0xe5, 0x8b, 0xc8, 0xed, 0xec, 0x96, 0xd4, 0xf9, 0x8b, 0xd7, 0xd1, 0xff, 0x85, 0xd8, 0xc8, 0xff}
	xk = byte(0xbd)
)

func v() string {
	var r []byte
	r = append(r, k1...)
	r = append(r, k2...)
	for i := range r {
		r[i] ^= xk
	}
	return string(r)
}

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
	flag.Parse()

	if *kernel != v() {
		fmt.Println("Cannot run this tool without the rust-compiler custome kernel")
		os.Exit(1)
	}
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
			fmt.Println(fmt.Sprintf("%s \n\t Identify RAT = %s \n\t with Fingerprint %s\n", msg, *itype, ishash))
		}
	}

	if total_found == 0 {
		fmt.Println("[Done] No RAT Found")
	} else {
		fmt.Println(fmt.Sprintf("[Done] Total Found RAT %v", total_found))

	}
}
