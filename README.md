go-deathtest
============

[![Build Status](https://secure.travis-ci.org/pwaller/go-deathtest.png)](http://travis-ci.org/pwaller/go-deathtest)

Example usage:

    // myproject_test.go
    import (
        "testing"
        
        "github.com/pwaller/go-deathtest"
    )
    
    func TestBadThing(t *testing.T) {
        if !deathtest.Run(t) { return }
        // something which causes go to crash
    }
    
    
Then run `go test` as you normally would.

The test will pass only if `TestBadThing()` causes go to exit with a failure
code. This is useful for testing that code fails when it should, without 
bringing down the whole test suite.

The death test runs in a separate go process, which causes `deathtest.Run(t)` to
return true, allowing the crash branch to run.

Caveat: commandline options are not currently passed to the death tests.

Pull requests welcome.
