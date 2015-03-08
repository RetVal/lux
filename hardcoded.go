package noname

var Postprocessfragmentshader_normalvisual = `#version 330
uniform sampler2D tex;
uniform vec2 resolution;
uniform float time;
in vec2 uv;
layout (location=0) out vec4 outputColor;
void main(){
	outputColor=texture(tex,uv);
}
` + "\x00"

var Postprocessfragmentshader_nothing = `#version 330
uniform sampler2D tex;
uniform vec2 resolution;
uniform float time;
in vec2 uv;
layout (location=0) out vec4 outputColor;

void main(){
	outputColor=texture(tex,uv);
}
` + "\x00"

var Postprocessfragmentshader_viewdepth = `#version 330
uniform sampler2D tex;
uniform vec2 resolution;
uniform float time;
in vec2 uv;
layout (location=0) out vec4 outputColor;

float LinearizeDepth(float z){
	float n = 0.1;
	float f = 5;
	return (2.0 * n) / (f + n-z*(f-n));
}

void main(){
	float depth = texture(tex,uv).r;
	outputColor=vec4(vec3(depth),1);
}
` + "\x00"

var Postprocessfragmentshader_inverse = `#version 330
uniform sampler2D tex;
uniform vec2 resolution;
uniform float time;
in vec2 uv;
layout (location=0) out vec4 outputColor;
void main(){
	outputColor=1-texture(tex,uv);
}
` + "\x00"

//mini FXAA
var Postprocessfragmentshader_fxaa = `#version 330
#define FXAA_REDUCE_MIN (1.0/128.0)
#define FXAA_REDUCE_MUL (1.0/8.0)
#define FXAA_SPAN_MAX 8.0
uniform sampler2D tex;
uniform vec2 resolution;
uniform float time;
in vec2 uv;

layout (location=0) out vec4 outputColor;

void main(){
	vec2 inverse_resolution=vec2(1.0/resolution.x,1.0/resolution.y);
	vec3 rgbNW = texture(tex, (gl_FragCoord.xy + vec2(-1.0,-1.0)) * inverse_resolution).xyz;
	vec3 rgbNE = texture(tex, (gl_FragCoord.xy + vec2(1.0,-1.0)) * inverse_resolution).xyz;
	vec3 rgbSW = texture(tex, (gl_FragCoord.xy + vec2(-1.0,1.0)) * inverse_resolution).xyz;
	vec3 rgbSE = texture(tex, (gl_FragCoord.xy + vec2(1.0,1.0)) * inverse_resolution).xyz;
	vec3 rgbM  = texture(tex,  gl_FragCoord.xy  * inverse_resolution).xyz;
	vec3 luma = vec3(0.299, 0.587, 0.114);
	float lumaNW = dot(rgbNW, luma);
	float lumaNE = dot(rgbNE, luma);
	float lumaSW = dot(rgbSW, luma);
	float lumaSE = dot(rgbSE, luma);
	float lumaM  = dot(rgbM,  luma);
	float lumaMin = min(lumaM, min(min(lumaNW, lumaNE), min(lumaSW, lumaSE)));
	float lumaMax = max(lumaM, max(max(lumaNW, lumaNE), max(lumaSW, lumaSE))); 
	vec2 dir;
	dir.x = -((lumaNW + lumaNE) - (lumaSW + lumaSE));
	dir.y =  ((lumaNW + lumaSW) - (lumaNE + lumaSE));
	float dirReduce = max((lumaNW + lumaNE + lumaSW + lumaSE) * (0.25 * FXAA_REDUCE_MUL),FXAA_REDUCE_MIN);
	float rcpDirMin = 1.0/(min(abs(dir.x), abs(dir.y)) + dirReduce);
	dir = min(vec2( FXAA_SPAN_MAX,  FXAA_SPAN_MAX),max(vec2(-FXAA_SPAN_MAX, -FXAA_SPAN_MAX),dir * rcpDirMin)) * inverse_resolution;
	vec3 rgbA = 0.5 * (texture(tex,   gl_FragCoord.xy  * inverse_resolution + dir * (1.0/3.0 - 0.5)).xyz + texture(tex,   gl_FragCoord.xy  * inverse_resolution + dir * (2.0/3.0 - 0.5)).xyz);
	vec3 rgbB = rgbA * 0.5 + 0.25 * (texture(tex,  gl_FragCoord.xy  * inverse_resolution + dir *  - 0.5).xyz + texture(tex,  gl_FragCoord.xy  * inverse_resolution + dir * 0.5).xyz);
	float lumaB = dot(rgbB, luma);
	if((lumaB < lumaMin) || (lumaB > lumaMax)) {
		outputColor = vec4(rgbA,1.0);
	} else {
		outputColor = vec4(rgbB,1.0);
	}
}` + "\x00"

var Postprocessfragmentshader_wobbly = `#version 330
uniform sampler2D tex;
uniform vec2 resolution;
uniform float time;
in vec2 uv;

layout (location=0) out vec4 outputColor;

void main() {
	vec2 coord = uv;
	coord.y = coord.y+sin(coord.y*100)/200;
	coord.x = coord.x+cos(coord.x*100)/200;
	outputColor = texture(tex, coord);
}
` + "\x00"

var Postprocessfragmentshader_flip = `#version 330
uniform sampler2D tex;
uniform vec2 resolution;
uniform float time;
in vec2 uv;
layout (location=0) out vec4 outputColor;

void main(){
	outputColor=texture(tex,vec2(uv.x,1-uv.y));
}
` + "\x00"
