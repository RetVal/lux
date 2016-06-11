package lux

import (
	"strings"
	"unsafe"

	gl2 "github.com/luxengine/lux/gl"
	"github.com/luxengine/lux/glm"

	"github.com/go-gl/gl/v3.3-core/gl"
)

// ParticleSystem is a particle system, it generates tiny points in space and
// renders them
type ParticleSystem struct {
	root, direction         glm.Vec3
	transformfeedbacks      []gl2.TransformFeedback //size 2
	buffers                 []gl2.Buffer            //size 2
	program                 gl2.Program
	isFirst                 bool
	currentVB, currentTFB   int
	particles               []Particle
	time                    float64
	gDeltaTimeMillis        gl2.UniformLocation
	gTime                   gl2.UniformLocation
	gRandomTexture          gl2.UniformLocation
	gLauncherLifetime       gl2.UniformLocation
	gShellLifetime          gl2.UniformLocation
	gSecondaryShellLifetime gl2.UniformLocation
}

// Particle is a single particle.
type Particle struct {
	Type               int32
	Position, Velocity glm.Vec3
	Lifetime           float32
}

// the types of particle there are
const (
	Launcher int32 = 0
)

// NewParticleSystem builds a particle system
func NewParticleSystem(position, direction glm.Vec3, size int) *ParticleSystem {
	ps := ParticleSystem{}
	ps.root = position
	ps.direction = direction
	ps.transformfeedbacks = gl2.GenTransformFeedbacks(2)
	MustNotGLError()
	ps.buffers = gl2.GenBuffers(2)
	MustNotGLError()
	ps.particles = make([]Particle, size)
	ps.isFirst = true

	ps.program = gl2.CreateProgram()
	vss, err := CompileShader(vs, gl2.VERTEX_SHADER)
	if err != nil {
		D(err)
	}
	ps.program.AttachShader(vss.Loc)
	gss, err := CompileShader(gs, gl2.GEOMETRY_SHADER)
	if err != nil {
		D(err)
	}
	ps.program.AttachShader(gss.Loc)
	fss, err := CompileShader(fs, gl2.FRAGMENT_SHADER)
	if err != nil {
		D(err)
	}
	ps.program.AttachShader(fss.Loc)
	ps.program.Link()
	if !ps.program.GetLinkStatus() {
		var logLength int32
		gl.GetProgramiv(uint32(ps.program), gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(uint32(ps.program), logLength, nil, gl.Str(log))
	}

	ps.gDeltaTimeMillis = ps.program.GetUniformLocation("gDeltaTimeMillis")
	ps.gTime = ps.program.GetUniformLocation("gTime")
	ps.gRandomTexture = ps.program.GetUniformLocation("gRandomTexture")
	ps.gLauncherLifetime = ps.program.GetUniformLocation("gLauncherLifetime")
	ps.gShellLifetime = ps.program.GetUniformLocation("gShellLifetime")
	ps.gSecondaryShellLifetime = ps.program.GetUniformLocation("gSecondaryShellLifetime")

	ps.particles[0].Lifetime = 0
	ps.particles[0].Position = position
	ps.particles[0].Velocity = glm.Vec3{}
	ps.particles[0].Type = Launcher

	for i := 0; i < len(ps.buffers); i++ {
		ps.transformfeedbacks[i].Bind()
		MustNotGLError()
		ps.buffers[i].Bind(gl2.ARRAY_BUFFER)
		MustNotGLError()
		ps.buffers[i].Data(gl2.ARRAY_BUFFER, int(unsafe.Sizeof(Particle{}))*len(ps.particles), unsafe.Pointer(&ps.particles[0]), gl2.DYNAMIC_DRAW)
		MustNotGLError()
		ps.transformfeedbacks[i].BindBufferBase(gl2.TRANSFORM_FEEDBACK_BUFFER, 0, ps.buffers[i])
		MustNotGLError()
	}
	return &ps
}

// Render renders this particle system
func (ps *ParticleSystem) Render(delta float64, VP glm.Mat4, camera glm.Vec3) {
	ps.time += delta

	//do update

	ps.gTime.Uniform1f(float32(ps.time) * 1000)
	ps.gDeltaTimeMillis.Uniform1f(float32(delta) * 1000)

	//bind the rand texture
	//m_randomTexture.Bind(RANDOM_TEXTURE_UNIT);

	gl.Enable(gl.RASTERIZER_DISCARD)
	ps.buffers[ps.currentVB].Bind(gl.ARRAY_BUFFER)
	ps.transformfeedbacks[ps.currentTFB].Bind()
	gl.EnableVertexAttribArray(0)
	gl.EnableVertexAttribArray(1)
	gl.EnableVertexAttribArray(2)
	gl.EnableVertexAttribArray(3)

	gl.VertexAttribPointer(0, 1, gl.FLOAT, false, int32(unsafe.Sizeof(Particle{})), IntToUnsafePointer(0))
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, int32(unsafe.Sizeof(Particle{})), IntToUnsafePointer(4))
	gl.VertexAttribPointer(2, 3, gl.FLOAT, false, int32(unsafe.Sizeof(Particle{})), IntToUnsafePointer(16))
	gl.VertexAttribPointer(3, 1, gl.FLOAT, false, int32(unsafe.Sizeof(Particle{})), IntToUnsafePointer(28))

	gl.BeginTransformFeedback(gl.POINTS)

	if ps.isFirst {
		gl.DrawArrays(gl.POINTS, 0, 1)
		ps.isFirst = false
	} else {
		gl.DrawTransformFeedback(gl.POINTS, uint32(ps.transformfeedbacks[ps.currentVB]))
	}
	gl.EndTransformFeedback()
	gl.DisableVertexAttribArray(0)
	gl.DisableVertexAttribArray(1)
	gl.DisableVertexAttribArray(2)
	gl.DisableVertexAttribArray(3)

	//do render
	ps.currentVB = ps.currentTFB
	ps.currentTFB = (ps.currentTFB + 1) & 0x1
}

