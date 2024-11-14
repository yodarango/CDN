package main

import "cdn/src/utils"

func main(){
disPathCSS := "./src/dist/css"
disPathJS := "./src/dist/js"
utils.BundleAndMinifyAllCSS(disPathCSS)
utils.BundleAndMinifySingleCSSFile("./src/CSS/tokens.css", disPathCSS)
utils.BundleAndMinifySingleCSSFile("./src/CSS/utils.css", disPathCSS)
// TODO: Fix this. It's not working
utils.BundleAndMinifyAllJS(disPathJS)
//utils.LoopOverIcons("./src/CSS/icons.css", "./src/JS/icons.js")
}
