package event

import "github.com/go-gl/glfw/v3.3/glfw"


func HandleKeyBoard(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	println("handleKeyBoard~")
	// 如果按Esc 那么关闭窗口
	if key == glfw.KeyEscape && action == glfw.Press {
		glfw.GetCurrentContext().SetShouldClose(true)
	}
}