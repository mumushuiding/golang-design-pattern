package command

import "fmt"

// ICommand 命令接口
type ICommand interface {
	execute()
}

// LightSource 灯/接收者
type LightSource struct{}

func (*LightSource) off() {
	fmt.Println("关灯 light off")
}
func (*LightSource) on() {
	fmt.Println("开灯 Light on ")
}

// LightSwitch 开关/调用者
type LightSwitch struct {
	iCommand ICommand
}

// setCommand 初始化用于执行命令的方法
func (l *LightSwitch) setCommand(command ICommand) {
	l.iCommand = command
}

// activate 执行命令
func (l *LightSwitch) activate() {
	l.iCommand.execute()
}

// LightOffCommand 关灯命令
type LightOffCommand struct {
	light *LightSource
}

func (l *LightOffCommand) setReceiver(receiver *LightSource) {
	l.light = receiver
}
func (l *LightOffCommand) execute() {
	l.light.off()
}

// LightOnCommand 开灯命令
type LightOnCommand struct {
	light *LightSource
}

func (l *LightOnCommand) setReceiver(receiver *LightSource) {
	l.light = receiver
}
func (l *LightOnCommand) execute() {
	l.light.on()
}
