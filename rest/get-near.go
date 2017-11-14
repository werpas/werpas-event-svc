package rest

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/werpas/werpas-event-svc/model"
	"github.com/werpas/werpas-event-svc/mongodb"
	"encoding/json"
)

func HandleGetNear(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req model.RequestGeo
	var res model.ResponseGeo
	var err error

	// set process date
	res.Summary.Date = time.Now().UTC()

	// validate longitude
	long := r.URL.Query().Get("long")
	req.Longitude, err = strconv.ParseFloat(long, 64)
	if err != nil {
		http.Error(w, "longitude is not valid.", http.StatusBadRequest)
		return
	}
	if req.Longitude < -180.0 || req.Longitude > 180.0 {
		http.Error(w, "longitude is not within range.", http.StatusBadRequest)
		return
	}

	// validate latitude
	lat := r.URL.Query().Get("lat")
	req.Latitude, err = strconv.ParseFloat(lat, 64)
	if err != nil {
		http.Error(w, "latitude is not valid.", http.StatusBadRequest)
		return
	}
	if req.Latitude < -90.0 || req.Latitude > 90.0 {
		http.Error(w, "latitude is not within range.", http.StatusBadRequest)
		return
	}

	// validate sort param
	sortParam := r.URL.Query().Get("sort")
	if len(sortParam) > 0 {
		switch sp := strings.ToLower(sortParam); sp {
		case strings.ToLower(model.SORT_PARAM_EVENTTYPE):
			req.SortParam = model.SORT_PARAM_EVENTTYPE
		case strings.ToLower(model.SORT_PARAM_EVENTNAME):
			req.SortParam = model.SORT_PARAM_EVENTNAME
		case strings.ToLower(model.SORT_PARAM_EVENTDESC):
			req.SortParam = model.SORT_PARAM_EVENTDESC
		case strings.ToLower(model.SORT_PARAM_DATETIME):
			req.SortParam = model.SORT_PARAM_DATETIME
		case strings.ToLower(model.SORT_PARAM_DISTANCE):
			req.SortParam = model.SORT_PARAM_DISTANCE
		default:
			http.Error(w, "sort param is not valid.", http.StatusBadRequest)
			return
		}
	} else {
		req.SortParam = model.SORT_PARAM_DISTANCE
	}

	// validate sort param
	sortOrder := r.URL.Query().Get("order")
	if len(sortOrder) > 0 {
		switch so := strings.ToLower(sortOrder); so {
		case model.SORT_ORDER_ASC, model.SORT_ORDER_DESC:
			req.SortOrder = so
		default:
			http.Error(w, "sort order is not valid.", http.StatusBadRequest)
			return
		}
	} else {
		req.SortOrder = model.SORT_ORDER_ASC
	}

	// validate limit
	limit := r.URL.Query().Get("limit")
	if len(limit) > 0 {
		req.Limit, err = strconv.Atoi(limit)
		if err != nil {
			http.Error(w, "limit is not valid.", http.StatusBadRequest)
			return
		}
	} else {
		req.Limit = 100
	}

	// validate radius
	radius := r.URL.Query().Get("radius")
	if len(radius) > 0 {
		req.Radius, err = strconv.Atoi(radius)
		if err != nil {
			http.Error(w, "radius is not valid.", http.StatusBadRequest)
			return
		}
	} else {
		req.Radius = 1000
	}

	// get params
	res.Summary.Params = req

	// query mongodb
	err = mongodb.GetNear(&res)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting events. %v", err.Error()), http.StatusInternalServerError)
		return
	}

	// get count
	res.Summary.Count = len(res.Events)

	// get duration
	res.Summary.Duration = time.Now().Sub(res.Summary.Date).String()

	jsonOutput, err := json.Marshal(res)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error marshaling events. %v", err.Error()), http.StatusInternalServerError)
		return
	}

	// Always set content type and status code
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonOutput))
}
