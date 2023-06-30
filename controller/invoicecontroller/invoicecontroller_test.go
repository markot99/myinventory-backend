package invoicecontroller

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/markot99/myinventory-backend/testhelper"
)

func setupTest(t *testing.T) (*gin.Engine, *testhelper.TestHelper, func()) {
	testHelper := testhelper.CreateTestHelper(t)
	router := testHelper.Router

	routes := router.Group("/api")

	RegisterRoutes(routes, testHelper.Authenticator, testHelper.ItemTable, testHelper.InvoiceStorage)
	return router, &testHelper, func() {
		testHelper.Teardown()
	}
}
