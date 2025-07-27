package cleanup

import (
	"expiredex/cmd/config"
	"fmt"
	"strings"
	"time"

	aero "github.com/aerospike/aerospike-client-go/v8"
)

func CleanupExpiredKeys(client *aero.Client, namespace string, set string, cleanupConfig *config.AerospikeCleanUp, dryRun bool) error {
	policy := aero.NewScanPolicy()

	recordset, err := client.ScanAll(policy, namespace, set)

	if err != nil {
		return err
	}

	for rec := range recordset.Results() {
		if rec.Err != nil {
			continue
		}

		keyObject := rec.Record.Key
		fmt.Printf("Key: %s \n", keyObject)
		var keyString string
		if keyObject.Value() != nil {
			keyString = keyObject.Value().String()
		} else {
			// fallback to digest if needed
			keyString = keyObject.String() // prints full key: ns/set/digest
		}
		fmt.Println("✅ Found key with prefix:", keyObject.Value())

		if strings.HasPrefix(keyString, cleanupConfig.Key_Prefix) {
			datePart := strings.Split(strings.TrimPrefix(keyString, cleanupConfig.Key_Prefix), ":")[0]
			deleteDate, err := time.Parse(cleanupConfig.Date_Format, datePart)

			if err != nil {
				return err
			}
			fmt.Println("NOW:", time.Now().Format("20060102"), "KEY_DATE:", deleteDate.Format("20060102"))

			if time.Now().After(deleteDate) {
				fmt.Println("⏰ Expired date:", deleteDate)

				if dryRun {
					fmt.Println("Would Delete: ", keyString)
				}

				deleted, err := client.Delete(nil, keyObject)
				if err != nil {
					fmt.Println("❌ Failed to delete:", err)
				} else if deleted {
					fmt.Println("🧹 Successfully deleted:", keyObject.Value().String())
				} else {
					fmt.Println("⚠️ Key not found:", keyObject.Value().String())
				}

			}
		}

	}
	return nil
	// You can now use:
	// recordset, err := client.ScanAll(policy, namespace, set)
}
