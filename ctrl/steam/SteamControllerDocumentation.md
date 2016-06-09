# Disclaimer
This documentation is available at https://partner.steamgames.com/documentation/steamcontroller, it requires you to be logged in to a valid steamworks account. I don't know if it's legal to actually post this here but since not all lux dev have a steamworks account I'll keep it here until Valve asks me to take it down.

# Overview
The Steam Controller API is designed to allow you to easily enable full Steam controller support in your game. We define full support as the follows:

- Your game uses the Steam Controller glyphs when showing in-game input prompts.
- The Controller Configuration screen in Steam uses in-game actions that the player performs in your game, instead of keys or buttons.
- You've published an Official configuration for the controller.
- Your game doesn't restrict the user's ability to customize their controls. This means it allows any mix of mouse, keyboard, or gamepad input simultaneously.
- When your game wants keyboard input (e.g. when naming avatars), you call ISteamUtils::ShowGamepadTextInput() to automatically bring up the text entry UI.
- Your game has no launchers that require mouse or KB input - or even better, no launcher at all.

To ensure users have a good experience from the couch, we also recommend the following:
- Make your UI readable from several feet away. Our rule of thumb: when your game is running at 1920x1080, your fonts should be a minimum of 24px in size.
- Start your game in fullscreen by default when the user is running Steam Big Picture (the "SteamTenfoot" environment variable will be set)
- For bonus points, at first launch detect the user's screen resolution and set your resolution to match it.


The implementation process is straightforward, and shouldn't take more than a few days of work. Four steps are involved:

