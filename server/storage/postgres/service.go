package postgres

import (
	"context"

	"github.com/lib/pq"

	pb "gophkeeper/proto"
)

// InsertDataIntoDataTable производит запись пользовательских данных в БД.
func (s DBConnect) InsertDataIntoDataTable(ctx context.Context, userID int, userData *pb.SendDataRequestArray) (dataID int, error error) {

	tx, err := s.DBConnect.BeginTx(ctx, nil)
	if err != nil {
		s.Logger.Error().Msg(err.Error())
		return 0, err
	}
	defer tx.Rollback()

	sqlInsertData, err := tx.Prepare("INSERT INTO user_data (user_ID, type_ID, title, metadata, checksum, data) VALUES ((SELECT user_ID from users WHERE user_ID=$1), (SELECT type_ID from data_types WHERE type_title=$2), $3, $4, $5, $6) RETURNING data_id;")

	if err != nil {
		s.Logger.Error().Msg(err.Error())
		return 0, err
	}
	defer sqlInsertData.Close()

	err = sqlInsertData.QueryRow(userID, userData.DataType, userData.Title, userData.Metadata, userData.Checksum, userData.Data).Scan(&dataID)
	if err != nil {
		s.Logger.Error().Msg(err.Error())
		return 0, err
	}

	tx.Commit()

	return dataID, nil
}

// GetUpdatesByChecksums получает на вход массив с контрольными суммами и возвращает новые или измененные данные.
func (s DBConnect) GetUpdatesByChecksums(ctx context.Context, userID int, checkSums []string) ([]*pb.GetDataResponseArray, error) {
	var result []*pb.GetDataResponseArray

	rows, err := s.DBConnect.QueryContext(ctx,
		"select data_id, data_types.type_title, title, metadata, create_date, checksum, data "+
			"from user_data "+
			"inner join data_types on data_types.type_id = user_data.type_id "+
			"where is_delete = False and user_data.user_id = $1 and user_data.checksum <> ALL ($2)", userID, pq.Array(checkSums))

	if err != nil {
		s.Logger.Error().Msg(err.Error())
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		var v pb.GetDataResponseArray

		err = rows.Scan(&v.DataID, &v.DataType, &v.Title, &v.Metadata, &v.CreateDate, &v.CheckSum, &v.Userdata)
		if err != nil {
			s.Logger.Error().Msg(err.Error())
			return []*pb.GetDataResponseArray{}, err
		}

		result = append(result, &v)
	}

	err = rows.Err()
	if err != nil {
		s.Logger.Error().Msg(err.Error())
		return []*pb.GetDataResponseArray{}, err
	}

	return result, nil
}

// UpdateData обновляет пользовательские данные.
func (s DBConnect) UpdateData(ctx context.Context, userData *pb.UpdateDataRequestArray) (error error) {

	tx, err := s.DBConnect.BeginTx(ctx, nil)
	if err != nil {
		s.Logger.Error().Msg(err.Error())
		return err
	}
	defer tx.Rollback()

	sqlInsertData, err := tx.Prepare("UPDATE user_data SET title = $1, metadata = $2, checksum = $3, data = $4 WHERE data_id = $5;")

	if err != nil {
		s.Logger.Error().Msg(err.Error())
		return err
	}
	defer sqlInsertData.Close()

	_, err = sqlInsertData.Exec(userData.Title, userData.Metadata, userData.Checksum, userData.Data, userData.ServerID)
	if err != nil {
		s.Logger.Error().Msg(err.Error())
		return err
	}

	tx.Commit()

	return nil
}
