# Music Library API

## Overview
The Music Library API is designed to manage a collection of music tracks and playlists.

## Endpoints
1. `GET /tracks`: Get a list of all tracks
2. `GET /tracks/:id`: Get details of a specific track
3. `POST /tracks`: Add a new track
4. `PUT /tracks/:id`: Update details of a specific track
5. `DELETE /tracks/:id`: Delete a specific track
6. `GET /playlists`: Get a list of all playlists
7. `GET /playlists/:id`: Get details of a specific playlist
8. `POST /playlists`: Create a new playlist
9. `PUT /playlists/:id`: Update details of a specific playlist
10. `DELETE /playlists/:id`: Delete a specific playlist
11. `POST /playlists/:id/add/:trackId`: Add a track to a playlist
12. `POST /playlists/:id/remove/:trackId`: Remove a track from a playlist
