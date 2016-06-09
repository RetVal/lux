#version 330
uniform mat4 P;
uniform mat4 V;
uniform mat4 M;
layout (location=0) in vec3 vert;
layout (location=1) in vec2 vertTexCoord;
layout (location=2) in vec3 vertNormal;
out vec2 fragTexCoord;
out vec3 normal;
out vec3 world_pos;
void main() {
	normal = vertNormal;
    fragTexCoord = vertTexCoord;
    world_pos=(M*vec4(vert,1)).xyz;
    gl_Position = P * V * M * vec4(vert, 1);
}