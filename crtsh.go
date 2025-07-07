package sources

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

func Collect(domain string) []string {
    var results []map[string]interface{}
    subdomains := []string{}
    seen := map[string]bool{}

    url := fmt.Sprintf("https://crt.sh/?q=%%25.%s&output=json", domain)
    resp, err := http.Get(url)
    if err != nil {
        return subdomains
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    json.Unmarshal(body, &results)

    for _, entry := range results {
        name := entry["name_value"].(string)
        if !seen[name] {
            seen[name] = true
            subdomains = append(subdomains, name)
        }
    }

    return subdomains
}
