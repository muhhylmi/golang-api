package repositories

import (
	"database/sql"
	"golang-api/modules/books/models/domain"
	"golang-api/utils/database"
	"golang-api/utils/logger"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TestSuite struct {
	suite.Suite
	mock   sqlmock.Sqlmock
	db     *sql.DB
	gormDb *gorm.DB
}

func (suite *TestSuite) SetupSuite() {
	var err error
	suite.db, suite.mock, err = sqlmock.New()
	if err != nil {
		suite.T().Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	dialector := postgres.New(postgres.Config{Conn: suite.db})
	suite.gormDb, _ = gorm.Open(dialector, &gorm.Config{})
}

func (suite *TestSuite) TearDownSuite() {
	suite.db.Close()
}

func (suite *TestSuite) TestFindAll() {
	suite.T().Skip()
	suite.T().Run("should return nil if data not found", func(t *testing.T) {
		suite.mock.ExpectQuery(regexp.QuoteMeta(
			`SELECT "logee_warehouses"."id","logee_warehouses"."is_deleted","logee_warehouses"."created_at","logee_warehouses"."created_by","logee_warehouses"."updated_at","logee_warehouses"."updated_by","logee_warehouses"."app_code","logee_warehouses"."building_code","logee_warehouses"."warehouse_code","logee_warehouses"."warehouse_type_code","logee_warehouses"."name","logee_warehouses"."description","logee_warehouses"."length","logee_warehouses"."width","logee_warehouses"."height","logee_warehouses"."warehouse_pinpoint","logee_warehouses"."warehouse_pinpoint_detail","logee_warehouses"."warehouse_pinpoint_address","Building"."id" AS "Building__id","Building"."is_deleted" AS "Building__is_deleted","Building"."created_at" AS "Building__created_at","Building"."created_by" AS "Building__created_by","Building"."updated_at" AS "Building__updated_at","Building"."updated_by" AS "Building__updated_by","Building"."app_code" AS "Building__app_code","Building"."building_code" AS "Building__building_code","Building"."building_type_code" AS "Building__building_type_code","Building"."address_id" AS "Building__address_id","Building"."company_id" AS "Building__company_id","Building"."name" AS "Building__name","Building"."description" AS "Building__description" FROM "logee_warehouses" INNER JOIN "logee_buildings" "Building" ON "logee_warehouses"."building_code" = "Building"."building_code" AND "Building"."company_id" = $1 WHERE id = $2 ORDER BY "logee_warehouses"."id" LIMIT 1`)).
			WithArgs()

		defaultLogger := logger.Newlogger()

		repo := NewRepositoryImpl(defaultLogger, &database.DBService{
			Gorm: suite.gormDb,
		})
		result, _ := repo.FindAll()
		require.NoError(suite.T(), suite.mock.ExpectationsWereMet())
		assert.Equal(t, ([]*domain.Book)(nil), result)
	})
}
