package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Chatbot/logger"
	"github.com/Chatbot/service"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestChatbotHandler(t *testing.T) {
	logger.Init()

	ctrl := gomock.NewController(t)
	mockBot := service.NewMockBotService(ctrl)

	e := echo.New()

	input := Message{
		Input: "Hello",
	}

	jsonReq, err := json.Marshal(input)
	if err != nil {
		t.Fatal(err)
	}

	testcases := []struct {
		description string
		mocks       *gomock.Call
		req         []byte
		expOutput   *Response
		expectedErr error
	}{
		{
			description: "success",
			req:         jsonReq,
			mocks:       mockBot.EXPECT().GetReply(input.Input).Return("Hi there! How can I help you today?").Times(1),
			expOutput:   &Response{Reply: "Hi there! How can I help you today?"},
		},
		{
			description: "error: Invalid Input",
			req:         []byte{6},
			expOutput:   nil,
			expectedErr: NewErrorResponse("invalid input", http.StatusBadRequest),
		},
		{
			description: "error: Missing Input",
			req:         nil,
			expOutput:   nil,
			expectedErr: NewErrorResponse("input is empty", http.StatusBadRequest),
		},
	}

	for _, v := range testcases {
		req := httptest.NewRequest(http.MethodPost, "/chat", bytes.NewReader(v.req))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		h := NewChatbotHandler(mockBot)

		resp, err := h.HandleChat(ctx)
		assert.Equal(t, v.expOutput, resp)
		assert.Equal(t, v.expectedErr, err)
	}
}
