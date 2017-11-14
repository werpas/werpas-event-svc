package mongodb

import (
	"github.com/werpas/werpas-event-svc/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

func GetNear(resp *model.ResponseGeo) (err error) {

	session, err := mgo.Dial(connectionString)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// get sort order
	order := 1
	if strings.EqualFold(resp.Summary.Params.SortOrder, model.SORT_ORDER_DESC) {
		order = -1
	}

	// query the database
	c := session.DB(database).C(collection)
	err = c.Pipe([]bson.M	{
					{
						"$geoNear":
							bson.M	{
									"near": bson.M{ "type": "Point", "coordinates": []float64{resp.Summary.Params.Longitude, resp.Summary.Params.Latitude} },
									"distanceField": "distance",
									"maxDistance": resp.Summary.Params.Radius,
									"query": "{}",
									"limit": resp.Summary.Params.Limit,
									"spherical": true,
								},
					},
					{
						"$sort": bson.M{ resp.Summary.Params.SortParam: order },
					},
				},
		).All(&resp.Events)
	return
}


