package lux

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	gl2 "github.com/luxengine/gl"
	"github.com/luxengine/glm"
	"github.com/luxengine/lux/utils"
)

//Mesh is an interface to represent any renderable mesh
type Mesh interface {
	Bind()
	Delete()
	Size() int32
	DrawCall()
}

//VUNMesh is a Vertex-Uv-Normal mesh
type VUNMesh struct { //Vertex, Uv, Normal Model
	VAO                              gl2.VertexArray
	Indices, Positions, Uvs, Normals gl2.Buffer
	Msize                            int32
}

//NewWavefrontModelFromFile loads a wavefront from the given file. Can only load files that are triangulated and with UV. Does not do anythign with material property.
func NewWavefrontModelFromFile(file string) Mesh {
	//load object
	meshObj := utils.LoadObject(file, false)

	//prepare indices //to optimise later
	indices, indexedVertices, indexedUvs, indexedNormals := utils.IndexVBOSlow(meshObj.Vertices, meshObj.UVs, meshObj.Normals)
	return NewVUNModel(indices, indexedVertices, indexedUvs, indexedNormals)
}

//NewVUNModel process and uploads the data to the GPU.
func NewVUNModel(indices []uint16, indexedVertices []glm.Vec3, indexedUvs []glm.Vec2, indexedNormals []glm.Vec3) Mesh {

	m := VUNMesh{}
	m.VAO = gl2.GenVertexArray()
	m.VAO.Bind()
	defer m.VAO.Unbind()

	m.Msize = int32(len(indices))
	//create a bunch of buffers and fill them
	//Positions
	m.Positions = gl2.GenBuffer()
	m.Positions.Bind(gl.ARRAY_BUFFER)
	m.Positions.Data(gl.ARRAY_BUFFER, len(indexedVertices)*3*4, ptr(indexedVertices), gl.STATIC_DRAW)

	//Uvs
	m.Uvs = gl2.GenBuffer()
	m.Uvs.Bind(gl.ARRAY_BUFFER)
	m.Uvs.Data(gl.ARRAY_BUFFER, len(indexedUvs)*2*4, ptr(indexedUvs), gl.STATIC_DRAW)

	//Normals
	m.Normals = gl2.GenBuffer()
	m.Normals.Bind(gl.ARRAY_BUFFER)
	m.Normals.Data(gl.ARRAY_BUFFER, len(indexedNormals)*3*4, ptr(indexedNormals), gl.STATIC_DRAW)

	//indices
	m.Indices = gl2.GenBuffer()
	m.Indices.Bind(gl.ELEMENT_ARRAY_BUFFER)
	m.Indices.Data(gl.ELEMENT_ARRAY_BUFFER, len(indices)*2, ptr(indices), gl.STATIC_DRAW)

	return &m
}

//NewVUNModelGlm process and uploads the data to the GPU.
func NewVUNModelGlm(indices []uint16, indexedVertices []glm.Vec3, indexedUvs []glm.Vec2, indexedNormals []glm.Vec3) Mesh {

	m := VUNMesh{}
	m.VAO = gl2.GenVertexArray()
	m.VAO.Bind()
	defer m.VAO.Unbind()

	m.Msize = int32(len(indices))
	//create a bunch of buffers and fill them
	//Positions
	m.Positions = gl2.GenBuffer()
	m.Positions.Bind(gl.ARRAY_BUFFER)
	m.Positions.Data(gl.ARRAY_BUFFER, len(indexedVertices)*3*4, ptr(indexedVertices), gl.STATIC_DRAW)

	//Uvs
	m.Uvs = gl2.GenBuffer()
	m.Uvs.Bind(gl.ARRAY_BUFFER)
	m.Uvs.Data(gl.ARRAY_BUFFER, len(indexedUvs)*2*4, ptr(indexedUvs), gl.STATIC_DRAW)

	//Normals
	m.Normals = gl2.GenBuffer()
	m.Normals.Bind(gl.ARRAY_BUFFER)
	m.Normals.Data(gl.ARRAY_BUFFER, len(indexedNormals)*3*4, ptr(indexedNormals), gl.STATIC_DRAW)

	//indices
	m.Indices = gl2.GenBuffer()
	m.Indices.Bind(gl.ELEMENT_ARRAY_BUFFER)
	m.Indices.Data(gl.ELEMENT_ARRAY_BUFFER, len(indices)*2, ptr(indices), gl.STATIC_DRAW)

	return &m
}

//Bind the vertex array and all vertex attrib required to render this mesh.
func (m *VUNMesh) Bind() {
	vao := m.VAO
	vao.Bind()

	vao.EnableVertexAttribArray(0)
	m.Positions.Bind(gl.ARRAY_BUFFER)
	vao.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	vao.EnableVertexAttribArray(1)
	m.Uvs.Bind(gl.ARRAY_BUFFER)
	vao.VertexAttribPointer(1, 2, gl.FLOAT, false, 0, nil)

	vao.EnableVertexAttribArray(2)
	m.Normals.Bind(gl.ARRAY_BUFFER)
	vao.VertexAttribPointer(2, 3, gl.FLOAT, false, 0, nil)

	m.Indices.Bind(gl.ELEMENT_ARRAY_BUFFER)
}

//Unbind all the resources.
func (m *VUNMesh) Unbind() {
	m.VAO.Unbind()
}

//Delete all allocated resources (buffers, vertexarray,etc).
func (m VUNMesh) Delete() {
	defer m.Positions.Delete()
	defer m.Uvs.Delete()
	defer m.Normals.Delete()
	defer m.Indices.Delete()
	defer m.VAO.Delete()
}

//Size returns the amount of verteices to be drawn.
func (m *VUNMesh) Size() int32 {
	return m.Msize
}

//DrawCall send a single draw call
func (m *VUNMesh) DrawCall() {
	gl.DrawElements(gl.TRIANGLES, m.Size(), gl.UNSIGNED_SHORT, nil)
}
