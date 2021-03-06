## 참고
- [Tucker의 GO 언어 프로그래밍](yes24.com/Product/Goods/99108736?OzSrank=1)
- [예제로 배우는 Go 프로그래밍](http://golang.site/)
- Slice
  - [[Go] Slice 집중 탐구](https://velog.io/@kimmachinegun/Go-Slice-%EC%A7%91%EC%A4%91-%ED%83%90%EA%B5%AC-t2jn1kd1gc)

## 구조체
- Go에서 대입은 기본적으로 "복사"다
- 구조체 변수를 다른 변수에 대입해도 값이 복사되고, 함수의 파라미터로 전달하는 값도 복사된다. (call by value)
- 구조체는 필드의 메모리 크기를 합한 크기를 갖는다. 그러나 8의 배수가 아닌 경우 **메모리 정렬(Memory Alignment)**이 발생해 8의 배수로 맞춰진다.
  - 64비트 컴퓨터는 레지스터 크기가 8바이트이고, 데이터 크기도 8바이트면 효율적으로 연산이 가능해진다.
- 빈 구조체를 생성하려면
  ```go
  type Human struct {
    // ...  
  }
  s := Human{}
  ```
- 구조체를 다른 변수에 대입하면 새로운 주소가 할당되고 값이 복사된다.
  
## 포인터
- 포인터 변수의 기본값은 nil이다
- 포인터 비교는 assert.Equal이 아닌 assert.Same으로 비교해야 한다.
  - 동일성과 동등성의 차이.
- 빈 구조체 포인터를 생성하려면
  ```go
  type Human struct {
    // ...  
  }
  var s *Human = &Human{}
  ```
- 구조체 포인터를 다른 변수에 대입하면 같은 주소를 가리킨다.

## 문자열
- 문자열 대소 비교는 문자열 길이와 관계 없이 앞글자부터 비교. UTF-8 값으로 비교한다.
- string 내부 구조.
  ```go
  // value.go
  type StringHeader struct {
    Data uintptr // 문자열을 가리키는 포인터
    Len int      // 문자열 길이
  }
  ```
- string은 불변
  - slice로 형변환시 새로운 메모리 공간 연결
  - string concat시 새로 메모리 공간 연결
  
## Slice
- contains 함수가 없다.
- 구조
  ```go
  type SliceHeader struct {
    Data uintptr
    Len int
    Cap int
  }
  ```
- 슬라이스는 배열의 주소(Data)와 길이(len), 용량(cap)을 갖는다.
  - 용량과 길이를 확인하려면 cap(slice), len(slice)을 사용한다.
- 생성시에 용량을 생략하면 길이와 용량이 같은 슬라이스가 생성된다.
  