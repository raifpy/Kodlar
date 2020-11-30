package main

import (
	"os"

	"github.com/webview/webview"
)

func main() {
	web := webview.New(true) // debug
	web.SetTitle("WebView")
	web.Navigate(os.Args[1])
	//fmt.Println()

	web.Run()

}

const loading = `
/* Function to add style element */ 
function addStyle(styles) { 
	  
	/* Create style document */ 
	var css = document.createElement('style'); 
	css.type = 'text/css'; 
  
	if (css.styleSheet)  
		css.styleSheet.cssText = styles; 
	else  
		css.appendChild(document.createTextNode(styles)); 
	  
	/* Append style to the tag name */ 
	document.getElementsByTagName("head")[0].appendChild(css); 
} 
  
/* Set the style */ 
var styles = 'h1 { color: green }'; 
styles += ' body { text-align: center }'; 
  
/* Function call */ 
window.onload = function() { addStyle(styles) }; 
`
