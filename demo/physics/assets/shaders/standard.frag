#version 330
#define cellshadelevel 3
#define numlight 2
uniform sampler2D diffuse;
uniform vec3 pointlight[8];

uniform vec3 eye;
uniform mat4 N;

in vec2 fragTexCoord;
in vec3 normal;
in vec3 world_pos;
layout (location=0) out vec4 outputColor;
void main() {
	vec3 world_normal = (N*vec4(normal,1)).xyz;
	vec3 eyedir = normalize(eye-world_pos);
	float lum =0;
	for(int x = 0; x < 2; x++){
		vec3 lightdir = normalize(world_pos-pointlight[x]);
		lum +=clamp(dot(world_normal,lightdir), 0.0,1.0);
	}
	lum = min(lum+0.3,1.0);
	//lum = int(lum*cellshadelevel)/cellshadelevel;//cell shade
	outputColor = texture(diffuse, fragTexCoord)*lum;
}