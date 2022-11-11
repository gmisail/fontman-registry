-- +goose Up
-- +goose StatementBegin
CREATE TABLE FontStyle (
	id TEXT NOT NULL PRIMARY KEY,
	family TEXT NOT NULL,
	type TEXT NOT NULL,
	url TEXT NOT NULL,
	FOREIGN KEY (family) REFERENCES FontFamily(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE FontStyle;
-- +goose StatementEnd
