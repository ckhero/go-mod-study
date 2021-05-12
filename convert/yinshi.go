/**
 *@Description
 *@ClassName yinshi
 *@Date 2021/5/12 下午7:01
 *@Author ckhero
 */

package convert

type Reader interface {
	Read(p []byte) (n int, err error)
}
type ReadCloser interface {
	Reader
	Close() error
}
