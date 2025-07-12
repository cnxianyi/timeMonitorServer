package utils

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"time"
	"timeMonitorServer/models"
)

func InitCronJobs() {
	// Create a new cron scheduler.
	// cron.WithSeconds() is important if your schedule needs second precision,
	// like "0 0 0 * * *" which includes the '0' for seconds.
	c := cron.New(cron.WithSeconds())

	// Schedule the daily 0:00 AM task
	// Cron expression: "秒 分 时 日 月 星期"
	// "0 0 0 * * *" means:
	// - 0 seconds past the minute
	// - 0 minutes past the hour
	// - 0 hours (midnight)
	// - every day of the month (*)
	// - every month (*)
	// - every day of the week (*)
	_, err := c.AddFunc("0 56 14 * * *", func() {
		fmt.Printf("Daily 0:00 AM task executed! Current time: %s\n", time.Now().Format("2006-01-02 15:04:05"))
		err := models.UpdateAllDailyTime()
		if err != nil {
			fmt.Printf("UpdateAllDailyTime err: %v\n", err)
		}
		// --- Your daily task logic goes here ---
		// Example:
		// global.DB.Exec("DELETE FROM old_records WHERE created_at < CURDATE() - INTERVAL 30 DAY")
		// Another example: Call a function in another package:
		// timeMonitorServer.jobs.CleanOldData()
		// ------------------------------------
	})
	if err != nil {
		log.Printf("Error adding daily 0:00 AM cron job: %v\n", err)
		// Depending on severity, you might want to panic or exit here
	}

	// You can add other cron jobs here as needed:
	// _, err = c.AddFunc("0 */30 * * * *", func() { // Every 30 minutes
	// 	fmt.Println("This runs every 30 minutes.")
	// })
	// if err != nil {
	// 	log.Printf("Error adding 30-min cron job: %v\n", err)
	// }

	// Start the cron scheduler in a separate goroutine.
	// This allows the main goroutine to continue to router.Init().
	c.Start()
}
