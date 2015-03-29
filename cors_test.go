package cors_test

import (
	"math/rand"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jamie-stackhouse/gin-cors"
)

func ExampleMiddleware() {
	// Initialize the gin-gonic router
	router := gin.New()

	// Set up CORS middleware options
	config := cors.Config{
		Origins:         "*",
		RequestHeaders:  "Authorization",
		Methods:         "GET, POST, PUT",
		Credentials:     true,
		ValidateHeaders: false,
		MaxAge:          1 * time.Minute,
	}

	// Apply the middleware to the router (works on groups too)
	router.Use(cors.Middleware(config))
}

var sHeaders = "TestOne, TestTwo, TestThree, TestFour, TestFive"

func BenchmarkSortFive(b *testing.B) {
	ssHeaders := sort.StringSlice(strings.Split(sHeaders, ", "))
	sort.Sort(ssHeaders)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		index := rand.Intn(10)
		search := ""

		switch {
		case index >= 5 && index < 10:
			search = "NotFound"
		default:
			search = ssHeaders[index]
		}

		if idx := sort.SearchStrings(ssHeaders, search); ssHeaders[idx] == search {

		}
	}
}

func BenchmarkRangeFive(b *testing.B) {
	ssHeaders := strings.Split(sHeaders, ", ")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		index := rand.Intn(10)
		search := ""

		switch {
		case index >= 5 && index < 10:
			search = "NotFound"
		default:
			search = ssHeaders[index]
		}

		for _, value := range ssHeaders {
			if value == search {
				break
			}
		}
	}
}
