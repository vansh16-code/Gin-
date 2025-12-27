package main

import (
	"fmt"
)

type EmailJob struct {
	To      string
	Subject string
	Body    string
}

var emailJobs = make(chan EmailJob, 20) //creates a channel queue with a buffer size of 20

func StartEmailWorker() {
	go func() {
		for job := range emailJobs {
			fmt.Println(" Sending email to:", job.To)

			err := SendEmail(job.To, job.Subject, job.Body)
			if err != nil {
				fmt.Println(" Error sending:", err)
				continue
			}

			fmt.Println(" Email sent:", job.To)
		}
	}()
}
