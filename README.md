# pubg-heat-drop

This is a website and backend service to let multiple people vote on a map location in PUBG.
The votes should be shown via a heat map overlay.
It is mostly intended for twitch streamers, so their viewer can decide where the streamer should jump to.

## Idea

The Idea for this was proposed by [Heawin](https://twitch.tv/heawin)

## Setup

### Requirements

- Bun installed

### Initial setup

1. Clone the repository

   ```bash
   git clone git@github.com:tiluk/pubg-heat-drop.git
   ```

1. Add High_Res files to /src/assets/maps/<map_name>

   - The files can be found in the [pubg api assetes](https://github.com/pubg/api-assets/tree/master/Assets/Maps)

1. Install the dependencies
   ```bash
   bun install
   ```

### Running the server

```bash
bun dev
```

##### Attributions

- [Leaflet](https://leafletjs.com/)
