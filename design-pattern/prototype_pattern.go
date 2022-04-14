package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// prototype 패턴은 매번 객체를 새로 생성하기보다는 original 객체를 복사해서 필요한 부분만 수정해서 사용하는 패턴이다.
// 1. 복제할 객체 생성 프로세스가 복잡할 때 프로토타입 패턴을 사용
// 2. 새 인스턴스를 처음부터 만드는 대신 개체의 복사본이 만들어 지므로 데이터베이스 작업과 같은 새로운 개체를 생성하는 동안 관련된 비용이 드는 작업 방지

// 간단한 예시
// 이러한 구조체를 만들었을 때

type Address struct {
	City          string
	AddressNumber uint
}

type Person struct {
	Name    string
	Address *Address
}

func main() {
	alex := Person{"alex", &Address{"Seoul", 1234}}

	sarah := alex
	sarah.Address.City = "Busan"
	fmt.Println(alex.Address)  // Busan
	fmt.Println(sarah.Address) // Busan
	// 위처럼 코드 작성시에 Address가 포인터 타입이기 때문에 alex의 주소도 Busan으로 바뀌어 버립니다.

	hyeon := Person{"alex", &Address{"Seoul", 1234}}

	geun := hyeon.DeepCopy()
	geun.Address.City = "Busan"
	fmt.Println(hyeon.Address) // Seoul
	fmt.Println(geun.Address)  // Busan
	// 위 코드를 보면 deepcopy가 되어서 geun의 주소는 여전히 Seoul로 되어있습니다.
}

func (p *Person) DeepCopy() *Person { // 여러가지 방법이 있겠지만 gob 패키지를 이용하는 방법으로 예시
	buf := bytes.Buffer{}
	e := gob.NewEncoder(&buf)
	_ = e.Encode(p)

	decoder := gob.NewDecoder(&buf)
	result := Person{}
	_ = decoder.Decode(&result)

	return &result // encoding과 decoding 과정을 거치게되면 포인터 형태로 이루어진	구조체도 deepcopy가 가능합니다.
}
