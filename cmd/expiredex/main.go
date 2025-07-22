// main.go
package main

import (
	"expiredex/internal/utils"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	utils.InitLogging()
	showBanner()
	utils.LogInfo("########## 	Starting ExpireDex Cleanup Service...     ##########")

}

func showBanner() {
	utils.LogInfo(`
.---..   ..--. --.--.--. .---..--. .---..   .
|     \ / |   )  |  |   )|    |   :|     \ / 
|---   /  |--'   |  |--' |--- |   ||---   /  
|     / \ |      |  |  \ |    |   ;|     / \ 
'---''   ''    --'--'   \'---''--' '---''   '
`)
}
