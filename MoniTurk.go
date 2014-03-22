package main

import (
    "fmt"
    "net/http"
    "regexp"
    "github.com/miekg/dns"
    "log"
)

func main() {
    http.HandleFunc("/moniturk/", dnsQueryHandler)
    http.ListenAndServe(":8080", nil)
}

// DNS Utility

func dnsQueryHandler(w http.ResponseWriter, r *http.Request) {
    domain := r.URL.Path[len("/moniturk/"):]
    if(isValidDomain(domain)){
      success, answers := getDNSAnswer(domain, "8.8.8.8")

      if (success){
        if (len(answers) < 1){
          // Reply saying that there is no entry
          log.Println("No Answers")
        } else {
          // Generate reply page with all possible
          fmt.Fprintf(w, "We succesfully processed : %s", answers)
        }
      } else{
        // The DNSQuery failed
      }

    } else{
      // Reply saying non valid domain
    }

    return
}

// DNS Query Handler

func getDNSAnswer(queryDomain string, serverAddress string) (bool, []dns.RR){
  m := new(dns.Msg)
  m.SetQuestion(dns.Fqdn(queryDomain), dns.TypeANY)
  in, err := dns.Exchange(m, serverAddress+":53")

  if (err != nil && in.Answer != nil) {
      return false, nil
  }

  return true, in.Answer
}

// Helper methods

func isValidDomain(domain string) bool{
    var validDomainRegex = regexp.MustCompile("^(?:[-A-Za-z0-9]+\\.)+[A-Za-z]{2,6}$")
    m := validDomainRegex.FindStringSubmatch(domain)
    if (len(m) == 1) {
      return true
    } else{
      return false
    }
}
