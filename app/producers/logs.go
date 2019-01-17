package producers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/iandkenzt/kafka-golang/utils"
)

// UserActivity ...
func UserActivity(res http.ResponseWriter, req *http.Request) {
	var err error

	// decode requests JSON to map
	var reqJSON map[string]interface{}
	err = json.NewDecoder(req.Body).Decode(&reqJSON)
	if err != nil {
		utils.Response(res, 1, "Error decode requests JSON on UserActivity", "", http.StatusBadRequest, nil)
		return
	}

	// marshal request JSON
	var message []byte
	message, err = json.Marshal(reqJSON)
	if err != nil {
		msg := fmt.Errorf("Error when marshal requests JSON on UserActivity | Error: %s", err.Error())

		utils.Response(res, 1, msg, "", http.StatusBadRequest, nil)
		return
	}

	// publish to Kafka
	p, err := utils.Publish(nil, message)
	if p == false {
		utils.Response(res, 1, err, "", http.StatusBadRequest, nil)
	}

	// return success response
	utils.Response(res, 0, "Success", "", http.StatusOK, nil)

}
