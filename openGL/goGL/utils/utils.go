package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"
)

func getStatus(shader uint32,program uint32,errorHeader string) {
	var success int32
	gl.GetShaderiv(shader,gl.COMPILE_STATUS,&success)
	if success == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader,gl.INFO_LOG_LENGTH,&logLength)

		log := strings.Repeat("\x00",int(logLength+1))
		gl.GetShaderInfoLog(shader,logLength,nil,gl.Str(log))

		fmt.Printf("%v\n\t failed to compile %v",errorHeader,log)
	}
}

func InitShader(vertexPath,fragmentPath string) uint32 {
	var vertexShader,fragmentShader uint32
	program := gl.CreateProgram()
	vertexShader = gl.CreateShader(gl.VERTEX_SHADER)
	fragmentShader = gl.CreateShader(gl.FRAGMENT_SHADER)
	vertexSource,err := readAsString(vertexPath)
	if err != nil {
		log.Fatal(err)
	}
	fragmentSource,err := readAsString(fragmentPath)
	if err != nil {
		log.Fatal(err)
	}
	pvertexSource,freeV := gl.Strs(vertexSource+"\x00")
	pfragmentSource,freeF := gl.Strs(fragmentSource+"\x00")

	gl.ShaderSource(vertexShader,1,pvertexSource,nil)
	gl.ShaderSource(fragmentShader,1,pfragmentSource,nil)
	gl.CompileShader(vertexShader)
	getStatus(vertexShader,program,"VertexShader")
	gl.CompileShader(fragmentShader)
	getStatus(fragmentShader,program,"FragmentShader")
	freeV()
	freeF()

	gl.AttachShader(program,vertexShader)
	gl.AttachShader(program,fragmentShader)
	gl.LinkProgram(program)

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)
	return program
}

func readAsString(path string) (string,error) {
	f,err := ioutil.ReadFile(path)
	if err != nil {
		return "",err
	}
	return string(f),nil
}


func GetVBOFloat32(data []float32) uint32 {
	var VBO uint32
	gl.GenBuffers(1,&VBO)
	gl.BindBuffer(gl.ARRAY_BUFFER,VBO)
	/*
		GL_STATIC_DRAW ：数据不会或几乎不会改变。
		GL_DYNAMIC_DRAW：数据会被改变很多。
		GL_STREAM_DRAW ：数据每次绘制时都会改变。

		如果数据经常改变那么会被放在GPU的高速显存中
	*/
	gl.BufferData(gl.ARRAY_BUFFER,4 * len(data),gl.Ptr(data),gl.STATIC_DRAW)
	return VBO
}


func InjectUniform4F(program uint32,variableName string,v0 ,v1,v2,v3 float32) {
	location := gl.GetUniformLocation(program, gl.Str(variableName+"\x00"))
	if location == -1 {
		log.Fatalf("Could not Find the uniform Variable %v",variableName)
	}
 	gl.UseProgram(program)
	gl.Uniform4f(location,v0,v1,v2,v3)
}