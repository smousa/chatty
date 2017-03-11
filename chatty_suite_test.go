package main_test

import (
	"context"
	"math/rand"

	"github.com/goadesign/goa"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestChatty(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Chatty Suite")
}

var (
	ctx context.Context
	svc *goa.Service
)

var _ = BeforeSuite(func() {
	rand.Seed(GinkgoRandomSeed())
})

var _ = BeforeEach(func() {
	ctx = context.Background()
	svc = goa.New("chatty")
})
