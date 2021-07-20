package usecases_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/golang/mock/gomock"
	mock_usecases "github.com/peterstirrup/messages/internal/messages/usecases/mock"

	"github.com/peterstirrup/messages/internal/messages/usecases"
)

var (
	ctx = context.Background()

	testUserID = int64(10)
)

type setupUseCasesTestConfig struct {
	mockCtrl *gomock.Controller
	usecases *usecases.Messages
	whatsapp *mock_usecases.MockWhatsApp
}

func setupUseCasesTest(t *testing.T) *setupUseCasesTestConfig {
	ctrl := gomock.NewController(t)
	w := mock_usecases.NewMockWhatsApp(ctrl)

	u := usecases.New(w)

	return &setupUseCasesTestConfig{
		mockCtrl: ctrl,
		usecases: u,
		whatsapp: w,
	}
}

func teardownWhatsAppTest(cfg *setupUseCasesTestConfig) {
	cfg.mockCtrl.Finish()
}

func gotWant(t *testing.T, got, want interface{}) {
	if diff := cmp.Diff(got, want); diff != "" {
		fmt.Println(diff)
		t.Fail()
	}
}
