package ctrl

import (
	"bytes"
	"testing"
)

func TestLoadCtrlFormat(t *testing.T) {
	r := bytes.NewReader([]byte(`
{
	"actionsets":[
		{
			"name": "gameplay",
			"stickpadgyro":[
				{
					"name": "move",
					"inputmode": "joystick_move"
				}
			],
			"buttons":["jump", "fire"]
		}
	]
}`))
	_, err := LoadCtrlFormat(r)
	if err != nil {
		t.Error(err)
	}
}
