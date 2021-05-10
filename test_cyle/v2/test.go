/**
 *@Description
 *@ClassName test
 *@Date 2021/5/4 下午4:06
 *@Author ckhero
 */

package v2

import (
	"fmt"
	"go-mod-study/test_cyle/domain"
)

type V2S struct {
	V domain.V1
}

func(v *V2S) Test() {
	v.V.Test()
	fmt.Println("v2")
}
