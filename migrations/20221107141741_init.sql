-- +goose Up
-- +goose StatementBegin
CREATE TABLE FontFamily (
	id TEXT NOT NULL PRIMARY KEY,
	name TEXT NOT NULL,
	license TEXT,
	creator TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE FontFamily;
-- +goose StatementEnd
