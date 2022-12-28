package variable

import (
	"fmt"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
)

func OnReady() {
	systray.SetTemplateIcon(icon.Data, icon.Data)

	go func() {
		systray.AddMenuItem("Ignored", "Ignored")

		quit := systray.AddMenuItem("退出", "")
		systray.AddSeparator()

		for {
			select {
			case <-quit.ClickedCh:
				systray.Quit()
				fmt.Println("Quit2 now...")
				return
			}
		}
	}()
}
