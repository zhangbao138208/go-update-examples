package main

import (
	"net/http"

	"time"

	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"gue/bindata"

	"github.com/inconshreveable/go-update"
)

func restart() {
	// s := []string{"cmd.exe", "/C", "start", `restart.exe`}

	// cmd := exec.Command(s[0], s[1:]...)
	s := []string{"restart.exe"}

	cmd := exec.Command(s[0])
	if err := cmd.Run(); err != nil {
		fmt.Println("Error:", err)
	}
	os.Exit(0)
}

var version string

func doUpdate(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	err = update.Apply(
		//resp.Body, update.Options{Patcher: update.NewBSDiffPatcher()})
		resp.Body, update.Options{})
	if err != nil {
		// error handling
		panic(err)
	}

	return err
}

func restore() {
	dirs := []string{"asserts", "asset", "config.yaml", "restart.sh", "restart.bat", "restart.exe"} // 设置需要释放的目录
	// for _, dir := range dirs {
	// 	os.RemoveAll(filepath.Join("./", dir))
	// }
	isSuccess := true
	for _, dir := range dirs {
		// 解压dir目录到当前目录
		if err := bindata.RestoreAssets("./", dir); err != nil {
			isSuccess = false
			break
		}
	}
	if !isSuccess {
		for _, dir := range dirs {
			os.RemoveAll(filepath.Join("./", dir))
		}
	}
}

func main() {

	ch := make(chan bool)
	done := make(chan struct{} )
	
	go run(ch, done)
	ShowDialog(ch)
	//isUpdate := <-ch
	
	<-done

}
func run(ch chan bool, done chan struct{}) {
	restore()
	updateFlag := flag.Bool(
		"update", false,
		"Update to specified version")

	versionFlag := flag.String(
		"version", "",
		"Update to this version")

	flag.Parse()
	isUpdate := <-ch
	if *updateFlag || isUpdate {
		url := fmt.Sprintf(
			//"http://localhost:8080/static/patch-%v-%v",
			//version,
			"http://localhost:8080/static/ex-%v.exe",
			*versionFlag)
		fmt.Println(fmt.Sprintf("Updating to %v", url))
		doUpdate(url)
		restart()
	}

	fmt.Println(fmt.Sprintf("version %v", version))
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hha1")
	}
	done <- struct{}{}
}
