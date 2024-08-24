package main

import (
	"os/exec"
	"time"
	"github.com/robfig/cron"
)



func main() {
	scheduler := cron.New()

	scheduler.AddFunc("0 0 10 * * 1-5", break1)
	scheduler.AddFunc("0 50 11 * * 1-5", break2)
	scheduler.AddFunc("0 50 13 * * 1-5", break3)
}

func break1() {
	PlayRadio(20)
}

func break2() {
	PlayRadio(30)
}

func break3() {
	PlayRadio(10)
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