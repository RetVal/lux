package main

import (
	"fmt"
	"github.com/luxengine/steam"
	"os"
)

func main() {
	fmt.Println(os.Getwd())
	fmt.Printf("steam init: %t\n", steam.Init())
	controller := steam.SteamController()
	fmt.Printf("Controller init: %t\n", controller.Init())
	handles := make([]steam.ControllerHandle, 16)
	controller.RunFrame()
	n := controller.GetConnectedControllers(&handles[0])
	fmt.Printf("connected controllers: %d\n", n)
	fmt.Printf("controller handles: %v\n", handles)
	if n == 0 {
		fmt.Println("steam controller not found, exiting.")
		return
	}
	funcs := []func(string) int{
		func(s string) int { return int(controller.GetDigitalActionHandle(s)) },
		func(s string) int { return int(controller.GetAnalogActionHandle(s)) },
		func(s string) int { return int(controller.GetActionSetHandle(s)) },
	}

	str := []string{
		"In Game Actions",
		"actions",
		"InGameControls",
		"title",
		"Set_Ingame",
		"StickPadGyro",
		"Move",
		"title",
		"Action_Move",
		"input_mode",
		"joystick_move",
		"#Set_Ingame",
		"#Action_Move",
	}
	for x, f := range funcs {
		for y, s := range str {
			i := f(s)
			if i != 0 {
				fmt.Println(x, y)
			}
		}
	}

	//steam.Shutdown()
}
