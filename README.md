# headandshoulders

*helps keep your tests free of flakes*

## usage

``` go
import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"

  . "github.com/xoebus/headandshoulders"
)

var _ = Describe("Usage", func() {
  Describe("helps you make sure all flakiness is gone", func() {
    RIt("by running repeated tests 3 times by default", func() {
      // flaky test...
    })

    RIt("but optionally many times for those more tricky ones", func() {
      // annoyingly not very flaky test...
    }, 15)
  })
})
```

## faq

> But what if I use the timeout feature of Ginkgo's `It` function?

Tough.
