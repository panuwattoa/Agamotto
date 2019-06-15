package monitoring

import (
	"Agamoto/linenotify"
	"log"
	"net/http"
	"strconv"
	"time"
)

//
// type CCUType struct {
// 	AllUser      	 string
// 	IOSCCU           string
// 	AndriodCCU       string
// }
var allCCU  = 500
var sendTimeStamp  = time.Now()
var countRecvUpdate int64 = 0
var lastCount int64 = 1
var lastCCU  = 0
var timeDiff float64 = 600

func RecvUpdateCCU(w http.ResponseWriter, r *http.Request)  {

	r.ParseForm()
	var err error
	allCCU, err = strconv.Atoi(r.FormValue("allccu"))
	if err != nil {
		log.Print("Can't read data.. RecvUpdateCCU ", err)
	}
    log.Print("allCCU ", allCCU)
	if countRecvUpdate > 500 {
		countRecvUpdate  = 0
	}else{
		countRecvUpdate ++
	}
}

func CheckServerAlive(){
	for {
		var locationTime, _ = time.LoadLocation("Asia/Bangkok")
		timeNow := time.Now().In(locationTime)
		dayOfWeek := timeNow.Weekday()
		if int(dayOfWeek) != 3 && (timeNow.Hour() <= 9 || timeNow.Hour() >= 13) {  // Check Maintenance day
			if lastCount == countRecvUpdate {
				if timeDiff >= 600  {  // delay time send to Line
					sendTimeStamp  = time.Now()
					linenotify.SendLineNotifyMessage("ผิดปกติ >> ไม่ได้รับ CCU จาก Server! : " + sendTimeStamp.Format("2006-01-02 15:04:05") ,"2","145")
				}

			}

			if  allCCU < lastCCU/2 {
				if timeDiff >= 600 { // delay time send to Line
					sendTimeStamp  = time.Now()
					linenotify.SendLineNotifyMessage("ผพบ CCU ลดลงผิดปกติ  : CCU เหลือ " + strconv.Itoa(allCCU) + sendTimeStamp.Format("2006-01-02 15:04:05") ,"2","149")
				}
			}
			timeDiff = timeNow.Sub(sendTimeStamp).Seconds()
			lastCCU = allCCU
			lastCount = countRecvUpdate
			delay := 15 * time.Second
			time.Sleep(delay)
		}

	}
}




