## 테스트 코드 작성하기
- 참고
    - https://golang.org/doc/code#Testing
- Go는 가벼운 테스트 프레임워크를 갖고있다.
    * `go test` 명령어로 실행 가능하다.
- 테스트 코드를 작성하는 방법
    * 파일 명: suffix로 `_test.go`
    * 함수 명: prefix로 `TestXxx`
    * 파라미터로 `t *testing.T`를 받는다.
    * `t.Error` 또는 `t.Fail`을 호출하면 테스트가 실패하게 된다.
  ```go
  package main

  import "testing"
  
  func TestCalcSuccess(t *testing.T) {
    expected := 3
    actual := Calculate(1, 2)
  
    if expected != actual {
        t.Error("not matched!")
    }
  }
  
  func TestCalcFail(t *testing.T) {
    expected := 2
    actual := Calculate(1, 1)
    
    if expected != actual {
        t.Error("not matched!")
    }
  }
  ```
    - console result
      ```
      --- FAIL: TestCalcFail (0.00s)
          calculator_test.go:20: not matched!
      FAIL
      exit status 1
      FAIL    example/test    0.110s
      ```
  