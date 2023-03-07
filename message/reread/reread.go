package reread

import (
	"regexp"
	"robot/common"
	"robot/request"
)

func Reread(groupID int64, message string) {

	// 将当前消息加入到记录中
	newMessage := regReplace(message)
	common.LastMessages[groupID] = append(common.LastMessages[groupID], newMessage)
	// 如果记录超过了四条，删除最早的一条
	if len(common.LastMessages[groupID]) > 4 {
		common.LastMessages[groupID] = common.LastMessages[groupID][1:]
	}
	// 判断是否有连续四条相同的消息
	if len(common.LastMessages[groupID]) == 4 {
		isEqual, err := allMapValuesEqual(common.LastMessages[groupID])
		if err != nil {
			return
		}
		if isEqual {
			// 发送警告消息
			delete(common.LastMessages, groupID)
			request.M().Send("你复读尼玛呢")
		}
	}
}

func regReplace(message string) string {
	pattern := `\[(CQ:image,file=[^,\]]+),subType=1,url=([^,\]]+)\]`
	replace := `[$1,subType=1,url=XXX]`

	re := regexp.MustCompile(pattern)
	output := re.ReplaceAllString(message, replace)
	return output
}

// 判断 是否全等
func allMapValuesEqual(m []string) (bool, error) {
	var value string
	first := true
	for _, v := range m {
		if first {
			value = v
			first = false
		} else if v != value {
			return false, nil
		}
	}
	return true, nil
}
