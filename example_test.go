package uniq_test

import (
	"fmt"

	"github.com/rwxrob/uniq-go"
)

func Example() {

	uid32 := uniq.Base32()
	uuid := uniq.UUID()
	hexid := uniq.Hex(18)
	rgb := uniq.Hex(3)
	b := uniq.Hex(1)
	sec := uniq.Second()
	isosec := uniq.IsoSecond()

	fmt.Println(uid32)  // BCF1KFRJMSAQ1G9HCQ3L25CQOHFIGNQF
	fmt.Println(uuid)   // e98ee42a-d820-4bff-9b0e-67bcff639a17
	fmt.Println(hexid)  // 98af788e67de0032b86bb3a3b04f935e72bb
	fmt.Println(rgb)    // 35ba0f
	fmt.Println(b)      // b9
	fmt.Println(sec)    // 1648009213
	fmt.Println(isosec) // 20060102150405

}
