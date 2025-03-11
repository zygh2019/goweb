package res

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

/*
*
 */
func CoverJson(c *gin.Context, any any) (err error) {
	switch c.GetHeader("Content-Type") {
	case "application/json":
		data, _ := c.GetRawData()
		err = json.Unmarshal(data, &any)
		if err != nil {
			logrus.Info("bind fail")
			return nil
		}
		return nil
	default:
		return nil
	}

}
