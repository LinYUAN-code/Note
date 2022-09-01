package main

import (
	"log"
	"math"
	"runtime"

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

	var VAO uint32
	gl.GenVertexArrays(1,&VAO)
	gl.BindVertexArray(VAO)

	vertices := []float32{
		-0.5, -0.5, 0.0,
		0.5, -0.5, 0.0,
		0.0,  0.5, 0.0,
	}
	VBO := utils.GetVBOFloat32(vertices)

	// 绑定顶点属性 layout(location = 0)
	gl.VertexAttribPointer(0,3,gl.FLOAT,false,3 * 4, nil)
	gl.EnableVertexAttribArray(0)

	// 取消绑定
	gl.BindBuffer(gl.ARRAY_BUFFER,0)
	gl.BindVertexArray(0)

	program := utils.InitShader("./shader/vertexShader.glsl","./shader/fragmentShader.glsl")

	// 设置背景色
	for !window.ShouldClose() {
		// 处理事件
		glfw.PollEvents()
		
		// 设置一个颜色
		gl.ClearColor(0.2,0.3,0.3,1.0)
		// 使用设置的颜色来清空屏幕
		gl.Clear(gl.COLOR_BUFFER_BIT)
		gl.UseProgram(program)

		println(glfw.GetTime())

		blue := math.Sin(glfw.GetTime()) / 2.0 + 0.5;
		utils.InjectUniform4F(program,"time_color",1.0,1.0,float32(blue),1.0)

		gl.BindVertexArray(VAO)
		// Do OpenGL stuff.
		gl.DrawArrays(gl.TRIANGLES,0,3)

		// 交换缓冲区-双缓存机制
		window.SwapBuffers()
	}

	gl.DeleteVertexArrays(1, &VAO)
	gl.DeleteBuffers(1, &VBO)
	// 程序结束 释放资源
	glfw.Terminate()
}