1. In a text editor, create an In-Game Actions file, which tells Steam what in-game actions your players can bind to the controller.
2. In Steam, use the Controller Configuration UI to create your default configuration.
3. In your game, use the Steamworks Controller API to read actions from the controller, and to retrieve appropriate glyphs for display.
4. Update your game depot with the new binaries, and publish your configuration as the Official config.
[Learn more about Steam Controllers](http://store.steampowered.com/livingroom/SteamController/)



# Step 1 - Creating an In-Game Actions File
Start by downloading the starting [In-Game Actions (IGA) file](https://steamcdn-a.akamaihd.net/steam/partner/controller/game_actions_X.vdf) . Place it in your "<Steam Install Directory>\controller_config" directory (create the directory if it doesn't exist). Rename the file to the following: "game_actions_X.vdf", where X is your game's Steam AppID. If you don't know what your game's Steam AppID is, it can be found by logging into your Steam partner site.

Now open up the file in your favorite text editor. The file is in a standard Valve format called KeyValues, which is a simple & easily read format. You can read about it [here](https://developer.valvesoftware.com/wiki/KeyValues), but you'll likely understand it just by reading the starting IGA file. You might also find it useful to download the [Portal 2 IGA File](https://steamcdn-a.akamaihd.net/steam/partner/controller/game_actions_example.vdf) so you can see an example.

## File Format

IGA files contain an "actions" section, which should list all the In-Game Action Sets (IGAS) in your game. An IGAS describes all the actions a player can take within some game context - such as when the player is in a vehicle, or on foot, or navigating the menu system. For each IGAS, the Steam Controller Configuration UI will provide a tab that allows a player to customize how those actions are bound to the controller.

An IGAS entry in the IGA file should contain a "title" key & value, and the following subsections: "StickPadGyro", "AnalogTrigger", and "Button". The "StickPadGyro" and "AnalogTrigger" sections each contain a list of IGAs that the player can only assign to the Stick/Pad/Gyro and Analog Triggers respectively. The "Buttons" section contains IGAs that can only be bound to digital inputs (like the physical ABXY buttons, or a trackpad that's in ABXY mode, or a Trigger that's not being used as an AnalogTrigger action).

## Button Actions

The format for "Button" actions is as follows:

	"<action name>" "#<localization key>"
The <action name> is the internal name you want to refer to this action by in your game code, when talking to the controller API. The <localization key> should be the name of an entry in your localization section (see Adding Localization). Make sure to include the '#' character at the start of your localization key.
## Analog Trigger Actions

The format for "AnalogTrigger" actions is the same as for "Button". "AnalogTrigger" actions should be actions that your game will be interpreting as full analog inputs, like vehicle acceleration. If you don't have any of that kind of input, just leave the section empty. The Controller Configuration UI will allow players to assign any "Button" action to the physical triggers if they're not being used as an "AnalogTrigger" action.

## StickPadGyro Actions

The format for "StickPadGyro" actions is as follows:

	"<action name>"
	{
		"title" 	"#<localization key>"
		"input_mode" 	"<analog mode>"
	}
	
<action name> and <localization key> are the same as the matching keys in the "Button" format. <analog mode> tells us how to interpret the data coming from the physical controls, before we pass it to your game via the controller API.
Valid <analog modes> are as follows:

- "absolute_mouse" - For when you're expecting the action to behave like a mouse. Useful for first or third person camera, or an actual mouse cursor.
- "joystick_move" - For when you're using the action to move a character around.

There is one optional setting for StickPadGyro actions that are using "absolute_mouse" as their input. If you set the "os_mouse" key to "1", we'll pass the input from the player into the OS as well as your game. This is useful if you have a visible OS mouse cursor that should be controlled by this action. Here's an example:

	"menu_mouse"
	{
		"title"		"#Menu_Mouse_Title"
		"input_mode"	"absolute_mouse"
		"os_mouse"	"1"
	}
	


# Step 1.1 - Adding Localization
The "localization" section in your IGA file contains a list of languages, each of which is a section containing localization keys & values. For example:

	"localization"
	{
		"english"
		{
			"Action_Jump" 	"Jump"
			"Action_Camera"	"Camera"
		}

		"german"
		{
			"Action_Jump" 	"Springen"
			"Action_Camera"	"Kameraansicht"
		}
	}
	
The actions in your IGA file should then specify their names by referring to the desired localization key, preceded by the '#' character. Here are example actions using the above localization keys:

	"StickPadGyro"
	{
		"Camera"
		{
			"title" 	"#Action_Camera"
			"input_mode"	"absolute_mouse"
		}
	}
	"Button"
	{
		"Jump" "#Action_Jump"
	}
	
If the language the game is running under is not found in the localization section, English will be used as a fallback. If English is not found, the string will be shown as is. Supported languages are as follows:

- brazilian
- bulgarian
- czech
- danish
- dutch
- english
- finnish
- french
- german
- greek
- hungarian
- italian
- japanese
- koreana
- korean
- norwegian
- polish
- portuguese
- romanian
- russian
- schinese
- spanish
- swedish
- tchinese
- thai
- turkish
- ukrainian

# Step 1.2 - Titles & Descriptions
Configurations also need a localized title and description. If you're only making a single official configuration for your game, you don't need to worry about this - we'll provide a default title and description, and you can skip this step entirely. But if you want to have multiple official configurations for your game, then you'll also need to provide localized titles and descriptions for them.

Configuration titles and descriptions should be listed along with the rest of your localization keys. The title key must start with "Title_", and the description key must start with "Description_". Here's an example:

	"localization"
	{
		"english"
		{
			"Title_Config1"		"Official Configuration"
			"Description_Config1"	"This config was created by the developers of Game X."
			"Title_Config2"		"Official Southpaw Configuration"
			"Description_Config2"	"This config was created by the developers of Game X, and is setup for Southpaw users."

			"Action_Jump" 	"Jump"
			"Action_Camera"	"Camera"
		}
	}
	
When you publish a configuration (see [Step 4 - Publishing](https://partner.steamgames.com/documentation/steamcontroller#step4)), you'll be able to select which of these localized titles and descriptions you want to use.


# Step 2 - Creating a Default Configuration
Once you've created your IGA file, and ensured it's in the right directory & named to match your game's AppID, you're ready to create a configuration. Run Steam in Big Picture mode, and navigate to your game's Game Details page. Select Manage Game, and then Configure Controller. If you receive any errors at this point, they'll be identifying issues in your IGA file and you'll need to go fix them - the most common mistake is a missing closing quote or brace.

If you don't receive any errors, you're now looking at an empty controller configuration for your game, and it should be fully aware of your in-game actions. Use the UI to create a default configuration. Make sure you set defaults for all your In-Game Action Sets, not just the first one.

Once you've got a configuration, save it privately. Don't publish it, because your game is not yet ready to receive IGAs.



# Step 3 - The Steam Controller API
Make sure you have the latest version of the Steamworks API - Grab it from the [Getting Started With SteamWorks](https://partner.steamgames.com/documentation/getting_started) page.

Add the following command line parameters to your steam shortcut:

	-forcecontrollerappid <your game's AppID>
This will tell Steam to keep the controller locked to your game. Usually, the controller flips its configuration as focus shifts from your game to Steam or the desktop, because we use different configurations for each of those states. This can make it trickier to debug the controller while you're working on it, because the controller will change its configuration when you hit a breakpoint in your debugger. Setting this -forcecontrollerappid option will keep the controller locked to your game. This may have the side effect of making it less/non-functional in the desktop or Steam overlay.
Download the official controller glyphs.



## Step 3.1 - API Overview
The Steam Controller API is built around the In-Game Actions & Sets that you specified in your In-Game Actions File in Step 1. The functionality overview is as follows:

- Use GetConnectedControllers() to detect when a Steam controller exists. Use the handles it returns to identify controllers throughout the rest of the API.
- On game startup, use the GetActionSetHandle(), GetDigitalActionHandle(), and GetAnalogActionHandle() functions to resolve the action & set names into handles. Use these handles to identify actions & sets throughout the rest of the API.
- Use GetDigitalActionData() and GetAnalogActionData() to read the input state of all your in-game actions.
- Use GetDigitalActionOrigins() and GetAnalogActionOrigins() to get the origin(s) that the player has assigned to your in-game actions, and show the appropriate glyphs in your onscreen UI.
- Use ShowBindingPanel() to pop up the Steam Overlay directly into the controller configuration for a specified controller. This means you won't need to make a customization UI.


## Step 3.2 - Input Handling
For digital actions, the data returned by GetDigitalActionOrigins() is straightforward:

	bState: will be true if the action is being sent by the controller (button is pressed, trigger pulled, etc)
For analog actions, the data returned by GetAnalogActionData() is only slightly more complex:
	x,y: will depend on the mode you assigned the game action to in your IGA File.
		- "absolute_mouse": x & y will be deltas from the previous mouse position.
		- "joystick_move": x & y will be values ranging between -1 & 1, representing the current position of the joystick.
	
While you're in your code, make sure you're not preventing the user from combining different kinds of input simultaneously. A common mistake we've seen is for games to start ignoring gamepad input when they see mouse & KB, or vice versa.



## Step 3.3 - On-screen Glyphs
To display an on-screen prompt for the steam controller, you need to get the physical origins that are bound to an in-game action. The player may have bound more than one physical origin to the same action, so it's best to have your UI cycle through displaying each origin. GetDigitalActionOrigins() and GetAnalogActionOrigins() will return the count of origins for the specified action, and fill out the passed in originsOut array.

For each origin, you can use the EControllerActionOrigin enum to remap the origin to the matching .png file in the official controller glyphs. You'll need to include the [official controller glyphs](https://steamcdn-a.akamaihd.net/steam/partner/controller/SteamControllerGlyphs_v1.zip) in your game depot, packaged in whatever way you need to be able to render them in-game.

Please note that the user can change their configuration at any time. To help with this, we've ensured the GetDigitalActionOrigins() and GetAnalogActionOrigins() functions are extremely cheap to call. When you display an onscreen prompt, don't cache the origins and keep displaying the first results. Instead, we recommend you re-gather the origins each frame, and display the matching prompts. That way, if a user decides to change their configuration as a result of seeing the prompt, when they return from the configuration screen the prompts will automatically update to match the new origins.



# Step 4 - Publishing
Once your game is working with the controller, you're ready to publish. You'll need to release your new game update and make your configuration the Official one. Official configurations are automatically loaded when a player launches your game for the first time. This allows users to simply fire up your game and play without needing to go into the configuration screen at all.

Here's the recommended process:

4.1 - Update your game
- Update your steam depot with your new version of the game.
4.2 - Publish your configuration
- Run Steam in Big Picture mode and navigate to the default configuration you created in the controller configuration.
- Hit the (Y) button, or click the [Save As] button at the bottom of the screen.
- Select the desired localized title & description in the popup, change it to a PUBLIC profile, and click the PUBLISH button.
- You will get a confirmation, and the file ID of the public configuration. Copy/save the ID.

4.3 - Make the configuration official
- Open up your game's page on the Steam Partner website.
- Select 'Edit Steamworks Settings' in the Technical Tools section.
- Scroll to the section called 'Controller Template'. Select the 'Custom Configuration' option in the dropdown.
- Paste the file ID of your public configuration into the box, and click the Save button.
- Publish your app changes in the partner site as usual.
You can keep making changes to your official configuration even after you've published and made it official. All you have to do is re-publish the configuration through the Steam Controller Configuration screen, and users will get the updated one automatically. You won't need to go through step 4C of making it official again.

If you'd like, you can also specify multiple official configurations for your game. You may want to create an official Southpaw mode, for example. To do this, simply create multiple configurations and publish them each via step 4.2. Then, in step 4.3, paste all the file IDs for your configurations into the box, separated by the comma character (,) characters. The first one on the list will be considered the highest priority, and will be picked by default for new players. Don't forget to make Titles & Descriptions for each of them.



# Hints & Tips
We recommend creating a separate In-Game Action Set for your menu controls, instead of just re-using actions in your main game set. Most customers won't need to modify this menu controls set, but it's an easy way to provide the capability to players who actually need it (as can be the case for some disabled gamers, for example).
