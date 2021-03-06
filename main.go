package main

import (
	"fmt"
	"github.com/Azure/batch-insights/pkg"
	"os"
	"strings"
)

func main() {
	var appInsightsKey = os.Getenv("APP_INSIGHTS_INSTRUMENTATION_KEY")
	var poolId = os.Getenv("AZ_BATCH_POOL_ID")
	var nodeId = os.Getenv("AZ_BATCH_NODE_ID")
	var processNamesStr = os.Getenv("AZ_BATCH_MONITOR_PROCESSES")

	if len(os.Args) > 2 {
		poolId = os.Args[1]
		nodeId = os.Args[2]
	}

	if len(os.Args) > 3 {
		appInsightsKey = os.Args[3]
	}

	if len(os.Args) > 4 {
		processNamesStr = os.Args[4]
	}

	processNames := strings.Split(processNamesStr, ",")
	for i := range processNames {
		processNames[i] = strings.TrimSpace(processNames[i])
	}

	batchinsights.PrintSystemInfo()
	fmt.Printf("   Pool ID: %s\n", poolId)
	fmt.Printf("   Node ID: %s\n", nodeId)

	hiddenKey := "-"
	if appInsightsKey != "" {
		hiddenKey = "xxxxx"
	}

	fmt.Printf("   Instrumentation Key: %s\n", hiddenKey)

	fmt.Printf("   Monitoring processes: %s\n", strings.Join(processNames, ", "))

	batchinsights.ListenForStats(poolId, nodeId, appInsightsKey, processNames)
}
