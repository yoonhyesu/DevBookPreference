package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type User struct {
	UserID   string `json:"USER_ID"`
	Password string `json:"USER_PW"`
}

var user User // 사용자 정보 (나중에 보안을 어떻게 유지시킬것인지 생각할것)

var PORT = ":8080" // 웹 서버가 사용할 TCP 포트 번호

var DATA = make(map[string]string) // DATA는 User 레코드를 갖고 있는 맵

// 기본 핸들러
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host) // 요청된 URL경로, 요청보낸 Host 기록
	w.WriteHeader(http.StatusNotFound)                  // HTTP 응답 상태 코드 설정 -> 404
	Body := "Thanks for visiting!\n"
	fmt.Fprint(w, "%s", Body)
}

// 현재 날짜와 시간 반환 핸들러(보통 서버의 상태를 테스트할 때 사용하고 실제 프로덕션 환경에선 제거)
func timeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	t := time.Now().Format(time.RFC1123)
	Body := "The current time is: " + t + "\n"
	fmt.Fprintf(w, "%s", Body)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host, r.Method)
	if r.Method != http.MethodPost {
		http.Error(w, "Error:", http.StatusMethodNotAllowed) // 에러메시지 및 HTTP 상태코드와 함께 클라이언트 요청에 대한 응답줌
		// 응답을 보내는 에러 메시지는 일반 텍스트여야함
		fmt.Fprintf(w, "%s\n", "Method not allowed!")
		return
	}

	d, err := io.ReadAll(r.Body) // 클라이언트 요청에서 데이터를 문제 없이 읽었는지 체크
	if err != nil {              // error변수의 값이 nil이면 에러가 없다는 뜻
		http.Error(w, "Error:", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(d, &user) // 클라이언트에서 데이터를 읽은 뒤에는 user 전역 변수에 데이터 넣음
	if err != nil {
		log.Println(err)
		http.Error(w, "Error:", http.StatusBadRequest)
		return
	}

	// 주어진 UserID필드가 빈 값이 아니라면 DATA맵에 새로운 구조체 추가
	// 현재는 데이터의 영속성을 보장하지 않으므로 RESTful 서버 재시작할 때마다 DATA 맵 초기화됨
	if user.UserID != "" {
		DATA[user.UserID] = user.Password
		log.Println(DATA)
		w.WriteHeader(http.StatusOK) // 응답헤더
	} else { // UserID필드가 빈값이면 DATA 맵에 새로운 데이터를 추가할 수 없으므로 실패요청
		http.Error(w, "Error:", http.StatusBadRequest)
		return
	}
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host, r.Method) //요청정보 로그 기록
	if r.Method != http.MethodGet {                               // 요청 메서드가 get인지 확인
		http.Error(w, "Error:", http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "%s\n", "Method not allowed!")
		return
	}
	// 1. 클라이언트 요청에서 데이터를 문제없이 읽었는지 체크
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "ReadAll - Error", http.StatusBadRequest)
	}
	// 2. 클라이언트 데이터를 User 구조체에 넣음
	err = json.Unmarshal(d, &user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unmarshal-Error", http.StatusBadRequest)
		return
	}
	fmt.Println(user)

	// 원하는 레코드를 찾았다면 d 변수의 데이터와 함께 요청이 성공했다는 사실을 클라이언트로 응답
	_, ok := DATA[user.UserID] //UserID를 key값으로 찾음. '_'는 실제 값 무시하고 'ok'는 키의존재여부 저장
	if ok && user.UserID != "" {
		log.Println("Found!")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s\n", d)
	} else {
		log.Println("Not found!")
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, "Map - Resource not found!", http.StatusNotFound)
	}
	return
}

// 리소스 삭제
func deleteHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host, r.Method)
	if r.Method != http.MethodDelete {
		http.Error(w, "Error:", http.StatusMethodNotAllowed)
		fmt.Fprint(w, "%s\n", "Method not allowed!")
		return
	}

	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "ReadAll - Error", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(d, &user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unmarshal - Error", http.StatusBadRequest)
		return
	}
	log.Println(user) // 리소스를 삭제할 때는 추가적으로 관련 정보를 로깅하는 것이 좋음

	// 실제 데이터 삭제하기 전에 주어진 사용자 아이디와 비밀번호값이 DATA맵에 저장된 내용과 같은지 확인
	_, ok := DATA[user.UserID]
	if ok && user.UserID != "" {
		if user.Password == DATA[user.UserID] {
			delete(DATA, user.UserID)
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%s\n", d)
			log.Println(DATA)
		} else {
			log.Println("User", user.UserID, "Not found!")
			w.WriteHeader(http.StatusNotFound)
			http.Error(w, "Delete - Resource not found!", http.StatusNotFound)
		}
		log.Println("After:", DATA)
		return
	}
}

func main() {
	arguments := os.Args
	if len(arguments) != 1 {
		PORT = ":" + arguments[1] // 커맨드라인 인수가 없을 때 기본 값 사용
	}

	mux := http.NewServeMux()
	s := &http.Server{
		Addr:         PORT,
		Handler:      mux,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	// 웹 서버 엔드포인트 정의
	mux.Handle("/time", http.HandlerFunc(timeHandler))
	mux.Handle("/add", http.HandlerFunc(addHandler))
	mux.Handle("/get", http.HandlerFunc(getHandler))
	mux.Handle("/delete", http.HandlerFunc(deleteHandler))
	mux.Handle("/", http.HandlerFunc(defaultHandler))

	fmt.Println("Ready to serve at", PORT)
	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// 기본 값이 있을 때 웹 서버의 TCP 포트 번호를 정의하는
