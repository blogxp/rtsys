package trans

import "fmt"

//TranFileSize 文件大小格式化
func TranFileSize(size int64) string {
	if size< 1<<10 {
		return fmt.Sprintf("%.2f B",float64(size))
	} else if size < 1<<20{
		return fmt.Sprintf("%.2f KB",float64(size)/float64(1<<10))
	} else if size < 1<<30{
		return fmt.Sprintf("%.2f MB",float64(size)/float64(1<<20))
	} else if size < 1<<40{
		return fmt.Sprintf("%.2f GB",float64(size)/float64(1<<30))
	} else {
		return fmt.Sprintf("%.2f TB",float64(size)/float64(1<<40))
	}
}
