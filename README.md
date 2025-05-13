# Concentric Rings Puzzle

There are two rings, _outer_ and _inner_, inner has _n_ pegs and outer has _n_ holes.
Player moves one peg from inner ring to the same position in the outer ring and then rotates the inner ring clockwise.
The game is over when there are no available moves. Game is considered a victory if there are no pegs in the inner ring and a loss otherwise.

## Example Game

To better demonstrate the game, here is an example of a winning game with _n = 3_. Positions are labeled 0, 1, 2, position labels do not move when the ring rotates.

#### Initial setup

![Initial setup](assets/step0.png?raw=true "Initial setup")

The game always starts with move at position 0 (`Move(0)`), moving anywhere else is an isomorphism of `Move(0)`.

![Position after the first move](assets/step1.png?raw=true "Position after the first move")

Next we do `Move(2)`.

![Winning position](assets/step2.png?raw=true "Position after 2 moves")

Lastly the only available move is `Move(1)`.

![Winning position](assets/step3.png?raw=true "Winning position")

After `Move(1)` all the inner pegs have been moved to the outer ring, we have won the game.

# Programmatic solution

This repository recreates this puzzle (`package puzzle`) and a solver (`package solver`), which brute forces the puzzle to find solutions.
