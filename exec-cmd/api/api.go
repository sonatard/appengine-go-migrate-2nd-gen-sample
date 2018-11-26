package api

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func init() {
	// wget
	out0, err := exec.Command("wget").CombinedOutput()
	log.Printf("wget : %v\n", string(out0))
	log.Printf("wget : error %v\n", err)

	out1, err := exec.Command("wget", "https://github.com/sonatard/ghs/releases/download/0.0.10/ghs-0.0.10-linux_amd64.tar.gz", "-O", "/tmp/ghs.tar.gz").CombinedOutput()
	log.Printf("wget : %v\n", string(out1))
	log.Printf("wget : error %v\n", err)

	// tar
	out2, err := exec.Command("tar", "xvf", "/tmp/ghs.tar.gz", "-C", "/tmp/").CombinedOutput()
	log.Printf("tar: %v\n", string(out2))
	log.Printf("tar: error %v\n", err)
}

func IndexHandle(w http.ResponseWriter, r *http.Request) {
	// exec binary
	out, err := exec.Command("/tmp/ghs-0.0.10-linux_amd64/ghs", "golang/go", "-m1").CombinedOutput()
	log.Printf("ghs: %v\n", string(out))
	log.Printf("ghs: %v\n", err)
	fmt.Fprintf(w, "%v", string(out))
}
