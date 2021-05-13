package main

import (
	"fmt"
	"github.com/sdxstar/bass"
	"time"
)

var devices = make(map[uint32]*bass.DeviceInfo)

func init() {
	var i uint32 = 1
	var deviceInfo bass.DeviceInfo
	for ; bass.BASS_GetDeviceInfo(i, &deviceInfo); i++ {
		fmt.Printf("%d: %+v\n", i, deviceInfo.Name())
		devices[i] = &deviceInfo
	}
}

func main() {
	defer bass.BassFree()

	bassPlayer, err := bass.NewBass(3)
	if err != nil {
		panic(err)
	}
	defer bassPlayer.Close()

	fmt.Printf("curr device: %+v\n", devices[bass.BASS_GetDevice()].Name())

	err = bassPlayer.OpenFile("F:\\workspace\\gomod\\ehotel\\bass\\example\\Alarm08试.wav")
	if err != nil {
		panic(err)
	}
	//i , err := bassPlayer.GetDevice()
	//fmt.Printf("curr device: %+v\n", devices[i].Name())

	err = bassPlayer.SetVolume(80)
	if err != nil {
		panic(err)
	}
	length, err := bassPlayer.GetLength()
	if err != nil {
		panic(err)
	}
	fmt.Println(length)
	bassPlayer.Play(false)
	time.Sleep(time.Duration(length) * time.Millisecond)
}
