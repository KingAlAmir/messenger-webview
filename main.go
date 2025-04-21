package main

import webview "github.com/webview/webview_go"

func main() {
    w := webview.New(false)
    defer w.Destroy()
    w.SetTitle("Messenger Lite")
    w.SetSize(1280, 800, webview.HintNone)
    w.Navigate("https://www.messenger.com/t/")
    w.Run()
}
