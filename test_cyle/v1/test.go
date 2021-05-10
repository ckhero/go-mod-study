/**
 *@Description
 *@ClassName test
 *@Date 2021/5/4 下午4:06
 *@Author ckhero
 */

package v1

import (
	"fmt"
	"go-mod-study/test_cyle/domain"
)

type V1S struct {
	v22 domain.V2
}

func(v *V1S) Test() {
	fmt.Println("v1")
}

func(v *V1S) Test2() {
	v.v22.Test()
}