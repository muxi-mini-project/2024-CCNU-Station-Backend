package CCNU

import (
	"encoding/json"
	"fmt"
	"github.com/anaskhan96/soup"
	"log"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"time"
)

type T struct {
	User struct {
		DeptId       string `json:"deptId"`
		DeptName     string `json:"deptName"`
		Email        string `json:"email"`
		Id           string `json:"id"`
		Mobile       string `json:"mobile"`
		Name         string `json:"name"`
		Status       int    `json:"status"`
		UserFace     string `json:"userFace"`
		Username     string `json:"username"`
		Usernumber   string `json:"usernumber"`
		Usertype     string `json:"usertype"`
		UsertypeName string `json:"usertypeName"`
		Xb           string `json:"xb"`
	} `json:"user"`
}

func CCNULogin(studentID string, password string) *http.Client {
	htmlBody, _ := soup.Get("https://account.ccnu.edu.cn/cas/login?service=http%3A%2F%2Fone.ccnu.edu.cn%2Fcas%2Flogin_portal")
	doc := soup.HTMLParse(htmlBody)
	links1 := doc.Find("body", "id", "cas").FindAll("script")
	js := links1[2].Attrs()["src"][26:]
	links2 := doc.Find("div", "class", "logo").FindAll("input")

	st := links2[2].Attrs()["value"]
	jar, _ := cookiejar.New(&cookiejar.Options{})

	client := &http.Client{
		Jar:     jar,
		Timeout: 5 * time.Second,
	}

	url := fmt.Sprintf("https://account.ccnu.edu.cn/cas/login;jsessionid=%v?service=http", js) + "%3A%2F%2Fone.ccnu.edu.cn%2Fcas%2Flogin_portal"
	//text := fmt.Sprintf("username=%v&password=%v&lt=%v&execution=e1s1&_eventId=submit&submit=登录", studentID, password, st) + "%E7%99%BB%E5%BD%95"
	text := fmt.Sprintf("username=%s&password=%s&lt=%s&execution=e1s1&_eventId=submit&submit=", studentID, password, st) + "%E7%99%BB%E5%BD%95"
	body := strings.NewReader(text)
	req, _ := http.NewRequest("POST", url, body)
	req.Header.Set("Cookie", "JSESSIONID="+js)
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Origin", "https://account.ccnu.edu.cn")
	req.Header.Set("Referer", "https://account.ccnu.edu.cn/cas/login?service=http%3A%2F%2Fone.ccnu.edu.cn%2Fcas%2Flogin_portal")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 Edg/120.0.0.0")
	req.Header.Set("sec-ch-ua", `"Not_A Brand";v="8", "Chromium";v="120", "Microsoft Edge";v="120"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	if resp.Header.Get("Pragma") == "" {
		return client
	}
	return nil
}

func LoginSuccess(username string, password string) bool {
	return CCNULogin(username, password) != nil
}
func GetLoginToken(client *http.Client) string {
	req, err := http.NewRequest("GET", "http://one.ccnu.edu.cn/index", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "one.ccnu.edu.cn")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 Edg/120.0.0.0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Referer", "http://one.ccnu.edu.cn/cas/login_portal")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	token := resp.Header.Get("Set-Cookie")
	return token[13:]
}
func GetUserNameAndCollegeAndGender(client *http.Client) map[string]string {
	fmt.Println("Getting user name and College and Gender")
	info := make(map[string]string)
	URL := "http://one.ccnu.edu.cn/user_portal/userDetailCcnu"
	// 创建一个HTTP客户端
	c := &http.Client{}
	// 创建一个GET请求
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil
	}
	// 添加Bearer Token到Authorization头
	token := GetLoginToken(client)
	fmt.Println(token)
	req.Header.Add("Authorization", "Bearer "+token)
	// 发送请求
	resp, err := c.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil
	}
	defer resp.Body.Close()
	// 解析JSON响应
	var responseData T
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	fmt.Println(responseData)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil
	}
	info["Name"] = responseData.User.Name
	info["College"] = responseData.User.DeptName
	if responseData.User.Xb == "1" {
		info["Gender"] = "男"
	}
	if responseData.User.Xb == "2" {
		info["Gender"] = "女"
	}
	fmt.Println("info:", info)
	return info
}
