/**
 *@Description
 *@ClassName main
 *@Date 2021/5/4 下午4:06
 *@Author ckhero
 */

package main

import (
	v12 "go-mod-study/test_cyle/v1"
	v22 "go-mod-study/test_cyle/v2"
)

func main()  {
	v2 := v22.V2S{V: &v12.V1S{}}
	v2.Test()
}
