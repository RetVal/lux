package lux

const (
	_gbufferVertexShaderSource = `#version 330
uniform mat4 M;
uniform mat4 MVP;

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
	gl_Position = MVP * vec4(vert, 1);
}
` + "\x00"

	_gbufferFragmentShaderSource = `#version 330
uniform sampler2D diffuse;
uniform mat4 N;

in vec2 fragTexCoord;
in vec3 normal;
in vec3 world_pos;
layout (location=0) out vec4 outAlbedo;
layout (location=1) out vec3 outNormal;
layout (location=2) out vec3 outPosition;
void main() {
	outAlbedo = vec4(texture(diffuse, fragTexCoord).rgb, 0);
	outNormal = (N*vec4(normal,1)).xyz;
	outPosition = world_pos;
}
` + "\x00"

	_gbufferAggregateFragmentShader = `#version 330

// GBuffer textures
uniform sampler2D albedoTex;
uniform sampler2D normaltex;
uniform sampler2D postex;
uniform sampler2D depthtex;

in vec2 uv;

layout (location=0) out vec4 outColor;

float Lux() {
	float lux = texture(albedoTex, uv).a;
	return lux;
}

void colMulLux(){
	vec4 t = texture(albedoTex, uv);
	outColor = vec4(t.rgb*max(Lux(), 1), 1);
}

void luxOnly() {
	float l = Lux();
	outColor = vec4(l, l, l, 1);
}

void main(){
	colMulLux();
}
` + "\x00"
)
