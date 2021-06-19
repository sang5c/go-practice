# go-practice

## 알아두면 좋은 내용
- GOPATH
  * 과거의 Go는 워크스페이스를 지정하고 해당 워크스페이스 하위에 코드가 존재해야 했다. 이 워크스페이스를 담는 환경변수가 GOPATH이다.
  - Windows의 default path는 C:\Users\<user name>\go 
  * Go modules를 사용하면 원하는 곳에 Go 소스를 저장해도 상관 없다. 
- GOROOT
  - Go의 컴파일러, 기본 모듈들이 담기는 공간.
  - Windows의 default path는 C:\Program Files\go 
- go modules
  - go mod init "모듈명"을 통해 초기화 할 수 있다. 이걸 하면 GOPATH를 사용하지 않아도 된다.   