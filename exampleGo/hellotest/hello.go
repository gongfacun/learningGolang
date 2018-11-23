package main

import "fmt"
import "time"
import "io"
import "io/ioutil"
import "errors"
import "os"
import "strings"
import "bufio"

import (
	"image"
	"image/color"
	"image/gif"
	"math"
	"math/rand"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	//pipeReader, pipeWriter := io.Pipe()

	//go PipeWrite(pipeWriter)
	//go PipeReade(pipeReader)
	//time.Sleep(30 * time.Second)

	//	dir := os.Args[1]
	//	ListAll(dir, 0)
	//	echoword()
	//	showInputWord()
	lissajous(os.Stdout)

}

func PipeWrite(write *io.PipeWriter) {
	data := []byte("新的一天新的世界开心每一天")
	for i := 1; i < 3; i++ {
		n, err := write.Write(data)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("写入了多少个字节%d\n", n)

	}
	write.CloseWithError(errors.New("写入数据完成"))

}

func PipeReade(read *io.PipeReader) {
	buf := make([]byte, 128)
	for {
		fmt.Println("接收开始阻塞5秒")
		time.Sleep(5 * time.Second)
		fmt.Println("开始接收")
		n, err := read.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("读取到的字节是%d\n", n)
		var s string = string(buf)
		fmt.Println(s)
	}

}

func ListAll(path string, curHier int) {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, fileinfo := range fileInfos {
		if fileinfo.IsDir() {
			for tmpHier := curHier; tmpHier > 0; tmpHier-- {

				fmt.Printf("|\t")

			}
			fmt.Println(fileinfo.Name() + "\\")

			ListAll(path+"/"+fileinfo.Name(), curHier+1)

		} else {

			for tmpHier := curHier; tmpHier > 0; tmpHier-- {
				fmt.Printf("|\t")
			}

			fmt.Println(fileinfo.Name())
		}

	}

}

func echoword() {

	fmt.Println(time.Now())
	var totals, step string
	for i := 1; i < len(os.Args); i++ {

		totals = totals + step + os.Args[i]
		step = "  "
	}
	step = ""
	fmt.Println(totals)

	fmt.Println(time.Now())
}
func echoword2() {
	s, sep := "", "  "
	for i, arg := range os.Args[1:] {
		s = fmt.Sprint(i) + sep + arg
		fmt.Println(s)

	}

}
func echoword3() {

	fmt.Println(strings.Join(os.Args[0:], " "))
}

func showInputWord() {

	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countsLines(os.Stdin, counts)
	} else {

		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup:%v\n", err)
				continue

			}

			countsLines(f, counts)

			f.Close()

		}
	}

	//input := bufio.NewScanner(os.Stdin)
	//for input.Scan() {
	//	counts[input.Text()]++
	//}

	for line, n := range counts {

		fmt.Println("start print")
		//if n > 1 {
		fmt.Printf("%d\t%s\n", n, line)
		//}

	}

}

func countsLines(f *os.File, counts map[string]int) {

	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	fmt.Println("countslines")

}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {

		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)

}
