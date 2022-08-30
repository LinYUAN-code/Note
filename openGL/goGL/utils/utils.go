package utils

import (
	"io/ioutil"
	"log"

	"github.com/go-gl/gl/v3.3-core/gl"
)


func InitShader(vertexPath,fragmentPath string) {
	var vertexShader,fragmentShader uint32
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
	gl.CompileShader(fragmentShader)
	freeV()
	freeF()

	program := gl.CreateProgram()
	gl.AttachShader(program,vertexShader)
	gl.AttachShader(program,fragmentShader)
	gl.LinkProgram(program)

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)
}

func readAsString(path string) (string,error) {
	f,err := ioutil.ReadFile(path)
	if err != nil {
		return "",err
	}
	return string(f),nil
}