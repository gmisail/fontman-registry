-- +goose Up
-- +goose StatementBegin
CREATE TABLE FontStatistics (
	id TEXT NOT NULL PRIMARY KEY,
	family TEXT NOT NULL,
	downloads INTEGER NOT NULL,
	FOREIGN KEY (family) REFERENCES FontFamily(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE FontStatistics;
-- +goose StatementEnd
