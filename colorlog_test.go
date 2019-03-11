package colorlog

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func ExampleDebug() {
	Info("      饮酒·其五     ")
	Info("              -----%s", "陶渊明")
	Debug("结庐在人境，而无车马喧。")
	Info("问君何能尔？心远地自偏。")
	Warning("采菊东篱下，悠然见南山。")
	Error("山气日夕佳，飞鸟相与还。")
	Debug("此中有真意，欲辨已忘言。")
}
func ExampleLogger_SetLevel() {
	lg := New(os.Stdout, true)
	lg.SetLevel(INFO)
	lg.Info("人生如逆旅，我亦是行人-------------%s", "出自宋代苏轼")
	lg.Debug("输不出来的")
	lg.SetLevel(DEBUG)
	lg.Debug("我胡汉三又回来了")
}

func ExampleNew() {
	//把结果输入到buffer里
	buff := new(bytes.Buffer)
	lg := New(buff, true)
	lg.Warning("人间正道是沧桑")
	//输出结果
	fmt.Print(buff.String())
}

//把日志写到文件里去
func ExampleFile() {
	file := "./log/color_log.txt"
	f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	buff := new(bytes.Buffer)
	lg := New(buff, true)
	lg.Warning("人间正道是沧桑")
	//输出结果
	n, err := f.Write(buff.Bytes())
	if err != nil {
		panic(err)
	}
	fmt.Println("写入字节: ", n)
}
func TestError(t *testing.T) {
	ExampleFile()
}

//
func TestExampleColorLog(t *testing.T) {
	ExampleDebug()
	ExampleLogger_SetLevel()
	ExampleNew()
	ExampleFile()
}

func BenchmarkDebug(b *testing.B) {
	b.StopTimer()
	buff := new(bytes.Buffer)
	lg := New(buff, true)
	//输出结果
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		lg.Debug("人间正道是沧桑 ~ debug")
	}
}
func BenchmarkInfo(b *testing.B) {
	b.StopTimer()
	buff := new(bytes.Buffer)
	lg := New(buff, true)
	//输出结果
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		lg.Info("人间正道是沧桑 ~ info")
	}
}
func BenchmarkWarning(b *testing.B) {
	b.StopTimer()
	buff := new(bytes.Buffer)
	lg := New(buff, true)
	//输出结果
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		lg.Warning("人间正道是沧桑 ~ warning")
	}
}

func BenchmarkError(b *testing.B) {
	b.StopTimer()
	buff := new(bytes.Buffer)
	lg := New(buff, true)
	//输出结果
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		lg.Error("人间正道是沧桑 ~ err")
	}
}
