package ko

import "fmt"

type PostSender struct {

}

func (p *PostSender) Send(parcel string) {
	fmt.Printf("우체국에서 택배 %v를 보냅니다\n", parcel)
}
