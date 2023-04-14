package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	tld "github.com/jpillora/go-tld"
	"github.com/projectdiscovery/httpx/common/fileutil"
)

// Timeout to check for a connection - in seconds
var timout time.Duration = 2
var wg sync.WaitGroup

func getips(args []string) {
	var input []string

	if fileutil.HasStdin() && len(args) == 1 {
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			text := strings.TrimSpace(scanner.Text())
			if len(text) != 0 {
				input = append(input, text)
			}
		}
	} else if len(args) == 2 {
		filename := args[1]
		input = fileutil.LoadFile(filename)
	}

	for _, ip := range input {
		wg.Add(1)
		go process_ips(ip)
	}
	wg.Wait()
}

func gethostname(ip_port string) string {
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}
	var domainname string
	conn, err := net.DialTimeout("tcp", ip_port, timout*time.Second)
	if err == nil {
		tlsconn := tls.Client(conn, conf)
		handshake := tlsconn.Handshake()
		if handshake == nil {
			state := tlsconn.ConnectionState()
			hostname := state.PeerCertificates[0].Subject.CommonName
			hostname = "https://" + hostname
			u, errr := tld.Parse(hostname)
			if errr == nil {
				if u.Subdomain == "*" || u.Subdomain == "" {
					domainname = u.Domain + "." + u.TLD
				} else {
					domainname = u.Subdomain + "." + u.Domain + "." + u.TLD
				}
			}
			tlsconn.Close()
		}
		conn.Close()
	}

	return domainname
}

func process_ips(ip string) {
	defer wg.Done()

	var ip_port string
	port := "443"
	if strings.Contains(ip, "https") {
		ip_port = strings.TrimPrefix(ip, "https://")
	} else {
		ip_port = ip
	}

	if strings.Count(ip_port, ":") == 0 {
		ip_port = ip_port + ":" + port
	}

	hostname := gethostname(ip_port)
	fmt.Println(ip, strings.ToLower(hostname))
}

func main() {
	args := os.Args
	if !fileutil.HasStdin() && len(args) != 2 {
		fmt.Println("Please provide one file with list of IPs")
	} else {
		getips(args)
	}
}
