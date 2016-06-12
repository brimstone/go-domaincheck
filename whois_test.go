package main_test

import (
	"fmt"
	"log"
	"testing"

	dg "github.com/brimstone/go-domainglass"
)

var rawWhois = `% IANA WHOIS server
% for more information on IANA, visit http://www.iana.org
% This query returned 1 object

domain:       GLASS

organisation: Black Cover, LLC
address:      c/o Donuts Inc., 10500 NE 8th Street, Suite 350
address:      Bellevue, Washington 98004
address:      United States

contact:      administrative
name:         Richard Tindal
organisation: Donuts Inc.
address:      c/o Donuts Inc., 10500 NE 8th Street, Suite 350
address:      Bellevue, Washington 98004
address:      United States
phone:        +1.425.283.8248
fax-no:       +1.425.671.0020
e-mail:       richard@donuts.co

contact:      technical
name:         Chris Cowherd
organisation: Donuts Inc.
address:      c/o Donuts Inc., 10500 NE 8th Street, Suite 350
address:      Bellevue, Washington 98004
address:      United States
phone:        +1.425.296.6802
fax-no:       +1.425.671.0020
e-mail:       chris@donuts.co

nserver:      DEMAND.ALPHA.ARIDNS.NET.AU 2001:dcd:1:0:0:0:0:7 37.209.192.7
nserver:      DEMAND.BETA.ARIDNS.NET.AU 2001:dcd:2:0:0:0:0:7 37.209.194.7
nserver:      DEMAND.DELTA.ARIDNS.NET.AU 2001:dcd:4:0:0:0:0:7 37.209.198.7
nserver:      DEMAND.GAMMA.ARIDNS.NET.AU 2001:dcd:3:0:0:0:0:7 37.209.196.7
ds-rdata:     12799 8 1 9C1F9CB19B734D8E151F72C54FA50C46A561F937
ds-rdata:     12799 8 2 F1D217D39494BB123262FDC44F437C9D708653B0A83E9F43AB302091DAC72865

whois:        whois.donuts.co

status:       ACTIVE
remarks:      Registration information: http://donuts.co/

created:      2013-12-19
changed:      2015-06-29
source:       IANA

`

func TestGetWhoisRaw(*testing.T) {
	result, err := dg.GetWhoisRaw("whois.iana.org", "glass")
	if err != nil {
		log.Fatal("getWhoisRaw returned:", err)
	}

	if len(result) == 0 {
		log.Fatal("getWhoisRaw returned an empty string")
	}
}

func TestGetWhoisMap(*testing.T) {
	result, err := dg.GetWhoisMap(rawWhois)
	if err != nil {
		log.Fatal(err)
	}
	if len(result) == 0 {
		log.Fatal("Expected result of have multiple items")
	}
}

func TestGetWhoisInfo(*testing.T) {
	result, err := dg.GetWhoisInfo("domain.glass")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
