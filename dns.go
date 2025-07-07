package resolver

import (
    "net"
    "sync"
)

func ResolveSubdomains(subs []string) []string {
    var resolved []string
    var wg sync.WaitGroup
    var mu sync.Mutex

    for _, sub := range subs {
        wg.Add(1)
        go func(s string) {
            defer wg.Done()
            ips, err := net.LookupHost(s)
            if err == nil && len(ips) > 0 {
                mu.Lock()
                resolved = append(resolved, s)
                mu.Unlock()
            }
        }(sub)
    }

    wg.Wait()
    return resolved
}
