package lux

import (
	"github.com/luxengine/gl"
	"log"
	"strings"
)

// AssetManager keeps track of loaded textures, models and programs
type AssetManager struct {
	modelsDir, shadersDir, texturesDir string
	Models                             map[string]Mesh
	Textures                           map[string]gl.Texture2D
	Programs                           map[string]gl.Program
}

// NewAssetManager makes a new asset manager
//	-root: the root of all the other folders. eg. "assets"
//	-models: location of models. eg. "models", located at "assets/models"
//	-shaders: location of shaders. Not really used right now because everything is hard coded :\. eg. "shaders", located at "assets/shaders"
//	-textures: location of texture. eg. "textures", located at "assets/textures"
func NewAssetManager(root, models, shaders, textures string) AssetManager {
	return AssetManager{
		modelsDir:   root + models,
		shadersDir:  root + shaders,
		texturesDir: root + textures,
		Models:      make(map[string]Mesh),
		Textures:    make(map[string]gl.Texture2D),
		Programs:    make(map[string]gl.Program),
	}
}

// LoadModel loads a single model. Only wavefront available for now. iname is
// the internal name to be set in the map.
func (am *AssetManager) LoadModel(name, iname string) {
	if strings.HasSuffix(name, ".obj") {
		am.Models[iname] = NewWavefrontModelFromFile(am.modelsDir + name)
	} else {
		log.Fatal("cannot find " + name)
	}
}

// LoadRenderProgram is suppose to load a render program, although with the
// geometry buffer takes care of most of it. Do not use.
func (am *AssetManager) LoadRenderProgram(vertexShader, fragmentShader, iname string) {
	//program, err := LoadProgram(am.shadersDir+"standard.vert", am.shadersDir+"standard.frag")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//am.Programs[iname] = program
}

// LoadTexture Load an image as a texture2D. iname is the internal name to be
// set in the map.
func (am *AssetManager) LoadTexture(name, iname string) {
	if strings.HasSuffix(name, ".png") {
		pic, err := LoadPng(am.texturesDir + name)
		if err != nil {
			log.Fatal(err)
		}
		am.Textures[iname] = pic
	} else {
		log.Fatal("unable to find texture " + (am.modelsDir + name))
	}
}

// Clean releases every asset.
func (am *AssetManager) Clean() {
	for name, model := range am.Models {
		model.Delete()
		delete(am.Models, name)
	}
	for name, tex := range am.Textures {
		tex.Delete()
		delete(am.Textures, name)
	}
	for name, prog := range am.Programs {
		prog.Delete()
		delete(am.Programs, name)
	}
}
