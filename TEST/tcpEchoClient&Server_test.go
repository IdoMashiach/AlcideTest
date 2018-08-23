package main

import (
    "testing"
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
)

func TestEcho(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "Tcp Echo Server&Client")
}

var _ = Describe("Gadi", func() {

	var data string = "Gadi is the king"
	
	go client(&data)
	go server()
     
    It("should be gadi", func() {
        Expect(data).To(Equal("Gadi is the king"))
    })

})

