package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strconv"
	"time"
)

type Summary struct {
	Pagesize      int `json:"pagesize"`
	Pageno        int `json:"pageno"`
	Pagecount     int `json:"pagecount"`
	Totalmilestat int `json:"totalmilestat"`
	Totaloutspeed int `json:"totaloutspeed"`
	Totalstop     int `json:"totalstop"`
}
type Keys struct {
	Id       int `json:"id"`
	Name     int `json:"name"`
	Milestat int `json:"milestat"`
	Outspeed int `json:"outspeed"`
	Stop     int `json:"stop"`
}

type GPSResult struct {
	Errcode int        `json:"errcode"`
	Success bool       `json:"success"`
	Msg     string     `json:"msg"`
	Summary Summary    `json:"summary"`
	Keys    Keys       `json:"keys"`
	Data    [][]string `json:"data"`
}

var dataStart = false

func main() {
	data := [][]string{}
	//month := getMonth(2018)  //2018至今

	month := getMonthByYear(2023)

	for _, timestamp := range month {

		start := timestamp["start"]
		end := timestamp["end"]

		gpsResultData := TestGetGPS(start, end)

		if gpsResultData == nil {
			continue
		}

		if len(data) == 0 {
			data = gpsResultData
		} else {
			//累加数据
			for outIndex, datum := range data {
				for innerIndex, resultDatum := range gpsResultData {
					if outIndex == innerIndex && datum[0] == resultDatum[0] {
						fmt.Printf("%s  %s\n", datum[2], resultDatum[2])
						fmt.Printf("%s  %s\n", datum[3], resultDatum[3])
						fmt.Printf("%s  %s\n", datum[4], resultDatum[4])
						datum[2] = addFloat(datum[2], resultDatum[2])
						datum[3] = addInt(datum[3], resultDatum[3])
						datum[4] = addInt(datum[4], resultDatum[4])
						break
					}
				}
			}
		}
		time.Sleep(time.Second)
	}

	writeCSV(data)
}

// [startYear,endYear]
func getMonthByYear(year int) []map[string]int64 {
	shai, _ := time.LoadLocation("Asia/Shanghai")

	timestamps := make([]map[string]int64, 0)

	for month := 0; month < 12; month++ {
		// 获取当前月份的月初时间戳
		firstDay := time.Date(year, time.Month(month+1), 1, 0, 0, 0, 0, shai)
		firstDayTimestamp := firstDay.Unix()

		// 获取下个月的第一天
		lastDay := firstDay.AddDate(0, 1, 0)
		lastDayTimestamp := lastDay.Unix()

		timestamp := map[string]int64{
			"start": firstDayTimestamp,
			"end":   lastDayTimestamp,
		}
		timestamps = append(timestamps, timestamp)
	}
	// 打印结果数组
	for _, timestamp := range timestamps {
		fmt.Printf("Month: %v-%v\n", timestamp["start"], timestamp["end"])
	}
	return timestamps

}

// "10.2" + "20.3" = "30.5"
func addFloat(a, b string) string {
	af, _ := strconv.ParseFloat(a, 64)
	bf, _ := strconv.ParseFloat(b, 64)
	cf := af + bf
	return strconv.FormatFloat(cf, 'f', -1, 64)
}

// "5"+"15"="20"
func addInt(a, b string) string {
	aInt, _ := strconv.ParseInt(a, 10, 64)
	bInt, _ := strconv.ParseInt(b, 10, 64)
	cInt := aInt + bInt
	return strconv.Itoa(int(cInt))
}

// [startYear,now]
func getMonth(startYear int) []map[string]int64 {

	shai, _ := time.LoadLocation("Asia/Shanghai")

	endYear := time.Now().Year()

	timestamps := make([]map[string]int64, 0)

	for year := startYear; year <= endYear; year++ {
		for month := 0; month < 12; month++ {
			// 获取当前月份的月初时间戳
			firstDay := time.Date(year, time.Month(month+1), 1, 0, 0, 0, 0, shai)
			firstDayTimestamp := firstDay.Unix()

			// 获取下个月的第一天
			lastDay := firstDay.AddDate(0, 1, 0)
			lastDayTimestamp := lastDay.Unix()

			timestamp := map[string]int64{
				"start": firstDayTimestamp,
				"end":   lastDayTimestamp,
			}
			timestamps = append(timestamps, timestamp)
		}
	}

	// 打印结果数组
	for _, timestamp := range timestamps {
		fmt.Printf("Month: %v-%v\n", timestamp["start"], timestamp["end"])
	}
	return timestamps
}

// TestGetGPS 注:2021-07-01之前是不返回数据的
func TestGetGPS(start, end int64) [][]string {

	// 创建一个新的cookie jar
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}

	// 假设我们已经有了一个或多个cookie（这里仅作演示）
	cookie := &http.Cookie{
		Name:  "sign",
		Value: "20007387847710170703352232c4bdba82e55d3d6f0887cf461c966cfd0000010018010",
	}

	// 将cookie添加到jar中
	jar.SetCookies(&url.URL{}, []*http.Cookie{cookie})

	// 创建一个client，并设置其cookiejar
	client := &http.Client{
		Jar: jar,
	}

	// 创建GET请求并设置User-Agent
	base := "https://www.gpsoo.net/gpsoo-api/api/report/v1/runstatus"
	params := url.Values{}
	params.Add("eid", "4465999")
	params.Add("timezone", "8")
	params.Add("maptype", "google")
	params.Add("pageno", "0")
	params.Add("pagesize", "100")
	params.Add("method", "runOverViewNew")
	//params.Add("beginTime", "1685548800") //2023-06-01 00:00:00
	params.Add("beginTime", strconv.Itoa(int(start)))
	//params.Add("endTime", "1688140800")   //2023-07-01 00:00:00
	params.Add("endTime", strconv.Itoa(int(end)))

	reqURL := base + "?" + params.Encode()
	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36 Edg/121.0.0.0")
	req.Header.Set("Cookie", "sign=20007387847710170703352232c4bdba82e55d3d6f0887cf461c966cfd0000010018010")

	// 使用带有cookie和User-Agent的客户端发送请求
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("Response Status:", resp.Status)
	// 如果需要处理响应体内容，例如读取HTML
	bodyBytes, _ := io.ReadAll(resp.Body)

	var results GPSResult

	json.Unmarshal(bodyBytes, &results)
	//fmt.Println("Response Body:", string(bodyBytes))
	//fmt.Printf("%#v\n", results)
	//fmt.Println("----------------------------------------------------")

	if results.Success {

		if !dataStart {

			for _, datum := range results.Data {
				d2, _ := strconv.ParseFloat(datum[2], 64)
				if d2 > 0 {
					fmt.Printf("[%s]有公里数的开始时间: start=%d end=%d\n", datum[0], start, end)
					dataStart = true
				}
			}
		}

		return results.Data
	} else {
		fmt.Printf("请求失败:入参[start=%v,end=%v] 响应:[%#v\n]", start, end, results)
		return nil
	}

}

func writeCSV(data [][]string) {

	format := time.Now().Format("20060102150405")
	file, err := os.Create(fmt.Sprintf("test-%s.csv", format))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//写入UTF-8 BOM,此处如果不写入就会导致写入的汉字乱码
	file.WriteString("\xEF\xBB\xBF")
	write := csv.NewWriter(file)

	//w.Write(data) //保存slice一维数据

	write.WriteAll(data)
	write.Flush()
}
