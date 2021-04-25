/**
 *@Description
 *@ClassName main
 *@Date 2021/4/22 下午3:45
 *@Author ckhero
 */

package hystrix_go

import (
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
)

func main() {
	hystrix.Go("my_command", func() error {
		// talk to other services
		fmt.Println(222222)
		return nil
	}, nil)

	hystrix.DoC()
}
