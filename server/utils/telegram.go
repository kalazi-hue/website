package utils

import "net/http"

func SendMsgToGroup(message string) {
	groupInfo := "https://api.telegram.org/bot1421246646:AAHAxKDMPsgRWp1Msf7xmRYVJUi4ob3OXSg/sendMessage?chat_id=-475483587&text="
	rsp, _ := http.Get(groupInfo + message)
	rsp.Body.Close()
}
