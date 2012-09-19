package deathtest

import (
	"flag"
	"log"
	"os/exec"
	"reflect"
	"strings"
	"testing"
)

var deathtest_running = flag.Bool("deathtest.running", false, "death test is running")

type deathtestWriter struct{}

func (m deathtestWriter) Write(in []byte) (int, error) {
	for _, l := range strings.Split(string(in), "\n") {
		log.Print("  @DT: ", l)
	}
	return len(in), nil
}

func Run(t *testing.T) bool {
	if *deathtest_running {
		return true
	}
	name := reflect.ValueOf(*t).FieldByName("name").String()

	args := []string{"go", "test", "-deathtest.running", "-run", name}
	log.Printf("deathtest.Run(%v)", args)
	p := exec.Command(args[0], args[1:]...)
	w := deathtestWriter{}
	p.Stdout, p.Stderr = w, w
	result := p.Run()

	if result == nil {
		log.Print("deathtest.Run: Failure to fail")
		t.Fail()
	}
	return false
}
