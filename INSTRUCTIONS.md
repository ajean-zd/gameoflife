## What is game of life?

Conway’s Game of Life is a beloved computer simulation and thought experiment, first invented in 1970. It demonstrates that a system with a few simple rules can produce complex and unpredictable outcomes. The world of Life is a 2 dimensional grid of cells that are either alive or dead. From one generation to the next there are 4 rules that define which cells live and die:
* Any live cell with fewer than two live neighbours dies, as if by underpopulation.
* Any live cell with two or three live neighbours lives on to the next generation.
* Any live cell with more than three live neighbours dies, as if by overpopulation.
* Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.

## The challenge

Your challenge for Code Club Q1 2022 is to implement the Game of Life with a user interface of your choice. You can complete this challenge with a pair, or by yourself if you prefer.

User Interface Inspiration
All of these examples, and many more that you think up, are perfectly valid. Choose an interface that you think will best allow you to practice the software engineering skills that you’re most interested in developing.

A purely server side app that uses a table full of checkboxes in a html form tag to accept the initial state, the next state is rendered server side and presented in a html table

A command line app that reads the initial state from a csv and generates a gif animation

A purely client side React app that uses interactive SVG to let users draw the game state and watch it autoplay

An app that communicates with a physical midi controller, or digital midi controller, to light buttons and accept user input

A purely audio playback system where users chose from a library of initial states and listen to the simulation

## Things you may consider

What happens when a cell reaches the edge?
Does it wrap around to the other side?
Could you make an infinite grid with no edges?
How do users input the initial state?
Can users pause the game and make changes? Or control the playback speed?
What data structures could you use to represent the game state?
What will be the most fun to implement, or beautiful to watch?
Can users step backwards, as well as forwards, through time?
How fast is your solution? Can it handle big games?
