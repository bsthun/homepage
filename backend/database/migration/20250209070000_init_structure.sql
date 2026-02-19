-- +goose Up
-- +goose StatementBegin

CREATE TABLE users
(
    id          BIGSERIAL PRIMARY KEY,
    username    VARCHAR(255) NOT NULL UNIQUE,
    firstname   VARCHAR(255) NOT NULL,
    lastname    VARCHAR(255) NOT NULL,
    email       VARCHAR(255) UNIQUE,
    picture_url VARCHAR(255) NULL,
    is_admin    BOOLEAN      NOT NULL,
    metadata    JSONB        NOT NULL,
    created_at  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE dairies
(
    id                    BIGSERIAL PRIMARY KEY,
    user_id               BIGINT REFERENCES users (id) ON DELETE CASCADE NOT NULL,
    content               TEXT                                           NOT NULL,
    graph_emotional       JSONB                                          NOT NULL,
    graph_productivity    JSONB                                          NOT NULL,
    graph_interaction     JSONB                                          NOT NULL,
    graph_wellness        JSONB                                          NOT NULL,
    created_at            TIMESTAMP                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at            TIMESTAMP                                      NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE dairy_attachments
(
    id          BIGSERIAL PRIMARY KEY,
    dairy_id    BIGINT       REFERENCES dairies (id) ON DELETE CASCADE NOT NULL,
    type        VARCHAR(255) NOT NULL,
    referred_at  VARCHAR(255) NOT NULL,
    created_at  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE blog_tags
(
    id         BIGSERIAL PRIMARY KEY,
    name       VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE blog_series
(
    id         BIGSERIAL PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    created_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE blog_items
(
    id               BIGSERIAL PRIMARY KEY,
    blog_series_id   BIGINT       REFERENCES blog_series (id) ON DELETE SET NULL NULL,
    title            VARCHAR(255) NOT NULL,
    description      TEXT         NOT NULL,
    is_video         BOOLEAN      NOT NULL,
    publish_duration INTEGER      NOT NULL,
    referred_at       VARCHAR(255) NOT NULL,
    created_at       TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE blog_sections
(
    id                BIGSERIAL PRIMARY KEY,
    blog_item_id      BIGINT       REFERENCES blog_items (id) ON DELETE CASCADE NOT NULL,
    title             VARCHAR(255) NOT NULL,
    content           TEXT         NOT NULL,
    spotify_track_id  VARCHAR(255) NULL,
    spotify_start     INTEGER      NULL,
    created_at        TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at        TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE blog_section_revisions
(
    id              BIGSERIAL PRIMARY KEY,
    blog_section_id BIGINT REFERENCES blog_sections (id) ON DELETE CASCADE NOT NULL,
    change          TEXT                                                   NOT NULL,
    created_at      TIMESTAMP                                              NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE sessions
(
    id         BIGSERIAL PRIMARY KEY,
    user_id    BIGINT REFERENCES users (id) ON DELETE CASCADE NOT NULL,
    metadata   JSONB                                          NOT NULL,
    created_at TIMESTAMP                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP                                      NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- * auto-update function for updated_at timestamps
CREATE OR REPLACE FUNCTION auto_updated_at()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- * triggers to automatically update updated_at
CREATE TRIGGER auto_updated_at_users
    BEFORE UPDATE
    ON users
    FOR EACH ROW
EXECUTE FUNCTION auto_updated_at();

CREATE TRIGGER auto_updated_at_dairies
    BEFORE UPDATE
    ON dairies
    FOR EACH ROW
EXECUTE FUNCTION auto_updated_at();

CREATE TRIGGER auto_updated_at_dairy_attachments
    BEFORE UPDATE
    ON dairy_attachments
    FOR EACH ROW
EXECUTE FUNCTION auto_updated_at();

CREATE TRIGGER auto_updated_at_blog_tags
    BEFORE UPDATE
    ON blog_tags
    FOR EACH ROW
EXECUTE FUNCTION auto_updated_at();

CREATE TRIGGER auto_updated_at_blog_series
    BEFORE UPDATE
    ON blog_series
    FOR EACH ROW
EXECUTE FUNCTION auto_updated_at();

CREATE TRIGGER auto_updated_at_blog_items
    BEFORE UPDATE
    ON blog_items
    FOR EACH ROW
EXECUTE FUNCTION auto_updated_at();

CREATE TRIGGER auto_updated_at_blog_sections
    BEFORE UPDATE
    ON blog_sections
    FOR EACH ROW
EXECUTE FUNCTION auto_updated_at();

CREATE TRIGGER auto_updated_at_sessions
    BEFORE UPDATE
    ON sessions
    FOR EACH ROW
EXECUTE FUNCTION auto_updated_at();

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS sessions;
DROP TABLE IF EXISTS blog_section_revisions;
DROP TABLE IF EXISTS blog_sections;
DROP TABLE IF EXISTS blog_items;
DROP TABLE IF EXISTS blog_series;
DROP TABLE IF EXISTS blog_tags;
DROP TABLE IF EXISTS dairy_attachments;
DROP TABLE IF EXISTS dairies;
DROP TABLE IF EXISTS users;
DROP FUNCTION IF EXISTS auto_updated_at();
-- +goose StatementEnd
