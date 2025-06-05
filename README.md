# Welcome to the Picture-Pokedex
This is my implementation of the boot dot dev pokedex project. Currently it is as was designed in the aforementioned guided project but I had an idea.

I wanted to implement transforming the images that are sent down from pokeapi from their original PNG state to ascii art and learn how these technologies work under the hood.

It took awhile but I read the spec knocked out a very ***Incomplete Parser*** and now I can parse the PNG data for the pokemon png's that are provided : - ).

This currently includes decoding images that use indexed and true color with alpha images at this point. It also only currently includes defiltering images that have been either no scan line filtering or sub scanline filtering.


### TODO
1. [ ] Handling pokemon forms/species

2. [x] Defiltering Scanlines
- [x] None
- [x] Sub
- [x] Up (not tested)
- [x] Average
- [x] Paeth

3. [ ] Color Types
- [x] Indexed colors
- [x] Truecolor w/ alpha
- [ ] Truecolor
- [ ] grayscale
- [ ] grayscale w/ alpha

## Building the Project

Currently there are no plans to provide builds for this as it is really just something intended for learning but thanks to go it's really simple to build.

First make sure you have go installed by entering `go --version` and seeing if it's installed.  If it isn't then I recommend installing it using [webi](https://webinstall.dev/golang/). it makes installing things wayyyyy easier.

Once you have go installed just git clone the repo into the dir you want to build in and type `go build` into the root dir of the project and you will have a working pokedex. 

To explore the commands simply type in the help command and it will give an explanation of all of the commands.  It has yet to write to a file so all runs are just saved to memory and dumped at program termination but that's another commit for another day.

## Thanks for checking out my project!
