package common

import (
	"testing"
	"fmt"
)

func TestMD5(t *testing.T) {
	md5 := MD5("1231212312321")
	md5_16 := MD5_16("1231212312321")
	fmt.Println(md5)
	fmt.Println(md5_16)
	fmt.Println(Md5File("./test.txt"))
}

func TestSHA1(t *testing.T) {
	sha1 := SHA1("1231212312321")
	fmt.Println(sha1)
	fmt.Println(HashFile("./test.txt"))
}

func TestToInt(t *testing.T) {
	var strint, strfloat, strBool = "-56", "0.001", "true"
	fmt.Println(ToInt(strint))
	fmt.Println(ToInt8(strint))
	fmt.Println(ToInt32(strint))
	fmt.Println(ToInt64(strint))
	fmt.Println(ToUint8(strint))
	fmt.Println(ToUint8Hex(strint))
	fmt.Println(ToUint32(strint))
	fmt.Println(ToUint64(strint))

	fmt.Println(ToFloat32(strfloat))
	fmt.Println(ToFloat64(strfloat))

	fmt.Println(ToBool(strBool))
}

func TestRandonInt(t *testing.T) {

	for i := 0; i <= 100; i++ {
		fmt.Printf("TestRandonInt:%d \n", RandonInt(i+100, 1000))
	}

	for i := 0; i <= 100; i++ {
		fmt.Printf("TestRandonFloat:%f \n",RandonFloat(100.0, 1000.01))
	}
}
