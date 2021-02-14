# Real time ranking of players

Swagger REST API where clients can register absolute and relatives scores for players and that can return absoulte and relative tops


## Endpoints

```bash
POST: v1/add/score/playerId/{playerId}/score/{points}/isRelative/{isRelative}  -- setting/modifing score of the players
```
Parameters:
playerId (int) - the id of the player to be modified
points (string) - format can be either "+number", either "-number" to increase/decrease the number of point, either "number" to set the number of points if the parameter "isRelative" is set to false
isRelative (boolean) - to see if we add/substract or set points

```bash

GET:  v1/top/{number}  -- Getting TOP N players
```
number (int) - the number of player to show in the top

```bash
GET: /top/position/{position}/players/{players} -- Getting the relative rankings
```
position (int) - the relative position for the top to be displayed (the centre)
players (int) - how many players to display up and down according to the relative position

## Technical description

In the package "handlers" there is implemented the logic of the endpoints

In the package "playerstorage" we have a thread-safe singleton object to act as "storage unity"

In the package "utils" there are some helpers for the logic

In the package "tests" there are the unit tests for the methods (Not 100% coverage)

Functional testing was made with POSTMAN. In the main.go file there is some data added to the "storage"



