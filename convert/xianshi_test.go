/**
 *@Description
 *@ClassName xianshi_test
 *@Date 2021/5/12 下午7:00
 *@Author ckhero
 */

package convert

import (
	"fmt"
	"testing"
)

func TestXianShi(t *testing.T) {
	a := TestA{Name:"222"}
	b := TestB(a)
	fmt.Println(b)
}
