package factory

import (
	"{{.moduleName}}/pkg/domains/errorcodes"
	"{{.moduleName}}/pkg/domains/errs"

	"github.com/karim-w/stdlib/sqldb"
	"go.uber.org/zap"
)

func (s *sf) PSQL() sqldb.DB {
	return deps.PSQL
}

func (s *sf) BeginPsqlTx() (*sqldb.Tx, bool, error) {
	s.sql_mtx.Lock()
	defer s.sql_mtx.Unlock()

	if s.psqlTx != nil {
		return s.psqlTx, false, nil
	}

	var err error

	s.psqlTx, err = deps.PSQL.BeginTx(s.ctx, nil)
	if err != nil {
		s.logger.Error("failed to begin psql tx", zap.Error(err))

		err = errs.SqlError(err, errorcodes.FAILED_TO_BEGIN_SQL_TRANSACTION, s.traceparent)

		return nil, false, err
	}

	return s.psqlTx, true, nil
}

func (s *sf) CommitPsqlTx() error {
	s.sql_mtx.Lock()
	defer s.sql_mtx.Unlock()

	if s.psqlTx == nil {
		return nil
	}

	err := s.psqlTx.Commit()
	if err != nil {
		s.logger.Error("failed to commit psql tx", zap.Error(err))
		return err
	}

	s.psqlTx = nil
	return nil
}
