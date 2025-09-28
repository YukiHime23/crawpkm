# pkmcraw
Craw data from website [www.pokemonspecial.com](http://www.pokemonspecial.com)

## Installation

1. Install Go: https://golang.org/dl/
2. Run `go get -u github.com/YukiHime23/pkmcraw/cmd/pkmcraw`

## Usage

- Run `pkmcraw` in terminal
- Wait for crawling done
- Check the result in folder `Pokemon`

## Note

- This project use [gocolly](https://github.com/gocolly/colly) for crawling
- This project use [crawpkm](https://github.com/YukiHime23/crawpkm) for download image
- If you want to change the start chapter, you can change the variable `LinkPkmCraw` in file `cmd/pkmcraw/main.go`
- If you want to change the folder to save the data, you can change the variable `pathPkm` in file `cmd/pkmcraw/main.go`
