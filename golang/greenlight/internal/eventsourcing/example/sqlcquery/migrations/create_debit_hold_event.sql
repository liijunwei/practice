CREATE TABLE debit_hold_event (
  aggregate_id uuid NOT NULL,
  version int NOT NULL,
  parent_id uuid NOT NULL,
  event_type VARCHAR (50),
  payload jsonb NOT NULL,
  created_at timestamp without time zone NOT NULL,
  PRIMARY KEY (aggregate_id, version)
);
