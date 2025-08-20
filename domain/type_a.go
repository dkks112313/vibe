package domain

import (
	"fmt"
	"net"

	"github.com/miekg/dns"
)

func LookupA(url string, channel chan []net.IP, errors chan error) {
	var msg dns.Msg
	answer := dns.Fqdn(url)
	msg.SetQuestion(answer, dns.TypeA)

	in, err := dns.Exchange(&msg, "8.8.8.8:53")
	if err != nil {
		errors <- err
		return
	}

	if len(in.Answer) < 1 {
		errors <- fmt.Errorf("no records at %s", url)
		return
	}

	var records []net.IP
	for _, answer := range in.Answer {
		if a, ok := answer.(*dns.A); ok {
			records = append(records, a.A)
		}
	}

	if len(records) > 0 {
		channel <- records
		return
	}

	errors <- fmt.Errorf("no records A")
}
