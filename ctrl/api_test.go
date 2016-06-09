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
			"i18nkey": "gameplay",
			"stickpadgyro":[
				{
					"name": "move",
					"i18nkey": "move",
					"inputmode": "joystick_move"
				}
			],
			"buttons":[
				{
					"name": "jump",
					"i18nkey": "jump"
				},
				{
					"name": "fire",
					"i18nkey": "fire"
				}
			]
		}
	],
	"localization":[
		{
			"language": "english",
			"mapping":[
				{
					"i18nkey": "jump",
					"localkey": "Jump"
				},
				{
					"i18nkey": "fire",
					"localkey": "Fire"
				}]
		}
	]
}`))
	err := LoadCtrlFormat(r)
	if err != nil {
		t.Error(err)
	}
}
