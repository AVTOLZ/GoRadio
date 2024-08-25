package main

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/robfig/cron"
)



func main() {
	fmt.Println("Starting Scheduler...")

	scheduler := cron.NewWithLocation(time.Now().Location())

	scheduler.AddFunc("0 0 10 * * 1-5", break1)
	scheduler.AddFunc("0 50 11 * * 1-5", break2)
	scheduler.AddFunc("0 50 13 * * 1-5", break3)

	scheduler.Start()

	fmt.Println("Scheduler Running")

	select {}
}

func break1() {
	fmt.Println("Break 1 Playing")
	PlayRadio(20)
	fmt.Println("Break 1 End")
}

func break2() {
	fmt.Println("Break 2 Playing")
	PlayRadio(30)
	fmt.Println("Break 2 End")
}

func break3() {
	fmt.Println("Break 3 Playing")
	PlayRadio(10)
	fmt.Println("Break 3 End")
}

func PlayRadio(minutes time.Duration) {
	var player *exec.Cmd = nil

	cmd := exec.Command("ffplay", "-nodisp", "https://playerservices.streamtheworld.com/api/livestream-redirect/RADIO538AAC.aac")
	cmd.Start()
	player = cmd

	time.Sleep(minutes * time.Minute)

	player.Process.Kill()
	player = nil
}