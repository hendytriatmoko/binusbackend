package def

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"runtime/debug"
)

func DefFoo() {
	//fmt.Println("check error bos")

	if r := recover(); r != nil {
		fmt.Println("eh ada error bos ", r)
	}

	//fmt.Println("check error kelar bos")

}

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//Print error stack information
			log.Printf("panic: %v\n", r)
			fmt.Println("panic: ", fmt.Sprint(r))
			debug.PrintStack()
			//Package general json return
			//c.JSON(http.StatusOK, Result.Fail(errorToString(r)))
			//Result.Fail is not the focus of this example, so use the following code instead

			type Response struct {
				ApiMessage string      `json:"api_message"`
				Count      int64       `json:"count"`
				Data       interface{} `json:"data"`
			}

			res := Response{
				ApiMessage: errorToString(r),
			}

			c.JSON(400, res)

			// Terminate subsequent interface calls, if not added, after recovering to an exception, the subsequent code in the interface will continue to execute
			c.Abort()
		}
	}()
	//After loading defer recover, continue with subsequent interface calls
	c.Next()
}

// recover error, turn to string
func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
