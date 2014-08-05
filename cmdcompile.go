package main

import  (
	"fmt"
	"log"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
	
	"github.com/brunetto/goutils/debug"
)

func main () () {
	defer debug.TimeMe(time.Now())
	
	var(
		wdir string
		err error
		dirs []os.FileInfo
		folder os.FileInfo
	)
	
	if wdir, err = os.Getwd(); err != nil {
		log.Fatal("Can't read local folder")
	}
	
	if filepath.Base(wdir) != "cmd" {
		log.Fatal("Not in cmd folder, please, go to it. You are in: ", wdir)
	}

	if dirs, err = ioutil.ReadDir(wdir); err != nil {
		log.Fatal("Can't list folders here: ", wdir)
	}
	
	log.Println("Found")
	for _, folder = range dirs {
		fmt.Println(folder.Name())
	}
	
	for _, folder = range dirs {
		if strings.HasPrefix(folder.Name(), "_") {
			continue
		}
		if folder.IsDir() {
			fmt.Println("Entering ", folder.Name())
			if err = os.Chdir(folder.Name()); err != nil {
				log.Println(os.Getwd())
				log.Fatalf("Can't enter in %v with error %v\n", folder.Name(), err)
			}
			fmt.Println("Compiling and installing")
			BuildAndInstall()
			if err = os.Chdir(".."); err != nil {
				log.Fatalf("Can't enter in parent folder %v with error %v\n: ", filepath.Dir(folder.Name()), err)
			}
		}
	}
}


func BuildAndInstall() () {
	var (
		build *exec.Cmd
		install *exec.Cmd
		err error
	)
	
	build = exec.Command("go", "build")
	if build.Stdout = os.Stdout; err != nil {
		log.Fatal("Error connecting STDOUT: ", err)
	}
	if build.Stderr = os.Stderr; err != nil {
		log.Fatal("Error connecting STDERR: ", err)
	}
	
	if err = build.Start(); err != nil {
		log.Fatal("Error starting build: ", err)
	}
	
	if err = build.Wait(); err != nil {
		log.Fatal("Error while waiting for build: ", err)
	}
	
	install = exec.Command("go", "install")
	if install.Stdout = os.Stdout; err != nil {
		log.Fatal("Error connecting STDOUT: ", err)
	}
	if install.Stderr = os.Stderr; err != nil {
		log.Fatal("Error connecting STDERR: ", err)
	}
	
	if err = install.Start(); err != nil {
		log.Fatal("Error starting install: ", err)
	}
	
	if err = install.Wait(); err != nil {
		log.Fatal("Error while waiting for install: ", err)
	}
}

