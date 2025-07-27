// main.go
package main

import (
	"expiredex/cmd/cleanup"
	"expiredex/cmd/config"
	"expiredex/cmd/internal/utils"
	"fmt"
	"log"
	"time"

	aero "github.com/aerospike/aerospike-client-go/v8"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	utils.InitLogging()
	showBanner()
	utils.LogInfo("########## 	Starting ExpireDex Cleanup Service...     ##########")

	aeroCfg, cleanupCfg, err := config.ReadAerospikeConfig(`./config.yaml`)

	if err != nil {
		return
	}

	client := config.ClientConnect(aeroCfg)
	populateFakeData(client)
	cleanup.CleanupExpiredKeys(client, aeroCfg.Namespace, aeroCfg.Set, cleanupCfg, true)
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

func populateFakeData(client *aero.Client) {
	namespace := "test"
	set := "otp_data"

	// Add 5 fake keys: some expired, some future
	dates := []string{
		time.Now().AddDate(0, 0, -2).Format("20060102"), // expired
		time.Now().AddDate(0, 0, -1).Format("20060102"), // expired
		time.Now().Format("20060102"),                   // today
		time.Now().AddDate(0, 0, 1).Format("20060102"),  // future
		time.Now().AddDate(0, 0, 2).Format("20060102"),  // future
	}
	writePolicy := aero.NewWritePolicy(0, 0)
	writePolicy.SendKey = true
	for i, date := range dates {
		keyStr := fmt.Sprintf("delete_on:%s:key%d", date, i)
		key, _ := aero.NewKey(namespace, set, keyStr)

		bins := aero.BinMap{
			"otp":        fmt.Sprintf("%06d", 100000+i),
			"expires_on": date,
		}

		err := client.Put(writePolicy, key, bins)
		if err != nil {
			log.Printf("Failed to insert %s: %v", keyStr, err)
		} else {
			log.Printf("Inserted key: %s", keyStr)
		}
	}
}
