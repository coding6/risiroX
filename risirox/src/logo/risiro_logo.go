package logo

import (
	"fmt"
	"risirox/risirox/src/conf"
)

var risiroXLogo = `                                        
██████╗  ██████  ███████╗ ██████  ██████╗    ████╗    
██╔══██╗   ██║   ██╔════╝   ██║   ██╔══██╗ ██╗   ██╗ 
██████╔╝   ██║   ███████╗   ██║   ██████╔╝ ██║   ██║ 
██╔══██╗   ██║   ╚════██║   ██║   ██╔══██╗ ██║   ██║ 
██║  ██║ ██████  ███████║ ██████  ██║  ██║ ╚██████╔╝ 
╚═╝  ╚═╝ ╚═════╝ ╚══════╝ ╚═════╝ ╚═╝  ╚═╝  ╚═════╝
                                        `
var topLine = `┌──────────────────────────────────────────────────────┐`
var borderLine = `│`
var bottomLine = `└──────────────────────────────────────────────────────┘`

func PrintLogo() {
	fmt.Println(risiroXLogo)
	fmt.Println(topLine)
	fmt.Println(fmt.Sprintf("%s [Welcome] Welcome to use RisiroX			%s", borderLine, borderLine))
	fmt.Println(bottomLine)
	fmt.Printf("[RisiroX] Version: %s, MaxConn: %d, MaxPackageSize: %d\n",
		"RisiroX V1.0",
		conf.GlobalConfigObj.MaxConn,
		conf.GlobalConfigObj.MaxPackageSize)
}
