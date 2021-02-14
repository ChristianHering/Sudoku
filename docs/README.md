Sudoku
===========

This package implements all the functionality to make a sudoku game, board generator, etc.

It provides:

  * Automated (valid) puzzle generation
  * A basic value hint function that lets you make sure an entry isn't invalid
  * An automated solving function that displays changes as it bruteforces the puzzle
  * An example of how one could package a game for the web without using flash

Table of Contents:

  * [Installing and Compiling from Source](#installing-and-compiling-from-source)
  * [Contributing](#contributing)
  * [License](#license)
  * [About](#about)

Installing and Compiling from Source
------------

The easiest way to get started is to install the lastest release from the [releases](https://github.com/ChristianHering/sudoku/releases) tab.


If you're looking to compile from source, you'll need the following:

  * [Go](https://golang.org) installed and [configured](https://golang.org/doc/install)
  * [Lorca](https://github.com/zserge/lorca) installed properly

Contributing
------------

Contributions are always welcome. If you're interested in contributing, send me an email or submit a PR.

License
-------

Feel free to use this project in any way you like, so long as it's open source. Please refer to the [license](/LICENSE) file for more information.

About
-----

[Sudoku](https://en.wikipedia.org/wiki/Sudoku) is a puzzle game that requires you to fill in the remainder of a 9x9 grid without reusing digits. In any given row, column, or nonet (square) there will be numbers 1-9, but there will be no repeating digits.