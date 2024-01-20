package main

import (
	"cdn/utils"
)

func main(){
disPath := "./src/dist/css"
utils.BundleAndMinifyAllCSS(disPath)
utils.BundleAndMinifySingleCSSFile("./src/CSS/ds.css", disPath)
utils.BundleAndMinifySingleCSSFile("./src/CSS/icons.css", disPath)
utils.BundleAndMinifySingleCSSFile("./src/CSS/tokens.css", disPath)
utils.BundleAndMinifySingleCSSFile("./src/CSS/utils.css", disPath)
utils.LoopOverIcons("./src/CSS/icons.css", "./src/JS/icons.js")
}
