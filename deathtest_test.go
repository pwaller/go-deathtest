package deathtest

import (
    "log"
    "testing"
    "unsafe"
)

var run_successful_test = false

func TestDeathTest(t *testing.T) {
    if !Run(t) { return }
    log.Fatal("This is an intentional failure")
}

func TestBadMemoryPanic(t *testing.T) {
    if !Run(t) { return }
    // Try to allocate 100000GB ram (surely, this should fail....)
    slices := make([][]byte, 0)
    for i := 0; i < 100000; i++ {
        x := make([]byte, 1*1024*1024*1024)
        x[0] = 1
        slices = append(slices, x)
    }
}

func TestSegfault(t *testing.T) {
    if !Run(t) { return }
    x := uintptr(0xbaff1ed)
    p := unsafe.Pointer(x)
    var v *int = (*int)(p)
    log.Print("V has a value: ", *v)
}

func TestDeathTestWhenSuccess(t *testing.T) {
    if !run_successful_test && !*deathtest_running {
        log.Print("TestDeathTestWhenSuccess test is a NOOP outside TestSuccessShouldFail")
        return
    }
    if !Run(t) { return }
    log.Print("This test is an intentional failure and succeed")
}

func TestSuccessShouldFail(t *testing.T) {
    log.Print("--- START: FAILS inside this block are intentional and can be ignored ---")
    run_successful_test = true
    ok := testing.RunTests(func(pat, str string) (bool, error) { return true, nil },
        []testing.InternalTest{testing.InternalTest{"TestDeathTestWhenSuccess", TestDeathTestWhenSuccess}})
    run_successful_test = false
    log.Print("--- END: FAILS inside this block are intentional and can be ignored ---")
    if ok {
        t.Error("TestDeathTestWhenSuccess should fail")
    }
}
