package main

type Alarm struct {
	ObjectType   string  `json:"docType"`
	Id       string  `json:"Id"`	//站房
	SiteId       string  `json:"SiteId"`	//站房
	NoticeType       string  `json:"NoticeType"`	//通知类型
	AlarmDetail       string  `json:"AlarmDetail"`	//报警详情

	AlarmLevel       string  `json:"AlarmLevel"`	//报警级别

	SiteType       string  `json:"SiteType"`		//站房类型

	TriggerValue       string  `json:"TriggerValue"`	//触发值
	AlarmParam       string  `json:"AlarmParam"`		//报警参数
	AlarmTime       string  `json:"AlarmTime"`			//报警时间
	Auditor       string  `json:"Auditor"`				//作者
	Historys       []HistoryItem  `json:"alarmParam"`

}

type HistoryItem struct {
	TxId     string
	Alarm	 Alarm
}