var fs = `#version 330

void main()
{
}`

var vs = `#version 330                                                                        
																					
layout (location = 0) in float Type;                                                
layout (location = 1) in vec3 Position;                                             
layout (location = 2) in vec3 Velocity;                                             
layout (location = 3) in float Age;                                                 
																					
out float Type0;                                                                    
out vec3 Position0;                                                                 
out vec3 Velocity0;                                                                 
out float Age0;                                                                     
																					
void main()                                                                         
{                                                                                   
	Type0 = Type;                                                                   
	Position0 = Position;                                                           
	Velocity0 = Velocity;                                                           
	Age0 = Age;                                                                     
}`

var gs = `#version 330                                                                        
																					
layout(points) in;                                                                  
layout(points) out;                                                                 
layout(max_vertices = 30) out;                                                      
																					
in float Type0[];                                                                   
in vec3 Position0[];                                                                
in vec3 Velocity0[];                                                                
in float Age0[];                                                                    
																					
out float Type1;                                                                    
out vec3 Position1;                                                                 
out vec3 Velocity1;                                                                 
out float Age1;                                                                     
																					
uniform float gDeltaTimeMillis;                                                     
uniform float gTime;                                                                
uniform sampler1D gRandomTexture;                                                   
uniform float gLauncherLifetime;                                                    
uniform float gShellLifetime;                                                       
uniform float gSecondaryShellLifetime;                                              
																					
#define PARTICLE_TYPE_LAUNCHER 0.0f                                                 
#define PARTICLE_TYPE_SHELL 1.0f                                                    
#define PARTICLE_TYPE_SECONDARY_SHELL 2.0f                                          
																					
vec3 GetRandomDir(float TexCoord)                                                   
{                                                                                   
	 vec3 Dir = texture(gRandomTexture, TexCoord).xyz;                              
	 Dir -= vec3(0.5, 0.5, 0.5);                                                    
	 return Dir;                                                                    
}                                                                                   
																					
void main()                                                                         
{                                                                                   
	float Age = Age0[0] + gDeltaTimeMillis;                                         
																					
	if (Type0[0] == PARTICLE_TYPE_LAUNCHER) {                                       
		if (Age >= gLauncherLifetime) {                                             
			Type1 = PARTICLE_TYPE_SHELL;                                            
			Position1 = Position0[0];                                               
			vec3 Dir = GetRandomDir(gTime/1000.0);                                  
			Dir.y = max(Dir.y, 0.5);                                                
			Velocity1 = normalize(Dir) / 20.0;                                      
			Age1 = 0.0;                                                             
			EmitVertex();                                                           
			EndPrimitive();                                                         
			Age = 0.0;                                                              
		}                                                                           
																					
		Type1 = PARTICLE_TYPE_LAUNCHER;                                             
		Position1 = Position0[0];                                                   
		Velocity1 = Velocity0[0];                                                   
		Age1 = Age;                                                                 
		EmitVertex();                                                               
		EndPrimitive();                                                             
	}                                                                               
	else {                                                                          
		float DeltaTimeSecs = gDeltaTimeMillis / 1000.0f;                           
		float t1 = Age0[0] / 1000.0;                                                
		float t2 = Age / 1000.0;                                                    
		vec3 DeltaP = DeltaTimeSecs * Velocity0[0];                                 
		vec3 DeltaV = vec3(DeltaTimeSecs) * (0.0, -9.81, 0.0);                      
																					
		if (Type0[0] == PARTICLE_TYPE_SHELL)  {                                     
			if (Age < gShellLifetime) {                                             
				Type1 = PARTICLE_TYPE_SHELL;                                        
				Position1 = Position0[0] + DeltaP;                                  
				Velocity1 = Velocity0[0] + DeltaV;                                  
				Age1 = Age;                                                         
				EmitVertex();                                                       
				EndPrimitive();                                                     
			}                                                                       
			else {                                                                  
				for (int i = 0 ; i < 10 ; i++) {                                    
					 Type1 = PARTICLE_TYPE_SECONDARY_SHELL;                         
					 Position1 = Position0[0];                                      
					 vec3 Dir = GetRandomDir((gTime + i)/1000.0);                   
					 Velocity1 = normalize(Dir) / 20.0;                             
					 Age1 = 0.0f;                                                   
					 EmitVertex();                                                  
					 EndPrimitive();                                                
				}                                                                   
			}                                                                       
		}                                                                           
		else {                                                                      
			if (Age < gSecondaryShellLifetime) {                                    
				Type1 = PARTICLE_TYPE_SECONDARY_SHELL;                              
				Position1 = Position0[0] + DeltaP;                                  
				Velocity1 = Velocity0[0] + DeltaV;                                  
				Age1 = Age;                                                         
				EmitVertex();                                                       
				EndPrimitive();                                                     
			}                                                                       
		}                                                                           
	}                                                                               
}`
