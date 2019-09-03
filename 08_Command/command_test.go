package command

import "testing"

func TestTest(t *testing.T) {
	ls := new(LightSwitch)     // 开关
	source := new(LightSource) // 灯

	off := new(LightOffCommand) // 关灯命令
	off.setReceiver(source)
	on := new(LightOnCommand) // 开灯命令
	on.setReceiver(source)

	ls.setCommand(on)
	ls.activate()

	ls.setCommand(off)
	ls.activate()
}
