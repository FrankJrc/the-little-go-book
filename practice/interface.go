package main
import "fmt"

type IpAddr [4]byte

func (ip IpAddr) String() string {
    return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func main() {
    ips := map[string]IpAddr { 
        "google": {12,4,5,6}, 
        "facebook": {1,2,3,4},
        }
    for n,ip := range ips {
        fmt.Printf("%v: %v\n", n, ip)
    }
}
