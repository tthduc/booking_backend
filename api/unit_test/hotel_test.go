package unit_test

import (
	"booking-backed/api"
	mockdb "booking-backed/db/mock"
	db "booking-backed/db/sqlc"
	"booking-backed/util"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetHotelAPI(t *testing.T) {
	hotel := randomHotel()

	testCases := []struct {
		name          string
		hotelID       int64
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:    "OK",
			hotelID: hotel.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetHotel(gomock.Any(), gomock.Eq(hotel.ID)).
					Times(1).
					Return(hotel, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchHotel(t, recorder.Body, hotel)
			},
		},
		{
			name:    "NotFound",
			hotelID: hotel.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetHotel(gomock.Any(), gomock.Eq(hotel.ID)).
					Times(1).
					Return(db.Hotel{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:    "InternalError",
			hotelID: hotel.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetHotel(gomock.Any(), gomock.Eq(hotel.ID)).
					Times(1).
					Return(db.Hotel{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:    "InvalidID",
			hotelID: 0,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetHotel(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := api.NewServer(store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/hotel/%d", tc.hotelID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestCreateHotelAPI(t *testing.T) {
	hotel := randomHotel()

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"name":     hotel.Name,
				"address":  hotel.Address,
				"location": hotel.Location,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreateHotelParams{
					Name:     hotel.Name,
					Address:  hotel.Address,
					Location: hotel.Location,
				}

				store.EXPECT().
					CreateHotel(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(hotel, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchHotel(t, recorder.Body, hotel)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"name":     hotel.Name,
				"address":  hotel.Address,
				"location": hotel.Location,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateHotel(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Hotel{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := api.NewServer(store)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/hotel"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func TestListHotelsAPI(t *testing.T) {
	n := 5
	hotels := make([]db.Hotel, n)
	for i := 0; i < n; i++ {
		hotels[i] = randomHotel()
	}

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"page_id":   1,
				"page_size": n,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.ListHotelsParams{
					Limit:  int32(n),
					Offset: 0,
				}

				store.EXPECT().
					ListHotels(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(hotels, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchHotels(t, recorder.Body, hotels)
			},
		},
		{
			name: "InternalError",
			body: gin.H{
				"page_id":   1,
				"page_size": n,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListHotels(gomock.Any(), gomock.Any()).
					Times(1).
					Return([]db.Hotel{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InvalidPageID",
			body: gin.H{
				"page_id":   -1,
				"page_size": n,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListHotels(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InvalidPageSize",
			body: gin.H{
				"page_id":   1,
				"page_size": -1,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					ListHotels(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := api.NewServer(store)
			recorder := httptest.NewRecorder()

			// Marshal body data to JSON
			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/hotels"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.Router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func randomHotel() db.Hotel {
	return db.Hotel{
		ID:       util.RandomInt(1, 1000),
		Name:     util.RandomOwner(),
		Address:  util.RandomAddress(),
		Location: util.RandomAddress(),
	}
}

func requireBodyMatchHotel(t *testing.T, body *bytes.Buffer, account db.Hotel) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotHotel db.Hotel
	err = json.Unmarshal(data, &gotHotel)
	require.NoError(t, err)
	require.Equal(t, account, gotHotel)
}

func requireBodyMatchHotels(t *testing.T, body *bytes.Buffer, accounts []db.Hotel) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotHotels []db.Hotel
	err = json.Unmarshal(data, &gotHotels)
	require.NoError(t, err)
	require.Equal(t, accounts, gotHotels)
}
