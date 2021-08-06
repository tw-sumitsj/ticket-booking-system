CREATE TABLE tickets
(
    id         serial PRIMARY KEY NOT NULL,
    catalog_id serial             NOT NULL,
    slot_id    serial             NOT NULL,
    CONSTRAINT fk_catalog
        FOREIGN KEY (catalog_id)
            REFERENCES catalogs (id),
    CONSTRAINT fk_slot
        FOREIGN KEY (slot_id)
            REFERENCES slots (id)

);