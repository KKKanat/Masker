package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceRun(t *testing.T) {
	mockProducer := &MockProducer{}
	mockPresenter := &MockPresenter{}

	input := "Here's my spammy page: http://hehefouls.netHAHAHA see you."
	expectedOutput := "Here's my spammy page: http://******************* see you."

	mockProducer.On("Produce").Return(input, nil)
	mockPresenter.On("Present", expectedOutput).Return(nil)

	service := &Service{
		Prod: mockProducer,
		Pres: mockPresenter,
	}

	service.Run()

	mockProducer.AssertCalled(t, "Produce")
	mockPresenter.AssertCalled(t, "Present", expectedOutput)

	mockInput, err := mockProducer.Produce()
	assert.NoError(t, err, "Produce should not return an error")
	assert.Equal(t, input, mockInput, "Produced data should match the input")

	err = mockPresenter.Present(expectedOutput)
	assert.NoError(t, err, "Present should not return an error")

	mockProducer.AssertExpectations(t)
	mockPresenter.AssertExpectations(t)
}
