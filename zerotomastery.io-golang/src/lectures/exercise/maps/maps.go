//--Summary:
//  Write a program to display server status.
//
//--Requirements:
//* Create a function to print server status, including:
//  - Number of servers
//  - Number of servers for each status (Online, Offline, Maintenance, Retired)
//* Store the existing slice of servers in a map
//* Default all of the servers to `Online`
//* Perform the following status changes and display server info:
//  - display server info
//  - change `darkstar` to `Retired`
//  - change `aiur` to `Offline`
//  - display server info
//  - change all servers to `Maintenance`
//  - display server info

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)
var statusList = []string {
	"Online",
	"Offline",
	"Maintenance",
	"Retired",
}

//* Create a function to print server status, including:
//  - Number of servers
//  - Number of servers for each status (Online, Offline, Maintenance, Retired)

func printSummaryServerStatus(serverMap map[string]int){
	fmt.Println()
	fmt.Println("Number of servers:", len(serverMap))
	var smap = make(map[int]int)
	// for i:=0;i<4;i++{
	// 	smap[i]=0
	// }
	for _,stat := range serverMap {
		smap[stat] +=1
	}
	fmt.Println("Number of servers at each status:")
	for idx,statStr := range statusList {
		fmt.Println("  ", smap[idx], "at", statStr)
	}
}
func printServerStatusList(servers [] string, serverMap map[string]int){
	fmt.Println()
	for _,serverName := range servers{
		fmt.Println("  ", serverName, ":", serverMap[serverName],statusList[serverMap[serverName]])
	}
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	//* Store the existing slice of servers in a map
	serverStausMap := make(map[string]int)

	//* Default all of the servers to `Online`
	for _, server := range servers {
		serverStausMap[server] = Online
	}
	// for i:=0;i<len(servers);i++ {
	// 	serverStausMap[servers[i]] = Online
	// }

	//* Perform the following status changes and display server info:
	//  - display server info
	printSummaryServerStatus(serverStausMap)
	printServerStatusList(servers,serverStausMap)

	//  - change `darkstar` to `Retired`
	serverStausMap["darkstar"] = Retired

	//  - change `aiur` to `Offline`
	serverStausMap["aiur"] = Offline

	//  - display server info
	printSummaryServerStatus(serverStausMap)
	printServerStatusList(servers,serverStausMap)

	//  - change all servers to `Maintenance`
	for key,_ := range serverStausMap {
		serverStausMap[key] = Maintenance
	}

	//  - display server info
	printSummaryServerStatus(serverStausMap)
	printServerStatusList(servers,serverStausMap)
}
