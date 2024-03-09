package util

import "strings"

func ReadTGCmdFromMsgText(text string) string {
	if text[0] != '/' {
		return ""
	}
	strs := strings.Split(text, "@")
	strs = strings.Split(strs[0], " ")
	strs = strings.Split(strs[0], "?")
	return strs[0]
}

// ParseTGMsg parse tg cmd like "/cmd@bot something"
func ParseTGMsg(text string) (cmd, cmdParams string) {
	if text[0] != '/' {
		return "", ""
	}
	strs := strings.Split(text, "@")
	cmd = strs[0]
	if len(strs) == 1 {
		return cmd, ""
	}
	strs = strings.Split(strs[1], " ")
	if len(strs) > 1 {
		cmdParams = strs[1]
	}
	return
}
