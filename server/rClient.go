package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type User struct {
	UserID   string `json:"USER_ID"`
	Password string `json:"USER_PW"`
}

// 테스트용 User변수 정의
var u1 = User{"admin", "admin"}
var u2 = User{"tsoukalos", "pass"}
var u3 = User{"", "pass"}

// 엔드포인트 상수 정의
const addEndPoint = "/add"
const getEndPoint = "/get"
const deleteEndPoint = "/delete"
const timeEndPoint = "/time"

// 엔드포인트에 접근하기 위한 요청 생성
func deleteEndpoint(server string, user User) int {
	userMarshall, _ := json.Marshal(user)
	u := bytes.NewReader(userMarshall)

	req, err := http.NewRequest("DELETE", server+deleteEndPoint, u)

	if err != nil {
		fmt.Println("Error in req: ", err)
		return http.StatusInternalServerError
	}
	req.Header.Set("Content-Type", "application/json") //서버와 통신할 때 JSON 데이터 사용

	c := &http.Client{
		Timeout: 15 * time.Second, // 15초뒤에 타임아웃
	}

	resp, err := c.Do(req)  // Do메서드로 요청보냄(얘 역할은 클라이언트에 구성된 정책에 따라 HTTP요청을 보내고 응답을 반환함)
	defer resp.Body.Close() // defer - 특정한 함수를 실행할때, 해당부분을 함수가 종료되기 직전에 실행 (가장마지막에)

	if err != nil {
		fmt.Println("Error:", err)
	}
	if resp == nil {
		return http.StatusNotFound
	}
	data, err := io.ReadAll(resp.Body)
	fmt.Print("/delete returned: ", string(data))

	if err != nil {
		fmt.Println("Error:", err)
	}
	return resp.StatusCode // /delete요청을 수행한 결과를 알려쥼
}

func getEndpoint(server string, user User) int {
	userMarshall, _ := json.Marshal(user)
	u := bytes.NewReader(userMarshall)

	req, err := http.NewRequest("GET", server+getEndPoint, u)

	if err != nil {
		fmt.Println("Error in req: ", err)
		return http.StatusInternalServerError
	}
	req.Header.Set("Content-Type", "application/json")
	c := &http.Client{
		Timeout: 15 * time.Second,
	}
	resp, err := c.Do(req) // 서버응답은 resp변수에 저장, 에러값은 err변수에 저장
	defer resp.Body.Close()

	if err != nil {
		fmt.Println("Error:", err)
	}
	if resp == nil {
		return http.StatusNotFound
	}
	data, err := io.ReadAll(resp.Body)
	fmt.Print("/get returned: ", string(data))
	if err != nil {
		fmt.Println("Error:", err)
	}
	return resp.StatusCode // 요청이 성공했는지 알려줌
}

// /add 엔드포인트에 접근.
func addEndpoint(server string, user User) int {
	userMarshall, _ := json.Marshal(user)
	u := bytes.NewReader(userMarshall)
	req, err := http.NewRequest("POST", server+addEndPoint, u)

	if err != nil {
		fmt.Println("Error in req: ", err)
		return http.StatusInternalServerError
	}
	req.Header.Set("Content-Type", "application/json")

	c := &http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := c.Do(req)
	defer resp.Body.Close()

	if resp == nil || (resp.StatusCode == http.StatusNotFound) {
		return resp.StatusCode
	}
	return resp.StatusCode
}

func timeEndpoint(server string) (int, string) {
	req, err := http.NewRequest("POST", server+timeEndPoint, nil)

	if err != nil {
		fmt.Println("Error in req: ", err)
		return http.StatusInternalServerError, ""
	}

	c := &http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := c.Do(req)
	defer resp.Body.Close()

	if resp == nil || (resp.StatusCode == http.StatusNotFound) {
		return resp.StatusCode, ""
	}

	data, _ := io.ReadAll(resp.Body)
	return resp.StatusCode, string(data)
}

// 서버의 기본 엔드포인트 테스트용 (데이터 입력 필요 x)
func slashEndPoint(server, URL string) (int, string) {
	req, err := http.NewRequest("POST", server+URL, nil)

	if err != nil {
		fmt.Println("Error in req: ", err)
		return http.StatusInternalServerError, ""
	}

	c := &http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := c.Do(req)
	defer resp.Body.Close()

	if resp == nil {
		return resp.StatusCode, ""
	}

	data, _ := io.ReadAll(resp.Body)
	return resp.StatusCode, string(data)
}

func main() {
	if len(os.Args) != 2 { // 서버 ㅈ주소와 포트 번호
		fmt.Println("Wrong number of arguments!")
		fmt.Println("Need: Server")
		return
	}
	server := os.Args[1]
	fmt.Println("/add")
	HTTPcode := addEndpoint(server, u1)
	if HTTPcode != http.StatusOK {
		fmt.Println("u1 Return code:", HTTPcode)
	} else {
		fmt.Println("u1 Data added:", u1, HTTPcode)
	}

	HTTPcode = addEndpoint(server, u2)
	if HTTPcode != http.StatusOK {
		fmt.Println("u2 Return code:", HTTPcode)
	} else {
		fmt.Println("u2 Data added:", u2, HTTPcode)
	}

	HTTPcode = addEndpoint(server, u3)
	if HTTPcode != http.StatusOK {
		fmt.Println("u3 Return code:", HTTPcode)
	} else {
		fmt.Println("u3 Data added:", u3, HTTPcode)
	}

	fmt.Println("/get")
	HTTPcode = getEndpoint(server, u1)
	fmt.Println("/get u1 return code:", HTTPcode)
	HTTPcode = getEndpoint(server, u2)
	fmt.Println("/get u2 return code:", HTTPcode)
	HTTPcode = getEndpoint(server, u3)
	fmt.Println("/get u3 return code:", HTTPcode)

	fmt.Println("/delete")
	HTTPcode = deleteEndpoint(server, u1)
	fmt.Println("/delete u1 return code:", HTTPcode)
	HTTPcode = deleteEndpoint(server, u1)
	fmt.Println("/delete u1 return code:", HTTPcode)
	HTTPcode = deleteEndpoint(server, u2)
	fmt.Println("/delete u2 return code:", HTTPcode)
	HTTPcode = deleteEndpoint(server, u3)
	fmt.Println("/delete u3 return code:", HTTPcode)

	fmt.Println("/time")
	HTTPcode, myTime := timeEndpoint(server)
	fmt.Print("/time returned: ", HTTPcode, " ", myTime)
	time.Sleep(time.Second)
	HTTPcode, myTime = timeEndpoint(server)
	fmt.Print("/time returned: ", HTTPcode, " ", myTime)

	// 존재하지 않는 엔드포인트에 접속 시도해 기본 핸들러 함수가 잘 실행되는지 확인
	fmt.Println("/")
	URL := "/"
	HTTPcode, response := slashEndPoint(server, URL)
	fmt.Print("/ returned: ", HTTPcode, " with response: ", response)

	fmt.Println("/what")
	URL = "/what"
	HTTPcode, response = slashEndPoint(server, URL)
	fmt.Print(URL, " returned: ", HTTPcode, " with response: ", response)
}
