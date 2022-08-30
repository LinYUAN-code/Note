package main

import (

	"log"
	"runtime"
	"unsafe"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"linyuan.com/go-gl/event"
	"linyuan.com/go-gl/utils"
)

func init() {
	// 将当前协程锁在 当前执行的系统线程上，避免切换
	runtime.LockOSThread()
}

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	window, err := glfw.CreateWindow(640, 480, "LinYuan Learn OpenGL", nil, nil)
	if err != nil {
		panic(err)
	}
	// 使用OpenGL版本3.3
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	// 使用OpenGL的核心模式
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	// 不允许缩放
	glfw.WindowHint(glfw.Resizable, glfw.False)
	// Mac兼容模式
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	// 设置window为当前OpenGL context
	window.MakeContextCurrent()

	// 设置键盘事件处理函数
	glfw.GetCurrentContext().SetKeyCallback(event.HandleKeyBoard)

	//Call gl.Init only under the presence of an active OpenGL context
	if err := gl.Init(); err != nil {
		log.Fatalln(err)
	}
	// 获取当前Window的宽高，并设置其为OpenGL的ViewPort大小
	width,height := glfw.GetCurrentContext().GetFramebufferSize()
	// 设置视口变化
	gl.Viewport(0,0,int32(width),int32(height))

	var VBO uint32
	gl.GenBuffers(1,&VBO)
	gl.BindBuffer(gl.ARRAY_BUFFER,VBO)
	vertices := []float32{
		-0.5, -0.5, 0.0,
		0.5, -0.5, 0.0,
		0.0,  0.5, 0.0,
	}
	/*
		GL_STATIC_DRAW ：数据不会或几乎不会改变。
		GL_DYNAMIC_DRAW：数据会被改变很多。
		GL_STREAM_DRAW ：数据每次绘制时都会改变。

		如果数据经常改变那么会被放在GPU的高速显存中
	*/
	gl.BufferData(gl.ARRAY_BUFFER,len(vertices),unsafe.Pointer(&vertices[0]),gl.STATIC_DRAW)

	utils.InitShader("./shader/vertexShader.glsl","./shader/fragmentShader.glsl")
	// 设置背景色
	for !window.ShouldClose() {
		// 处理事件
		glfw.PollEvents()
		// Do OpenGL stuff.
		// 设置一个颜色
		gl.ClearColor(0.2,0.3,0.3,1.0)
		// 使用设置的颜色来清空屏幕
		gl.Clear(gl.COLOR_BUFFER_BIT)
		// 交换缓冲区-双缓存机制
		window.SwapBuffers()
	}
	// 程序结束 释放资源
	glfw.Terminate()
}
