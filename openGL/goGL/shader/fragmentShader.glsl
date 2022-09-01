#version 330 core

out vec4 color;

uniform vec4 time_color;

void main()
{
    color = time_color;
}