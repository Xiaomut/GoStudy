package mylogger

import (
	"path"
	"os"
	"fmt"
	"time"
)

type FileLogger struct{
	Level LogLevel
	filepath string // 日志目录
	filename string // 保存名字
	fobj *os.File
	err_file *os.File
	max_file_size int64
	log_chan chan *logMsg
}

type logMsg struct{
	level LogLevel
	msg string
	func_name string
	file_name string
	timestamp string
	line int
}

// NewFileLogger构造函数
func NewFileLogger(level, fp, fn string, maxsize int64) *FileLogger  {
	loglevel, err := parse_log_level(level)
	if err != nil{
		panic(err)
	}
	fl := &FileLogger{
		Level: loglevel,
		filepath: fp,
		filename: fn,
		max_file_size: maxsize,
		log_chan: make(chan *logMsg, 50000),
	}
	err = fl.initFile()
	if err != nil{
		panic(err)
	}
	return fl
}

func (f *FileLogger) initFile() (error) {
	log_path := path.Join(f.filepath, f.filename)
	err_path := path.Join(f.filepath, "err_" + f.filename)
	log_file, err := os.OpenFile(log_path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil{
		fmt.Println("open log failed: ", err)
		return err
	}
	err_file, err := os.OpenFile(err_path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil{
		fmt.Println("open err log failed: ", err)
		return err
	}
	// 日志都已打开
	f.fobj = log_file
	f.err_file = err_file
	// 开启一个后台的goroutine写日志
	// for i := 0; i < 5; i++{
	// 	go f.log_back()
	// }
	go f.log_back()
	return nil
}

// 切割文件
func (f *FileLogger) fileSplit(file *os.File) (*os.File, error) {
	/*	切割
		1. 关闭当前日志
		2. 备份一下 rename x.log -> xx.log.bak2022*****
		3. 打开一个新的日志
		4. 将打开的新日志对象赋值给f.fobj
	*/
	now := time.Now().Format("20160102150405000")
	file_info, err := file.Stat()
	if err != nil{
		fmt.Println("get file info failed: ", err)
		return nil, err
	}
	log_name := path.Join(f.filepath, file_info.Name())
	log_new := fmt.Sprintf("%s.bak%s", log_name, now)
	
	file.Close()
	os.Rename(log_name, log_new)

	fobj_new, err := os.OpenFile(log_name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil{
		fmt.Println("open new log failed: ", err)
		return nil, err
	}
	return fobj_new, nil
}

// 判断是否需要记录
func (f *FileLogger) enable(loglevel LogLevel) bool  {
	return loglevel >= f.Level
}

// 根据给定的size判断是否切割
func (f *FileLogger) checkSize(file *os.File) bool {
	file_info, err := file.Stat()
	if err != nil{
		fmt.Println("get file info failed: ", err)
		return false
	}
	return file_info.Size() >= f.max_file_size
}

func (f *FileLogger) log_back()  {
	for {
		select{
		case log_tmp := <- f.log_chan:
			log_info := fmt.Sprintf("[%s] [%s] [%s:%s:%d] %s\n", log_tmp.timestamp, status, log_tmp.func_name, log_tmp.file_name, log_tmp.line, log_tmp.msg)
			fmt.Fprintf(f.fobj, log_info)
			var status string
			switch log_tmp.level{
			case DEBUG:
				status = "DEBUG"
			case TRACE:
				status = "TRACE"
			case INFO:
				status = "INFO"
			case WARNING:
				status = "WARNING"
			case ERROR:
				status = "ERROR"
			case FATAL:
				status = "FATAL"
			default:
				status = "UNKNOWN"
			}
			if log_tmp.level >= ERROR{
				if f.checkSize(f.err_file){
					new_log, err := f.fileSplit(f.err_file)
					if err != nil{
						return 
					}
					f.err_file = new_log
				}
				// 大于error级别的再单独记录一遍，个人认为可以不需要	
				fmt.Fprintf(f.err_file, log_info)
			}
		default:
			time.Sleep(time.Millisecond * 500)
		}
	}
}

// 记录日志
func (f *FileLogger) get_log(lv LogLevel, format string, a ...interface{}){
	if f.enable(lv){
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcname, filename, linenum := getInfo(3)
		// 先把日志发送到通道中
		log_tmp := &logMsg{
			level: lv,
			msg: msg,
			func_name: funcname,
			file_name: filename,
			timestamp: now.Format("2006-01-02 15:04:05"),
			line: linenum,
		}
		select {
		case f.log_chan <- log_tmp:
		default:
		}
		if f.checkSize(f.fobj){
			new_log, err := f.fileSplit(f.fobj)
			if err != nil{
				return 
			}
			f.fobj = new_log
		}
	}	
}


func (f *FileLogger) Debug(format string, a ...interface{}){
	f.get_log(DEBUG, format, a...)
}

func (f *FileLogger) Info(format string, a ...interface{}){
	f.get_log(INFO, format, a...)
}

func (f *FileLogger) Warning(format string, a ...interface{}){
	f.get_log(WARNING, format, a...)
}

func (f *FileLogger) Error(format string, a ...interface{})  {
	f.get_log(ERROR, format, a...)
}

func (f *FileLogger) Fatal(format string, a ...interface{})  {
	f.get_log(FATAL, format, a...)
}

func (f *FileLogger) Close()  {
	f.fobj.Close()
	f.err_file.Close()
}