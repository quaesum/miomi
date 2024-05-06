package v1

import (
	"github.com/gavv/httpexpect"
	"net/http"
	"testing"
)

func TestGetAnimalByIDHandler(t *testing.T) {
	tests := []struct {
		name     string
		idHeader string
		id       string
		status   int
	}{
		{
			name:     "bad id header",
			idHeader: "fsd",
			id:       "my",
			status:   http.StatusInternalServerError,
		},
		{
			name:     "all ok",
			idHeader: "123",
			id:       "123",
			status:   http.StatusOK,
		},
	}
	//h := http.Handler
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			e := httpexpect.WithConfig(httpexpect.Config{
				Client: &http.Client{
					Transport: httpexpect.NewBinder(h),
					Jar:       httpexpect.NewJar(),
				},
				Reporter: httpexpect.NewAssertReporter(t),
			})
			e.GET("/animals/"+tc.id).
				WithHeader("id", tc.idHeader).
				Expect().
				Status(tc.status)
		})
	}
}

func TestGetAnimalsHandler(t *testing.T) {
	m := &animalServiceMock{}
	defer m.AssertExpectations(t)

	h := ControllerHandler()

	tests := []struct {
		name   string
		page   int
		status int
	}{
		{
			name:   "bad page",
			page:   -1,
			status: http.StatusBadRequest,
		},
		{
			name:   "all ok",
			page:   1,
			status: http.StatusOK,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			e := httpexpect.WithConfig(httpexpect.Config{
				Client: &http.Client{
					Transport: httpexpect.NewBinder(h),
					Jar:       httpexpect.NewJar(),
				},
				Reporter: httpexpect.NewAssertReporter(t),
			})
			e.GET("/animals").
				WithJSON(map[string]interface{}{"page": tc.page, "perPage": 10}).
				Expect().
				Status(tc.status)
		})
	}
}

func TestCreateAnimalHandler(t *testing.T) {
	m := &animalServiceMock{}
	defer m.AssertExpectations(t)

	h := ControllerHandler()

	tests := []struct {
		name   string
		status int
	}{
		{
			name:   "bad request",
			status: http.StatusBadRequest,
		},
		{
			name:   "all ok",
			status: http.StatusOK,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			e := httpexpect.WithConfig(httpexpect.Config{
				Client: &http.Client{
					Transport: httpexpect.NewBinder(h),
					Jar:       httpexpect.NewJar(),
				},
				Reporter: httpexpect.NewAssertReporter(t),
			})
			e.POST("/animals").
				WithJSON(map[string]interface{}{}).
				Expect().
				Status(tc.status)
		})
	}
}

func TestUpdateAnimalHandler(t *testing.T) {
	m := &animalServiceMock{}
	defer m.AssertExpectations(t)

	h := ControllerHandler()

	tests := []struct {
		name   string
		id     string
		status int
	}{
		{
			name:   "bad id",
			id:     "my",
			status: http.StatusBadRequest,
		},
		{
			name:   "all ok",
			id:     "123",
			status: http.StatusOK,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			e := httpexpect.WithConfig(httpexpect.Config{
				Client: &http.Client{
					Transport: httpexpect.NewBinder(h),
					Jar:       httpexpect.NewJar(),
				},
				Reporter: httpexpect.NewAssertReporter(t),
			})
			e.PATCH("/animals/" + tc.id).
				WithJSON(map[string]interface{}{}).
				Expect().
				Status(tc.status)
		})
	}
}

func TestRemoveAnimalHandler(t *testing.T) {
	m := &animalServiceMock{}
	defer m.AssertExpectations(t)

	h := ControllerHandler()

	tests := []struct {
		name   string
		id     string
		status int
	}{
		{
			name:   "bad id",
			id:     "my",
			status: http.StatusBadRequest,
		},
		{
			name:   "all ok",
			id:     "123",
			status: http.StatusOK,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			e := httpexpect.WithConfig(httpexpect.Config{
				Client: &http.Client{
					Transport: httpexpect.NewBinder(h),
					Jar:       httpexpect.NewJar(),
				},
				Reporter: httpexpect.NewAssertReporter(t),
			})
			e.DELETE("/animals/" + tc.id).
				Expect().
				Status(tc.status)
		})
	}
}
