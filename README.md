# Get Domain Name
__gdn__ is a Go module to get domain name from SSL certificates given an IP address

## Installation Instructions

### From Source
gdn requires go1.14+ to install successfully. Run the following command to get the repo -

```
▶ GO111MODULE=on go get -v github.com/kmskrishna/gdn
```
### From Github

```
▶ git clone https://github.com/kmskrishna/gdn.git; cd gdn; go build; mv gdn /usr/local/bin/
```

## Usage

You can input a list of IPs in two ways.

Filename as an argument
```
▶  gdn ./ips.txt

173.0.84.29 cloudmonitor.paypal.com
173.0.84.43 3ph.paypalcorp.com
173.0.84.31 pics.paypal.com
173.0.84.44 t.paypal.com
173.0.84.12 t.paypal.com
173.0.84.4 securepayments.paypal.com
173.0.84.36 securepayments.paypal.com
173.0.84.45 business.paypal.com
173.0.84.14 t.paypal.com
173.0.84.25 pics.paypal.com
173.0.84.46 t.paypal.com
173.0.84.24 demo.paypal.com
173.0.84.32 py.pl
173.0.84.9
173.0.84.13 business.paypal.com
173.0.84.6 www.paypal.com
173.0.84.16 www.paypal.com
173.0.84.34 www.paypal.com
```

Piping the content
```
▶  cat ips.txt | gdn

```

If you only want the Domain names in respose, it is not implemented right now but you can use the following command
```
▶  gdn ips.txt | awk '{print $2}' | sort -u

3ph.paypalcorp.com
business.paypal.com
cloudmonitor.paypal.com
demo.paypal.com
pics.paypal.com
py.pl
securepayments.paypal.com
t.paypal.com
www.paypal.com
```
This will give you unique domain names for all given IPs.

You can also use __gdn__ along side __Project Discovery's__ [httpx](https://github.com/projectdiscovery/httpx). You can directly pipe __httpx__ input into __gdn__ and get the domain/subdomain names directly.

```
▶  cat ips.txt | httpx --silent | gdn

https://173.0.84.25 pics.paypal.com
https://173.0.84.36 securepayments.paypal.com
https://173.0.84.29 cloudmonitor.paypal.com
https://173.0.84.4 securepayments.paypal.com
https://173.0.84.32 py.pl
https://173.0.84.34 www.paypal.com
https://173.0.84.16 www.paypal.com
https://173.0.84.14 t.paypal.com
https://173.0.84.44 t.paypal.com
https://173.0.84.46 t.paypal.com
https://173.0.84.45 business.paypal.com
https://173.0.84.12 t.paypal.com
https://173.0.84.43 3ph.paypalcorp.com
https://173.0.84.24 demo.paypal.com
https://173.0.84.13 business.paypal.com
```



### Limitations and further work

Will add a feature to just get the domain name and not just the subdomain name in future.


