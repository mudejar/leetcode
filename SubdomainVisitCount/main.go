package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

}

func subdomainVisits(cpdomains []string) []string {
	ans := []string{}
	domainCounts := map[string]int{}
	for _, str := range cpdomains {
		cpdomain := strings.Split(str, " ")
		subdomains := getAllSubdomains(cpdomain[1])
		count, _ := strconv.Atoi(cpdomain[0])
		for _, subdomain := range subdomains {
			domainCounts[subdomain] = domainCounts[subdomain] + count
		}
	}

	for domain, count := range domainCounts {
		ans = append(ans,fmt.Sprintf("%d %s", count, domain))
	}

	return ans
}

func getAllSubdomains(str string) []string {
	names := strings.Split(str, ".")
	subdomains := []string{}
	for i := 0; i < len(names); i++ {
		subdomains = append(subdomains, strings.Join(names[i:len(names)], "."))
	}

	return subdomains
}