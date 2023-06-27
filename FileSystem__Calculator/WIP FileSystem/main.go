// COMPOSITE DESIGN PATTERN
// used when in case of tree like structure
package main

import "filesystem/Fs"

func main() {
	movieDirectory := Fs.NewDirectory("Movie")

	border := Fs.NewFile("Border")
	movieDirectory.Add(border)

	comedyMovieDirectory := Fs.NewDirectory("ComedyMovie")
	hulchul := Fs.NewFile("Hulchul")
	comedyMovieDirectory.Add(hulchul)
	movieDirectory.Add(comedyMovieDirectory)

	movieDirectory.LS()
}
