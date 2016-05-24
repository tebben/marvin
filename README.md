# marvin
Home automation project written in Go, easily add functionality by plugging in new modules into Marvin. The idea is to create a client which can display all events and actions of the running modules and let a user add triggers to Marvin.
For instance: When the module BigRedButton sends the event pushed trigger module.Spotify.play. Marvin will support a rest endpoint for gathering information and websockets so all actions can be triggered easily from other applications. 

## Modules planned
- Philips HUE
- Arduino 
- [GOST](https://github.com/Geodan/gost)
- Spotify
- .....

## Progress
- Initial setup work
- Setting up Hue module

## Roadmap
- Rest interface for retrieving info from Marvin on modules and triggers
- Targetted websocket messages
- Client
- IFTT like options to create triggers
- Documentation
- Modules and more modules
