package main

import (
	"cdn/utils"
)

func main(){
disPathCSS := "./src/dist/css"
disPathJS := "./src/dist/js"
utils.BundleAndMinifyAllCSS(disPathCSS)
utils.BundleAndMinifySingleCSSFile("./src/CSS/tokens.css", disPathCSS)
utils.BundleAndMinifySingleCSSFile("./src/CSS/utils.css", disPathCSS)
utils.BundleAndMinifyAllJS(disPathJS)
//utils.LoopOverIcons("./src/CSS/icons.css", "./src/JS/icons.js")
}
