package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleHelloWorld))
	defer testServer.Close()

	testClient := testServer.Client()

	t.Run("success", func(t *testing.T) {

		res, err := testClient.Get(testServer.URL)
		if err != nil {
			t.Error(err)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			t.Error(err)
		}

		assert.NoError(t, err)
		assert.Equal(t, 200, res.StatusCode)
		assert.Equal(t, "Hello World!", string(body))
	})

	t.Run("should failed", func(t *testing.T) {
		body := strings.NewReader("New Body")

		res, err := testClient.Post(testServer.URL, "application/json", body)
		if err != nil {
			t.Error(err)
		}

		assert.NoError(t, err)
		assert.Equal(t, 405, res.StatusCode)
	})
}

func TestHealth(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleHealth))
	defer testServer.Close()

	testClient := testServer.Client()

	t.Run("success", func(t *testing.T) {

		res, err := testClient.Get(testServer.URL)
		if err != nil {
			t.Error(err)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			t.Error(err)
		}

		assert.NoError(t, err)
		assert.Equal(t, 200, res.StatusCode)
		assert.Equal(t, "OK!", string(body))
	})

	t.Run("should failed", func(t *testing.T) {
		body := strings.NewReader("New Body")

		res, err := testClient.Post(testServer.URL, "application/json", body)
		if err != nil {
			t.Error(err)
		}

		assert.NoError(t, err)
		assert.Equal(t, 405, res.StatusCode)
	})
}

func TestNewEnpoint(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleNewEndpoint))
	defer testServer.Close()

	testClient := testServer.Client()

	t.Run("success", func(t *testing.T) {

		res, err := testClient.Get(testServer.URL)
		if err != nil {
			t.Error(err)
		}

		body, err := io.ReadAll(res.Body)
		if err != nil {
			t.Error(err)
		}

		assert.NoError(t, err)
		assert.Equal(t, 200, res.StatusCode)
		assert.Equal(t, "OK!", string(body))
	})

	t.Run("should failed", func(t *testing.T) {
		body := strings.NewReader("New Body")

		res, err := testClient.Post(testServer.URL, "application/json", body)
		if err != nil {
			t.Error(err)
		}

		assert.NoError(t, err)
		assert.Equal(t, 405, res.StatusCode)
	})
}
