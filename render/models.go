package lux

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	gl2 "github.com/luxengine/lux/gl"
	"github.com/luxengine/lux/glm"
	"github.com/luxengine/lux/render/utils"
	"unsafe"
)

// Mesh is an interface to represent any renderable mesh
type Mesh interface {
	Bind()
	Delete()
	Size() int32
	DrawCall()
}

// VUNMesh is a Vertex-Uv-Normal mesh
type VUNMesh struct { //Vertex, Uv, Normal Model
	VAO                              gl2.VertexArray
	Indices, Positions, Uvs, Normals gl2.Buffer
	Msize                            int32
}

// NewWavefrontModelFromFile loads a wavefront from the given file. Can only
// load files that are triangulated and with UV. Does not do anything with
// material property.
func NewWavefrontModelFromFile(file string) Mesh {
	//load object
	meshObj := utils.LoadObject(file, false)

	//prepare indices //TODO(hydroflame): optimise
	indices, indexedVertices, indexedUvs, indexedNormals := utils.IndexVBOSlow(meshObj.Vertices, meshObj.UVs, meshObj.Normals)
	return NewVUNModel(indices, indexedVertices, indexedUvs, indexedNormals)
}

// NewVUNModel process and uploads the data to the GPU.
func NewVUNModel(indices []uint16, indexedVertices []glm.Vec3, indexedUvs []glm.Vec2, indexedNormals []glm.Vec3) Mesh {
	var m VUNMesh

	m.VAO = gl2.GenVertexArray()
	m.VAO.Bind()

	m.Msize = int32(len(indices))
	//create a bunch of buffers and fill them
	//Positions
	m.Positions = gl2.GenBuffer()
	m.Positions.Bind(gl2.ARRAY_BUFFER)
	m.Positions.Data(gl2.ARRAY_BUFFER, len(indexedVertices)*3*4, unsafe.Pointer(&indexedVertices[0]), gl2.STATIC_DRAW)

	//Uvs
	m.Uvs = gl2.GenBuffer()
	m.Uvs.Bind(gl2.ARRAY_BUFFER)
	m.Uvs.Data(gl2.ARRAY_BUFFER, len(indexedUvs)*2*4, unsafe.Pointer(&indexedUvs[0]), gl2.STATIC_DRAW)

	//Normals
	m.Normals = gl2.GenBuffer()
	m.Normals.Bind(gl2.ARRAY_BUFFER)
	m.Normals.Data(gl2.ARRAY_BUFFER, len(indexedNormals)*3*4, unsafe.Pointer(&indexedNormals[0]), gl2.STATIC_DRAW)

	//indices
	m.Indices = gl2.GenBuffer()
	m.Indices.Bind(gl2.ELEMENT_ARRAY_BUFFER)
	m.Indices.Data(gl2.ELEMENT_ARRAY_BUFFER, len(indices)*2, unsafe.Pointer(&indices[0]), gl2.STATIC_DRAW)

	m.VAO.Unbind()
	return &m
}

// NewVUNModelGlm process and uploads the data to the GPU.
func NewVUNModelGlm(indices []uint16, indexedVertices []glm.Vec3, indexedUvs []glm.Vec2, indexedNormals []glm.Vec3) Mesh {
	var m VUNMesh

	m.VAO = gl2.GenVertexArray()
	m.VAO.Bind()

	m.Msize = int32(len(indices))
	//create a bunch of buffers and fill them
	//Positions
	m.Positions = gl2.GenBuffer()
	m.Positions.Bind(gl2.ARRAY_BUFFER)
	m.Positions.Data(gl2.ARRAY_BUFFER, len(indexedVertices)*3*4, unsafe.Pointer(&indexedVertices[0]), gl2.STATIC_DRAW)

	//Uvs
	m.Uvs = gl2.GenBuffer()
	m.Uvs.Bind(gl2.ARRAY_BUFFER)
	m.Uvs.Data(gl2.ARRAY_BUFFER, len(indexedUvs)*2*4, unsafe.Pointer(&indexedUvs[0]), gl2.STATIC_DRAW)

	//Normals
	m.Normals = gl2.GenBuffer()
	m.Normals.Bind(gl2.ARRAY_BUFFER)
	m.Normals.Data(gl2.ARRAY_BUFFER, len(indexedNormals)*3*4, unsafe.Pointer(&indexedNormals[0]), gl2.STATIC_DRAW)

	//indices
	m.Indices = gl2.GenBuffer()
	m.Indices.Bind(gl2.ELEMENT_ARRAY_BUFFER)
	m.Indices.Data(gl2.ELEMENT_ARRAY_BUFFER, len(indices)*2, unsafe.Pointer(&indices[0]), gl2.STATIC_DRAW)

	m.VAO.Unbind()
	return &m
}

// Bind the vertex array and all vertex attrib required to render this mesh.
func (m *VUNMesh) Bind() {
	vao := m.VAO
	vao.Bind()

	vao.EnableVertexAttribArray(0)
	m.Positions.Bind(gl2.ARRAY_BUFFER)
	vao.VertexAttribPointer(0, 3, gl2.FLOAT, false, 0, nil)

	vao.EnableVertexAttribArray(1)
	m.Uvs.Bind(gl2.ARRAY_BUFFER)
	vao.VertexAttribPointer(1, 2, gl2.FLOAT, false, 0, nil)

	vao.EnableVertexAttribArray(2)
	m.Normals.Bind(gl2.ARRAY_BUFFER)
	vao.VertexAttribPointer(2, 3, gl2.FLOAT, false, 0, nil)

	m.Indices.Bind(gl2.ELEMENT_ARRAY_BUFFER)
}

// Unbind all the resources.
func (m *VUNMesh) Unbind() {
	m.VAO.Unbind()
}

// Delete all allocated resources (buffers, vertexarray, etc).
func (m VUNMesh) Delete() {
	defer m.Positions.Delete()
	defer m.Uvs.Delete()
	defer m.Normals.Delete()
	defer m.Indices.Delete()
	defer m.VAO.Delete()
}

// Size returns the amount of vertices to be drawn.
func (m *VUNMesh) Size() int32 {
	return m.Msize
}

// DrawCall send a single draw call
func (m *VUNMesh) DrawCall() {
	gl.DrawElements(gl2.TRIANGLES, m.Size(), gl2.UNSIGNED_SHORT, nil)
}
