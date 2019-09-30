package main

import (
	"os"
	"github.com/pkg/errors"
	"fmt"
)

// *********  日志对外接口  **********

// 声明日志写入器接口
type LogWriter interface {
	Write(data interface{}) error
}

// 日志器
type Logger struct {
	// 这个日志器用到的日志写入器
	writerList []LogWriter
}

// 注册一个日志写入器
func (l *Logger) RegisterWriter(writer LogWriter){
	l.writerList = append(l.writerList, writer)
}

// 将一个data类型的数据写入日志
func (l *Logger) Log(data interface{})  {

	// 遍历所有注册的写入器
	for _, writer := range l.writerList{

		// 将日志输入到每个写入器中
		writer.Write(data)
	}
}

// 创建日志器的实例
func NewLogger() *Logger{
	return &Logger{}
}


// *********  文件写入器  **********
type fileWriter struct {
	file *os.File
}

// 设置文件写入器写入的文件名
func (f *fileWriter) SetFile(filename string)(err error)  {

	// 如果文件已经打开，关闭前一个文件
	if f.file != nil {
		f.file.Close()
	}

	// 创建一个文件并保存文件句柄
	f.file, err = os.Create(filename)

	// 如果创建过程出现异常，则返回错误
	return err
}

// 实现LogWriter的Write()方法
func (f *fileWriter) Write(data interface{}) error{

	// 日志文件可能没有创建成功
	if f.file == nil {

		return errors.New("file not created")
	}

	// 将数据序列化为字符串
	str := fmt.Sprintf("%v\n",data)

	// 将数据以字节数组写入文件中
	_, err := f.file.Write([]byte(str))

	return err
}

func newFileWriter() *fileWriter {
	return &fileWriter{}
}


// *********  命令行写入器  **********

// 命令行写入器
type consoleWriter struct {
}

// 实现LogWriter的Write()方法
func (f *consoleWriter) Write(data interface{}) error {

	// 将数据序列化为字符串
	str := fmt.Sprintf("%v\n", data)

	// 将数据以字节数组写入命令行中
	_, err := os.Stdout.Write([]byte(str))

	return err
}

// 创建命令行写入器实例
func newConsoleWriter() *consoleWriter {
	return &consoleWriter{}
}

// 创建日志器
func createLogger() *Logger {

	// 创建日志器
	l := NewLogger()

	// 创建命令行写入器
	cw := newConsoleWriter()

	// 注册命令行写入器到日志器中
	l.RegisterWriter(cw)


	// 创建文件写入器
	fw := newFileWriter()

	// 设置文件名
	if err := fw.SetFile("log.log"); err != nil {
		panic(err)
	}

	// 注册文件写入器到日志器中
	l.RegisterWriter(fw)

	return l
}

func main() {

	// 准备日志器
	l := createLogger()

	// 写一个日志
	l.Log("hello")

}


